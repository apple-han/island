package worker

import (
	"errors"

	"fmt"
	"log"

	"reptiles/crawler/config"
	"reptiles/crawler/engine"
	xcar "reptiles/crawler/xcar/parser"
	zhenai "reptiles/crawler/zhenai/parser"
	pb "reptiles/crawler_distributed/proto"
)


func SerializeRequest(r engine.Request) *pb.ProcessRequest {
	name, args := r.Parser.Serialize()
	return &pb.ProcessRequest{
		Url: r.Url,
		SerializedParser: &pb.SerializedParser{
			Name: name,
			Args: args,
		},
	}
}

func SerializeResult(
	r engine.ParseResult) pb.ProcessResult {
	result := pb.ProcessResult{
		Item: r.Items,
	}

	for _, req := range r.Requests {
		result.Request = append(result.Request,
			SerializeRequest(req))
	}
	return result
}

func DeserializeRequest(
	r *pb.ProcessRequest) (engine.Request, error) {
	parser, err := deserializeParser(r.SerializedParser)
	if err != nil {
		return engine.Request{}, err
	}
	return engine.Request{
		Url:    r.Url,
		Parser: parser,
	}, nil
}

func DeserializeResult(
	r pb.ProcessResult) pb.ProcessResult {
	result := pb.ProcessResult{
		Item: r.Item,
	}

	for _, req := range r.Request {
		engineReq, err := DeserializeRequest(req)
		if err != nil {
			log.Printf("error deserializing "+
				"request: %v", err)
			continue
		}
		result.Request = append(result.Request,
			engineReq)
	}
	return result
}

func deserializeParser(
	p *pb.SerializedParser) (engine.Parser, error) {
	switch p.Name {
	case config.ParseCityList:
		return engine.NewFuncParser(
			zhenai.ParseCityList,
			config.ParseCityList), nil
	case config.ParseCity:
		return engine.NewFuncParser(
			zhenai.ParseCity,
			config.ParseCity), nil

	case config.ParseProfile:
		if userName, ok := p.Args.(string); ok {
			return zhenai.NewProfileParser(
				userName), nil
		} else {
			return nil, fmt.Errorf("invalid "+
				"arg: %v", p.Args)
		}
	case config.ParseCarDetail:
		return engine.NewFuncParser(
			xcar.ParseCarDetail,
			config.ParseCarDetail), nil
	case config.ParseCarModel:
		return engine.NewFuncParser(
			xcar.ParseCarModel,
			config.ParseCarModel), nil
	case config.ParseCarList:
		return engine.NewFuncParser(
			xcar.ParseCarList,
			config.ParseCarList), nil
	case config.NilParser:
		return engine.NilParser{}, nil
	default:
		return nil, errors.New(
			"unknown parser name")
	}
}
