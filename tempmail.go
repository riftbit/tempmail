package tempmail

import (
	"encoding/json"
	"errors"

	"gopkg.in/resty.v0"
)

const host = "http://api.temp-mail.ru"

// GetDomains returns slice of available domains.
func GetDomains() (domains []string, err error) {
	resp, err := resty.New().R().
		Get(host + "/request/domains/format/json")

	if err != nil {
		return
	}

	err = json.Unmarshal([]byte(resp.String()), &domains)

	return
}

// GetEmails returns slice of emails, if you have any.
func GetEmails(email MD5Encoder) (emails []Email, err error) {
	resp, err := resty.New().R().
		Get(host + "/request/mail/id/" + email.MD5() + "/format/json/")

	if err != nil {
		return
	}

	switch resp.StatusCode() {
	case 200:
		err = json.Unmarshal([]byte(resp.String()), &emails)
		return

	case 404:
		var tempMailErrMap map[string]string

		if err = json.Unmarshal([]byte(resp.String()), &tempMailErrMap); err != nil {
			return
		}

		tempMailErr, ok := tempMailErrMap["error"]

		if ok {
			err = errors.New(tempMailErr)
		}

		return

	default:
		return emails, errors.New("Unexpected status code")
	}
}
