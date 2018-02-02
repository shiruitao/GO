/*
 * MIT License
 *
 * Copyright (c) 2018 SmartestEE Co., Ltd..
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

/*
 * Revision History:
 *     Initial: 2018/01/30        Shi Ruitao
 */

package segmentf

import (
	"fmt"
	"os"
	"time"

	"github.com/asciimoo/colly"
	"github.com/russross/blackfriday"

	"github.com/fengyfei/gu/libs/crawler"
	"github.com/fengyfei/gu/libs/logger"
)

type over struct{}

type segmentCrawler struct {
	collectorUrl  *colly.Collector
	collectorBlog *colly.Collector
	chUrl         chan string
	overClawler   chan over
}

// NewGoCNCrawler generates a crawler for segment blog.
func NewSegmentCrawler() crawler.Crawler {
	return &segmentCrawler{
		collectorUrl:  colly.NewCollector(),
		collectorBlog: colly.NewCollector(),
		chUrl:         make(chan string),
	}
}

// Crawler interface Init
func (c *segmentCrawler) Init() error {
	c.collectorUrl.OnHTML("article.ArticleInList.fade-in.ArticleInList--cardWhenBig", c.parse)
	c.collectorBlog.OnHTML("article.Article.fade-in.Article--featured", c.parseBlog)

	return nil
}

// Crawler interface Start
func (c *segmentCrawler) Start() error {
	go c.startUrl()
	for {
		select {
		case url := <-c.chUrl:
			err := c.startBlog(url)
			if err != nil {
				return err
			}
		case <-time.NewTimer(3 * time.Second).C:
			goto EXIT
		}
	}

EXIT:
	return nil
}

func (c *segmentCrawler) parse(e *colly.HTMLElement) {
	url, ok := e.DOM.Find("h2").Find("a").Attr("href")
	url = "https://segment.com" + url
	fmt.Println(url, ok)
	c.chUrl <- url
}

func (c *segmentCrawler) parseBlog(e *colly.HTMLElement) {
	title := e.DOM.Find("h2").Find("a").Text()
	html, _ := e.DOM.Html()
	by := []byte(html)
	output := blackfriday.MarkdownBasic(by)
	fmt.Println(string(output))

	f, _ := os.OpenFile("./blog/"+title+".md", os.O_CREATE|os.O_RDWR, 0644)
	f.Write([]byte(output))
}

func (c *segmentCrawler) startUrl() {
	for i := 1; ; i++ {
		if i == 1 {
			c.collectorUrl.Visit("https://segment.com/blog/")
		} else {
			var url []string
			url = append(url, fmt.Sprint("https://segment.com/blog/page/", i))
			err := c.collectorUrl.Visit(url[0])
			if err != nil {
				logger.Error("Visit Error:", err)
				break
			}
		}
	}
}

func (c *segmentCrawler) startBlog(url string) error {
	err := c.collectorBlog.Visit(url)
	if err != nil {
		logger.Error("Visit Error:", err)
		return err
	}
	return nil
}
