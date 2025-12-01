package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"bgray/taskApi/internal/domain"
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
	createTaskUseCase  *task.CreateTask
	getTaskUseCase     *task.GetByIdTaskUseCase
	getAllTasksUseCase *task.GetAllTasksUseCase
	deleteTaskByID     *task.DeleteTaskById
	completedTask      *task.CompletedTask
	getTasksByUserID   *task.GetTasksByUserID
	userService        *user.Service
}

func NewTaskHandler(
	createTaskUseCase *task.CreateTask,
	getTaskUseCase *task.GetByIdTaskUseCase,
	getAllTasksUseCase *task.GetAllTasksUseCase,
	deleteTaskByID *task.DeleteTaskById,
	completedTask *task.CompletedTask,
	getTasksByUserID *task.GetTasksByUserID,
	userService *user.Service,
) *TaskHandler {
	return &TaskHandler{
		createTaskUseCase:  createTaskUseCase,
		getTaskUseCase:     getTaskUseCase,
		getAllTasksUseCase: getAllTasksUseCase,
		deleteTaskByID:     deleteTaskByID,
		completedTask:      completedTask,
		getTasksByUserID:   getTasksByUserID,
		userService:        userService,
	}
}

func (h *TaskHandler) CreateTask(ctx *gin.Context) {
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
	}

	createdTask, err := h.createTaskUseCase.Execute(task)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, createdTask)
}

func (h *TaskHandler) GetTask(ctx *gin.Context) {
	taskID, ok := parseTaskID(ctx)
	if !ok {
		return
	}

	task, err := h.getTaskUseCase.Execute(taskID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, task)
}

func (h *TaskHandler) GetAllTasks(ctx *gin.Context) {
	tasks, err := h.getAllTasksUseCase.Execute()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"tasks": tasks})
}

func (h *TaskHandler) DeleteTaskById(ctx *gin.Context) {
	taskID, ok := parseTaskID(ctx)
	if !ok {
		return
	}

	message, err := h.deleteTaskByID.Execute(taskID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": message})
}

func (h *TaskHandler) CompletedTask(ctx *gin.Context) {
	taskID, ok := parseTaskID(ctx)
	if !ok {
		return
	}

	// Obtener status del body
	var req struct {
		Status bool `json:"status"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateTask, err := h.completedTask.Execute(taskID, req.Status)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, updateTask)
}

func (h *TaskHandler) CreatedUser(ctx *gin.Context) {

	var req struct {
		Email    string `json:"email" binding:"required"`
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := domain.User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}

	createdUser, err := h.userService.CreatedUser(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, createdUser)

}

func (h *TaskHandler) GetTasksByUserID(ctx *gin.Context) {
	userIDParam := ctx.Param("id")
	userID, err := strconv.Atoi(userIDParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	tasks, err := h.getTasksByUserID.Execute(userID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"tasks": tasks})
}

func (h *TaskHandler) SignIn(ctx *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.userService.SingIn(req.Username, req.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "login successful",
		"user": gin.H{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
		},
	})
}
