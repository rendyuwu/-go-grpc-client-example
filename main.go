package main

import (
	"context"
	"fmt"
	pb "github.com/rendyuwu/go-grpc-client-example/model"
	"google.golang.org/grpc"
	"log"
)

func main() {
	var username string
	fmt.Print("Input a username: ")
	fmt.Scan(&username)
	conn, err := grpc.Dial("localhost:1200", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("error when dial %v", err)
	}

	client := pb.NewUserClient(conn)

	req := pb.UserRequest{Username: username}

	user, err := client.FindUserByEmail(context.Background(), &req)
	if err != nil {
		log.Fatalf("error when find user %v", err)
	}

	fmt.Println(user)
}
