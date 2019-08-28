package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "../pb"
	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial("localhost:55555", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer connection.Close()

	client := pb.NewIncrementServiceClient(connection)

	// タイムアウトを1秒に設定する
	context, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// 適当な数を送信してみる
	var number int32 = 1000

	// 送信処理を呼び出す
	response, err := client.Increment(context, &pb.IncrementRequest{Number: number})
	if err != nil {
		log.Println(err)
	}

	fmt.Println(response.GetNumber()) // 1001
}
