package zap

import (
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestFunc(t *testing.T) {
	InitLogger()
	defer logger.Sync()
	//for i := 1; i < 50000; i++ {
	//	logger.Info("this is test log")
	//}
	simpleHttpGet("www.google.com")
	simpleHttpGet("http://www.google.com")
}

func TestGinLogger(t *testing.T) {
	InitLogger()
	defer logger.Sync()

	r := gin.New()
	r.Use(GinLogger(logger), GinRecovery(logger, true))
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, "hello world.")
	})

	r.Run()
}
