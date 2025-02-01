package hello

import (
	"context"
	"monitor-dashboard/api/hello/v1"
)

func (c *ControllerV1) Hello(ctx context.Context, req *v1.HelloReq) (res *v1.HelloRes, err error) {
	return nil, nil
}
