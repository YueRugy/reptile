package avtb

import (
	"bufio"
	"fmt"
	"github.com/antchfx/htmlquery"
	"github.com/colly"
	"github.com/reptile/pool"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
	"sync"
)

var (
	imgPath="/home/yue/Downloads/avtb/"
	w sync.WaitGroup
)

func Run()  {

	var handler=downLoadImg
	task:=pool.NewTask(100,handler)
	go task.DisposeData()
	c := colly.NewCollector(
		//colly.Async(true),

		colly.UserAgent("Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.105 Safari/537.36"),
	)
	//_ = c.Limit(&colly.LimitRule{DomainGlob: "*avtb2043*", Parallelism: 5})

	c.OnRequest(func(request *colly.Request) {
		//fmt.Println("request:  ", request.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("something is wrong ", err)
	})

	c.OnResponse(func(response *colly.Response) {
		doc, err := htmlquery.Parse(strings.NewReader(string(response.Body)))
		if err != nil {
			log.Fatal("parse error")
		}
		nodes := htmlquery.Find(doc, `//ul[@class="videos"]/
									li/div[@class="video"]/a/div[@class="video-thumb"]`)
		//nodes := htmlquery.Find(doc, `//ul[@class="videos"]/
		//							li/div[@class="video"]/div[@class="video-thumb"]`)
		go func() {
			for _, node := range nodes {
				task.AddData(htmlquery.InnerText(htmlquery.FindOne(node, `./img/@src`)),
					htmlquery.InnerText( htmlquery.FindOne(node, `./img/@alt`)))
			}
		}()
		//	fmt.Println("response: ", response.Request.URL)
	})

	c.OnHTML(".prevnext", func(e *colly.HTMLElement) {

		p := e.Attr("href")
		if p == "/recent/" {
			return
		}
		activePath := e.Request.URL.Path
		pathNum := strings.Split(p, "/")[1]
		activeNum := strings.Split(activePath, "/")[1]
		if activeNum > pathNum {
			return
		}
		//fmt.Println(e.Request.URL)
		_ = e.Request.Visit(e.Attr("href"))
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
	})

	_ = c.Visit("http://www.avtb2043.com/recent/")
	w.Add(1)
	w.Wait()

}
func downLoadImg(job *pool.Data)  {
	res,err:=http.Get(job.Url)
	if err!=nil{
		log.Fatal("downloading image url error")
	}
	defer res.Body.Close()
	fileName:=path.Base(job.Title)
	reader:=bufio.NewReaderSize(res.Body,32*1024)

	file,err:=os.Create(imgPath+fileName)
	if err!=nil{
		log.Fatal("file is exist or img path error")
	}
	wr:=bufio.NewWriter(file)
	_, _ = io.Copy(wr, reader)
	_ = wr.Flush()

}
