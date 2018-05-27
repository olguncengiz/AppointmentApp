package main

import (
	"log"
	"net"
	"sync"
	"time"
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
var mutex = &sync.Mutex{}
var appointmentDb = make(map[string]pb.AppointmentInfo)

// RequestAppointment implements appointment.RequestAppointment
func (s *server) RequestAppointment(ctx context.Context, in *pb.AppointmentReq) (*pb.AppointmentRep, error) {
	clientName := in.AppInfo.Client.Name

	mutex.Lock()
	appInfo, chk := appointmentDb[clientName]
	mutex.Unlock()

	if !chk || appInfo.Status != "a" {
		mutex.Lock()
		// This value can be increased to see mutex is working well
		time.Sleep(time.Millisecond)
		appointmentDb[clientName] = *in.AppInfo
		mutex.Unlock()
		
		log.Printf("DB: %s", appointmentDb)

		return &pb.AppointmentRep{Message: "Appointment Request Received From " + clientName}, nil
	} else {
		return &pb.AppointmentRep{Message: clientName + " Already Has An Appointment"}, nil
	}
}

// GetAppointments implements appointment.GetAppointments
func (s *server) GetAppointments(ctx context.Context, in *pb.ClientInfo) (*pb.AppointmentList, error) {
	appInfo := &pb.AppointmentInfo{Client: nil, Date: "a", Time: "b", Status: "c"}
	appList := []*pb.AppointmentInfo{appInfo}
	return &pb.AppointmentList{Appointments: appList}, nil
}

// MoveAppointment implements appointment.MoveAppointment
func (s *server) MoveAppointment(ctx context.Context, in *pb.AppointmentReq) (*pb.AppointmentRep, error) {
	clientName := in.AppInfo.Client.Name

	mutex.Lock()
	_, chk := appointmentDb[clientName]
	mutex.Unlock()

	if chk {		
		// This value can be increased to see mutex is working well
		time.Sleep(time.Millisecond)

		mutex.Lock()
		appointmentDb[clientName] = *in.AppInfo
		mutex.Unlock()
		
		log.Printf("DB: %s", appointmentDb)

		return &pb.AppointmentRep{Message: "Appointment Moved For " + clientName}, nil
	} else {
		return &pb.AppointmentRep{Message: clientName + " Doesn't Have An Appointment"}, nil
	}
}

// DeleteAppointment implements appointment.DeleteAppointment
func (s *server) DeleteAppointment(ctx context.Context, in *pb.ClientInfo) (*pb.AppointmentRep, error) {
	clientName := in.Name

	mutex.Lock()
	_, chk := appointmentDb[clientName]
	mutex.Unlock()

	if chk {
		mutex.Lock()
		// This value can be increased to see mutex is working well
		time.Sleep(time.Millisecond)
		delete(appointmentDb, clientName)
		mutex.Unlock()
		
		log.Printf("DB: %s", appointmentDb)

		return &pb.AppointmentRep{Message: "Appointment Deleted For " + clientName}, nil
	} else {
		return &pb.AppointmentRep{Message: clientName + " Doesn't Have An Appointment"}, nil
	}
}

// ApproveAppointment implements appointment.ApproveAppointment
func (s *server) ApproveAppointment(ctx context.Context, in *pb.ClientInfo) (*pb.AppointmentRep, error) {
	clientName := in.Name

	mutex.Lock()
	appInfo, chk := appointmentDb[clientName]
	mutex.Unlock()

	if chk {		
		// This value can be increased to see mutex is working well
		time.Sleep(time.Millisecond)

		appInfo.Status = "a"

		mutex.Lock()
		appointmentDb[clientName] = appInfo
		mutex.Unlock()
		
		log.Printf("DB: %s", appointmentDb)

		return &pb.AppointmentRep{Message: "Appointment Approved For " + clientName}, nil
	} else {
		return &pb.AppointmentRep{Message: clientName + " Doesn't Have An Appointment"}, nil
	}
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
