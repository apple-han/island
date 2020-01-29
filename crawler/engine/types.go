package engine

import (
	"reptiles/crawler/config"
	pb "reptiles/crawler_distributed/proto"
)

type ParserFunc func(
	contents []byte, url string) pb.ProcessResult

type Parser interface {
	Parse(contents []byte, url string) pb.ProcessResult
	Serialize() (name string, args string)
}

type Request struct {
	Url    string
	Parser Parser
}

type ParseResult struct {
	Requests []Request
	Items    []*pb.Item
}

type NilParser struct{}

func (NilParser) Parse(
	_ []byte, _ string) pb.ProcessResult {
	return pb.ProcessResult{}
}

func (NilParser) Serialize() (
	name string, args interface{}) {
	return config.NilParser, nil
}

type FuncParser struct {
	parser ParserFunc
	name   string
}

func (f *FuncParser) Parse(
	contents []byte, url string) pb.ProcessResult {
	return f.parser(contents, url)
}

func (f *FuncParser) Serialize() (
	name string, args interface{}) {
	return f.name, nil
}

func NewFuncParser(
	p ParserFunc, name string) *FuncParser {
	return &FuncParser{
		parser: p,
		name:   name,
	}
}
