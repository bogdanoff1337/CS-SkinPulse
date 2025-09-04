package steam

import (
	"regexp"
	"strings"
)

var (
	steamID64Re = regexp.MustCompile(`^https?://steamcommunity\.com/profiles/(\d{17})/?$`)
	vanityRe    = regexp.MustCompile(`^https?://steamcommunity\.com/id/([a-zA-Z0-9_-]{2,32})/?$`)
)

func IsValidSteamURL(url string) bool {
	url = strings.TrimSpace(url)
	return steamID64Re.MatchString(url) || vanityRe.MatchString(url)
}

func ExtractSteamID64(url string) (string, bool) {
	m := steamID64Re.FindStringSubmatch(url)
	if len(m) == 2 {
		return m[1], true
	}
	return "", false
}
