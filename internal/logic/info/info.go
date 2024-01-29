package info

import (
	"context"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/util/grand"
	"log"
	"okgopro/internal/dao"
	"strings"
	"time"
)

type (
	sInfo struct{}
)

func init() {

}
func New() {

}

/**
小刀娱乐网
https://www.xiaodao1.com/
我爱资源网
https://www.52zyw1.com/
小k网
https://www.xkwo.com/
**/

func (s *sInfo) GetXiaodao(ctx context.Context) {
	glog.Info(ctx, "get GetXiaodao")

	//	http get请求获取https://www.xiaodao1.com/的响应
	url := "https://www.xiaodao1.com"
	rr, err1 := g.Client().Get(ctx, url)
	if err1 != nil {
		panic(err1)
	}

	defer rr.Close()

	fmt.Print(rr.ReadAllString())
	//	获取当前的日期，例如12-11
	now := time.Now()
	day := now.Format("01-02")
	glog.Info(ctx, day)

}

func (s *sInfo) Get52zyw1(ctx context.Context) {
	url := "https://www.52zyw1.com/"
	rr, err1 := g.Client().Get(ctx, url)
	if err1 != nil {
		panic(err1)
	}

	defer rr.Close()

	fmt.Print(rr.ReadAllString())
}

func (s *sInfo) Getxk(ctx context.Context) {

	strTest := "日曝光1W+ 快手引流获客全套脚本+详细教程"
	space := strings.TrimSpace(strTest)
	log.Println(space)
	log.Println("11111~~~~~~~")

	glog.Info(ctx, "get Getxk https://www.xkwo.com/")

	url := "https://www.xkwo.com/"
	rr, err1 := g.Client().Get(ctx, url)
	if err1 != nil {
		panic(err1)
	}

	defer rr.Close()

	//fmt.Print(rr.ReadAllString())

	// 使用 goquery 解析 HTML
	doc, err := goquery.NewDocumentFromReader(rr.Response.Body)
	if err != nil {
		log.Fatal(err)
	}

	// 获取div class为n-r的元素
	nrDiv := doc.Find("div.n-r")

	// 在n-r元素中查找div class为content的元素
	contentDiv := nrDiv.Find("div.content")

	// 在contentDiv元素中查找ul class为media media-ul的内容
	mediaUl := contentDiv.Find("ul.media.media-ul")

	mediaUl.Find("li a").Each(func(i int, a *goquery.Selection) {
		log.Println("a 11", a.Text())
		val, exists := a.Attr("href")
		if exists {
			log.Println("findAHref", val)
		}
		findDivCol := a.Find("div.col")
		findDivTextStr := findDivCol.Find("div.text").Text()
		findDivText := strings.TrimSpace(findDivTextStr)
		findDivTime := findDivCol.Find("div.time")
		log.Println("findDivText", findDivText)
		log.Println("findDivTime", findDivTime.Text())
	})
	log.Println("------------11")

	// 打印ul中的内容
	mediaUl.Each(func(i int, ul *goquery.Selection) {
		fmt.Println(ul.Text())
		log.Println("--------")
		log.Println(ul.Get(i))
	})

	log.Println("end 11")
}
func (s *sInfo) GetxkInfoDB(ctx context.Context) {
	// 创建 MySQL 连接对象
	db := g.DB()

	// 开启事务
	tx, err := db.Begin(ctx)
	if err != nil {
		fmt.Println("开启事务失败:", err)
		return
	}
	insert, err := dao.Content.Ctx(ctx).Insert(ctx, "content", g.Map{"id": grand.Intn(9999999999999999), "title": "111"})
	if err != nil {
		panic(err)
	}
	err2 := tx.Commit()
	if err2 != nil {
		panic(err2)
	}
	log.Println("insert", insert)
}
