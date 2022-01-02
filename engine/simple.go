package engine

import "fmt"

type SimpleEngine struct{}

func (e SimpleEngine) Run(seeds ...Request) {
	Q := []Request{}
	for _, seed := range seeds {
		Q = append(Q, seed)
	}
	for len(Q) > 0 {
		//time.Sleep(time.Second * 6)
		r := Q[0]
		Q = Q[1:]
		parseResult, err := work(r)
		if err != nil {
			fmt.Printf("url: %s , error:%s\n", r.Url, err)
			continue
		}

		for _, v := range parseResult.Request {
			fmt.Printf("add request: %s\n", v.Url)
			Q = append(Q, Request{
				Url:        v.Url,
				ParserFunc: v.ParserFunc,
			})
		}
	}
}
