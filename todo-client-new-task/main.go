package main

import (
	"encoding/json"
	pb "github.com/baddin/todo-service/proto"
	"io/ioutil"
	"google.golang.org/grpc"
	"os"
	"context"
	"log"
)
const (
	address = "localhost:50051"
	defaultFileName = "defaultTask.json"
)

func parseJsonFile(fileName string)(*pb.Task, error){
	var task *pb.Task
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(data, &task)
	return task, nil
}


func main(){

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := pb.NewTodoTasksClient(conn)

	fileName := defaultFileName
	if len(os.Args) > 1 {
		fileName = os.Args[1]
	}

	task, err := parseJsonFile(fileName)
	if err != nil {
		panic(err)
	}


	r, err := client.NewTask(context.Background(), task)
	if err != nil {
		panic(err)
	}
	log.Printf("Created : %t", r.Created)


}
