package main

import (
	pb "github.com/baddin/todo-service/proto"
	"google.golang.org/grpc"
	"context"
	"log"
	"fmt"
	"os"
)
const (
	address = "localhost:50051" //the server address
)



func main(){

	// connecting the grpc server
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	//making a new client for the service
	client := pb.NewTodoTasksClient(conn)


	//getting all the tasks from the client
	getAll, err := client.GetTasks(context.Background(), &pb.GetRequest{})

	//iterating in the tasks
	for _, v := range getAll.Tasks {
		log.Println(v)
	}
	//testing the "DoneTask"
	fmt.Print("type the task ID if you done it or type 0 to quite: ")
	var input int32
	fmt.Scanln(&input)
	if input == 0 {
		os.Exit(0)
	} else {
		done, err := client.DoneTask(context.Background(), &pb.DoneRequest{Id:input})
		if err != nil {
			panic(err)
		}
		log.Println("Task %v is Done!", done.Task.Title)
	}

}
