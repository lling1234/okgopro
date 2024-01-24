package info

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
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
func (s *sInfo) Get52zyw(ctx context.Context) {
	glog.Info(ctx, "get Get52zyw")
}
func (s *sInfo) Getxk(ctx context.Context) {
	glog.Info(ctx, "get Getxk")
}
