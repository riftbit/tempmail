package tempmail

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
)


type Mailbox string

// New returns pointer to the temprorary mailbox.
func New(email string) *Mailbox {
	m := Mailbox(email)
	return &m
}

func Random() (*Mailbox, error) {
	domains, err := GetDomains()
	
	if err != nil {
		return new(Mailbox), err
	}
	
	var m Mailbox
	
	for _, i := range rand.Perm(len(domains)) {
		m = Mailbox("test" + domains[i])
		break
	}
	
	return &m, nil
}

// GetEmails is is an alias for GetEmails(string)
func (m *Mailbox) GetEmails() ([]Email, error) {
	return GetEmails(m)
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
