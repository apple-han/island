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


type ItemSaverService struct {
	Client *elastic.Client
	Index  string
}


func (s *ItemSaverService) Process(
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


func (s *ItemSaverService) SaveItem(
	ctx context.Context, item *pb.SaveItemRequest) (*pb.SaveItemResult, error) {
	err := persist.Save(s.Client, s.Index, item)
	log.Printf("Item %v saved.", item)
	if err != nil {
		log.Printf("Error saving item %v: %v",
			item, err)
	}
	return nil, err
}
