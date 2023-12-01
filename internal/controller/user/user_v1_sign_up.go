package user

import (
	"context"
	"fmt"

	"okgopro/api/user/v1"
	"okgopro/internal/service"
)

func (c *ControllerV1) SignUp(ctx context.Context, req *v1.SignUpReq) (res *v1.SignUpRes, err error) {
	fmt.Println("111333--------")
	service.User().Create(ctx)
	// service.RegisterUser(service.User())
	// service.IUser.Create(service.User(),ctx)
	return nil, nil
}
