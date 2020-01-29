package main

import (
	"reptiles/crawler/config"
	"reptiles/crawler/engine"
	"reptiles/crawler/persist"
	"reptiles/crawler/scheduler"
	"reptiles/crawler/xcar/parser"
)

func main() {
	itemChan, err := persist.ItemSaver(
		config.ElasticIndex)
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      100,
		ItemChan:         itemChan,
		RequestProcessor: engine.Worker,
	}
	e.Run(engine.Request{
		Url: "http://www.starter.url.here",
		Parser: engine.NewFuncParser(
			parser.ParseCarList,
			config.ParseCarList),
	})
}
