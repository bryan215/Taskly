package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"

	"bgray/taskApi/internal/controller/http"
	"bgray/taskApi/internal/infrastructure/config"
	"bgray/taskApi/internal/infrastructure/databases"
	"bgray/taskApi/internal/infrastructure/middleware"
	"bgray/taskApi/internal/infrastructure/repository"
	"bgray/taskApi/internal/infrastructure/security"
	"bgray/taskApi/internal/services/task"
	"bgray/taskApi/internal/services/user"
)

func main() {
	cfg, err := config.LoadConfig("internal/infrastructure/config/config.yaml")
	if err != nil {
		log.Fatal("Error loading config:", err)
	}

	db, err := databases.NewPostgresDB(cfg)
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	//repository
	taskRepo := repository.NewPostgresTaskRepository(db)
	userRepo := repository.NewPostgresUserRepository(db)
	hasher := security.NewBcryptHasher()
	jwtService := security.NewJWTService(cfg.JWTSecret, time.Hour*24)

	//Service
	taskService := task.NewService(taskRepo)
	userService := user.NewService(userRepo, hasher)

	taskHandler := http.NewTaskHandler(
		taskService,
		userService,
		jwtService,
	)

	router := gin.Default()

	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	api := router.Group("/api/v1")
	{
		api.POST("/users/register", taskHandler.CreatedUser)
		api.POST("/users/login", taskHandler.Login)

		protected := api.Group("")
		protected.Use(middleware.ProtectHandler(jwtService))
		{
			protected.POST("/tasks", taskHandler.CreateTask)
			protected.DELETE("/tasks/:id", taskHandler.DeleteTaskById)
			protected.PATCH("/tasks/:id/completed", taskHandler.CompletedTask)
			protected.GET("/tasks", taskHandler.GetTasksByUserID)
		}
	}

	fmt.Println("Server starting on :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Error starting server:", err)
	}
}
