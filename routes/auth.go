package routes

import (
	// "fmt"
	"fproj/controllers"
	// "html/template"
	"log"

	// "fmt"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	// "gorm.io/gorm"
)

// var requestCount = prometheus.NewCounter(
// 	prometheus.CounterOpts{
// 		Name: "HandlerCount",
// 		Help: "Count of counters that count requests",
// 	},
// )

func AuthRoutes(r *gin.Engine) {
	// Регистрация маршрутов
	r.GET("/", func(c *gin.Context) {
		controllers.IndexHandler(c.Writer, c.Request)
	})

	r.POST("/register", controllers.RegisterHandler)
	// Регистрация метрики Prometheus
	// prometheus.MustRegister(requestCount)
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// Запуск сервера
	log.Println("Server is running on port 8080")
	log.Println("Запуск веб-сервера на http://127.0.0.1:8080")
	r.Run(":8080")
}
