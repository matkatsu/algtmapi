package search

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("Search", func() {
	BasePath("/v1")
	Action("GetSearch", func() {
		Routing(GET("/search"))
		Description("キーワード検索の結果を返します")
		Params(func() {
			Param("word", String, func() {
				Description("キーワード")
				Default("")
				Example("霧矢あおい")
			})
		})
		Response(OK, func() {
			Description("正常に取得できた場合に返却")
			Media(SearchMT)
		})
	})
})

// SearchMT 検索結果MediaType
var SearchMT = MediaType("application/vnd.romiogaku.com.algtmapi.search+json", func() {
	ContentType("application/vnd.romiogaku.com.algtmapi.search+json; charset=utf8")
	Reference(SearchResultType)
	Attributes(func() {
		Attribute("result")
		Required("result")
	})
	View("default", func() {
		Attribute("result")
	})
})

// SearchResultType 検索結果Type
var SearchResultType = Type("SearchResultType", func() {
	Member("result", ArrayOf(searchItemType))
	Required("result")
})

var searchItemType = Type("SearchItemType", func() {
	Member("id", Integer, func() {
		Description("ID")
	})
	Member("url", String, func() {
		Description("画像URL")
	})
	Member("word", String, func() {
		Description("ワード")
	})
	Required("id", "url", "word")
})
