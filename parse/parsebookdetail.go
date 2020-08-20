package parse

import "redspider/model"

const autoRe = `<span class="pl"> 作者</span>:[\d\D]*?<a.*?>([^<]+)</a>`
const public = `<span class="pl">出版社:</span>([^<]+)<br/>`
const pageRe = `<span class="pl">页数:</span>([^<]+)<br/>`
const priceRe = `<span class="pl">定价:</span>([^<]+)<br/>`
const scoreRe =`<strong class="ll rating_num " property="v:average">([^<]+)</strong>`
const introRe = `<div class="intro">[\d\D]*?<p>([^<]+)</p></div>`

func ParseBookDetail(contents []byte)  {
	bookdetail:=model.BookDetail{}
	//match:=autoRe.F

}
