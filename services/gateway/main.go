package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

// Gateway Service
//
// This service acts as an API gateway, routing requests to appropriate backend services

// Clients for backend services
type Clients struct {
	// userClient  *client.UserClient   // TODO: Uncomment when client is ready
	// orderClient *client.OrderClient  // TODO: Uncomment when client is ready
}

func main() {
	// Initialize clients
	clients := &Clients{}

	// TODO: Initialize gRPC clients
	// userClient, err := client.NewUserClient("localhost:9001")
	// if err != nil {
	//     log.Fatalf("Failed to create user client: %v", err)
	// }
	// defer userClient.Close()
	// clients.userClient = userClient

	// orderClient, err := client.NewOrderClient("localhost:9002")
	// if err != nil {
	//     log.Fatalf("Failed to create order client: %v", err)
	// }
	// defer orderClient.Close()
	// clients.orderClient = orderClient

	// Create Gin router
	r := gin.Default()

	// Health check endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":    "healthy",
			"service":   "gateway",
			"timestamp": time.Now().Unix(),
		})
	})

	// User service proxy routes
	userGroup := r.Group("/api/users")
	{
		userGroup.GET("/:id", makeGetUserHandler(clients))
		userGroup.POST("/", makeCreateUserHandler(clients))
		userGroup.PUT("/:id", makeUpdateUserHandler(clients))
		userGroup.DELETE("/:id", makeDeleteUserHandler(clients))
	}

	// Order service proxy routes
	orderGroup := r.Group("/api/orders")
	{
		orderGroup.GET("/:id", makeGetOrderHandler(clients))
		orderGroup.POST("/", makeCreateOrderHandler(clients))
		orderGroup.PUT("/:id/status", makeUpdateOrderStatusHandler(clients))
		orderGroup.GET("/user/:userId", makeListUserOrdersHandler(clients))
	}

	// Authentication endpoint
	r.POST("/api/auth/login", makeLoginHandler(clients))

	// Start server with graceful shutdown
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	// Start server in a goroutine
	go func() {
		log.Println("Gateway service starting on port 8080...")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// Give outstanding requests 30 seconds to complete
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exited")
}

// Handler factory functions
func makeGetUserHandler(clients *Clients) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		log.Printf("Getting user: %s", id)

		// TODO: Use actual gRPC client
		// user, err := clients.userClient.GetUser(c.Request.Context(), id)
		// if err != nil {
		//     c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		//     return
		// }

		// Mock response for now
		c.JSON(http.StatusOK, gin.H{
			"id":       id,
			"email":    "user@example.com",
			"username": "testuser",
			"service":  "user-service",
		})
	}
}

func makeCreateUserHandler(clients *Clients) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("Creating user")

		var req map[string]interface{}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// TODO: Use actual gRPC client
		// user, err := clients.userClient.CreateUser(c.Request.Context(), req)
		// if err != nil {
		//     c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		//     return
		// }

		// Mock response for now
		c.JSON(http.StatusCreated, gin.H{
			"id":      "user-123",
			"email":   req["email"],
			"status":  "created",
			"service": "user-service",
		})
	}
}

func makeUpdateUserHandler(clients *Clients) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		log.Printf("Updating user: %s", id)

		var req map[string]interface{}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// TODO: Use actual gRPC client
		// user, err := clients.userClient.UpdateUser(c.Request.Context(), req)
		// if err != nil {
		//     c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		//     return
		// }

		// Mock response for now
		c.JSON(http.StatusOK, gin.H{
			"id":      id,
			"status":  "updated",
			"service": "user-service",
		})
	}
}

func makeDeleteUserHandler(clients *Clients) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		log.Printf("Deleting user: %s", id)

		// TODO: Use actual gRPC client
		// success, err := clients.userClient.DeleteUser(c.Request.Context(), id)
		// if err != nil {
		//     c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		//     return
		// }

		// Mock response for now
		c.JSON(http.StatusOK, gin.H{
			"id":      id,
			"status":  "deleted",
			"service": "user-service",
		})
	}
}

// Order handlers
func makeGetOrderHandler(clients *Clients) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		log.Printf("Getting order: %s", id)

		// TODO: Use actual gRPC client
		// order, err := clients.orderClient.GetOrder(c.Request.Context(), id)
		// if err != nil {
		//     c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		//     return
		// }

		// Mock response for now
		c.JSON(http.StatusOK, gin.H{
			"id":      id,
			"user_id": "user-123",
			"status":  "confirmed",
			"total":   99.99,
			"service": "order-service",
		})
	}
}

func makeCreateOrderHandler(clients *Clients) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("Creating order")

		var req map[string]interface{}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// TODO: Use actual gRPC client
		// order, err := clients.orderClient.CreateOrder(c.Request.Context(), req)
		// if err != nil {
		//     c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		//     return
		// }

		// Mock response for now
		c.JSON(http.StatusCreated, gin.H{
			"id":      "order-123",
			"user_id": req["user_id"],
			"status":  "pending",
			"total":   req["total"],
			"service": "order-service",
		})
	}
}

func makeUpdateOrderStatusHandler(clients *Clients) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		log.Printf("Updating order status: %s", id)

		var req map[string]interface{}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// TODO: Use actual gRPC client
		// order, err := clients.orderClient.UpdateOrderStatus(c.Request.Context(), id, req["status"].(string))
		// if err != nil {
		//     c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		//     return
		// }

		// Mock response for now
		c.JSON(http.StatusOK, gin.H{
			"id":      id,
			"status":  req["status"],
			"service": "order-service",
		})
	}
}

func makeListUserOrdersHandler(clients *Clients) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Param("userId")
		log.Printf("Listing orders for user: %s", userID)

		// TODO: Use actual gRPC client
		// orders, err := clients.orderClient.ListUserOrders(c.Request.Context(), userID, 1, 10)
		// if err != nil {
		//     c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		//     return
		// }

		// Mock response for now
		c.JSON(http.StatusOK, gin.H{
			"orders": []gin.H{
				{
					"id":      "order-123",
					"user_id": userID,
					"status":  "confirmed",
					"total":   99.99,
				},
			},
			"service": "order-service",
		})
	}
}

func makeLoginHandler(clients *Clients) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("User login")

		var req map[string]interface{}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		email, _ := req["email"].(string)
		password, _ := req["password"].(string)

		if email == "" || password == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Email and password required"})
			return
		}

		// TODO: Use actual gRPC client
		// auth, err := clients.userClient.AuthenticateUser(c.Request.Context(), email, password)
		// if err != nil {
		//     c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
		//     return
		// }

		// Mock response for now
		c.JSON(http.StatusOK, gin.H{
			"token": "jwt-token-123",
			"user": gin.H{
				"id":       "user-123",
				"email":    email,
				"username": "testuser",
			},
			"service": "user-service",
		})
	}
}

// User service handlers
func getUserHandler(c *gin.Context) {
	id := c.Param("id")
	// TODO: Forward request to user service
	c.JSON(http.StatusOK, gin.H{
		"message": "Get user " + id,
		"service": "user-service",
	})
}

func createUserHandler(c *gin.Context) {
	// TODO: Forward request to user service
	c.JSON(http.StatusCreated, gin.H{
		"message": "Create user",
		"service": "user-service",
	})
}

func updateUserHandler(c *gin.Context) {
	id := c.Param("id")
	// TODO: Forward request to user service
	c.JSON(http.StatusOK, gin.H{
		"message": "Update user " + id,
		"service": "user-service",
	})
}

func deleteUserHandler(c *gin.Context) {
	id := c.Param("id")
	// TODO: Forward request to user service
	c.JSON(http.StatusOK, gin.H{
		"message": "Delete user " + id,
		"service": "user-service",
	})
}

// Order service handlers
func getOrderHandler(c *gin.Context) {
	id := c.Param("id")
	// TODO: Forward request to order service
	c.JSON(http.StatusOK, gin.H{
		"message": "Get order " + id,
		"service": "order-service",
	})
}

func createOrderHandler(c *gin.Context) {
	// TODO: Forward request to order service
	c.JSON(http.StatusCreated, gin.H{
		"message": "Create order",
		"service": "order-service",
	})
}

func updateOrderStatusHandler(c *gin.Context) {
	id := c.Param("id")
	// TODO: Forward request to order service
	c.JSON(http.StatusOK, gin.H{
		"message": "Update order status " + id,
		"service": "order-service",
	})
}
