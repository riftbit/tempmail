package tempmail

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGetDomains(t *testing.T) {
	domains, err := GetDomains()

	Convey("GetDomains() should not return any errors", t, func() {
		So(err, ShouldBeNil)
	})

	Convey("Domains slice should not be empty", t, func() {
		So(domains, ShouldNotBeEmpty)
	})
}

func TestGetEmails(t *testing.T) {
	Convey("GetEmails() should return no emails", t, func() {
		emails, err := GetEmails(TemporaryUser{"fqwfwqrwqr", "@mailzi.ru"})

		So(emails, ShouldBeEmpty)
		So(err, ShouldNotBeNil)
	})

// 	Convey("GetEmails() should return non-empty []Email", t, func() {
// 		emails, err := GetEmails("mojob@mailzi.ru")

// 		So(err, ShouldBeNil)
// 		So(emails, ShouldNotBeEmpty)
// 	})
}
