package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"bgray/taskApi/internal/handler/http"
	"bgray/taskApi/internal/infrastructure/config"
	"bgray/taskApi/internal/infrastructure/databases"
	"bgray/taskApi/internal/infrastructure/repository"
	"bgray/taskApi/internal/infrastructure/security"
	"bgray/taskApi/internal/usecase/task"
	uSerCase "bgray/taskApi/internal/usecase/user"
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

	taskRepo := repository.NewPostgresTaskRepository(db)
	UserRepo := repository.NewPostgresUserRepository(db)

	createTaskUseCase := task.NewCreateTask(taskRepo)
	getTaskUseCase := task.NewGetByIdTaskUseCase(taskRepo)
	getAllTasksUseCase := task.NewGetAllTasksUseCase(taskRepo)
	deleteTaskById := task.NewDeleteTaskById(taskRepo)
	completedTask := task.NewCompletedTask(taskRepo)
	hasher := security.NewBcryptHasher()
	createUser := uSerCase.NewCreateUser(UserRepo, hasher)
	getTasksByUserID := task.NewGetTasksByUserID(taskRepo)

	taskHandler := http.NewTaskHandler(createTaskUseCase, getTaskUseCase, getAllTasksUseCase, deleteTaskById, completedTask, createUser, getTasksByUserID)

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
		api.POST("/tasks", taskHandler.CreateTask)
		api.GET("/tasks/:id", taskHandler.GetTask)
		api.GET("/tasks", taskHandler.GetAllTasks)
		api.DELETE("/tasks/:id", taskHandler.DeleteTaskById)
		api.PATCH("/tasks/:id/completed", taskHandler.CompletedTask)
		api.POST("/users/register", taskHandler.CreatedUser)
		api.GET("/users/:id/tasks", taskHandler.GetTasksByUserID)
	}

	fmt.Println("Server starting on :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Error starting server:", err)
	}
}
