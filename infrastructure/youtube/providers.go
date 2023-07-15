package youtube

import (
	"net/http"

	"github.com/google/wire"
)

// ProvideHTTPClient provides an HTTP client.
func ProvideHTTPClient() *http.Client {
	return http.DefaultClient
}

// Set is a Wire provider set that provides a YouTube API client.
var Set = wire.NewSet(NewAPI, ProvideHTTPClient)
