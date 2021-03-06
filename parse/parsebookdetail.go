package parse

import (
	"redspider/engine"
	"redspider/model"
	"regexp"
	"strconv"
)

var autoRe = regexp.MustCompile(`<span class="pl"> 作者</span>:[\d\D]*?<a.*?>([^<]+)</a>`)
var public = regexp.MustCompile(`<span class="pl">出版社:</span>([^<]+)<br/>`)
var pageRe = regexp.MustCompile(`<span class="pl">页数:</span> ([^<]+)<br/>`)
var priceRe = regexp.MustCompile(`<span class="pl">定价:</span>([^<]+)<br/>`)
var scoreRe = regexp.MustCompile(`<strong class="ll rating_num " property="v:average">([^<]+)</strong>`)
var introRe = regexp.MustCompile(`<div class="intro">[\d\D]*?<p>([^<]+)</p></div>`)

func ParseBookDetail(url string, contents []byte, bookname string) engine.ParseResult {
	bookdetail := model.BookDetail{}
	bookdetail.Author = ExtraString(contents, autoRe)
	page, err := strconv.Atoi(ExtraString(contents, pageRe))
	if err == nil {
		bookdetail.Bookpages = page
	}
	bookdetail.Bookname = bookname
	bookdetail.Publicer = ExtraString(contents, public)
	bookdetail.Intro = ExtraString(contents, introRe)
	bookdetail.Score = ExtraString(contents, scoreRe)
	bookdetail.Price = ExtraString(contents, priceRe)
	result := engine.ParseResult{
		Items: []engine.Item{
			//bookdetail
			{
				Url:     url,
				Type:    "douban",
				Payload: bookdetail,
			},
		},
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

type BookDetailParse struct {
	bookName string
}

func (b *BookDetailParse) Parser(contents []byte, url string) engine.ParseResult {
	return ParseBookDetail(url, contents, b.bookName)
}

func (b *BookDetailParse) Serialize() (name string, args interface{}) {
	return "BookDetailParse", b.bookName
}

func NewBookDetailParse(name string) *BookDetailParse {
	return &BookDetailParse{
		bookName: name,
	}
}
