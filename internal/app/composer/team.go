package composer

import (
	"context"
	"github.com/statistico/statistico-web-gateway/internal/app"
	"github.com/statistico/statistico-web-gateway/internal/app/grpc"
	"github.com/statistico/statistico-web-gateway/internal/app/grpc/proto"
)

type TeamComposer interface {
	TeamById(id uint64) (*app.Team, error)
	TeamsBySeasonId(seasonId uint64) ([]*app.Team, error)
}

type teamComposer struct {
	client grpc.TeamClient
}

func (t teamComposer) TeamById(id uint64) (*app.Team, error) {
	request := &proto.TeamRequest{TeamId: id}

	return t.client.TeamById(context.Background(), request)
}

func (t teamComposer) TeamsBySeasonId(seasonId uint64) ([]*app.Team, error) {
	return t.client.TeamsBySeasonId(context.Background(), seasonId)
}

func NewTeamComposer(c grpc.TeamClient) TeamComposer {
	return &teamComposer{client: c}
}
