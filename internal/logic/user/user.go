package user

import (
	"context"
	"fmt"
	"okgopro/internal/dao"
	"okgopro/internal/model/do"
	"okgopro/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/grand"
)

type (
	sUser struct{}
)

func init() {
	service.RegisterUser(New())
}

func New() service.IUser {
	return &sUser{}
}

// Create creates user account.
func (s *sUser) Create(ctx context.Context) (err error) {
	fmt.Println("1111~~~~~~")
	// return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
	// 	glog.Info(ctx, "create user start")

	// 	_, err = dao.User.Ctx(ctx).Data(do.User{
	// 		Id:         grand.Intn(999999999999999999),
	// 		Username:   "user1",
	// 		Password:   "123123",
	// 		Phone:      13298316872,
	// 		CreateUser: 1,
	// 		CreateTime: gtime.Now(),
	// 		IsAdmin:    0,
	// 		IsDeleted:  0,
	// 	}).Insert()
	// 	glog.Info(ctx, err.Error())
	// 	glog.Info(ctx, "create user end")
	// 	tx.Commit()
	// 	defer func() {
	// 		if err != nil {
	// 			tx.Rollback()
	// 			return
	// 		} else {
	// 			tx.Commit()
	// 		}
	// 	}()

	// 	return err
	// })
	// -------------------
	// return dao.User.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
	// 	glog.Info(ctx, "create user start")

	// 	_, err = dao.User.Ctx(ctx).Data(do.User{
	// 		Id:         grand.Intn(999999999999999999),
	// 		Username:   "user1",
	// 		Password:   "123123",
	// 		Phone:      13298316872,
	// 		CreateUser: 1,
	// 		CreateTime: gtime.Now(),
	// 		IsAdmin:    0,
	// 		IsDeleted:  0,
	// 	}).Insert()
	// 	glog.Info(ctx, err.Error())
	// 	glog.Info(ctx, "create user end")
	// 	tx.Commit()
	// 	return nil
	// })
	// 创建 MySQL 连接对象
	db := g.DB()

	// 开启事务
	tx, err := db.Begin(ctx)
	if err != nil {
		fmt.Println("开启事务失败:", err)
		return
	}
	_, err = dao.User.Ctx(ctx).Data(do.User{
		Id:         grand.Intn(999999999999999999),
		Username:   "user1",
		Password:   "123123",
		Phone:      13298316872,
		CreateUser: 1,
		CreateTime: gtime.Now(),
		IsAdmin:    0,
		IsDeleted:  0,
	}).Insert()
	// 提交事务
	if err1 := tx.Commit(); err != nil {
		fmt.Println("提交事务失败:", err1)
		return
	}
	fmt.Println("22222!!!!!!!!!!!!")
	//
	return nil
}
