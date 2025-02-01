package main

// コマンドから「curl http://localhost:8080/ping」で「{"message":"pong"}」が返ってくることを確認
import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run(":8080") // ポート8080で起動
}
