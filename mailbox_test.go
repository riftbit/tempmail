package tempmail

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNew(t *testing.T) {
    Convey("New() should return pointer to Mailbox", t, func() {
        So(New("user@test.tld"), ShouldHaveSameTypeAs, new(Mailbox))
    })
}

func TestRandom(t *testing.T) {
    Convey("Random() should return pointer to Mailbox with random username", t, func() {
        m, err := Random()
        
        So(err, ShouldBeNil)
        So(m, ShouldHaveSameTypeAs, new(Mailbox))
    })
}

func TestMD5(t *testing.T) {
    Convey("MD5() should return valid md5", t, func() {
        encodedEmail := Mailbox("user@test.tld").MD5()
        
        So(encodedEmail, ShouldHaveLength, 32)
        // So(encodedEmail, ShouldEqual, "3d3b4ecf96c195bbd6649711825d2d63")
    })
}
