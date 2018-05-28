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
var appointmentDb = make(map[string]appointmentStruct)
type appointmentStruct struct {
	sync.Mutex
	appointmentInfo pb.AppointmentInfo	
}

// RequestAppointment implements appointment.RequestAppointment
func (s *server) RequestAppointment(ctx context.Context, in *pb.AppointmentReq) (*pb.AppointmentRep, error) {
	clientName := in.AppInfo.Client.Name

	mutex.Lock()
	appStruct, chk := appointmentDb[clientName]
	mutex.Unlock()

	if chk && appStruct.appointmentInfo.Status != "a"{
		as := &appointmentStruct{appointmentInfo: *in.AppInfo}
		appStruct.Lock()
		// This value can be increased to see mutex is working well
		time.Sleep(time.Millisecond)
		//mutex.Lock()
		appointmentDb[clientName] = *as
		//mutex.Unlock()
		appStruct.Unlock()
		
		log.Printf("DB: %s", appointmentDb)

		return &pb.AppointmentRep{Message: "Appointment Request Received From " + clientName}, nil
	} else if !chk {
		as := &appointmentStruct{appointmentInfo: *in.AppInfo}
		mutex.Lock()
		// This value can be increased to see mutex is working well
		time.Sleep(time.Millisecond)
		appointmentDb[clientName] = *as
		mutex.Unlock()
		
		log.Printf("DB: %s", appointmentDb)

		return &pb.AppointmentRep{Message: "Appointment Request Received From " + clientName}, nil
	} else {
		return &pb.AppointmentRep{Message: "Can't Request An Appointment For " + clientName}, nil
	}
}

// GetAppointments implements appointment.GetAppointments
func (s *server) GetAppointments(ctx context.Context, in *pb.ClientInfo) (*pb.AppointmentList, error) {
	clientName := in.Name

	if clientName != "" { // Appointment for a client
		mutex.Lock()
		appStruct, chk := appointmentDb[clientName]
		mutex.Unlock()

		if chk {
			appList := []*pb.AppointmentInfo{&appStruct.appointmentInfo}
			log.Printf("DB: %s", appointmentDb)
			return &pb.AppointmentList{Appointments: appList}, nil
		} else {
			log.Printf("DB: %s", appointmentDb)
			return &pb.AppointmentList{Appointments: nil}, nil
		}
	} else { // All appointments
		// Convert map to slice of values.
	    appointments := []*pb.AppointmentInfo{}
	    for _, value := range appointmentDb {
	    	var appInfo pb.AppointmentInfo = value.appointmentInfo
	        appointments = append(appointments, &appInfo)
	    }
		log.Printf("DB: %s", appointmentDb)
		return &pb.AppointmentList{Appointments: appointments}, nil
	}
}

// MoveAppointment implements appointment.MoveAppointment
func (s *server) MoveAppointment(ctx context.Context, in *pb.AppointmentReq) (*pb.AppointmentRep, error) {
	clientName := in.AppInfo.Client.Name

	mutex.Lock()
	appStruct, chk := appointmentDb[clientName]
	mutex.Unlock()

	if chk {
		as := &appointmentStruct{appointmentInfo: *in.AppInfo}
		appStruct.Lock()
		// This value can be increased to see mutex is working well
		time.Sleep(time.Millisecond)
		appointmentDb[clientName] = *as
		appStruct.Unlock()
		
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
		//appStruct.Lock()
		// This value can be increased to see mutex is working well
		time.Sleep(time.Millisecond)
		delete(appointmentDb, clientName)
		//appStruct.Unlock()
		
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
	appStruct, chk := appointmentDb[clientName]
	mutex.Unlock()

	if chk {		
		// This value can be increased to see mutex is working well
		time.Sleep(time.Millisecond)
		appStruct.appointmentInfo.Status = "a"

		appStruct.Lock()
		appointmentDb[clientName] = appStruct
		appStruct.Unlock()
		
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
