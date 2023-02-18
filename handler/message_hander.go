package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	linebot "github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

type LineBot struct {
	Viper              *viper.Viper
	Bot                *linebot.Client
	ChannelSecret      string
	ChannelAccessToken string
}

func NewLineBot() (*LineBot, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath("./config")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file: %v", err)
	}

	channelSecret := viper.GetString("channel_secret")
	channelAccessToken := viper.GetString("channel_access_token")

	bot, err := linebot.New(channelSecret, channelAccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize linebot: %v", err)
	}

	return &LineBot{
		Viper:              viper.GetViper(),
		Bot:                bot,
		ChannelSecret:      channelSecret,
		ChannelAccessToken: channelAccessToken,
	}, nil
}

func HandleMessage(c *gin.Context) {
	lineBot, err := NewLineBot()
	if err != nil {
		log.Fatalf("Failed to initialize linebot: %v", err)
	}

	events, err := lineBot.Bot.ParseRequest(c.Request)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			c.AbortWithStatus(400)
		} else {
			c.AbortWithStatus(500)
		}
		return
	}

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			log.Printf("Received message: %+v\n", event)
			// handle diff type event
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				fmt.Printf("Received text message: %s\n", message.Text)
			case *linebot.ImageMessage:
				fmt.Println("Received image message")
			case *linebot.VideoMessage:
				fmt.Println("Received video message")
			case *linebot.AudioMessage:
				// 收到音訊訊息
				fmt.Println("Received audio message")
			case *linebot.LocationMessage:
				// 收到位置訊息
				fmt.Println("Received location message")
			case *linebot.StickerMessage:
				// 收到貼圖訊息
				fmt.Println("Received sticker message")
			}
		}
	}
	c.JSON(http.StatusOK, gin.H{})
}
