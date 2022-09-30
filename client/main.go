package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ArchitBhonsle/grpc-example/schema"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	c := schema.NewTodoManagerClient(conn)
	log.Println("Client started at", conn.Target())

	for {
		fmt.Print("\n", "Give an instruction. [A/a]dd, [R/r]ead: ")
		choice := ""
		_, err := fmt.Scanln(&choice)
		if err != nil {
			log.Fatalln(err)
		}

		switch choice {
		case "add", "a", "Add", "A":
			fmt.Print("Enter the name of todo: ")
			name := ""
			scanner := bufio.NewScanner(os.Stdin)
			if scanner.Scan() {
				name = scanner.Text()
			} else {
				log.Fatalln(scanner.Err())
			}

			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()

			_, err := c.AddTodo(ctx, &schema.Todo{Name: name})
			if err != nil {
				log.Fatalln(err)
			}
		case "read", "r", "Read", "R":
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()

			todos, err := c.ReadTodos(ctx, &empty.Empty{})
			if err != nil {
				log.Fatalln(err)
			}

			fmt.Println("Current todos are:")
			for _, t := range todos.GetTodos() {
				fmt.Printf("%v,\n", t.GetName())
			}
		default:
			log.Println("Exiting")
		}
	}
}
