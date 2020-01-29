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


func (ItemSaverService) Process(
	ctx context.Context, req *pb.ProcessRequest,
	result *pb.ProcessResult) error {
	engineReq, err := t.DeserializeRequest(req)
	if err != nil {
		return err
	}

	engineResult, err := engine.Worker(engineReq)
	if err != nil {
		return err
	}

	*result = t.SerializeResult(engineResult)
	return nil
}


func (s *ItemSaverService) Save(
	item pb.Item, result *string) error {
	err := persist.Save(s.Client, s.Index, item)
	log.Printf("Item %v saved.", item)
	if err == nil {
		*result = "ok"
	} else {
		log.Printf("Error saving item %v: %v",
			item, err)
	}
	return err
}
