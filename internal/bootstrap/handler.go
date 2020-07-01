package bootstrap

import "github.com/statistico/statistico-web-gateway/internal/app/rest"

func (c Container) RestTeamHandler() *rest.TeamHandler {
	return rest.NewTeamHandler(c.TeamComposer())
}
