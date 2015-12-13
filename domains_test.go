package tempmail

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDomains(t *testing.T) {
	domains, err := Domains()

	Convey("Domains() should not return any errors", t, func() {
		So(err, ShouldBeNil)
	})

	Convey("Domains slice should not be empty", t, func() {
		So(domains, ShouldNotBeEmpty)
	})
}

func TestRandomDomain(t *testing.T) {
	Convey("RandomDomain() should return random domain", t, func() {
		domain, err := RandomDomain()

		So(err, ShouldBeNil)
		So(domain, ShouldStartWith, "@")
		So(domain, ShouldContainSubstring, ".")
	})
}
