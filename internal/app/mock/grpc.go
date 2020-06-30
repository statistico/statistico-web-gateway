package mock

import (
	"context"
	"github.com/statistico/statistico-web-gateway/internal/app/grpc/proto"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
)

type TeamClient struct {
	mock.Mock
}

func (t *TeamClient) GetTeamByID(ctx context.Context, in *proto.TeamRequest, opts ...grpc.CallOption) (*proto.Team, error) {
	args := t.Called(ctx, in, opts)
	return args.Get(0).(*proto.Team), args.Error(1)
}
