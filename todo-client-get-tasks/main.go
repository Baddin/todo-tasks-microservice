package main

import (
	pb "github.com/baddin/todo-service/proto"
	"google.golang.org/grpc"
	"context"
	"log"
)
const (
	address = "localhost:50051"
)



func main(){

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := pb.NewTodoTasksClient(conn)

	getAll, err := client.GetTasks(context.Background(), &pb.GetRequest{})
	for _, v := range getAll.Tasks {
		log.Println(v)
	}

}
