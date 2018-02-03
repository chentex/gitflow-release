package version

import (
	"fmt"
	"regexp"
	"strconv"

	fm "github.com/chentex/go-fm"
	"github.com/pkg/errors"
)

//Versioner defines what a version control should do
type Versioner interface {
	BumpVersion(versionFile string, bumpType string, alpha bool, beta bool) error
}

const (
	patchBump = "patch"
	minorBump = "minor"
	majorBump = "major"
)

//NewVersioner returns an implementation of a Versioner
func NewVersioner() Versioner {
	return &Manager{}
}

//Manager implementation for a Versioner
type Manager struct {
}

//BumpVersion modifies the file with the bumped version
//Examples:
//File contains: 0.1.0
//major bump will result in file containing: 1.1.0
//minor bump will result in file containing: 0.2.0
//patch bump will result in file containing: 0.1.1
//sending alpha true bump will result in file containing: 0.1.1-alpha
//sending beta true bump will result in file containing: 0.1.1-beta
//alpha and beta cannot be sent both true
func (m *Manager) BumpVersion(versionFile string, bumpType string, alpha bool, beta bool) error {
	if alpha && beta {
		return errors.New("Cannot have both alpha and beta")
	}
	f := fm.NewFileManager()
	content, err := f.OpenFile(versionFile)
	if err != nil {
		return errors.Wrap(err, "while opening version file")
	}
	re := regexp.MustCompile("^([0-9]+).([0-9]+).([0-9]+)(.+)?")
	matches := re.FindStringSubmatch(content)

	var version string
	switch bumpType {
	case majorBump:
		major, err := strconv.Atoi(matches[1])
		if err != nil {
			return errors.Wrap(err, "while parsing major version")
		}
		major++
		version = fmt.Sprintf("%d.0.0", major)
		break
	case minorBump:
		minor, err := strconv.Atoi(matches[2])
		if err != nil {
			return errors.Wrap(err, "while parsing minor version")
		}
		minor++
		version = fmt.Sprintf("%s.%d.0", matches[1], minor)
		break
	case patchBump:
		patch, err := strconv.Atoi(matches[3])
		if err != nil {
			return errors.Wrap(err, "while parsing patch version")
		}
		patch++
		version = fmt.Sprintf("%s.%s.%d", matches[1], matches[2], patch)
		break
	default:
		return errors.New("Incorrect bump type")
	}

	if alpha {
		version = version + "-alpha"
	}

	if beta {
		version = version + "-beta"
	}

	f.WriteFile(versionFile, []byte(version), 0644)
	return nil
}
