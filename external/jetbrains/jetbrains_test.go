package jetbrains

import (
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

func TestGetProjectsFromXML(t *testing.T) {
	a := assert.New(t)

	path, err := filepath.Abs("../../utest/data/TestJetBrains/1_input.xml")
	a.NoError(err)

	projects, err := GetProjectsFromXML([]string{path})
	a.NoError(err)
	a.Equal(12, len(projects))
	a.Equal("awesomeProject", projects[0].Name)
	a.Equal("/Users/passer/go/src/awesomeProject", projects[0].Dir)
	a.Equal(int64(1690633942632), projects[0].LastOpenTimestamp)
	a.Equal(int64(1690634440044), projects[0].LastActivationTimestamp)
}

func TestGetRecentProjects(t *testing.T) {
	if _, set := os.LookupEnv("DEBUG"); !set {
		return
	}

	a := assert.New(t)
	projects, err := GetRecentProjects(GoLand)
	a.NoError(err)
	a.NotEmpty(projects)

	projects, err = GetRecentProjects(Idea)
	a.NoError(err)
	a.NotEmpty(projects)

	projects, err = GetRecentProjects(PyCharm)
	a.NoError(err)
	a.NotEmpty(projects)
}
