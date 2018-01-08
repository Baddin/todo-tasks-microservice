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
	address = "localhost:50051" //the server address
	defaultFileName = "defaultTask.json"
)



func parseJsonFile(fileName string)(*pb.Task, error){
	/*a simple function that decode the json file into a task struct*/
	var task *pb.Task
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(data, &task)
	return task, nil
}


func main(){
	// connecting the grpc server
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	//making a new client for the service
	client := pb.NewTodoTasksClient(conn)


	fileName := defaultFileName
	//getting the json file name from the command line first argument
	if len(os.Args) > 1 {
		fileName = os.Args[1]
	}

	//decoding the json file using the parseJsonFile function
	task, err := parseJsonFile(fileName)
	if err != nil {
		panic(err)
	}

	//making a request to make a new task
	r, err := client.NewTask(context.Background(), task)
	if err != nil {
		panic(err)
	}
	log.Printf("Created : %t", r.Created)


}
