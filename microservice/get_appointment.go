package main

import (
	"fmt"
	"log"
	"os"
	"time"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "github.com/olguncengiz/AppointmentApp/microservice/appointment"
)

const (
	address     = "localhost:50888"
	defaultName = ""
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

	if len(os.Args) > 1 {
		rName = os.Args[1]
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetAppointments(ctx, &pb.ClientInfo{Name: rName})
	if err != nil {
		log.Fatalf("could not get appointment: %v", err)
	} else {
		for i, appInfo := range r.Appointments {
			text := fmt.Sprintf("Client Name: %s, Date: %s, Time: %s, Status: %s", appInfo.Client.Name, appInfo.Date, appInfo.Time, appInfo.Status)
			log.Printf("%s: %s", i, text)
	    }
	}
}
