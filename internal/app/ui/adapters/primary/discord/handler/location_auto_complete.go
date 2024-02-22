package handler

import (
	"fmt"
	"strings"

	tempest "github.com/Amatsagu/Tempest"
	comctx "github.com/awlsring/texit/internal/app/ui/adapters/primary/discord/context"
	"github.com/awlsring/texit/internal/pkg/logger"
)

// TODO: Make this not discord specific so it can be used for the API and CLI validation
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
		{
			Name:        "ap-west",
			Geolocation: "Mumbai, India",
			Provider:    "Linode",
		},
		{
			Name:        "ca-central",
			Geolocation: "Toronto, Canada",
			Provider:    "Linode",
		},
		{
			Name:        "ap-southeast",
			Geolocation: "Sydney, Australia",
			Provider:    "Linode",
		},
		{
			Name:        "us-iad",
			Geolocation: "Virginia, USA",
			Provider:    "Linode",
		},
		{
			Name:        "us-ord",
			Geolocation: "Illinois, USA",
			Provider:    "Linode",
		},
		{
			Name:        "fr-par",
			Geolocation: "Paris, France",
			Provider:    "Linode",
		},
		{
			Name:        "us-sea",
			Geolocation: "Washington, USA",
			Provider:    "Linode",
		},
		{
			Name:        "br-gru",
			Geolocation: "Sao Paulo, Brazil",
			Provider:    "Linode",
		},
		{
			Name:        "nl-ams",
			Geolocation: "Amsterdam, Netherlands",
			Provider:    "Linode",
		},
		{
			Name:        "se-sto",
			Geolocation: "Stockholm, Sweden",
			Provider:    "Linode",
		},
		{
			Name:        "in-maa",
			Geolocation: "Chennai, India",
			Provider:    "Linode",
		},
		{
			Name:        "jp-osa",
			Geolocation: "Osaka, Japan",
			Provider:    "Linode",
		},
		{
			Name:        "it-mil",
			Geolocation: "Milan, Italy",
			Provider:    "Linode",
		},
		{
			Name:        "us-mia",
			Geolocation: "Florida, USA",
			Provider:    "Linode",
		},
		{
			Name:        "id-cgk",
			Geolocation: "Jakarta, Indonesia",
			Provider:    "Linode",
		},
		{
			Name:        "us-lax",
			Geolocation: "California (South), USA",
			Provider:    "Linode",
		},
		{
			Name:        "us-central",
			Geolocation: "Texas, USA",
			Provider:    "Linode",
		},
		{
			Name:        "us-west",
			Geolocation: "California (North), USA",
			Provider:    "Linode",
		},
		{
			Name:        "us-southeast",
			Geolocation: "Georgia, USA",
			Provider:    "Linode",
		},
		{
			Name:        "us-east",
			Geolocation: "New Jersey, USA",
			Provider:    "Linode",
		},
		{
			Name:        "eu-west",
			Geolocation: "London, UK",
			Provider:    "Linode",
		},
		{
			Name:        "ap-south",
			Geolocation: "Singapore",
			Provider:    "Linode",
		},
		{
			Name:        "eu-central",
			Geolocation: "Frankfurt, Germany",
			Provider:    "Linode",
		},
		{
			Name:        "ap-northeast",
			Geolocation: "Tokyo, Japan",
			Provider:    "Linode",
		},
		{
			Name:        "fsn1",
			Geolocation: "Falkenstein, Germany",
			Provider:    "Hetzner",
		},
		{
			Name:        "nbg1",
			Geolocation: "Nuremberg, Germany",
			Provider:    "Hetzner",
		},
		{
			Name:        "hel1",
			Geolocation: "Helsinki, Finland",
			Provider:    "Hetzner",
		},
		{
			Name:        "ash",
			Geolocation: "Virginia, USA",
			Provider:    "Hetzner",
		},
		{
			Name:        "hil",
			Geolocation: "Oregon, USA",
			Provider:    "Hetzner",
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
		filteredFields := []string{
			l.Name,
			l.Geolocation,
			strings.ToLower(l.Geolocation),
			l.Provider,
			strings.ToLower(l.Provider),
		}

		for _, field := range filteredFields {
			if strings.Contains(field, filter) {
				log.Debug().Str("location_name", l.Name).Msg("Adding location to choices")
				choices = append(choices, tempest.Choice{
					Name:  l.Display(),
					Value: l.Name,
				})
				break
			}
		}
	}

	if len(choices) >= 25 {
		return choices[:25]
	}

	return choices
}
