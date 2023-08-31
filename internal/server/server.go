package server

import (
	"context"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type LogsService interface {
	Insert(ctx context.Context, logString []byte) error
}

type Server struct {
	rabbitServer *amqp.Connection
	logsService  LogsService
}

func NewServer(logsService LogsService) (*Server, error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return nil, err
	}

	return &Server{rabbitServer: conn, logsService: logsService}, nil
}

func (s *Server) CloseConnection() error {
	err := s.rabbitServer.Close()
	return err
}

func (s *Server) Consume(name string) error {
	ch, err := s.rabbitServer.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(name, false, false, false, false, nil)
	if err != nil {
		return err
	}

	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		return err
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			s.logsService.Insert(context.Background(), d.Body)
			log.Printf("log added: %s", string(d.Body))
		}
	}()

	log.Print("Waiting for messages...")
	<-forever
	return nil
}
