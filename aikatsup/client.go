package aikatsup

import (
	"errors"
	"math/rand"

	"github.com/parnurzeal/gorequest"
)

// Client struct
type Client struct {
	BaseURL string
}

// SearchResult 検索結果
type SearchResult struct {
	Result []ResultUnit `json:"item"`
}

// ResultUnit 検索結果Unit
type ResultUnit struct {
	ID    int    `json:"id"`
	Word  string `json:"words"`
	Image Image  `json:"image"`
}

// Image 画像
type Image struct {
	Width  int    `json:"width"`
	Height int    `json:"height"`
	URL    string `json:"url"`
}

// GetSearchResult 検索結果を取得
func (c *Client) GetSearchResult(word string) ([]ResultUnit, error) {
	// wordが空ならi、infoからランダムに9件取得(chanel)
	var r SearchResult
	request := gorequest.New()
	resp, _, errs := request.
		Get(c.getRequestURL("v1/search")).
		Query("words=" + word).
		EndStruct(&r)
	if resp.StatusCode != 200 {
		return r.Result, errors.New("failed")
	}
	for _, err := range errs {
		if err != nil {
			return r.Result, err
		}
	}

	// 結果が9個より少なければそのまま返す
	if len(r.Result) < 9 {
		return r.Result, nil
	}

	var res []ResultUnit
	// 9個以上ならシャッフル
	shuffle(r.Result)
	res = r.Result[0:9]
	return res, nil
}

func (c *Client) getRequestURL(path string) string {
	return c.BaseURL + path
}

// シャッフルする
func shuffle(data []ResultUnit) {
	n := len(data)
	for i := n - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		data[i], data[j] = data[j], data[i]
	}
}
