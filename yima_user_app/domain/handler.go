package domain

import (
	"context"
	"git.horsecoder.com/chenpeiran/yima.demo.user/kitex_gen/yima/demo/user"
)

type YimaDemoUserImpl struct {
}

func (s YimaDemoUserImpl) Hello(ctx context.Context, req *user.MyReq) (resp string, err error) {
	resp = "hello"
	return resp, nil
}
