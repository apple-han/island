package client

import (
	"context"
	pb "island/crawler_distributed/proto"
	"time"

	"island/crawler/engine"
	"island/crawler_distributed/worker"
)

func CreateProcessor(
	clientChan chan pb.ReptilesClient) engine.Processor {
	return func(
		req engine.Request) (
		engine.ParseResult, error) {

		sReq := worker.SerializeRequest(req)
		c := <-clientChan
		// Call RPC to send work
		ctx, cancel := context.WithTimeout(context.Background(), 1 * time.Second)
		defer cancel()
		sResult, err := c.Process(ctx, &pb.ProcessRequest{
			Url: sReq.Url, SerializedParser: sReq.SerializedParser})
		if err != nil {
			return engine.ParseResult{}, err
		}
		return worker.DeserializeResult(*sResult),
			nil
	}
}
