package main

import (
	"log"
	"net"
	"time"
	"context"
	
	proto "tgbot/proto"
	bot "tgbot/telegram"

	"google.golang.org/grpc"
)

type Message struct {
	Text, Priority string
}

type server struct {
	proto.UnimplementedMyServiceServer
}

var Messages []Message

func (s *server) TgSend(ctx context.Context, req *proto.Request) (*proto.Response, error) {

	var newMessage Message

	newMessage.Text = req.GetText()
	newMessage.Priority = req.GetPriority()

	if newMessage.Priority == "high" {

		if len(Messages) > 0 {

			if Messages[0].Priority == "high" {

				var high int
				for _, i := range Messages {
					if i.Priority == "high" {
						high++
					}
				}

				if len(Messages) == high {
					Messages = append(Messages, newMessage)

				} else {
					Messages = append(Messages[:high+1], Messages[high:]...)
					Messages[high] = newMessage
				}
			} else {
				Messages = append([]Message{newMessage}, Messages...)
			}
		} else {
			Messages = append(Messages, newMessage)
		}

	} else if newMessage.Priority == "medium" {

		if len(Messages) > 0 {

			if Messages[0].Priority != "low" {

				var hm int
				for _, i := range Messages {
					if i.Priority != "low" {
						hm++
					}
				}

				if len(Messages) == hm {
					Messages = append(Messages, newMessage)

				} else {
					Messages = append(Messages[:hm+1], Messages[hm:]...)
					Messages[hm] = newMessage
				}

				
			} else {
				Messages = append([]Message{newMessage}, Messages...)
			}
		} else {
			Messages = append(Messages, newMessage)
		}

	} else if newMessage.Priority == "low" {

		Messages = append(Messages, newMessage)

	} else {

		log.Println("This priority has not in data:")
	}

	log.Println(Messages)

	return &proto.Response{Text: newMessage.Text}, nil
}

func Moment() {

	for {

		if len(Messages) > 0 {

			bot.SendToBot(Messages[0].Text)

			Messages = Messages[1:]

			time.Sleep(time.Second * 5)

		}
	}
}

func main() {

	log.Println("Welcome to the server!")

	listener, err := net.Listen("tcp", ":9000")

	if err != nil {
		log.Fatalf("Didn't run this server: %v", err)
	}

	go Moment()

	s := grpc.NewServer()

	proto.RegisterMyServiceServer(s, &server{})

	err = s.Serve(listener)

	if err != nil {
		log.Fatalf("Didn't connection server: %v", err)
	}
}
