package bootstrap

import "github.com/statistico/statistico-web-gateway/internal/app/rest"

func (c Container) RestResultHandler() *rest.ResultHandler {
	return rest.NewResultHandler(c.ResultComposer())
}

func (c Container) RestTeamHandler() *rest.TeamHandler {
	return rest.NewTeamHandler(c.TeamComposer())
}
