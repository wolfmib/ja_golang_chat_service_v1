package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"flag"
	"fmt"
	"os"
	"sync"
	"time"

	proto "github.com/wolfmib/ja_golang_chat_service_v1/proto"

	"log"

	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

var client proto.BroadcastClient
var waitAgent *sync.WaitGroup

func init() {
	waitAgent = &sync.WaitGroup{}
}

func connect(user *proto.User) error {
	var streamerror error

	stream, err := client.CreateStream(context.Background(), &proto.Connect{
		User:   user,
		Active: true,
	})

	if err != nil {
		return fmt.Errorf("Connection fail  %v", err)
	}

	waitAgent.Add(1)

	go func(str proto.Broadcast_CreateStreamClient) {
		defer waitAgent.Done()

		for {
			msg, err := str.Recv()
			if err != nil {
				streamerror = fmt.Errorf("Error reading message : %v", err)
				break
			}

			//Print
			fmt.Printf("%v:  %s\n", msg.Id, msg.Content)

		}
	}(stream)

	return streamerror
}

func main() {
	timestamp := time.Now()
	done := make(chan int)

	name := flag.String("N", "Anon", "The name of the user")
	flag.Parse()

	id := sha256.Sum256([]byte(timestamp.String() + *name))

	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Couldnt connect to service: %v", err)
	}

	client = proto.NewBroadcastClient(conn)
	user := &proto.User{
		Id:   hex.EncodeToString(id[:]),
		Name: *name,
	}

	connect(user)

	waitAgent.Add(1)
	go func() {
		defer waitAgent.Done()

		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			msg := &proto.Message{
				Id:        user.Id,
				Content:   scanner.Text(),
				Timestamp: timestamp.String(),
			}

			_, err := client.BroadcastMessage(context.Background(), msg)
			if err != nil {
				fmt.Printf("Error Sending Message: %v", err)
				break
			}
		}

	}()

	go func() {
		waitAgent.Wait()
		close(done)
	}()

	<-done
}
