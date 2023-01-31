package main

import (
	"gin/controller"
	"log"
	"os"
)

func main() {
	router := controller.GetRouter()

	// ポート番号を指定
	apiHost := os.Getenv("API_CONTAINER_IPV4") // APIコンテナIPv4
	apiPort := os.Getenv("API_CONTAINER_PORT") // APIコンテナポート番号
	router.Run(":" + apiPort)
	log.Print("server running at: http://" + apiHost + ":" + apiPort)
}
