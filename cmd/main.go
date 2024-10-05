package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"telegram_reminder_bot/scheduler"

	"github.com/gin-gonic/gin" // Используем Gin для создания веб - сервиса.
)

func main() {
	// Запускаем сервер планировщика задач
	scheduler.InitScheduler()

	// Создаем новый экземпляр Gin
	router := gin.Default()

	// Обрабатываем POST - запрос для создания задачи
	router.POST("/create-task", func(c *gin.Context) {
		var json struct {
			ChatID   int64  `json:"chat_id"`  // ID чата
			Task     string `json:"task"`     // Описание задачи
			Interval int    `json:"interval"` // Интервал времени
			Unit     string `json:"unit"`     // Единица зимерения (час, день и т.д)
			Username string `json:"username"` // Имя пользователя
		}

		// Связываем входящие данные с структурой json
		if err := c.BindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		// Возвращаем ответ
		c.JSON(http.StatusOK, gin.H{"status": "task scheduled"})
	})

	// Обрабатываем сигнал завершения работы сервера
	go func() {
		signals := make(chan os.Signal, 1)                      // Создаем канал для сигналов
		signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM) // Подписываемся на прерывание и завершение

		// Ждем получения сигнала
		<-signals
		fmt.Println("Shutting down gracefully...")
		scheduler.StopScheduler() // Останавливаем планировщик
		os.Exit(0)                // Выход из программы

	}()

	// Запускаем серевер на порту 8080
	log.Fatal(router.Run(":8080"))

}
