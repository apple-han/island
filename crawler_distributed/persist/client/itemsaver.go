package client

import (
	"context"
	"log"
	pb "island/crawler_distributed/proto"
	"time"

	"island/crawler_distributed/rpcsupport"
)

func ItemSaver(
	host string) (chan pb.Item, error) {
	c, err := rpcsupport.NewClient(host)
	if err != nil {
		return nil, err
	}
	out := make(chan pb.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver: got item "+
				"#%d: %s", itemCount, item.Car)
			itemCount++

			// Call RPC to save item
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			_, err := c.SaveItem(ctx, &pb.SaveItemRequest{Item: &item})

			if err != nil {
				log.Printf("Item Saver: error "+
					"saving item %v: %v",
					item, err)
				}
		}
	}()

	return out, nil
}
