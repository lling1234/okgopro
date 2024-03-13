package info

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/mitchellh/mapstructure"
	"log"
	"okgopro/internal/common/utils/snowflake"
	"okgopro/internal/consts"
	"okgopro/internal/dao"
	"okgopro/internal/model/entity"
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
小k网
https://www.xkwo.com/
小黑资源网
https://www.xiaoheizyw.com/
可可资源网
https://www.kekezyw.cn/
qq好基友
https://www.qqhjy1.xyz/

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

	//fmt.Print(rr.ReadAllString())
	//	获取当前的日期，例如12-11
	now := time.Now()
	day := now.Format("01-02")
	glog.Info(ctx, day)

	// 使用 goquery 解析 HTML
	doc, err := goquery.NewDocumentFromReader(rr.Response.Body)
	if err != nil {
		log.Fatal(err)
	}

	nrDiv := doc.Find("div.container")
	find := nrDiv.Find("div.news-article_container")
	//log.Println(find.Text())
	ul := find.Find("ul")
	selection := ul.Children().FilterFunction(func(i int, selections *goquery.Selection) bool {
		val, _ := selections.Attr("class")
		return val != "addd"
	})
	selection.Each(func(i int, selection *goquery.Selection) {
		s2 := selection.Find("li a")
		urlVal, exists := s2.Attr("href")
		if exists {
			log.Println(exists)
		}
		titleVal, titleErr := s2.Attr("title")
		if titleErr {
			log.Println(titleErr)
		}
		date := selection.Find("li span").Text()

		log.Println("--------------")
		log.Println(titleVal)
		log.Println(urlVal)
		log.Println(date)
		log.Println("--------------")

		InsterDBContent(ctx, entity.Content{
			Title:      titleVal,
			Extra:      date,
			PrefixUrl:  consts.WEBURLxiaodao1,
			Url:        urlVal,
			Link:       urlVal,
			WebUrlId:   consts.WEBURLIDxiaodao1,
			CreateTime: gtime.Now(),
		})

	})
	log.Println("1111111 end")

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
	//TODO  支持通过接口获取  https://www.xkwo.com/media-all?page=2&size=30
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
		urlVal, exists := a.Attr("href")
		if exists {
			log.Println("findAHref", urlVal)
		}
		findDivCol := a.Find("div.col")
		findDivTextStr := findDivCol.Find("div.text").Text()
		findDivText := strings.TrimSpace(findDivTextStr)
		findDivTime := findDivCol.Find("div.time").Text()
		log.Println("findDivText", findDivText)
		log.Println("findDivTime", findDivTime)

		// 保存到数据库 content表
		InsterDBContent(ctx, entity.Content{
			Title:      findDivText,
			Extra:      findDivTime,
			PrefixUrl:  "https://www.xkwo.com/",
			Url:        urlVal,
			WebUrlId:   consts.WEBURLIDxkwo,
			CreateTime: gtime.Now(),
		})
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
func InsterDBUser(ctx context.Context, in entity.User) (err error) {
	// 创建 MySQL 连接对象
	db := g.DB()

	// 开启事务
	tx, err := db.Begin(ctx)
	if err != nil {
		fmt.Println("开启事务失败:", err)
		return
	}

	in.Id = int64(snowflake.SnowGen.NextVal())
	//Password:okgopro base64:b2tnb3Bybw==
	in.Password = "b2tnb3Bybw=="
	in.Phone = "13800138001"
	insert, err := dao.User.DB().Insert(ctx, "user", in)
	if err != nil {
		panic(err)
	}
	err2 := tx.Commit()
	if err2 != nil {
		panic(err2)
	}

	log.Println("insert", insert)
	return err
}

func InsterDBContent(ctx context.Context, in entity.Content) (err error) {

	// 创建 MySQL 连接对象
	db := g.DB()

	// 开启事务
	tx, err := db.Begin(ctx)
	if err != nil {
		fmt.Println("开启事务失败:", err)
		return
	}

	in.Id = int64(snowflake.SnowGen.NextVal())

	insert, err := dao.Content.DB().Insert(ctx, "content", in)
	if err != nil {
		panic(err)
	}
	err2 := tx.Commit()
	if err2 != nil {
		panic(err2)
	}
	log.Println("insert", insert)
	return err
}

func InsterDBTitle(ctx context.Context, in entity.Title) (err error) {
	// 创建 MySQL 连接对象
	db := g.DB()

	// 开启事务
	tx, err := db.Begin(ctx)
	if err != nil {
		fmt.Println("开启事务失败:", err)
		return
	}

	in.Id = int64(snowflake.SnowGen.NextVal())
	insert, err := dao.Title.DB().Insert(ctx, "title", in)
	if err != nil {
		panic(err)
	}
	err2 := tx.Commit()
	if err2 != nil {
		panic(err2)
	}
	log.Println("insert", insert)
	return err
}

func QueryDBContent(ctx context.Context, in entity.Content) (r consts.R, err error) {
	var ctxList []entity.Content

	err = dao.Content.Ctx(ctx).Scan(&ctxList, in)
	//TODO 结果转换
	var data []consts.Data
	err = mapstructure.Decode(ctxList, &data)
	if err != nil {
		log.Println(err)
		return r, err
	}

	jsonData, err := json.Marshal(data)
	log.Println(string(jsonData))

	//TODO 分页
	//TODO 响应内容拼接
	log.Println(ctxList[0].Id)
	return r, nil
}
