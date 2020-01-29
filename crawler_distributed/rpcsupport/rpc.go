package rpcsupport

import (
	"google.golang.org/grpc"
	"log"
	"net"
	pb "reptiles/crawler_distributed/proto"
)

func ServeRpc(
	host string, service pb.ReptilesServer) error {

	Server := grpc.NewServer()
	pb.RegisterReptilesServer(Server, service)

	listener, err := net.Listen("tcp", host)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("Listening on %s", host)

	for {
		go func() {
			if err := Server.Serve(listener);err != nil{
				log.Fatalf("failed to listen: %v", err)
			}
		}()
	}
}

func NewClient(host string) (pb.ReptilesClient, error) {
	conn, err := grpc.Dial(host)
	if err != nil {
		return nil, err
	}
	client := pb.NewReptilesClient(conn)
	return client, nil
}
