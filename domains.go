package tempmail

import (
	"encoding/json"
	"math/rand"

	"gopkg.in/resty.v0"
)

// Domains returns slice of available domains.
func Domains() (domains []string, err error) {
	resp, err := resty.New().R().
		Get(host + "/request/domains/format/json")

	if err != nil {
		return
	}

	// TODO: error check. Sometimes 'unexpected end of JSON input' could be returned, even though json is valid.
	json.Unmarshal([]byte(resp.String()), &domains)

	return
}

// RandomDomain returns a random domain.
func RandomDomain() (domain string, err error) {
	domains, err := Domains()

	if err != nil {
		return
	}

	for _, i := range rand.Perm(len(domains)) {
		domain = domains[i]
		break
	}

	return
}
