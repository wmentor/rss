package rss

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRSS(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `<?xml version="1.0" encoding="utf-8"?>
<rss version="2.0">
	<channel>
	<title>WMentor.ru</title>
	<link>https://wmentor.ru</link>
	<description>WMentor.ru</description>
	<image>
	<url>https://wmentor.ru/images/rss.jpg?v=1577042225</url>
	<link>https://wmentor.ru</link>
	<title>WMentor.ru</title>
	</image>
	<lastBuildDate>12 Feb 2020 01:06:11 +0300</lastBuildDate>
<item><title>Embed</title>
<author>wmentor</author>
<link>https://wmentor.ru/embed</link>
<description>&lt;p&gt;Golang has&#39;t tools to embed external files to a application. But it&#39;s useful for my work under lemma project. So, I have created module for embedding static files…&lt;/p&gt;&lt;p&gt;&lt;a href="https://wmentor.ru/embed"&gt;Read more&lt;/a&gt;&lt;/p&gt;</description><pubDate>12 Feb 2020 01:06:11 +0300</pubDate>
<guid>https://wmentor.ru/embed</guid>
</item>
<item><title>New lemma</title>
<author>wmentor</author>
<link>https://wmentor.ru/lemma-v1.0.2</link>
<description>&lt;p&gt;Today is a good day. Lemma v1.0.2 is ready after month of development. That was not easy. There were problems with previous data structures (primarily dictionary) and concepts. But everything…&lt;/p&gt;&lt;p&gt;&lt;a href="https://wmentor.ru/lemma-v1.0.2"&gt;Read more&lt;/a&gt;&lt;/p&gt;</description><pubDate>30 Jan 2020 01:08:20 +0300</pubDate>
<guid>https://wmentor.ru/lemma-v1.0.2</guid>
</item>
<item><title>Lemma v1.0.1</title>
<author>wmentor</author>
<link>https://wmentor.ru/lemma-v1.0.1</link>
<description>&lt;p&gt;Lemma v1.0.1 was released today. Release notes…&lt;/p&gt;&lt;p&gt;&lt;a href="https://wmentor.ru/lemma-v1.0.1"&gt;Read more&lt;/a&gt;&lt;/p&gt;</description><pubDate>05 Jan 2020 02:10:11 +0300</pubDate>
<guid>https://wmentor.ru/lemma-v1.0.1</guid>
</item>
<item><title>Lemmatization and Hunspell</title>
<author>wmentor</author>
<link>https://wmentor.ru/hunspell</link>
<description>&lt;p&gt;Lemmatization, on the other hand, takes into consideration the morphological analysis of the words. To do so, it is necessary to have detailed dictionaries which…&lt;/p&gt;&lt;p&gt;&lt;a href="https://wmentor.ru/hunspell"&gt;Read more&lt;/a&gt;&lt;/p&gt;</description><pubDate>02 Jan 2020 01:53:59 +0300</pubDate>
<guid>https://wmentor.ru/hunspell</guid>
</item>
<item><title>Stemming and Lemmatization</title>
<author>wmentor</author>
<link>https://wmentor.ru/stemming-and-lemmatization</link>
<description>&lt;p&gt;I built a text tokenizer some time ago. Now results of tokenizer work must to be normalized before further processing…&lt;/p&gt;&lt;p&gt;&lt;a href="https://wmentor.ru/stemming-and-lemmatization"&gt;Read more&lt;/a&gt;&lt;/p&gt;</description><pubDate>28 Dec 2019 03:41:00 +0300</pubDate>
<guid>https://wmentor.ru/stemming-and-lemmatization</guid>
</item>
<item><title>First tokenizer</title>
<author>wmentor</author>
<link>https://wmentor.ru/first-tokenizer</link>
<description>&lt;p&gt;Today is a good day. First tokenizer version is ready. It allows return words, hashtags, signs, numbers, domains, ip addresses, complex words, urls, uuids and…&lt;/p&gt;&lt;p&gt;&lt;a href="https://wmentor.ru/first-tokenizer"&gt;Read more&lt;/a&gt;&lt;/p&gt;</description><pubDate>26 Dec 2019 02:48:18 +0300</pubDate>
<guid>https://wmentor.ru/first-tokenizer</guid>
</item>
<item><title>Start</title>
<author>wmentor</author>
<link>https://wmentor.ru/startup</link>
<description>&lt;p&gt;Hello, World! Today I decided to start work on the natural language processing system (NLP). The system will be developed on Pure Golang under the MIT license. Now I have created…&lt;/p&gt;&lt;p&gt;&lt;a href="https://wmentor.ru/startup"&gt;Read more&lt;/a&gt;&lt;/p&gt;</description><pubDate>23 Dec 2019 01:11:44 +0300</pubDate>
<guid>https://wmentor.ru/startup</guid>
</item>
</channel></rss>`)
	}))

	defer ts.Close()

	list, err := Get(ts.URL)

	if err != nil {
		t.Fatal("empty response")
	}

	if len(list) != 7 {
		t.Fatal("invalid list size")
	}
}
