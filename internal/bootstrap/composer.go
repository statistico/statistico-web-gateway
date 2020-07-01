package bootstrap

import "github.com/statistico/statistico-web-gateway/internal/app/composer"

func (c Container) TeamComposer() composer.TeamComposer {
	return composer.NewTeamComposer(c.GRPCTeamClient())
}
