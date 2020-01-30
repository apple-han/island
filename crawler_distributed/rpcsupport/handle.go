package rpcsupport

import (
	"context"
	"github.com/olivere/elastic/v7"
	"log"
	"reptiles/crawler/engine"
	"reptiles/crawler/persist"
	pb "reptiles/crawler_distributed/proto"
	t "reptiles/crawler_distributed/worker"
)


type RPCService struct {
	Client *elastic.Client
	Index  string
}


func (s *RPCService) Process(
	ctx context.Context, req *pb.ProcessRequest)(*pb.ProcessResult,error){
		engineReq, err := t.DeserializeRequest(req)
	if err != nil {
		return nil, err
	}

	engineResult, err := engine.Worker(engineReq)
	if err != nil {
		return nil, err
	}
	var result = t.SerializeResult(engineResult)

	return &result, nil
}


func (s *RPCService) SaveItem(
	ctx context.Context, item *pb.SaveItemRequest) (*pb.SaveItemResult, error) {
	err := persist.Save(s.Client, s.Index, item.Item)
	log.Printf("Item %v saved.", item.Item)
	if err != nil {
		log.Printf("Error saving item %v: %v",
			item, err)
	}
	return &pb.SaveItemResult{}, err
}
