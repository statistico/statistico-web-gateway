package bootstrap

import "github.com/statistico/statistico-web-gateway/internal/app/rest"

func (c Container) RestCompetitionHandler() *rest.CompetitionHandler {
	return rest.NewCompetitionHandler(c.CompetitionComposer())
}

func (c Container) RestResultHandler() *rest.ResultHandler {
	return rest.NewResultHandler(c.ResultComposer())
}

func (c Container) RestSeasonHandler() *rest.SeasonHandler {
	return rest.NewSeasonHandler(c.CompetitionComposer(), c.SeasonComposer())
}

func (c Container) RestTeamHandler() *rest.TeamHandler {
	return rest.NewTeamHandler(c.TeamComposer())
}

func (c Container) RestTeamStatHandler() *rest.TeamStatHandler {
	return rest.NewTeamStatHandler(c.TeamStatComposer())
}
