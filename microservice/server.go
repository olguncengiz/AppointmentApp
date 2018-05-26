package main

import (
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "github.com/olguncengiz/AppointmentApp/microservice/appointment"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50888"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

//var globalName []pb.AppointmentInfo = make([]pb.AppointmentInfo, 0)
var appointmentDb = make(map[string]pb.AppointmentInfo)

// SayHello implements helloworld.GreeterServer
func (s *server) RequestAppointment(ctx context.Context, in *pb.AppointmentReq) (*pb.AppointmentRep, error) {
	clientName := in.AppInfo.ClientName
	appointmentDb[clientName] = *in.AppInfo
	log.Printf("DB: %s", appointmentDb)

	_, chk := appointmentDb["user1"]
	if chk {
		log.Printf("user1 Requested An Appointment")
	}

	return &pb.AppointmentRep{Message: "Appointment Request Received"}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAppointmentServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
