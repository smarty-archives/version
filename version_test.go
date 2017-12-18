package version

import (
	"testing"

	"github.com/smartystreets/assertions"
	"github.com/smartystreets/assertions/should"
)

func TestVersionIncrementationFixture(t *testing.T) {
	assert := assertions.New(t)
	assert.So(New(1, 2, 3, true).IncrementMajor(), should.Resemble, New(2, 0, 0, false))
	assert.So(New(1, 2, 3, true).IncrementMinor(), should.Resemble, New(1, 3, 0, false))
	assert.So(New(1, 2, 3, true).IncrementPatch(), should.Resemble, New(1, 2, 4, false))

	assert.So(New(1, 2, 3, true).Increment("mAjOr"), should.Resemble, New(2, 0, 0, false))
	assert.So(New(1, 2, 3, true).Increment("MiNoR"), should.Resemble, New(1, 3, 0, false))
	assert.So(New(1, 2, 3, true).Increment("PATCH"), should.Resemble, New(1, 2, 4, false))
	assert.So(New(1, 2, 3, true).Increment(""), should.Resemble, New(1, 2, 4, false))
}

func TestVersionString(t *testing.T) {
	assert := assertions.New(t)
	assert.So(New(1, 2, 3, false).String(), should.Equal, "1.2.3")
	assert.So(New(1, 2, 3, true).String(), should.Equal, "1.2.3")
}
