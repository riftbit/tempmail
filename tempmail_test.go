package tempmail

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGetEmails(t *testing.T) {
	Convey("GetEmails() should return no emails", t, func() {
		emails, err := GetEmails(Mailbox("fqwfwqrwqr@mailzi.ru"))

		So(emails, ShouldBeEmpty)
		So(err, ShouldNotBeNil)
	})

	// 	Convey("GetEmails() should return non-empty []Email", t, func() {
	// 		emails, err := GetEmails("mojob@mailzi.ru")

	// 		So(err, ShouldBeNil)
	// 		So(emails, ShouldNotBeEmpty)
	// 	})
}
