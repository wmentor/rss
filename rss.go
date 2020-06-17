package rss

import (
	"bytes"
	"encoding/xml"
	"errors"

	"github.com/wmentor/ua"
	"golang.org/x/net/html/charset"
)

var (
	ErrRequestFailed error = errors.New("http request failed")
)

type Rss struct {
	Channel RssChannel `xml:"channel"`
}

type RssItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
}

type RssChannel struct {
	Title       string     `xml:"title"`
	Link        string     `xml:"link"`
	Description string     `xml:"description"`
	Items       []*RssItem `xml:"item"`
}

func Get(url string) ([]*RssItem, error) {

	client := ua.New()

	resp, err := client.Request("GET", url, nil, nil)

	if err != nil || resp == nil || len(resp.Content) == 0 {
		return nil, ErrRequestFailed
	}

	buffer := bytes.NewBuffer(resp.Content)
	xml := xml.NewDecoder(buffer)

	xml.CharsetReader = charset.NewReaderLabel

	rss := new(Rss)

	if err = xml.Decode(rss); err != nil {
		return nil, err
	}

	return rss.Channel.Items, nil
}
