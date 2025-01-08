package controllers

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/denyherianto/go-fiber-boilerplate/app/models/requests"
	"github.com/denyherianto/go-fiber-boilerplate/app/queries"
	"github.com/denyherianto/go-fiber-boilerplate/app/utils"
	"github.com/denyherianto/go-fiber-boilerplate/configs/cache"

	"github.com/gofiber/fiber/v2"
)

// RenewTokens method for renew access and refresh tokens.
// @Description Renew access and refresh tokens.
// @Summary renew access and refresh tokens
// @Tags Token
// @Accept json
// @Produce json
// @Param refresh_token body string true "Refresh token"
// @Success 200 {object} requests.SignInResponse
// @Security ApiKeyAuth
// @Router /v1/token/renew [post]
func RenewTokens(c *fiber.Ctx) error {
	// Get now time.
	now := time.Now().Unix()

	// Get claims from JWT.
	claims, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		// Return status 500 and JWT parse error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	tokenString := utils.ExtractToken(c)

	// Create a new renew refresh token struct.
	renew := &requests.Renew{}

	// Checking received data from JSON body.
	if err := c.BodyParser(renew); err != nil {
		// Return, if JSON data is not correct.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   fiber.StatusBadRequest,
			"message": err.Error(),
		})
	}

	// Set expiration time from Refresh token of current user.
	expiresRefreshToken, err := utils.ParseRefreshToken(renew.RefreshToken)
	if err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   fiber.StatusBadRequest,
			"message": err.Error(),
		})
	}

	// Checking, if now time greather than Refresh token expiration time.
	if now < expiresRefreshToken {
		// Define user ID.
		userID := claims.UserID

		// Get user by ID.
		foundedUser, err := queries.GetUserByID(userID)
		if err != nil {
			// Return, if user not found.
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error":   fiber.StatusNotFound,
				"message": "user with the given ID is not found",
			})
		}

		// Generate JWT Access & Refresh tokens.
		tokens, err := utils.GenerateNewTokens(userID)
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

		existingRefreshToken, errGet := connRedis.Get(context.Background(), fmt.Sprintf("user:refresh-token:%s", renew.RefreshToken)).Result()
		// Checking, if now time greather than expiration from JWT.
		if errGet != nil || existingRefreshToken != fmt.Sprintf("%d", userID) {
			// Return status 401 and unauthorized error message.
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":   true,
				"message": "unauthorized, check expiration time of your token",
			})
		}

		// Set expires hours count for secret key from .env file.
		accessTokenExpireHours, _ := strconv.Atoi(os.Getenv("JWT_SECRET_KEY_EXPIRE_HOURS_COUNT"))
		refreshTokenExpireHours, _ := strconv.Atoi(os.Getenv("JWT_REFRESH_KEY_EXPIRE_HOURS_COUNT"))

		// Delete access token to Redis.
		errDelAccessTokenFromRedis := connRedis.Del(context.Background(), fmt.Sprintf("user:access-token:%s", strings.Replace(tokenString, ".", "-", -1))).Err()
		if errDelAccessTokenFromRedis != nil {
			// Return status 500 and Redis deletion error.
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":   fiber.StatusInternalServerError,
				"message": errDelAccessTokenFromRedis.Error(),
			})
		}

		// Delete refresh token to Redis.
		errDelRefreshTokenFromRedis := connRedis.Del(context.Background(), fmt.Sprintf("user:refresh-token:%s", renew.RefreshToken)).Err()
		if errDelRefreshTokenFromRedis != nil {
			// Return status 500 and Redis deletion error.
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":   fiber.StatusInternalServerError,
				"message": errDelRefreshTokenFromRedis.Error(),
			})
		}

		// Save access token to Redis.
		errSaveAccessTokenToRedis := connRedis.Set(context.Background(), fmt.Sprintf("user:access-token:%s", strings.Replace(tokens.Access, ".", "-", -1)), userID, time.Hour*time.Duration(accessTokenExpireHours)).Err()
		if errSaveAccessTokenToRedis != nil {
			// Return status 500 and Redis connection error.
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":   fiber.StatusInternalServerError,
				"message": errSaveAccessTokenToRedis.Error(),
			})
		}

		// Save refresh token to Redis.
		errSaveRefreshTokenToRedis := connRedis.Set(context.Background(), fmt.Sprintf("user:refresh-token:%s", tokens.Refresh), userID, time.Hour*time.Duration(refreshTokenExpireHours)).Err()
		if errSaveRefreshTokenToRedis != nil {
			// Return status 500 and Redis connection error.
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":   fiber.StatusInternalServerError,
				"message": errSaveRefreshTokenToRedis.Error(),
			})
		}

		return c.JSON(fiber.Map{
			"data": fiber.Map{
				"name":          foundedUser.Name,
				"email":         foundedUser.Email,
				"access_token":  tokens.Access,
				"refresh_token": tokens.Refresh,
			},
		})
	} else {
		// Return status 401 and unauthorized error message.
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   fiber.StatusUnauthorized,
			"message": "unauthorized, your session was ended earlier",
		})
	}
}

// VerifyTokens method for renew access and refresh tokens.
// @Description Verify access token.
// @Summary verify access token
// @Tags Token
// @Accept json
// @Produce json
// @Success 200 {object} requests.SignInResponse
// @Security ApiKeyAuth
// @Router /v1/token/verify [post]
func VerifyToken(c *fiber.Ctx) error {
	// Get claims from JWT.
	token, err := utils.VerifyToken(c)
	if err != nil {
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

	// Get current access token
	_, errGet := connRedis.Get(context.Background(), fmt.Sprintf("user:access-token:%s", strings.Replace(token.Raw, ".", "-", -1))).Result()

	// Checking, if token not available, return error
	if errGet != nil {
		// Return status 401 and unauthorized error message.
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   fiber.StatusUnauthorized,
			"message": "unauthorized, check expiration time of your token",
		})
	}

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"message": "token is valid",
	})
}
