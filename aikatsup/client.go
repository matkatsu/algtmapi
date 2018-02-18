package aikatsup

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path"
	"runtime"

	"github.com/pkg/errors"
)

// Client struct
type Client struct {
	URL        *url.URL
	HTTPClient *http.Client
	Logger     *log.Logger
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

// NewClient constructor
func NewClient(urlStr string, logger *log.Logger) (*Client, error) {
	parsedURL, err := url.ParseRequestURI(urlStr)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse url: %s", urlStr)
	}
	var discardLogger = log.New(ioutil.Discard, "", log.LstdFlags)
	if logger == nil {
		logger = discardLogger
	}

	client := &Client{
		URL:        parsedURL,
		HTTPClient: &http.Client{},
		Logger:     logger,
	}

	return client, err
}

var userAgent = fmt.Sprintf("ALGTMGoClient/%s (%s)", "v1", runtime.Version())

func (c *Client) newRequest(method, spath string, query url.Values, body io.Reader) (*http.Request, error) {
	u := *c.URL
	u.Path = path.Join(c.URL.Path, spath)
	req, err := http.NewRequest(method, u.String(), body)
	req.URL.RawQuery = query.Encode()

	if err != nil {
		return nil, err
	}
	// req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", userAgent)

	return req, nil
}

func decodeBody(resp *http.Response, out interface{}) error {
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	return decoder.Decode(out)
}

// GetSearchResult 検索結果を取得
func (c *Client) GetSearchResult(word string) (*SearchResult, error) {
	values := url.Values{}
	values.Add("words", word)

	req, err := c.newRequest("GET", "search", values, nil)
	if err != nil {
		return nil, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	// Check status code here…

	var search SearchResult
	if err := decodeBody(res, &search); err != nil {
		return nil, err
	}

	return &search, nil
}
