package ide

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetVsCodeRecentProjects(t *testing.T) {
	a := assert.New(t)

	projects, err := GetVsCodeRecentProjects()
	a.NoError(err)

	a.NotEmpty(projects)
}
