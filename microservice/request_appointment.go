package main

import (
	"log"
	"os"
	"time"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "github.com/olguncengiz/AppointmentApp/microservice/appointment"
)

const (
	address     = "localhost:50888"
	defaultName = "world"
	defaultDate = "2018-05-26"
	defaultTime = "13:00"
	defaultStatus = "r"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewAppointmentClient(conn)

	// Contact the server and print out its response.
	rName := defaultName
	rDate := defaultDate
	rTime := defaultTime
	rStatus := defaultStatus

	if len(os.Args) > 1 {
		rName = os.Args[1]
	}
	if len(os.Args) > 2 {
		rDate = os.Args[2]
	}
	if len(os.Args) > 3 {
		rTime = os.Args[3]
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	clientInfo := &pb.ClientInfo{Name: rName}
	appInfo := &pb.AppointmentInfo{Client: clientInfo, Date: rDate, Time: rTime, Status: rStatus}
	r, err := c.RequestAppointment(ctx, &pb.AppointmentReq{AppInfo: appInfo})
	if err != nil {
		log.Fatalf("could not request appointment: %v", err)
	}
	log.Printf("Reply: %s", r.Message)
}
