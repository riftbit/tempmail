package tempmail

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDelete(t *testing.T) {
    Convey("Delete() shoud not return any errors", t, func() {
        So(Email{}.Delete(), ShouldBeNil)
    })
}
