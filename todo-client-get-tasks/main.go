package main

import (
	pb "github.com/baddin/todo-service/proto"
	"google.golang.org/grpc"
	"context"
	"log"
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

}
