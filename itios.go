package itios

import (
	"github.com/mssola/user_agent"
	"net/http"
	"strings"
)

// OSMatcher Store the schema to match with the request Path
type OSMatcher struct {
	osList []string
}

// New is the constructor to ItiOS
func New(os ...string) *OSMatcher {
	for i := range os {
		os[i] = strings.ToLower(os[i])
	}
	return &OSMatcher{
		osList: os,
	}
}

// Match returns if the request can be handled by this Route.
func (b *OSMatcher) Match(req *http.Request) bool {
	ua := user_agent.New(req.UserAgent())
	name := strings.ToLower(ua.OSInfo().Name)
	for _, v := range b.osList {
		if name == v {
			return true
		}
	}

	return false
}
