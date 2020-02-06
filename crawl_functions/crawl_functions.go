package crawl_functions

import (
	"fmt"
	"sync"
	"webmotor_crawler/query_handler"
)

func CrawlRoutine(wg *sync.WaitGroup, c chan int) {

	head := query_handler.QueryClient{
		Accept:                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8",
		Accept_encoding:           "gzip, deflate, br",
		Accept_language:           "en-US,en;q=0.5",
		Connection:                "keep-alive",
		Host:                      "www.webmotors.com.br",
		TE:                        "Trailers",
		Upgrade_insecure_requests: "1",
		User_agent:                "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:72.0) Gecko/20100101 Firefox/72.0",
		ProxyUrl:                  "socks5://127.0.0.1:9150",
	}

	client := head.GenerateNoProxiedClient()
	//wg.Add(1)

	for len(c) > 0 {

		link, path := CreateLink(<-c)
		fmt.Println(link)
		req := head.CreateRequest(link)
		resp, _ := client.Do(req)
		result := query_handler.ConvertGzipToString(resp)

		query_handler.SaveJsonString(result, path)

	}

	wg.Done()
}
