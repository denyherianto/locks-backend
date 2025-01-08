package middleware

import (
	"context"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"

	"github.com/denyherianto/go-fiber-boilerplate/app/queries"
	"github.com/denyherianto/go-fiber-boilerplate/app/utils"
	"github.com/denyherianto/go-fiber-boilerplate/configs/cache"
)

func TokenValidation(c *fiber.Ctx) error {
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

	return c.Next()
}

func BasicAuthValidation(c *fiber.Ctx) error {
	httpReq, err := adaptor.ConvertRequest(c, false)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	clientID, clientSecret, ok := httpReq.BasicAuth()
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   fiber.StatusInternalServerError,
			"message": "unauthorized, check your credentials",
		})
	}

	// Get Application by Client ID
	application, err := queries.GetApplicationByClientID(clientID)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   fiber.StatusUnauthorized,
			"message": "authorization failed, check your credentials",
		})
	}

	compareClientSecret := utils.ComparePasswords(application.ClientSecret, clientSecret)
	if !compareClientSecret {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   fiber.StatusBadRequest,
			"message": "authorization failed, incorrect Client Secret",
		})
	}

	return c.Next()
}
