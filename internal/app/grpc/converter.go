package grpc

import (
	"github.com/statistico/statistico-web-gateway/internal/app"
	"github.com/statistico/statistico-web-gateway/internal/app/grpc/proto"
	"time"
)

func convertCompetition(c *proto.Competition) *app.Competition {
	var x app.Competition
	x.ID = c.GetId()
	x.Name = c.GetName()
	x.IsCup = c.GetIsCup()
	x.CountryID = c.GetCountryId()

	return &x
}

func convertResult(result *proto.Result) (*app.Result, error) {
	d, err := time.Parse(time.RFC3339, result.GetDateTime().GetRfc())

	if err != nil {
		return nil, err
	}

	round, err := convertRound(result.GetRound())

	if err != nil {
		return nil, err
	}

	r := app.Result{
		ID:       result.GetId(),
		HomeTeam: convertTeam(result.GetHomeTeam()),
		AwayTeam: convertTeam(result.GetAwayTeam()),
		Season:   convertSeason(result.GetSeason()),
		Round:    round,
		Venue:    convertVenue(result.GetVenue()),
		DateTime: app.JsonDate(d),
		Stats:    convertResultStats(result.GetStats()),
	}

	return &r, nil
}

func convertResultStats(stats *proto.MatchStats) app.ResultStats {
	s := app.ResultStats{}

	if stats.GetHomeScore() != nil {
		s.HomeScore = uint8(stats.GetHomeScore().GetValue())
	}

	if stats.GetAwayScore() != nil {
		s.AwayScore = uint8(stats.GetAwayScore().GetValue())
	}

	return s
}

func convertRound(round *proto.Round) (*app.Round, error) {
	if round == nil {
		return nil, nil
	}

	start, err := time.Parse(time.RFC3339, round.GetStartDate())

	if err != nil {
		return nil, err
	}

	end, err := time.Parse(time.RFC3339, round.GetEndDate())

	if err != nil {
		return nil, err
	}

	r := app.Round{
		ID:        round.GetId(),
		Name:      round.GetName(),
		SeasonID:  round.GetSeasonId(),
		StartDate: app.JsonDate(start),
		EndDate:   app.JsonDate(end),
	}

	return &r, nil
}

func convertSeason(season *proto.Season) app.Season {
	s := app.Season{
		ID:        season.GetId(),
		Name:      season.GetName(),
		IsCurrent: season.GetIsCurrent().GetValue(),
	}

	return s
}

func convertTeam(team *proto.Team) app.Team {
	t := app.Team{
		ID:        team.GetId(),
		Name:      team.GetName(),
		CountryID: team.GetCountryId(),
		VenueID:   team.GetVenueId(),
	}

	if team.GetShortCode() != nil {
		t.ShortCode = &team.GetShortCode().Value
	}

	if team.GetIsNationalTeam() != nil {
		t.NationalTeam = team.GetIsNationalTeam().Value
	}

	if team.GetFounded() != nil {
		t.Founded = &team.GetFounded().Value
	}

	if team.GetLogo() != nil {
		t.Logo = &team.GetLogo().Value
	}

	return t
}

func convertTeamStat(stat *proto.TeamStat) *app.TeamStat {
	s := app.TeamStat{
		FixtureID: stat.FixtureId,
		Stat:      stat.Stat,
	}

	if stat.GetValue() != nil {
		s.Value = &stat.GetValue().Value
	}

	return &s
}

func convertVenue(venue *proto.Venue) app.Venue {
	return app.Venue{
		ID:   venue.GetId(),
		Name: venue.GetName(),
	}
}
