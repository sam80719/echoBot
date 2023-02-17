package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	line "github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

func HandleMessage(c *gin.Context) {
	viper.SetConfigName("config")
	viper.AddConfigPath("./config")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}

	channelSecret := viper.GetString("channel_secret")
	channelAccessToken := viper.GetString("channel_access_token")

	bot, err := line.New(channelSecret, channelAccessToken)
	if err != nil {
		log.Fatal(err)
	}

	events, err := bot.ParseRequest(c.Request)
	if err != nil {
		if err == line.ErrInvalidSignature {
			c.AbortWithStatus(400)
		} else {
			c.AbortWithStatus(500)
		}
		return
	}
	log.Printf("Received events: %+v\n", events)
	for _, event := range events {
		if event.Type == line.EventTypeMessage {
			log.Printf("Received message: %+v\n", event)
			// handle diff type event
			switch message := event.Message.(type) {
			case *line.TextMessage:
				fmt.Printf("Received text message: %s\n", message.Text)
			case *line.ImageMessage:
				fmt.Println("Received image message")
			case *line.VideoMessage:
				fmt.Println("Received video message")
			case *line.AudioMessage:
				// 收到音訊訊息
				fmt.Println("Received audio message")
			case *line.LocationMessage:
				// 收到位置訊息
				fmt.Println("Received location message")
			case *line.StickerMessage:
				// 收到貼圖訊息
				fmt.Println("Received sticker message")
			}
		}
	}
	c.JSON(http.StatusOK, gin.H{})
}
