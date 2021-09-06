package healthz

import (
	"net/http"

	"github.com/rs/zerolog/log"
)

func Init(port string) {
	log.Debug().Msg("Starting healthz server")

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	http.ListenAndServe(":"+port, nil)
}
