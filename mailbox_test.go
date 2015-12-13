package tempmail

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNew(t *testing.T) {
	Convey("New() should return a pointer to Mailbox", t, func() {
		So(New("user@test.tld"), ShouldHaveSameTypeAs, new(Mailbox))
	})
}

func TestRandom(t *testing.T) {
	Convey("Random() should return a pointer to Mailbox with random username", t, func() {
		m, err := Random()

		So(err, ShouldBeNil)
		So(m, ShouldHaveSameTypeAs, new(Mailbox))
	})
}

func TestMD5(t *testing.T) {
	Convey("MD5() should return a valid md5", t, func() {
		encodedEmail := New("user@test.tld").MD5()

		So(encodedEmail, ShouldHaveLength, 32)
		// So(encodedEmail, ShouldEqual, "3d3b4ecf96c195bbd6649711825d2d63")
	})
}

func TestEmails(t *testing.T) {
	Convey("Emails() should return no emails", t, func() {
		emails, err := New("fqwfwqrwqr@mailzi.ru").Emails()

		So(emails, ShouldBeEmpty)
		So(err, ShouldNotBeNil)
	})
}
