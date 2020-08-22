package parse

import (
	"redspider/engine"
	"redspider/model"
	"regexp"
	"strconv"
)

//const autoRe = `<span class="pl"> 作者</span>:[\d\D]*?<a.*?>([^<]+)</a>`
//const public = `<span class="pl">出版社:</span>([^<]+)<br/>`
//const pageRe = `<span class="pl">页数:</span>([^<]+)<br/>`
//const priceRe = `<span class="pl">定价:</span>([^<]+)<br/>`
//const scoreRe = `<strong class="ll rating_num " property="v:average">([^<]+)</strong>`
//const introRe = `<div class="intro">[\d\D]*?<p>([^<]+)</p></div>`

var autoRe = regexp.MustCompile(`<span class="pl"> 作者</span>:[\d\D]*?<a.*?>([^<]+)</a>`)
var public = regexp.MustCompile(`<span class="pl">出版社:</span>([^<]+)<br/>`)
var pageRe = regexp.MustCompile(`<span class="pl">页数:</span> ([^<]+)<br/>`)
var priceRe = regexp.MustCompile(`<span class="pl">定价:</span>([^<]+)<br/>`)
var scoreRe = regexp.MustCompile(`<strong class="ll rating_num " property="v:average">([^<]+)</strong>`)
var introRe = regexp.MustCompile(`<div class="intro">[\d\D]*?<p>([^<]+)</p></div>`)

func ParseBookDetail(contents []byte) engine.ParseResult{
	bookdetail := model.BookDetail{}
	bookdetail.Author = ExtraString(contents, autoRe)
	page, err := strconv.Atoi(ExtraString(contents, pageRe))
	if err == nil {
		bookdetail.Bookpages = page
	}
	bookdetail.Publicer = ExtraString(contents, public)
	bookdetail.Intro = ExtraString(contents, introRe)
	bookdetail.Score = ExtraString(contents, scoreRe)
	bookdetail.Price = ExtraString(contents,priceRe)
	result := engine.ParseResult{
		Items: []interface{}{bookdetail},
	}
	return result

}

func ExtraString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)

	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}
