package git

import (
	"testing"

	"github.com/smartystreets/assertions"
	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/version"
)

func TestParseGitVersion(t *testing.T) {
	assert := assertions.New(t)
	var parsed version.Number
	var err error

	parsed, err = parseGitDescribe("")
	assert.So(err, should.NotBeNil)
	assert.So(parsed, should.Resemble, version.New(0, 0, 0, false))

	parsed, err = parseGitDescribe("1")
	assert.So(err, should.NotBeNil)
	assert.So(parsed, should.Resemble, version.New(0, 0, 0, false))

	parsed, err = parseGitDescribe("1.2")
	assert.So(err, should.NotBeNil)
	assert.So(parsed, should.Resemble, version.New(0, 0, 0, false))

	parsed, err = parseGitDescribe("fatal: No names found, cannot describe anything.")
	assert.So(err, should.BeNil)
	assert.So(parsed, should.Resemble, version.New(0, 0, 0, true))

	parsed, err = parseGitDescribe("1.a.0")
	assert.So(err, should.NotBeNil)
	assert.So(parsed, should.Resemble, version.New(1, 0, 0, false))

	parsed, err = parseGitDescribe("1.2.0\n")
	assert.So(err, should.BeNil)
	assert.So(parsed, should.Resemble, version.New(1, 2, 0, false))

	parsed, err = parseGitDescribe("1.2.0-4-g3201d7a")
	assert.So(err, should.BeNil)
	assert.So(parsed, should.Resemble, version.New(1, 2, 0, true))
}
