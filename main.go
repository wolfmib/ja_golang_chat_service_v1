package main

import (
	"os"

	proto "github.com/wolfmib/ja_golang_chat_service_v1/proto"
	glog "google.golang.org/grpc/grpclog"

	"log"
	"net"
	"sync"

	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

var JAlog glog.LoggerV2

// 🎬: Initial
func init() {
	//🖨 Setting Loger: all (info, warining, error) message  is going to stdout
	JAlog = glog.NewLoggerV2(os.Stdout, os.Stdout, os.Stdout)
}

// 🕋 Create Connection Structure:
type Connection struct {
	stream proto.Broadcast_CreateStreamServer
	id     string
	active bool
	error  chan error
}

// A Collection of Connection
type Server struct {
	ConnectionList []*Connection
}

// Create Method
func (s *Server) CreateStream(pconn *proto.Connect, input_stream proto.Broadcast_CreateStreamServer) error {
	// 🦉 checking later , connections or connection
	conn := &Connection{
		stream: input_stream,
		id:     pconn.User.Id,
		active: true,
		error:  make(chan error),
	}

	// Append
	s.ConnectionList = append(s.ConnectionList, conn)

	//
	return <-conn.error
}

func (s *Server) BroadcastMessage(ctx context.Context, msg *proto.Message) (*proto.Close, error) {
	//Go routine
	waitAgent := sync.WaitGroup{}
	doneSignal := make(chan int)

	// Brocase to all the connections we had so far
	for _, conn := range s.ConnectionList {
		//Create wait-signal for each connection
		waitAgent.Add(1)

		// Excute go function
		go func(msg *proto.Message, conn *Connection) {
			defer waitAgent.Done()

			// Pass message to the client
			if conn.active {
				err := conn.stream.Send(msg)
				JAlog.Info("[💻]: Sending message to : ", conn.stream)

				if err != nil {
					JAlog.Errorf("[💻]: Error with stream:  %s - Error : %v", conn.stream, err)
					conn.active = false
					conn.error <- err
				}
			}
		}(msg, conn)

	}

	go func() {
		waitAgent.Wait()
		close(doneSignal)
	}()

	//Block until the previous go-function close the doneSignal
	<-doneSignal

	return &proto.Close{}, nil

}

func main() {

	var connections []*Connection

	// copy-paste connections to Server
	server := &Server{connections}

	grpcServer := grpc.NewServer()
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("error creating the server %v", err)
	}

	//🖨 Lib from grpclog
	JAlog.Info("Starting server at port :8080")

	proto.RegisterBroadcastServer(grpcServer, server)
	grpcServer.Serve(listener)

}
