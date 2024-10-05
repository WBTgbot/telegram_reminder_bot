package handler

import (
	"net/http"
	"telegram_reminder_bot/scheduler"

	"github.com/gin-gonic/gin"
)

// Структукра для представления задачи
type Task struct {
	ChatID   int64  `json:"chat_id"`
	Task     string `json:"task"`
	Interval int    `json:"interval"`
	Unit     string `json:"unit"`
}

// Handler для создания новой задачи
func CreateTask(c *gin.Context) {
	var newTask Task // создаем переменную для новой задачи

	// пробуем связать входящие данные с структурой Task
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // Возвращаем 400 Bad Request в случае ошибки
		return
	}

	// Здесь нужно доавить логику для планирования задачи
	// Примерно так: scheduler.ScheduleTask(newTask.ChatID, newTask.Task, newTask.Interval, newTask.Unit, newTask.Username)

	// Возвращаем успешный ответ с информацией о создании задачи
	c.JSON(http.StatusOK, gin.H{"status": "task scheduled", "task": newTask})
}

// Habdler для получения всех задач
func GetTasks(c *gin.Context) {
	// Тут нужно добавить логику для получения списка задач
	// Примрено: tasks := scheduler.GetAllTasks()

	// Возвращаем успешный ответ с списком задач
	c.JSON(http.StatusOK, gin.H{"tasks": []Task{}})
}

// Handler для удаления задачи по ID
func DeleteTask(c *gin.Context) {
	taskID := c.Param("id") // Получаем ID задачи из URL - параметра

	// Здесь нужно добавить логику для удаления задачи по ID
	// Например:
	err := scheduler.DeleteTask(taskID)

	// проверяем на ошибку при удалении.
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete task"}) // Возвращаем 500 Internal Server Error
		return
	}

	// Возвращаем успешный ответ об удалении задачи
	c.JSON(http.StatusOK, gin.H{"status": "task deleted"})

}
