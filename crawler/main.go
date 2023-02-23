package main

import (
	"xyy/learngo/crawler/engine"
	"xyy/learngo/crawler/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:        `https://www.zhenai.com/zhenghun`,
		ParserFunc: parser.ParseCityList,
	})
}
