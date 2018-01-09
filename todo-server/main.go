package main

import (
	"net"
	"context"
	pb "github.com/baddin/todo-service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)
const port = ":50051"

type ITask interface {
	Create(*pb.Task)(*pb.Task, error)
	GetAll() []*pb.Task
	GetTask(id int32) (*pb.Task, error)
}


type Task struct {
	tasks []*pb.Task
}

func (tsk *Task) Create(task *pb.Task)(*pb.Task, error){
	tsk.tasks = append(tsk.tasks, task)
	return task, nil
}

func (tsk *Task) GetAll() []*pb.Task {
	return tsk.tasks
}

func (tsk *Task)GetTask(id int32) (*pb.Task, error) {
	for _, v := range tsk.tasks {
		if v.Id ==id {
			return v, nil
		}
	}

	return nil, nil //for now
}



type server struct{
	tsk ITask
}

func (s *server) NewTask(ctx context.Context, request *pb.Task)(*pb.Response, error){
	tasks, err := s.tsk.Create(request)
	if err != nil {
		return nil, err
	}
	return &pb.Response{ Created:true, Task:tasks}, nil
}

func (s *server) GetTasks(ctx context.Context, request *pb.GetRequest)(*pb.Response, error){
	tasks := s.tsk.GetAll()
	return &pb.Response{Tasks:tasks}, nil
}

func (s *server) DoneTask(ctx context.Context, request *pb.DoneRequest)(*pb.Response, error){
	task, _ := s.tsk.GetTask(request.Id)
	task.Done = true
	return &pb.Response{}, nil
}

func main(){
	tsk := &Task{}
	lis, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	pb.RegisterTodoTasksServer(s, &server{tsk})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}

