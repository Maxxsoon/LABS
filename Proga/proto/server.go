package chat

import (
	"context"
	"log"
	"net"

	pb "github.com/Maxxsoon/LABS"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedChatServiceServer
}

func (s *server) SendMessage(ctx context.Context, msg *pb.Message) (*pb.Message, error) {
	log.Printf("Received message from %s: %s", msg.Sender, msg.Text)
	return msg, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterChatServiceServer(s, &server{})

	log.Println("Server is running on port 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
