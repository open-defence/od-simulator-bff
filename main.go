package main

import (
	"context"
	"log"
	"net"

	pb "./"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedConfigServiceServer
}

func (s *server) GetConfig(ctx context.Context, in *pb.ConfigRequest) (*pb.ConfigResponse, error) {
	configJSON := `{
    "interceptors": [
      {
        "name": "Interceptor1",
        "angle": 45,
        "count": 5,
        "isAutofire": true,
        "maxCount": 10,
        "speed": 2,
        "position": {
          "x": 400,
          "y": 300
        }
      },
      {
        "name": "Interceptor2",
        "angle": 90,
        "count": 3,
        "isAutofire": false,
        "maxCount": 8,
        "speed": 3,
        "position": {
          "x": 500,
          "y": 300
        }
      },
      {
        "name": "Interceptor3",
        "angle": 135,
        "count": 4,
        "isAutofire": true,
        "maxCount": 7,
        "speed": 4,
        "position": {
          "x": 600,
          "y": 300
        }
      }
    ]
  }`
	return &pb.ConfigResponse{ConfigJson: configJSON}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterConfigServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
