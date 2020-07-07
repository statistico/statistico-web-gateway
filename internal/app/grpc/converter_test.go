package grpc

import (
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/statistico/statistico-web-gateway/internal/app"
	"github.com/statistico/statistico-web-gateway/internal/app/grpc/proto"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConverter(t *testing.T) {
	t.Run("converts proto result struct into app result struct", func(t *testing.T) {
		t.Helper()

		result := newProtoResult()

		converted, err := convertResult(result)

		if err != nil {
			t.Fatalf("Expected nil, got %s", err.Error())
		}

		a := assert.New(t)

		a.Equal(uint64(78102), converted.ID)
		a.Equal(uint64(1), converted.HomeTeam.ID)
		a.Equal("West Ham United", converted.HomeTeam.Name)
		a.Equal("WHU", *converted.HomeTeam.ShortCode)
		a.Equal(uint64(8), converted.HomeTeam.CountryID)
		a.Equal(uint64(214), converted.HomeTeam.VenueID)
		a.Equal(false, converted.HomeTeam.NationalTeam)
		a.Equal(uint64(1895), *converted.HomeTeam.Founded)
		a.Equal("logo", *converted.HomeTeam.Logo)
		a.Equal(uint64(10), converted.AwayTeam.ID)
		a.Equal("Nottingham Forest", converted.AwayTeam.Name)
		a.Equal("NOT", *converted.AwayTeam.ShortCode)
		a.Equal(uint64(8), converted.AwayTeam.CountryID)
		a.Equal(uint64(300), converted.AwayTeam.VenueID)
		a.Equal(true, converted.AwayTeam.NationalTeam)
		a.Equal(uint64(1895), *converted.AwayTeam.Founded)
		a.Equal("logo", *converted.AwayTeam.Logo)
		a.Equal(uint64(16036), converted.Season.ID)
		a.Equal("2019/2020", converted.Season.Name)
		a.Equal(true, converted.Season.IsCurrent)
		a.Equal(uint64(38), converted.Round.ID)
		a.Equal("38", converted.Round.Name)
		a.Equal(uint64(16036), converted.Round.SeasonID)
		a.Equal("2020-07-07T12:00:00Z", converted.Round.StartDate.String())
		a.Equal("2020-07-23T23:59:59Z", converted.Round.EndDate.String())
		a.Equal(uint64(214), converted.Venue.ID)
		a.Equal("London Stadium", converted.Venue.Name)
		a.Equal("2020-07-07T15:00:00Z", converted.DateTime.String())
		a.Equal(uint8(5), converted.Stats.HomeScore)
		a.Equal(uint8(2), converted.Stats.AwayScore)
	})

	t.Run("can handle nullable fields", func(t *testing.T) {
		t.Helper()

		home := proto.Team{
			Id:        1,
			Name:      "West Ham United",
			CountryId: 8,
			VenueId:   214,
		}

		away := proto.Team{
			Id:        10,
			Name:      "Nottingham Forest",
			CountryId: 8,
			VenueId:   300,
		}

		season := proto.Season{
			Id:   16036,
			Name: "2019/2020",
		}

		round := proto.Round{
			Id:        38,
			Name:      "38",
			SeasonId:  16036,
			StartDate: "2020-07-07T12:00:00+00:00",
			EndDate:   "2020-07-23T23:59:59+00:00",
		}

		venue := proto.Venue{
			Id:   214,
			Name: "London Stadium",
		}

		date := proto.Date{
			Utc: 1594132077,
			Rfc: "2020-07-07T15:00:00+00:00",
		}

		result := proto.Result{
			Id:       78102,
			HomeTeam: &home,
			AwayTeam: &away,
			Season:   &season,
			Round:    &round,
			Venue:    &venue,
			DateTime: &date,
		}

		converted, err := convertResult(&result)

		if err != nil {
			t.Fatalf("Expected nil, got %s", err.Error())
		}

		a := assert.New(t)

		a.Equal(uint64(78102), converted.ID)
		a.Equal(uint64(1), converted.HomeTeam.ID)
		a.Equal("West Ham United", converted.HomeTeam.Name)
		a.Nil(converted.HomeTeam.ShortCode)
		a.Equal(uint64(8), converted.HomeTeam.CountryID)
		a.Equal(uint64(214), converted.HomeTeam.VenueID)
		a.False(converted.HomeTeam.NationalTeam)
		a.Nil(converted.HomeTeam.Founded)
		a.Nil(converted.HomeTeam.Logo)
		a.Equal(uint64(10), converted.AwayTeam.ID)
		a.Equal("Nottingham Forest", converted.AwayTeam.Name)
		a.Nil(converted.AwayTeam.ShortCode)
		a.Equal(uint64(8), converted.AwayTeam.CountryID)
		a.Equal(uint64(300), converted.AwayTeam.VenueID)
		a.False(converted.AwayTeam.NationalTeam)
		a.Nil(converted.AwayTeam.Founded)
		a.Nil(converted.AwayTeam.Logo)
		a.Equal(uint64(16036), converted.Season.ID)
		a.Equal("2019/2020", converted.Season.Name)
		a.False(converted.Season.IsCurrent)
		a.Equal(uint64(38), converted.Round.ID)
		a.Equal("38", converted.Round.Name)
		a.Equal(uint64(16036), converted.Round.SeasonID)
		a.Equal("2020-07-07T12:00:00Z", converted.Round.StartDate.String())
		a.Equal("2020-07-23T23:59:59Z", converted.Round.EndDate.String())
		a.Equal(uint64(214), converted.Venue.ID)
		a.Equal("London Stadium", converted.Venue.Name)
		a.Equal("2020-07-07T15:00:00Z", converted.DateTime.String())
		a.Equal(app.ResultStats{}, converted.Stats)
	})

	t.Run("returns error if unable to parse date time string", func(t *testing.T) {
		t.Helper()

		date := proto.Date{
			Rfc: "Wooooo Yeah",
		}

		result := proto.Result{
			DateTime: &date,
		}

		_, err := convertResult(&result)

		if err == nil {
			t.Fatal("Expected error got nil")
		}

		assert.Equal(
			t,
			"parsing time \"Wooooo Yeah\" as \"2006-01-02T15:04:05Z07:00\": cannot parse \"Wooooo Yeah\" as \"2006\"",
			err.Error(),
		)
	})
}

func newProtoResult() *proto.Result {
	home := proto.Team{
		Id:             1,
		Name:           "West Ham United",
		ShortCode:      &wrappers.StringValue{Value: "WHU"},
		CountryId:      8,
		VenueId:        214,
		IsNationalTeam: &wrappers.BoolValue{Value: false},
		Founded:        &wrappers.UInt64Value{Value: 1895},
		Logo:           &wrappers.StringValue{Value: "logo"},
	}

	away := proto.Team{
		Id:             10,
		Name:           "Nottingham Forest",
		ShortCode:      &wrappers.StringValue{Value: "NOT"},
		CountryId:      8,
		VenueId:        300,
		IsNationalTeam: &wrappers.BoolValue{Value: true},
		Founded:        &wrappers.UInt64Value{Value: 1895},
		Logo:           &wrappers.StringValue{Value: "logo"},
	}

	season := proto.Season{
		Id:        16036,
		Name:      "2019/2020",
		IsCurrent: &wrappers.BoolValue{Value: true},
	}

	round := proto.Round{
		Id:        38,
		Name:      "38",
		SeasonId:  16036,
		StartDate: "2020-07-07T12:00:00+00:00",
		EndDate:   "2020-07-23T23:59:59+00:00",
	}

	venue := proto.Venue{
		Id:   214,
		Name: "London Stadium",
	}

	date := proto.Date{
		Utc: 1594132077,
		Rfc: "2020-07-07T15:00:00+00:00",
	}

	stats := proto.MatchStats{
		HomeScore: &wrappers.UInt32Value{Value: 5},
		AwayScore: &wrappers.UInt32Value{Value: 2},
	}

	return &proto.Result{
		Id:       78102,
		HomeTeam: &home,
		AwayTeam: &away,
		Season:   &season,
		Round:    &round,
		Venue:    &venue,
		DateTime: &date,
		Stats:    &stats,
	}
}
