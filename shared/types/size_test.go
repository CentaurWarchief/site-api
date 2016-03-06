package types_test

import (
	"math/rand"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/txgruppi/site-api/shared/types"
)

func TestSize(t *testing.T) {
	Convey("Size", t, func() {
		Convey("NewSize", func() {
			Convey("it should return a Size struct with the correct values", func() {
				w := rand.Int()
				h := rand.Int()
				s := types.NewSize(w, h)
				So(s, ShouldHaveSameTypeAs, &types.Size{})
				So(s.Width, ShouldEqual, w)
				So(s.Height, ShouldEqual, h)
			})
		})
	})
}
