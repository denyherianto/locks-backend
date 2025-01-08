package controllers

import (
	"context"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/denyherianto/go-fiber-boilerplate/app/constants"
	"github.com/denyherianto/go-fiber-boilerplate/app/models/entities"
	"github.com/denyherianto/go-fiber-boilerplate/app/models/requests"
	"github.com/denyherianto/go-fiber-boilerplate/app/queries"
	"github.com/denyherianto/go-fiber-boilerplate/app/utils"
	"github.com/denyherianto/go-fiber-boilerplate/app/utils/logger"
	"github.com/denyherianto/go-fiber-boilerplate/configs/cache"
	"github.com/google/uuid"

	"github.com/gofiber/fiber/v2"
)

type UserPermissions struct {
	ModuleID   string   `json:"module_id"`
	ModuleName string   `json:"module_name"`
	Access     []string `json:"access"`
}
type UserRoles struct {
	ID          uint              `json:"id"`
	Name        string            `json:"name"`
	Permissions []UserPermissions `json:"permissions"`
}

// UserSignUp method to create a new user.
// @Description Create a new user.
// @Summary create a new user
// @Tags User
// @Accept json
// @Produce json
// @Param email body string true "Name"
// @Param email body string true "Email"
// @Param password body string true "Password"
// @Success 200 {object} entities.User
// @Router /v1/user/register [post]
func UserSignUp(c *fiber.Ctx) error {
	// Create a new user auth struct.
	signUp := &requests.SignUpRequest{}

	// Checking received data from JSON body.
	if err := c.BodyParser(signUp); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   fiber.StatusBadRequest,
			"message": err.Error(),
		})
	}

	// Create a new validator for a User model.
	validate := utils.NewValidator()

	// Validate sign up fields.
	if err := validate.Struct(signUp); err != nil {
		// Return, if some fields are not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   fiber.StatusBadRequest,
			"message": utils.ValidatorErrors(err),
		})
	}

	// Validate password
	err := utils.ValidatePassword(signUp.Password)
	if err != nil {
		// Return, if password is not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   fiber.StatusBadRequest,
			"message": "password should be min. 8 characters and include at least one uppercase letter, one lowercase letter, one number, and one special character",
		})
	}

	// Check existing User with same Email
	foundedUser, _ := queries.GetUserByEmailOrUsername(signUp.Email)
	if foundedUser != nil {
		// Return, if user not found.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   fiber.StatusBadRequest,
			"message": "user with the given email is already registered",
		})
	}

	// Create a new user struct.
	user := &entities.User{}

	// Set initialized default data for user:
	user.Name = signUp.Name
	user.Username = signUp.Username
	user.Email = signUp.Email
	user.PasswordHash = utils.GeneratePassword(signUp.Password)
	user.ResetToken = uuid.New().String()
	user.VerificationToken = uuid.New().String()
	user.UserStatus = constants.ActiveUserStatus // 0 == inactive, 1 == active, 2 == blocked

	// Validate user fields.
	if err := validate.Struct(user); err != nil {
		// Return, if some fields are not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   fiber.StatusBadRequest,
			"message": utils.ValidatorErrors(err),
		})
	}

	// Create a new user with validated data.
	if err := queries.CreateUser(user); err != nil {
		// Return status 500 and create user process error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Delete password hash field from JSON view.
	user.PasswordHash = ""

	// Log Activity
	logger.LogActivity(c, &entities.Activity{
		UserID:        user.ID,
		ApplicationID: 0,
		ModuleName:    "Auth",
		FunctionName:  "UserSignUp",
		TableName:     "users",
		ReferenceID:   user.ID,
		OperationType: "CREATE",
		EndpointURL:   "/v1/user/register",
	})

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"data": user,
	})
}

// UserSignIn method to auth user and return access and refresh tokens.
// @Description Auth user and return access and refresh token.
// @Summary auth user and return access and refresh token
// @Tags User
// @Accept json
// @Produce json
// @Param email body string true "User Email"
// @Param password body string true "User Password"
// @Success 200 {object} requests.SignInResponse
// @Router /v1/user/login [post]
func UserSignIn(c *fiber.Ctx) error {
	// Create a new user auth struct.
	signIn := &requests.SignInRequest{}

	// Checking received data from JSON body.
	if err := c.BodyParser(signIn); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   fiber.StatusBadRequest,
			"message": err.Error(),
		})
	}

	// Get user by email.
	foundedUser, err := queries.GetUserByEmailOrUsername(signIn.Identifier)
	if err != nil {
		// Return, if user not found.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   fiber.StatusNotFound,
			"message": "user with the given email or username is not found",
		})
	}

	// Compare given user password with stored in found user.
	compareUserPassword := utils.ComparePasswords(foundedUser.PasswordHash, signIn.Password)
	if !compareUserPassword {
		// Return, if password is not compare to stored in database.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   fiber.StatusBadRequest,
			"message": "wrong user email address or password",
		})
	}

	// Get User Roles by User ID.
	userRoles, _ := queries.GetUserRolesByUserID(foundedUser.ID)
	userRolesSlice := []string{}
	for _, role := range userRoles {
		userRolesSlice = append(userRolesSlice, role.RoleName)
	}

	// Get User Permissions by User ID.
	userPermissions, _ := queries.GetRolePermissionByUserID(foundedUser.ID)
	userPermissionsMap := make(map[string]UserPermissions)

	// Combine permissions by module name.
	for _, item := range userPermissions {
		access := []string{}

		// Check permission for each module.
		if item.PermissionCreate && !slices.Contains(access, "create") {
			access = append(access, "create")
		}
		if item.PermissionRead && !slices.Contains(access, "read") {
			access = append(access, "read")
		}
		if item.PermissionUpdate && !slices.Contains(access, "update") {
			access = append(access, "update")
		}
		if item.PermissionDelete && !slices.Contains(access, "delete") {
			access = append(access, "delete")
		}

		// Check if role name exists in map.
		if _, exists := userPermissionsMap[item.ModuleID]; !exists {
			userPermissionsMap[item.ModuleID] = UserPermissions{
				ModuleID:   item.ModuleID,
				ModuleName: item.ModuleName,
				Access:     access,
			}
		}
	}

	// Convert map to slice.
	userPermissionsSlice := []UserPermissions{}
	for _, value := range userPermissionsMap {
		userPermissionsSlice = append(userPermissionsSlice, value)
	}

	// Generate a new pair of access and refresh tokens.
	tokens, err := utils.GenerateNewTokens(foundedUser.ID)
	if err != nil {
		// Return status 500 and token generation error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Create a new Redis connection.
	connRedis, err := cache.RedisConnection()
	if err != nil {
		// Return status 500 and Redis connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Set expires hours count for secret key from .env file.
	accessTokenExpireHours, _ := strconv.Atoi(os.Getenv("JWT_SECRET_KEY_EXPIRE_HOURS_COUNT"))
	refreshTokenExpireHours, _ := strconv.Atoi(os.Getenv("JWT_REFRESH_KEY_EXPIRE_HOURS_COUNT"))

	// Save access token to Redis.
	errSaveAccessTokenToRedis := connRedis.Set(context.Background(), fmt.Sprintf("user:access-token:%s", strings.Replace(tokens.Access, ".", "-", -1)), foundedUser.ID, time.Hour*time.Duration(accessTokenExpireHours)).Err()
	if errSaveAccessTokenToRedis != nil {
		// Return status 500 and Redis connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   fiber.StatusInternalServerError,
			"message": errSaveAccessTokenToRedis.Error(),
		})
	}

	// Save refresh token to Redis.
	errSaveRefreshTokenToRedis := connRedis.Set(context.Background(), fmt.Sprintf("user:refresh-token:%s", tokens.Refresh), foundedUser.ID, time.Hour*time.Duration(refreshTokenExpireHours)).Err()
	if errSaveRefreshTokenToRedis != nil {
		// Return status 500 and Redis connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   fiber.StatusInternalServerError,
			"message": errSaveRefreshTokenToRedis.Error(),
		})
	}

	// Log Activity
	logger.LogActivity(c, &entities.Activity{
		UserID:        foundedUser.ID,
		ApplicationID: 1,
		ModuleName:    "Auth",
		FunctionName:  "UserSignIn",
		TableName:     "users",
		ReferenceID:   foundedUser.ID,
		OperationType: "READ",
		EndpointURL:   "/v1/user/login",
	})

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"data": fiber.Map{
			"id":            foundedUser.ID,
			"name":          foundedUser.Name,
			"username":      foundedUser.Username,
			"email":         foundedUser.Email,
			"roles":         userRolesSlice,
			"permissions":   userPermissionsSlice,
			"access_token":  tokens.Access,
			"refresh_token": tokens.Refresh,
		},
	})
}

// UserSignOut method to de-authorize user and delete refresh token from Redis.
// @Description De-authorize user and delete refresh token from Redis.
// @Summary de-authorize user and delete refresh token from Redis
// @Tags User
// @Accept json
// @Produce json
// @Success 204 {string} status "ok"
// @Security ApiKeyAuth
// @Router /v1/user/logout [post]
func UserSignOut(c *fiber.Ctx) error {
	// Get claims from JWT.
	tokenString := utils.ExtractToken(c)

	// Create a new Redis connection.
	connRedis, err := cache.RedisConnection()
	if err != nil {
		// Return status 500 and Redis connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Delete refresh token to Redis.
	errDelAccessTokenFromRedis := connRedis.Del(context.Background(), fmt.Sprintf("user:access-token:%s", strings.Replace(tokenString, ".", "-", -1))).Err()
	if errDelAccessTokenFromRedis != nil {
		// Return status 500 and Redis deletion error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   fiber.StatusInternalServerError,
			"message": errDelAccessTokenFromRedis.Error(),
		})
	}

	// Return status 204 no content.
	return c.SendStatus(fiber.StatusOK)
}
