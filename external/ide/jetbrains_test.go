package ide

import (
	"github.com/icankeep/simplego/setx"
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
	projects, err := GetJetBrainsRecentProjects(GoLand)
	a.NoError(err)
	a.NotEmpty(projects)
	projectDirs := setx.NewSet[string]()
	for _, project := range projects {
		if projectDirs.Contains(project.Dir) {
			a.FailNow("project duplicate: " + project.Dir)
		}

		projectDirs.Add(project.Dir)
	}

	projects, err = GetJetBrainsRecentProjects(IntelliJIdea)
	a.NoError(err)
	a.NotEmpty(projects)

	projects, err = GetJetBrainsRecentProjects(PyCharm)
	a.NoError(err)
	a.NotEmpty(projects)
}

func TestGetDataSpellRecentProjects(t *testing.T) {
	if _, set := os.LookupEnv("DEBUG"); !set {
		return
	}

	a := assert.New(t)
	path, err := filepath.Abs("../../utest/data/TestJetBrains/dataspell_input.xml")
	a.NoError(err)

	projects, err := GetProjectsFromXML([]string{path})
	a.NoError(err)
	a.Equal(1, len(projects))
	//a.Equal("awesomeProject", projects[0].Name)
	//a.Equal("/Users/passer/go/src/awesomeProject", projects[0].Dir)
	//a.Equal(int64(1690633942632), projects[0].LastOpenTimestamp)
	//a.Equal(int64(1690634440044), projects[0].LastActivationTimestamp)
}
