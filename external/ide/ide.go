package ide

type Project struct {
	Name                    string
	Dir                     string
	LastOpenTimestamp       int64
	LastActivationTimestamp int64
}

type IDEType string

const (
	GoLand       IDEType = "GoLand"
	IntelliJIdea IDEType = "IntelliJIdea"
	PyCharm      IDEType = "PyCharm"
	CLion        IDEType = "CLion"
	WebStorm     IDEType = "WebStorm"
	PhpStorm     IDEType = "PhpStorm"
	RustRover    IDEType = "RustRover"
	RubyMine     IDEType = "RubyMine"
	Rider        IDEType = "Rider"
	DataGrip     IDEType = "DataGrip"
	Aqua         IDEType = "Aqua"
	Fleet        IDEType = "Fleet"
	DataSpell    IDEType = "DataSpell"
	AppCode      IDEType = "AppCode"
	Writerside   IDEType = "Writerside"

	VsCode IDEType = "VsCode"
)

func GetRecentProjects(ideTypeStr string) ([]*Project, error) {
	ideType := IDEType(ideTypeStr)
	switch ideType {
	case VsCode:
		return GetVsCodeRecentProjects()
	default:
		return GetJetBrainsRecentProjects(ideType)
	}
}
