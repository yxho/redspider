package worker

import (
	"log"
	"redspider/engine"
)

type SerializeParser struct{
	Name string
	Args interface{}
}

type Request struct{
	Url string
	Parse SerializeParser
}

type ParseResult struct{
	items []engine.Item
	Requests []Request
}

func SerializeResult(r engine.ParseResult)ParseResult{
	result := ParseResult{items: r.Items}

	for _,req := range r.Requests{
		result.Requests=append(result.Requests,SerializeRequest(req))
	}
	return result
}

func SerializeRequest(r engine.Request) Request {
	name,args := r.Parse.Serialize()

	return Request{
		Url:r.Url,
		Parse: SerializeParser{
			Name: name,
			Args: args,
		},
	}
}

func DeserializeResult(r ParseResult)engine.ParseResult{
	result:=engine.ParseResult{
		Items: r.items,
	}
	for _,req:=range r.Requests{
		engineReq,err:=DeserializeRequest(req)
		if err!=nil{
			log.Printf("error deserialize:%v",err)
			continue
		}
		result.Requests=append(result.Requests,engineReq)
	}
	return result
}

func DeserializeRequest(r Request) (engine.Request,error) {
	parse,err:=deserializeParse(r.Parse)
	if err!=nil{
		return engine.Request{},err
	}
	return engine.Request{
		Url:r.Url,
		Parse: parse,
	},nil

}

func deserializeParse(parse SerializeParser) (engine.Parser,error) {
	switch parse.Name{
	case "booklist"
	}
}


