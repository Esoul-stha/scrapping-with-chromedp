package main

import (
	"context"
	"fmt"
	"log"
	"regexp"
	"strconv"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/chromedp"
)

func main() {
	// create chrome instance
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	// ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	// defer cancel()

	var nodes []*cdp.Node

	err := chromedp.Run(ctx,
		emulation.SetUserAgentOverride("WebScraper 1.0"),
		chromedp.Navigate("#URL"),
		// wait for #class element is visible (ie, page is loaded)
		chromedp.ScrollIntoView(`div.#class`),
		chromedp.Nodes("div.#class > ul > li > h4 > a", &nodes, chromedp.ByQueryAll))

	if err != nil {
		log.Fatal(err)
	}

	for _, n := range nodes {
		// extracting all name text from a (anchor tag)
		name := n.Children[0].NodeValue
		url := n.AttributeValue("href")
		// extracting only digits from href
		re := regexp.MustCompile("[0-9]+")
		specificNum := re.FindAllString(url, -1)

		nums := make([]int, len(specificNum))

		for i, numStr := range specificNum {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				log.Fatal(err)
			}
			nums[i] = num
		}
		fmt.Println(name, nums)
	}

}
