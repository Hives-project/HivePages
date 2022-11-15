package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Hives-project/HivePages/pkg/protobuf/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	var opts []grpc.DialOption = []grpc.DialOption{}
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial(":8080", opts...)
	if err != nil {
		log.Panic(err)
	}
	defer conn.Close()

	client := pb.NewPageServiceClient(conn)

	GetPage(client)
}

func GetPage(client pb.PageServiceClient) {
	page, err := client.GetPage(context.Background(), &pb.PageRequest{Uuid: "1"})
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("feature: %v\n", page)
}
