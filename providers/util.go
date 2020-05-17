package providers

import (
	"fmt"
	"net/http"
)

const (
	tokenTypeBearer = "Bearer"
	tokenTypeToken  = "token"

	acceptApplicationJSON       = "application/json"
	acceptApplicationGitHubJSON = "application/vnd.github.v3+json"
)

func getAuthorizationHeader(prefix, token, accept string, includeLi bool) http.Header {
	header := make(http.Header)
	if accept != "" {
		header.Set("Accept", accept)
	}
	if includeLi {
		header.Set("x-li-format", "json")
	}
	header.Set("Authorization", fmt.Sprintf("%s %s", prefix, token))
	return header
}
