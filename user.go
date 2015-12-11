package tempmail

import (
	"crypto/md5"
	"encoding/hex"
)

// TemporaryUser represents a temporary email data.
type TemporaryUser struct {
	Username string
	Domain string
}

// NewRandomUser returns TemporaryUser with randomly generated username.
func NewRandomUser(domain string) TemporaryUser {
	return TemporaryUser{"test", domain}
}

// MD5 returns md5-encoded email of current TemporaryUser.
func (t TemporaryUser) MD5() string {
	hash := md5.Sum([]byte(t.Username + t.Domain))
	return hex.EncodeToString(hash[:])
}


// MD5Encoder provives a function that return md5-encoded email address.
type MD5Encoder interface {
	MD5() string
}
