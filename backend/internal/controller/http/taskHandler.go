package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"bgray/taskApi/internal/domain"
	"bgray/taskApi/internal/dto"
	"bgray/taskApi/internal/services/task"
	"bgray/taskApi/internal/services/user"
)

func parseTaskID(ctx *gin.Context) (int, bool) {
	idStr := ctx.Param("id")
	taskID, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid task id"})
		return 0, false
	}
	return taskID, true
}

type TaskHandler struct {
	taskService    *task.Service
	userService    *user.Service
	tokenGenerator domain.TokenGenerator
}

func NewTaskHandler(
	taskService *task.Service,
	userService *user.Service,
	tokenGenerator domain.TokenGenerator,
) *TaskHandler {
	return &TaskHandler{
		taskService:    taskService,
		userService:    userService,
		tokenGenerator: tokenGenerator,
	}
}

func (h *TaskHandler) CreateTask(ctx *gin.Context) {
	// Obtener userID del token (ya validado por el middleware)
	userID, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
		return
	}

	var req struct {
		Title     string `json:"title" binding:"required"`
		Completed bool   `json:"completed"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task := domain.Task{
		Title:     req.Title,
		Completed: req.Completed,
		UserID:    userID.(int),
	}

	createdTask, err := h.taskService.CreateTask(task)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, createdTask)
}

func (h *TaskHandler) DeleteTaskById(ctx *gin.Context) {
	// Obtener userID del token
	userID, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
		return
	}

	taskID, ok := parseTaskID(ctx)
	if !ok {
		return
	}

	// Validar que la tarea pertenezca al usuario
	task, err := h.taskService.GetTaskById(taskID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
		return
	}

	if task.UserID != userID.(int) {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "you don't have permission to delete this task"})
		return
	}

	message, err := h.taskService.DeleteTaskById(taskID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": message})
}

func (h *TaskHandler) CompletedTask(ctx *gin.Context) {
	// Obtener userID del token
	userID, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
		return
	}

	taskID, ok := parseTaskID(ctx)
	if !ok {
		return
	}

	task, err := h.taskService.GetTaskById(taskID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
		return
	}

	if task.UserID != userID.(int) {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "you don't have permission to update this task"})
		return
	}

	var req struct {
		Status bool `json:"status"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateTask, err := h.taskService.CompletedTask(taskID, req.Status)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, updateTask)
}

func (h *TaskHandler) CreatedUser(ctx *gin.Context) {

	var req dto.CreateUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.userService.CreatedUser(ctx.Request.Context(), req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := dto.CreatedUserResponse{
		Message: "User created successfully",
	}

	ctx.JSON(http.StatusCreated, response)

}

func (h *TaskHandler) GetTasksByUserID(ctx *gin.Context) {
	userID, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
		return
	}

	tasks, err := h.taskService.GetTasksByUserID(userID.(int))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"tasks": tasks})
}

func (h *TaskHandler) Login(ctx *gin.Context) {
	var req dto.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.userService.SingIn(ctx.Request.Context(), req)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	token, err := h.tokenGenerator.GenerateToken(*user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
		return
	}

	response := dto.LoginResponse{
		Token: token,
	}

	ctx.JSON(http.StatusOK, response)
}
