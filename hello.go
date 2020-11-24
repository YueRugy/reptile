package main

import "github.com/reptile/avtb"

func main() {
	avtb.Run()
	//var handler = func(jobs  *pool.Data) {
	//	fmt.Println(jobs)
	//}
	//task := pool.NewTask(500, handler)

	//go func() {
	//	for i := 0; i < 100000; i++ {
	//		iStr := strconv.Itoa(i)
	//		url := "这里面去定义你的任务数据格式" + iStr
	//		task.AddData(url)
	//	}
	//}()

	//task.Consumer.DisposeData(task.Product.Jobs)
	/*dp := &pool.DownloadPool{
		Pool: pool.New(10),
	}

	for i := 0; i < 100; i++ {
		go dp.Execute()
	}

	dp.Wait()*/
	//	var handler=downLoadImg
	//	task:=pool.NewTask(100,handler)
	//	go task.DisposeData()
	//	c := colly.NewCollector(
	//		//colly.Async(true),
	//
	//		colly.UserAgent("Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.105 Safari/537.36"),
	//	)
	//	//_ = c.Limit(&colly.LimitRule{DomainGlob: "*avtb2043*", Parallelism: 5})
	//
	//	c.OnRequest(func(request *colly.Request) {
	//		//fmt.Println("request:  ", request.URL)
	//	})
	//
	//	c.OnError(func(_ *colly.Response, err error) {
	//		fmt.Println("something is wrong ", err)
	//	})
	//
	//	c.OnResponse(func(response *colly.Response) {
	//		doc, err := htmlquery.Parse(strings.NewReader(string(response.Body)))
	//		if err != nil {
	//			log.Fatal("parse error")
	//		}
	//		nodes := htmlquery.Find(doc, `//ul[@class="videos"]/
	//									li/div[@class="video"]/a/div[@class="video-thumb"]`)
	//		//nodes := htmlquery.Find(doc, `//ul[@class="videos"]/
	//		//							li/div[@class="video"]/div[@class="video-thumb"]`)
	//		go func() {
	//			for _, node := range nodes {
	//				task.AddData(htmlquery.InnerText(htmlquery.FindOne(node, `./img/@src`)),
	//					htmlquery.InnerText( htmlquery.FindOne(node, `./img/@alt`)))
	//			}
	//		}()
	//		//	fmt.Println("response: ", response.Request.URL)
	//	})
	//
	//	c.OnHTML(".prevnext", func(e *colly.HTMLElement) {
	//
	//		p := e.Attr("href")
	//		if p == "/recent/" {
	//			return
	//		}
	//		activePath := e.Request.URL.Path
	//		pathNum := strings.Split(p, "/")[1]
	//		activeNum := strings.Split(activePath, "/")[1]
	//		if activeNum > pathNum {
	//			return
	//		}
	//		//fmt.Println(e.Request.URL)
	//		_ = e.Request.Visit(e.Attr("href"))
	//	})
	//
	//	c.OnScraped(func(r *colly.Response) {
	//		fmt.Println("Finished", r.Request.URL)
	//	})
	//
	//	_ = c.Visit("http://www.avtb2043.com/recent/")
	//	w.Add(1)
	//	w.Wait()
	//
	//}
}
