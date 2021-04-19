/*
 *
 * stolen/adapted from the default greeter service
 */

// Package main implements a client for Greeter/PirateGreeter service.
package main

import (
  "context"
  "log"
  "time"
  "fmt"

  //opts
  "flag"

  "google.golang.org/grpc"
  pb "apigee/examples/grpcserver"
)

const (
  address     = "localhost:9090"
  defaultName = "world"
)

func main() {
  //key := flag.String("k", "", "an api key ") 
  host := flag.String("h", "localhost", "remote host to connect to") 
  port := flag.String("p", "9090", "remote port to connect to")
  pirate := flag.Bool("pirate", false, "set if you want a pirate greeting")
  
  flag.Parse()

  // Set up a connection to the server.
  address := fmt.Sprintf("%s:%s",*host,*port)
  conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
  if err != nil {
          log.Fatalf("did not connect: %v", err)
  }
  defer conn.Close()
  c := pb.NewMessengerServiceClient(conn)

  // Contact the server and print out its response.
  name := defaultName
  if ( len(flag.Args()) > 0) {
          name = flag.Args()[0]
  }

  if *pirate == true {
    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()
    r, err := c.GetPirateGreeting(ctx, &pb.MessengerRequest{Msg: name})
    if err != nil {
            log.Fatalf("could not greet: %v", err)
    }
    log.Printf("Greeting: %s", r.GetMsg())
  } else {
    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()
    r, err := c.GetGreeting(ctx, &pb.MessengerRequest{Msg: name})
    if err != nil {
            log.Fatalf("could not greet: %v", err)
    }
    log.Printf("Greeting: %s", r.GetMsg())
  }
}
