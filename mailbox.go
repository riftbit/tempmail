package tempmail

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"

	"gopkg.in/resty.v0"
)

// Mailbox represents temprorary email address.
type Mailbox string

// New returns pointer to the temprorary mailbox.
func New(email string) *Mailbox {
	m := Mailbox(email)
	return &m
}

// Random retuns randomly generated temporary mailbox.
func Random() (*Mailbox, error) {
	var m Mailbox

	domain, err := RandomDomain()

	if err != nil {
		return new(Mailbox), err
	}

	m = Mailbox("test" + domain)

	return &m, nil
}

// Emails returns slice of emails, if you have any.
func (m *Mailbox) Emails() (emails []Email, err error) {
	resp, err := resty.New().R().
		Get(host + "/request/mail/id/" + m.MD5() + "/format/json/")

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

// MD5 returns md5-encoded email of current TemporaryUser.
func (m Mailbox) MD5() string {
	hash := md5.Sum([]byte(m))
	return hex.EncodeToString(hash[:])
}

// MD5Encoder provives a function that return md5-encoded email address.
type MD5Encoder interface {
	MD5() string
}
