package main

import (
	"log"
	"context"

	proto "tgbot/proto"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

	_ "tgbot/docs"

	sF "github.com/swaggo/files"
	gS "github.com/swaggo/gin-swagger"
)

type Message struct {
	Text     string `json:"text"`
	Priority string `json:"priority"`
}

var c proto.MyServiceClient

// @Summary For Bot Message
// @Description Send Message To Bot
// @Tags sender
// @ID message_sender
// @Accept json
// @Produce json
// @Param input body Message true "message info"
// @Success 200
// @Failure 400
// @Router /send [post]
func SendMessage(ctx *gin.Context) {

	var newMessage Message

	err := ctx.BindJSON(&newMessage)

	if err != nil {

		log.Fatalf("Failed to accept message: %v", err)
	}

	res, err := c.TgSend(context.Background(), &proto.Request{
		Text:     newMessage.Text,
		Priority: newMessage.Priority,
	})

	if err != nil {

		log.Fatalf("Failed to call Sender RPC: %v", err)
	}

	ctx.JSON(200, res)
}

// @title Send Message To Bot
// @version 1.0
// @description Telegram Bot which sends messages to channels and groups
// @host localhost:9090
// @BasePath /

func main() {

	log.Println("Welcome to the client!")

	gin.SetMode(gin.ReleaseMode)

	conn, err := grpc.Dial(":9000", grpc.WithInsecure())

	if err != nil {
		log.Printf("Didn't run this client: %s", err)
		return
	}

	defer conn.Close()

	c = proto.NewMyServiceClient(conn)

	router := gin.Default()

	router.POST("/send", SendMessage)
	router.GET("/swagger/*any", gS.WrapHandler(sF.Handler))

	err = router.Run(":9090")

	if err != nil {
		log.Fatalf("Didn't not connection port: %v", err)
	}
}
