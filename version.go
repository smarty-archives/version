package version

import (
	"fmt"
	"strings"
)

type Number struct {
	Major int
	Minor int
	Patch int
	Dirty bool
}

func New(major, minor, patch int, dirty bool) Number {
	return Number{
		Major: major,
		Minor: minor,
		Patch: patch,
		Dirty: dirty,
	}
}

func (this Number) String() string {
	return fmt.Sprintf("%d.%d.%d", this.Major, this.Minor, this.Patch)
}

func (this Number) IncrementMajor() Number {
	return Number{Major: this.Major + 1}
}

func (this Number) IncrementMinor() Number {
	return Number{Major: this.Major, Minor: this.Minor + 1}
}

func (this Number) IncrementPatch() Number {
	return Number{Major: this.Major, Minor: this.Minor, Patch: this.Patch + 1}
}

func (this Number) Increment(how string) Number {
	switch strings.ToLower(how) {
	case "major":
		return this.IncrementMajor()
	case "minor":
		return this.IncrementMinor()
	default:
		return this.IncrementPatch()
	}
}
