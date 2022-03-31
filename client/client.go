package main

import (
	"context"
	"fmt"
	pb "github.com/hedeqiang/grpc-demo/api/user"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	fmt.Println("client start")
	conn, err := grpc.Dial("127.0.0.1:8012", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}

	defer conn.Close()
	c := pb.NewUserServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// 创建用户
	createUserRequest := &pb.CreateUserRequest{
		Name:     "hedeqiang",
		Email:    "hedeqiang@88.com",
		Password: "123456",
	}
	user, err := c.CreateUser(ctx, createUserRequest)
	if err != nil {
		log.Fatalf("create user failed: %v", err)
	}

	log.Printf("create user success: %v", user)

	// 获取用户信息 GetUser
	getUserRequest := &pb.GetUserRequest{
		UserId: "1",
	}
	getUser, err := c.GetUser(ctx, getUserRequest)
	if err != nil {
		log.Fatalf("get user failed: %v", err)
	}
	fmt.Printf("user: %v", getUser)

}
