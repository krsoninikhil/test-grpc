package main

import "fmt"
import "flag"
import "context"
import "net"
import "log"
import "io"

import "google.golang.org/grpc"

import pb "github.com/krsoninikhil/test-grpc/protos"


type interviewServer struct {
	pb.UnimplementedInterviewServer
}

func (is interviewServer) OneOnOne(ctx context.Context, user *pb.User) (*pb.Result, error) {
	print("One on one")
	return &pb.Result{Selected: true, Remark: fmt.Sprintf("Above average, %s", user.Name)}, nil
}

func (is interviewServer) OnlineScreening(batch *pb.Batch, stream pb.Interview_OnlineScreeningServer) error {
	for _, user := range batch.Users {
		result := &pb.Result{Selected: true, Remark: fmt.Sprintf("Good work %s", user.Name)}
		if err := stream.Send(result); err != nil {
			return err
		}
	}
	return nil
}

func (is interviewServer) FinalRound(stream pb.Interview_FinalRoundServer) error {
	var users []*pb.User
	for {
		user, err := stream.Recv()
		if err == io.EOF {
			resultReport := &pb.ResultReport{Results: []*pb.Result{}}
			for _, user := range users {
				resultReport.Results = append(
					resultReport.Results,
					&pb.Result{Selected: true, Remark: fmt.Sprintf("Good work streamed: %s", user.Name)},
				)
			}
			err := stream.SendAndClose(resultReport)
			return err
		}
		if err != nil {
			log.Fatalf("Failed to receive: %v", err)
		}
		users = append(users, user)
	}

}

func (is interviewServer) CampusDrive(stream pb.Interview_CampusDriveServer) error {
	for {
		user, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Printf("Failed to receive: %v", err)
			return err
		}
		err = stream.Send(&pb.Result{Selected: true, Remark: fmt.Sprintf("Well done %s", user.Name)})
		if err != nil {
			log.Printf("Failed to send: %v", err)
			return err
		}
	}
}

func main() {
	port := flag.Int("port", 8080, "The server port")
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("Failed to start server %v", err)
	}
	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterInterviewServer(grpcServer, interviewServer{})
	grpcServer.Serve(lis)
}
