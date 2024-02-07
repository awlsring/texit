package handler

import (
	"fmt"
	"strings"

	tempest "github.com/Amatsagu/Tempest"
	comctx "github.com/awlsring/texit/internal/app/ui/adapters/primary/discord/context"
	"github.com/awlsring/texit/internal/pkg/logger"
)

var (
	locations = []LocationChoice{
		{
			Name:        "us-east-1",
			Geolocation: "Virginia, USA",
			Provider:    "AWS",
		},
		{
			Name:        "us-east-2",
			Geolocation: "Ohio, USA",
			Provider:    "AWS",
		},
		{
			Name:        "us-west-1",
			Geolocation: "California, USA",
			Provider:    "AWS",
		},
		{
			Name:        "us-west-2",
			Geolocation: "Oregon, USA",
			Provider:    "AWS",
		},
		{
			Name:        "eu-west-1",
			Geolocation: "Dublin, Ireland",
			Provider:    "AWS",
		},
		{
			Name:        "eu-west-2",
			Geolocation: "London, UK",
			Provider:    "AWS",
		},
		{
			Name:        "eu-west-3",
			Geolocation: "Paris, France",
			Provider:    "AWS",
		},
		{
			Name:        "eu-central-1",
			Geolocation: "Frankfurt, Germany",
			Provider:    "AWS",
		},
		{
			Name:        "eu-central-2",
			Geolocation: "Zurich, Switzerland",
			Provider:    "AWS",
		},
		{
			Name:        "eu-north-1",
			Geolocation: "Stockholm, Sweden",
			Provider:    "AWS",
		},
		{
			Name:        "eu-south-1",
			Geolocation: "Milan, Italy",
			Provider:    "AWS",
		},
		{
			Name:        "eu-south-2",
			Geolocation: "Madrid, Spain",
			Provider:    "AWS",
		},
		{
			Name:        "ap-northeast-1",
			Geolocation: "Tokyo, Japan",
			Provider:    "AWS",
		},
		{
			Name:        "ap-northeast-2",
			Geolocation: "Seoul, South Korea",
			Provider:    "AWS",
		},
		{
			Name:        "ap-northeast-3",
			Geolocation: "Osaka, Japan",
			Provider:    "AWS",
		},
		{
			Name:        "ap-southeast-1",
			Geolocation: "Singapore",
			Provider:    "AWS",
		},
		{
			Name:        "ap-southeast-2",
			Geolocation: "Sydney, Australia",
			Provider:    "AWS",
		},
		{
			Name:        "ap-southeast-3",
			Geolocation: "Jakarta, Indonesia",
			Provider:    "AWS",
		},
		{
			Name:        "ap-southeast-4",
			Geolocation: "Melbourne, Australia",
			Provider:    "AWS",
		},
		{
			Name:        "ap-south-1",
			Geolocation: "Mumbai, India",
			Provider:    "AWS",
		},
		{
			Name:        "ap-south-2",
			Geolocation: "Hyderabad, India",
			Provider:    "AWS",
		},
		{
			Name:        "sa-east-1",
			Geolocation: "Sao Paulo, Brazil",
			Provider:    "AWS",
		},
		{
			Name:        "af-south-1",
			Geolocation: "Cape Town, South Africa",
			Provider:    "AWS",
		},
		{
			Name:        "me-south-1",
			Geolocation: "Bahrain",
			Provider:    "AWS",
		},
		{
			Name:        "me-central-1",
			Geolocation: "UAE",
			Provider:    "AWS",
		},
		{
			Name:        "il-central-1",
			Geolocation: "Israel",
			Provider:    "AWS",
		},
		{
			Name:        "ca-central-1",
			Geolocation: "Canada",
			Provider:    "AWS",
		},
		{
			Name:        "ca-west-1",
			Geolocation: "Calgary, Canada",
			Provider:    "AWS",
		},
		{
			Name:        "ap-east-1",
			Geolocation: "Hong Kong",
			Provider:    "AWS",
		},
	}
)

type LocationChoice struct {
	Name        string
	Geolocation string
	Provider    string
}

func (l LocationChoice) Display() string {
	// us-east-1 (Virginia, USA) - AWS
	return fmt.Sprintf("%s (%s) - %s", l.Name, l.Geolocation, l.Provider)
}

func (h *Handler) ProviderLocationAutoComplete(ctx *comctx.CommandContext, name, filter string) []tempest.Choice {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Auto completing location")
	choices := []tempest.Choice{}

	for _, l := range locations {
		if strings.Contains(l.Name, filter) || strings.Contains(l.Geolocation, filter) || strings.Contains(strings.ToLower(l.Provider), filter) {
			log.Debug().Str("location_name", l.Name).Msg("Adding location to choices")
			choices = append(choices, tempest.Choice{
				Name:  l.Display(),
				Value: l.Name,
			})
		}
	}

	if len(choices) >= 25 {
		return choices[:25]
	}

	return choices
}
