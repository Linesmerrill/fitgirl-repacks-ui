package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/gocolly/colly"
	"github.com/linesmerrill/fitgirl-repacks-ui/db"
	"golang.org/x/exp/slices"
)

type scraper struct {
	mdb db.Db
}

func (s scraper) execute() {
	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only urls that contain page numbers
		colly.URLFilters(
			regexp.MustCompile("https://fitgirl-repacks\\.site/all-my-repacks-a-z/\\?lcp_page0=\\d+"),
		),
	)

	// On every a element which has href attribute call callback
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if skipTitleText(e.Text) {
			// skip
			// fmt.Printf("Skipping: : %q -> %s\n", e.Text, link)
		} else {

			if skipLink(link) {
				// if we have a page # then don't write to the db
			} else {
				fmt.Printf("Link found, writing to db: %q -> %s\n", e.Text, link)

				document := map[string]interface{}{
					"Title": e.Text,
					"Link":  link,
				}
				// Write the title and link to the collection.
				s.mdb.Collection.Insert(document)
			}
			// Close the session.
			// defer s.mdb.Session.Close()
			// Visit link found on page
			// Only those links are visited which are matched by  any of the URLFilter regexps
			c.Visit(e.Request.AbsoluteURL(link))
		}
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on http://httpbin.org
	c.Visit("https://fitgirl-repacks.site/all-my-repacks-a-z/?lcp_page0=1")
}

func skipLink(link string) bool {
	skippableLinks := []string{
		"lcp_page0",
		"https://fitgirl-repacks.site/2016/01/",
		"https://fitgirl-repacks.site/2016/02/",
		"https://fitgirl-repacks.site/2016/03/",
		"https://fitgirl-repacks.site/2016/04/",
		"https://fitgirl-repacks.site/2016/05/",
		"https://fitgirl-repacks.site/2016/06/",
		"https://fitgirl-repacks.site/2016/07/",
		"https://fitgirl-repacks.site/2016/08/",
		"https://fitgirl-repacks.site/2016/09/",
		"https://fitgirl-repacks.site/2016/10/",
		"https://fitgirl-repacks.site/2016/11/",
		"https://fitgirl-repacks.site/2016/12/",
		"https://fitgirl-repacks.site/2017/01/",
		"https://fitgirl-repacks.site/2017/02/",
		"https://fitgirl-repacks.site/2017/03/",
		"https://fitgirl-repacks.site/2017/04/",
		"https://fitgirl-repacks.site/2017/05/",
		"https://fitgirl-repacks.site/2017/06/",
		"https://fitgirl-repacks.site/2017/07/",
		"https://fitgirl-repacks.site/2017/08/",
		"https://fitgirl-repacks.site/2017/09/",
		"https://fitgirl-repacks.site/2017/10/",
		"https://fitgirl-repacks.site/2017/11/",
		"https://fitgirl-repacks.site/2017/12/",
		"https://fitgirl-repacks.site/2018/01/",
		"https://fitgirl-repacks.site/2018/02/",
		"https://fitgirl-repacks.site/2018/03/",
		"https://fitgirl-repacks.site/2018/04/",
		"https://fitgirl-repacks.site/2018/05/",
		"https://fitgirl-repacks.site/2018/06/",
		"https://fitgirl-repacks.site/2018/07/",
		"https://fitgirl-repacks.site/2018/08/",
		"https://fitgirl-repacks.site/2018/09/",
		"https://fitgirl-repacks.site/2018/10/",
		"https://fitgirl-repacks.site/2018/11/",
		"https://fitgirl-repacks.site/2018/12/",
		"https://fitgirl-repacks.site/2019/01/",
		"https://fitgirl-repacks.site/2019/02/",
		"https://fitgirl-repacks.site/2019/03/",
		"https://fitgirl-repacks.site/2019/04/",
		"https://fitgirl-repacks.site/2019/05/",
		"https://fitgirl-repacks.site/2019/06/",
		"https://fitgirl-repacks.site/2019/07/",
		"https://fitgirl-repacks.site/2019/08/",
		"https://fitgirl-repacks.site/2019/09/",
		"https://fitgirl-repacks.site/2019/10/",
		"https://fitgirl-repacks.site/2019/11/",
		"https://fitgirl-repacks.site/2019/12/",
		"https://fitgirl-repacks.site/2020/01/",
		"https://fitgirl-repacks.site/2020/02/",
		"https://fitgirl-repacks.site/2020/03/",
		"https://fitgirl-repacks.site/2020/04/",
		"https://fitgirl-repacks.site/2020/05/",
		"https://fitgirl-repacks.site/2020/06/",
		"https://fitgirl-repacks.site/2020/07/",
		"https://fitgirl-repacks.site/2020/08/",
		"https://fitgirl-repacks.site/2020/09/",
		"https://fitgirl-repacks.site/2020/10/",
		"https://fitgirl-repacks.site/2020/11/",
		"https://fitgirl-repacks.site/2020/12/",
		"https://fitgirl-repacks.site/2021/01/",
		"https://fitgirl-repacks.site/2021/02/",
		"https://fitgirl-repacks.site/2021/03/",
		"https://fitgirl-repacks.site/2021/04/",
		"https://fitgirl-repacks.site/2021/05/",
		"https://fitgirl-repacks.site/2021/06/",
		"https://fitgirl-repacks.site/2021/07/",
		"https://fitgirl-repacks.site/2021/08/",
		"https://fitgirl-repacks.site/2021/09/",
		"https://fitgirl-repacks.site/2021/10/",
		"https://fitgirl-repacks.site/2021/11/",
		"https://fitgirl-repacks.site/2021/12/",
		"https://fitgirl-repacks.site/2022/01/",
		"https://fitgirl-repacks.site/2022/02/",
		"https://fitgirl-repacks.site/2022/03/",
		"https://fitgirl-repacks.site/2022/04/",
		"https://fitgirl-repacks.site/2022/05/",
		"https://fitgirl-repacks.site/2022/06/",
		"https://fitgirl-repacks.site/2022/07/",
		"https://fitgirl-repacks.site/2022/08/",
		"https://fitgirl-repacks.site/2022/09/",
		"https://fitgirl-repacks.site/2022/10/",
		"https://fitgirl-repacks.site/2022/11/",
		"https://fitgirl-repacks.site/2022/12/",
		"https://fitgirl-repacks.site/2023/01/",
		"https://fitgirl-repacks.site/2023/02/",
		"https://fitgirl-repacks.site/2023/03/",
		"https://fitgirl-repacks.site/2023/04/",
		"https://fitgirl-repacks.site/2023/05/",
		"https://fitgirl-repacks.site/2023/06/",
		"https://fitgirl-repacks.site/2023/07/",
		"https://fitgirl-repacks.site/2023/08/",
		"https://fitgirl-repacks.site/2023/09/",
		"https://fitgirl-repacks.site/2023/10/",
		"https://fitgirl-repacks.site/2023/11/",
		"https://fitgirl-repacks.site/2023/12/",
	}
	found, count := checkSubstrings(link, skippableLinks)
	if found || count > 0 {
		return true
	} else {
		return false
	}
}

func skipTitleText(titleText string) bool {
	skippableTitles := []string{
		"FitGirl Repacks",
		"\n\t\t\t\t\tSearch\t\t\t\t",
		"\n\t\t\t\t\tSkip to content\t\t\t\t",
		"All My Repacks, A-Z",
		"Games with my personal Pink Paw Award",
		"Switch Emulated Repacks",
		"PS3 Emulated Repacks",
		"FAQ",
		"Donate",
		"Contacts",
		"Repacks Troubleshooting",
		"Previous Page",
		"\n\t\t\t\t\tProudly powered by WordPress\t\t\t\t",
		"Donate to FitGirl",
		"this RSS feed",
	}
	return slices.Contains(skippableTitles, titleText)
}

func checkSubstrings(str string, subs []string) (bool, int) {

	matches := 0
	isCompleteMatch := true

	for _, sub := range subs {
		if strings.Contains(str, sub) {
			matches += 1
		} else {
			isCompleteMatch = false
		}
	}

	return isCompleteMatch, matches
}
