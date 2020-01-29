package engine

import (
	"reptiles/crawler/config"
	pb "reptiles/crawler_distributed/proto"
)

type ParserFunc func(
	contents []byte, url string) ParseResult

type Parser interface {
	Parse(contents []byte, url string) ParseResult
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
	_ []byte, _ string) ParseResult {
	return ParseResult{}
}

func (NilParser) Serialize() (
	name string, args string) {
	return config.NilParser, ""
}

type FuncParser struct {
	parser ParserFunc
	name   string
}

func (f *FuncParser) Parse(
	contents []byte, url string) ParseResult {
	return f.parser(contents, url)
}

func (f *FuncParser) Serialize() (
	name string, args string) {
	return f.name, ""
}

func NewFuncParser(
	p ParserFunc, name string) *FuncParser {
	return &FuncParser{
		parser: p,
		name:   name,
	}
}
