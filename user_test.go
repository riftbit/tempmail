package tempmail

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewRandomUser(t *testing.T) {
    Convey("NewRandomUser() should return TemporaryUser with random username", t, func() {
        So(NewRandomUser("@vkcode.ru"), ShouldHaveSameTypeAs, TemporaryUser{})
    })
}

func TestMD5(t *testing.T) {
    Convey("MD5() should return valid md5", t, func() {
        So(TemporaryUser{"user", "@test.tld"}.MD5(), ShouldEqual, "3d3b4ecf96c195bbd6649711825d2d63")
    })
}
