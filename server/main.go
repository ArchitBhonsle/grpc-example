package main

import (
	"context"
	"log"
	"net"

	"github.com/ArchitBhonsle/grpc-example/schema"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

var todos []string = make([]string, 0)

type server struct {
	schema.UnimplementedTodoManagerServer
}

func (s *server) AddTodo(ctx context.Context, input *schema.Todo) (*empty.Empty, error) {
	log.Println("Adding:", input.GetName())
	todos = append(todos, input.GetName())
	return &empty.Empty{}, nil
}

func (s *server) ReadTodos(ctx context.Context, input *empty.Empty) (*schema.Todos, error) {
	log.Println("Reading")
	t := make([]*schema.Todo, 0)
	for _, v := range todos {
		t = append(t, &schema.Todo{Name: v})
	}

	return &schema.Todos{
		Todos: t,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalln(err)
	}

	s := grpc.NewServer()
	schema.RegisterTodoManagerServer(s, &server{})

	log.Println("Listening at", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalln(err)
	}
}
