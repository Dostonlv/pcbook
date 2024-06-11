package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/Dostonlv/pcbook/pb"
	"github.com/Dostonlv/pcbook/sample"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	serverAddress := flag.String("address", "", "the server address")
	flag.Parse()
	log.Printf("dial server %s", *serverAddress)

	conn, err := grpc.Dial(*serverAddress, grpc.WithInsecure())
	if err != nil {
		fmt.Println("cannot dial server: %v", err)
		return
	}

	laptopClient := pb.NewLaptopServiceClient(conn)

	laptop := sample.NewLaptop()

	req := &pb.CreateLaptopRequest{
		Laptop: laptop,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := laptopClient.CreateLaptop(ctx, req)
	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.AlreadyExists {
			log.Println("laptop already exists")
		} else {
			log.Println("cannot  create laptop: ", err)
			return
		}
		return
	}
	log.Printf("created laptop with id: %s", res.Id)
}
