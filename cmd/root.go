package cmd

import (
	"github.com/rs/zerolog/log"
	"gitlab.playcourt.id/telkom-digital/svc/spbe/perizinan-event-user-management/cmd/rest"
)

func Execute() {
	if err := rest.ServeREST(); err != nil {
		log.Fatal().Err(err).Msg("error while starting http server")
	}
}
