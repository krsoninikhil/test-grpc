package main

import (
	"flag"
	"log"
	"fmt"
	"time"
	"context"
	"io"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/krsoninikhil/test-grpc/protos"
)

func oneOnOneResult(client pb.InterviewClient, user *pb.User) {
	log.Printf("Get result for user")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if result, err := client.OneOnOne(ctx, user); err != nil {
		log.Printf("%v", err)
	} else {
		fmt.Printf("Result: %t, %s\n", result.Selected, result.Remark)
	}
}

func onlineScreeningResult(client pb.InterviewClient, batch *pb.Batch) {
	log.Printf("Get result for the batch")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	stream, err := client.OnlineScreening(ctx, batch)
	if err != nil {
		log.Printf("%v", err)
		return
	}
	for {
		result, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error reading stream: %v", err)
		}
		fmt.Printf("Streamed Result: %t, %s\n", result.Selected, result.Remark)

	}
}

func finalRoundResult(client pb.InterviewClient, batch *pb.Batch) {
	log.Printf("Get result report for streamed user")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	stream, err := client.FinalRound(ctx)
	if err != nil {
		log.Fatalf("Cannot get stream: %v", err)
	}

	for _, user := range batch.Users {
		if err := stream.Send(user); err != nil {
			log.Fatalf("Error in sending: %v", err)
		}
	}
	if resultReport, err := stream.CloseAndRecv(); err != nil {
		log.Fatalf("Error in closing stream: %v", err)
	} else {
		for _, result := range resultReport.Results {
			fmt.Printf("Result %t, %s\n", result.Selected, result.Remark)
		}
	}
}


func campusDriveResult(client pb.InterviewClient, batch *pb.Batch) {
	log.Printf("Get streamed result for stream users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	stream, err := client.CampusDrive(ctx)
	if err != nil {
		log.Fatalf("Cannot get stream: %v", err)
	}

	for _, user := range batch.Users {
		if err := stream.Send(user); err != nil {
			log.Printf("Cannot send to stream: %v", err)
		}
	}

	waitc := make(chan struct{})
	// go routine to wait for messages from server
	go func() {
		for {
			result, err := stream.Recv()
			if err == io.EOF {
				close(waitc)
				break
			}
			if err != nil {
				fmt.Printf("Cannot read from stream: %v", err)
			}
			fmt.Printf("Result for duplex: %t, %s\n", result.Selected, result.Remark)
		}
	}()
	stream.CloseSend()
	<-waitc
}


func main() {
	flag.Parse()
	serverAddr := flag.String("addr", "localhost:8080", "Server address in format host:port")
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("Fail to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewInterviewClient(conn)

	batch := pb.Batch{
		Users: []*pb.User{
			&pb.User{Name: "Test User 1", YearOfExp: 5},
			&pb.User{Name: "Test User 2", YearOfExp: 5},
			&pb.User{Name: "Test User 3", YearOfExp: 5},
		},
	}

	oneOnOneResult(client, batch.Users[0])
	onlineScreeningResult(client, &batch)
	finalRoundResult(client, &batch)
	campusDriveResult(client, &batch)
}
