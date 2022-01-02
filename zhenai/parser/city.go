package parser

import (
	"regexp"
	"spider/engine"
	"spider/model"
	"strconv"
	"strings"
)

var cityRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)" [^>]*>([^<]+)</a>`)

//var nextPageRe = regexp.MustCompile(`<a href="(http://www.zhenai.com/[^"]*)">下一页</a>`)
var nextPageRe = regexp.MustCompile(`<a href="(http://www.zhenai.com/[^"]*)">.+</a>`)

//const infoRe =`<div class="content">[^<]*<table>[^<]*<tbody>[^<]*<tr>[^<]*<th>[^<]*<a href="http://album.zhenai.com/u/[0-9]+" target="_blank">([^<]*)</a>[^<]*</th>[^<]*</tr>[^<]*<tr>[^<]*<td width="180">[^<]*<span class="grayL">性别：</span>([^<]*)</td>[^<]*<td>[^<]*<span class="grayL">居住地：</span>([^<]*)</td>[^<]*</tr>[^<]*<tr>[^<]*<td width="180">[^<]*<span class="grayL">年龄：</span>([^<]*)</td>[^<]*<td>[^<]*<span class="grayL">学&nbsp;&nbsp;&nbsp;历：</span>([^<]*)</td>[^<]*<!---->[^<]*</tr>[^<]*<tr>[^<]*<td width="180">[^<]*<span class="grayL">婚况：</span>[^<]*</td>[^<]*<td width="180">[^<]*<span class="grayL">身&nbsp;&nbsp;&nbsp;高：</span>([^<]*)</td>[^<]*</tr>[^<]*</tbody>[^<]*</table>[^<]*<div class="introduce">([^<]*)</div>[^<]*</div>`
const infoRe1 = `<div class="list-item">[^<]*<div class="photo">[^<]*<a href="(http://album.zhenai.com/u/[0-9]+)" target="_blank">`
const infoRe2 = `[^<]*<img src="(https://photo.zastatic.com/images/photo/[^"]*)"[^<]*</a>[^<]*</div>[^<]*<div class="content">`
const infoRe3 = `[^<]*<table>[^<]*<tbody>[^<]*<tr>[^<]*<th>[^<]*<a href="http://album.zhenai.com/u/[0-9]+" target="_blank">([^<]*)</a>`
const infoRe4 = `[^<]*</th>[^<]*</tr>[^<]*<tr>[^<]*<td width="180">[^<]*<span class="grayL">性别：[^<]*</span>([^<]*)</td>[^<]*<td>`
const infoRe5 = `[^<]*<span class="grayL">居住地：[^<]*</span>([^<]*)</td>[^<]*</tr>[^<]*<tr>[^<]*<td width="180">`
const infoRe6 = `[^<]*<span class="grayL">年龄：[^<]*</span>([^<]*)</td>[^<]*<td>[^<]*<span class="grayL">学[^历]*历：[^<]*</span>([^<]*)</td>`
const infoRe7 = `[^</tr]*</tr>[^<]*<tr>[^<]*<td width="180">[^<]*<span class="grayL">婚况：[^<]*</span>([^<]*)</td>`
const infoRe8 = `[^<]*<td width="180">[^<]*<span class="grayL">身[^高]*高：[^<]*</span>([^<]*)</td>[^<]*</tr>[^<]*</tbody>[^<]*</table>`
const infoRe9 = `[^<]*<div class="introduce">([^</div>]*)</div>` //[^<]*</div>[^<]*<div class="item-btn">打招呼[^<]*</div>[^<]*</div>

var infoRe = regexp.MustCompile(infoRe1 + infoRe2 + infoRe3 + infoRe4 + infoRe5 + infoRe6 + infoRe7 + infoRe8 + infoRe9)

func ParseCity(contents []byte) engine.ParseResult {
	result := engine.ParseResult{}
	/*// 解析详情页URL
	submatch := infoRe.FindAllSubmatch(contents, -1)
	for _, v := range submatch {
		result.Request = append(result.Request, engine.Request{
			Url:        string(v[1]),
			ParserFunc: engine.NilParser,
		})
	}*/

	submatch := infoRe.FindAllSubmatch(contents, -1)
	for _, v := range submatch {
		// TODO 保存信息
		//v[1]主页信息, v[2]头像, v[3]昵称, v[4]性别, v[5]地区, v[6]年龄, v[7]学历, v[8]婚况, v[9]身高, v[10]个人介绍
		name := strings.ReplaceAll(string(v[3]), "\n", "")
		age, _ := strconv.Atoi(string(v[6]))
		height, _ := strconv.Atoi(string(v[9]))
		result.Item = append(result.Item, model.Profile{
			Name:      name,
			Gender:    string(v[4]),
			Age:       age,
			Height:    height,
			Income:    string(v[5]),
			Marriage:  string(v[8]),
			Education: string(v[7]),
		})
	}

	matchs := nextPageRe.FindAllSubmatch(contents, 1)
	for _, v := range matchs {
		result.Request = append(result.Request, engine.Request{
			Url:        string(v[1]),
			ParserFunc: ParseCity,
		})
	}
	return result
}
