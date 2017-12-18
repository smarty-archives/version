package git

import (
	"errors"
	"os/exec"
	"strconv"
	"strings"

	"github.com/smartystreets/assertions/assert"
	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/version"
)

type Repository struct{}

func (this *Repository) CurrentVersion() (version.Number, error) {
	output, err := exec.Command("git", "describe").CombinedOutput()
	if err != nil {
		return parseGitDescribe(string(output) + " " + string(err.Error()))
	}
	return parseGitDescribe(string(output))
}

func parseGitDescribe(raw string) (number version.Number, err error) {
	raw = strings.TrimSpace(raw)
	if strings.HasPrefix(raw, "fatal: No names found, cannot describe anything.") {
		number.Dirty = true
		return number, nil
	}
	fields := strings.Split(raw, "-")
	number.Dirty = len(fields) > 1

	parts := strings.Split(fields[0], ".")
	if len(parts) < 3 {
		return version.Number{}, errors.New("At least 3 version fields are required (major.minor.patch).")
	}
	number.Major, err = strconv.Atoi(parts[0])
	if err != nil {
		return number, err
	}
	number.Minor, err = strconv.Atoi(parts[1])
	if err != nil {
		return number, err
	}
	number.Patch, err = strconv.Atoi(parts[2])
	if err != nil {
		return number, err
	}
	return number, nil
}

func (this *Repository) UpdateVersion(version version.Number) error {
	_, err := exec.Command("git", "tag", "-a", version.String(), "-m", "''").CombinedOutput()
	if err != nil {
		return err
	}

	current, err := this.CurrentVersion()
	if err != nil {
		return err
	}

	return assert.So(current, should.Resemble, version).Error()
}
