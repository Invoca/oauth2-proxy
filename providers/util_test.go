package providers

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"
)

func TestGetAuhtorizationHeader(t *testing.T) {
	testCases := []struct {
		name      string
		prefix    string
		token     string
		accept    string
		includeLi bool
	}{
		{
			name:      "With an empty prefix, token, accept and includeLi false",
			prefix:    "",
			token:     "",
			accept:    "",
			includeLi: false,
		},
		{
			name:      "With a Bearer token type",
			prefix:    tokenTypeBearer,
			token:     "abcdef",
			accept:    "",
			includeLi: false,
		},
		{
			name:      "With a Token token type",
			prefix:    tokenTypeToken,
			token:     "123456",
			accept:    "",
			includeLi: false,
		},
		{
			name:      "With a Bearer token type and Accept application/json",
			prefix:    tokenTypeToken,
			token:     "abc",
			accept:    acceptApplicationJSON,
			includeLi: false,
		},
		{
			name:      "With a Bearer token type and Accept Github+JSON",
			prefix:    tokenTypeToken,
			token:     "abc123",
			accept:    acceptApplicationGitHubJSON,
			includeLi: false,
		},
		{
			name:      "With a Bearer token type and Include LI",
			prefix:    tokenTypeToken,
			token:     "123",
			accept:    "",
			includeLi: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			g := NewWithT(t)

			header := getAuthorizationHeader(tc.prefix, tc.token, tc.accept, tc.includeLi)
			g.Expect(header.Get("Authorization")).To(Equal(fmt.Sprintf("%s %s", tc.prefix, tc.token)))
			g.Expect(header.Get("Accept")).To(Equal(tc.accept))
			if tc.includeLi {
				g.Expect(header.Get("x-li-format")).To(Equal("json"))
			} else {
				g.Expect(header.Get("x-li-format")).To(Equal(""))
			}
		})
	}
}
