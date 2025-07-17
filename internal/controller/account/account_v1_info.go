package account

import (
	"context"

	v1 "easyword/api/account/v1"
)

func (c *ControllerV1) Info(ctx context.Context, req *v1.InfoReq) (res *v1.InfoRes, err error) {
	user, err := c.users.Info(ctx)
	if err != nil {
		return nil, err
	}
	return &v1.InfoRes{
		Username: user.Username,
		Email:    user.Email,
		CreateAt: user.CreatedAt,
		UpdateAt: user.UpdatedAt,
	}, nil
}
