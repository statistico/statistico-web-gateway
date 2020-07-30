package bootstrap

import "github.com/statistico/statistico-web-gateway/internal/app/composer"

func (c Container) CompetitionComposer() composer.CompetitionComposer {
	return composer.NewCompetitionComposer(c.GRPCCompetitionClient())
}

func (c Container) ResultComposer() composer.ResultComposer {
	return composer.NewResultComposer(c.GRPCResultClient())
}

func (c Container) TeamComposer() composer.TeamComposer {
	return composer.NewTeamComposer(c.GRPCTeamClient())
}
