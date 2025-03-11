package main

import (
	"context"
	pb "grpc-calculator/proto"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewCalculatorClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	addResp, err := client.Add(ctx, &pb.AddRequest{A: 10, B: 20})
	if err != nil {
		log.Fatalf("could not add: %v", err)
	}
	log.Printf("Sum response: %d", addResp.Result)

	mulResp, err := client.Multiply(ctx, &pb.MultiplyRequest{A: 10, B: 20})
	if err != nil {
		log.Fatalf("could not multiply: %v", err)
	}
	log.Printf("Multiple response: %d", mulResp.Result)

}
