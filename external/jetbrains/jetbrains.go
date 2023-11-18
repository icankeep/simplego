package jetbrains

import (
	"fmt"
	"github.com/beevik/etree"
	"github.com/icankeep/simplego/conv"
	"github.com/icankeep/simplego/utils"
	"log"
	"os"
	"path"
	"strings"
)

type JetBrainsType string

const (
	GoLand  JetBrainsType = "GoLand"
	Idea    JetBrainsType = "IntelliJIdea"
	PyCharm JetBrainsType = "PyCharm"

	BasePath = "$HOME/Library/Application Support/JetBrains"
)

type Project struct {
	Name                    string
	Dir                     string
	LastOpenTimestamp       int64
	LastActivationTimestamp int64
}

func GetBasePath() string {
	return os.ExpandEnv(BasePath)
}

func GetRecentProjects(jetBrainsType JetBrainsType) ([]*Project, error) {
	// 1. Get path
	basePath := GetBasePath()
	if exist, err := utils.PathExists(basePath); err != nil {
		return nil, err
	} else if !exist {
		return make([]*Project, 0), nil
	}

	dirs, err := os.ReadDir(basePath)
	if err != nil {
		return nil, fmt.Errorf("list dirs failed: %v", err)
	}

	walkFiles := GetAllXMLFiles(jetBrainsType, dirs)
	log.Printf("walkFiles: %v\n", walkFiles)

	// 2. read xml
	projects, err := GetProjectsFromXML(walkFiles)
	if err != nil {
		return nil, err
	}
	return projects, nil
}

func GetAllXMLFiles(jetBrainsType JetBrainsType, dirs []os.DirEntry) []string {
	walkFiles := make([]string, 0)
	for _, dir := range dirs {
		if !dir.IsDir() || !strings.HasPrefix(dir.Name(), string(jetBrainsType)) {
			continue
		}
		recentProjectXmlFilePath := fmt.Sprintf("%s/%s/options/recentProjects.xml", GetBasePath(), dir.Name())

		if exist, err := utils.PathExists(recentProjectXmlFilePath); err != nil || !exist {
			continue
		}
		walkFiles = append(walkFiles, recentProjectXmlFilePath)
	}
	return walkFiles
}

func GetProjectsFromXML(walkFiles []string) ([]*Project, error) {
	projects := make([]*Project, 0)
	for _, xmlFile := range walkFiles {
		doc := etree.NewDocument()
		if err := doc.ReadFromFile(xmlFile); err != nil {
			return nil, err
		}
		rootEle := doc.Root()
		if rootEle == nil {
			return nil, fmt.Errorf("root element is nil")
		}
		projectsEle := rootEle.FindElement("./component[@name='RecentProjectsManager']/option[@name='additionalInfo']/map")
		if projectsEle == nil {
			continue
		}
		for _, ele := range projectsEle.ChildElements() {
			projectPath := strings.Replace(GetElementAttr(ele.Attr, "key"), "$USER_HOME$", "$HOME", 1)
			projectPath = os.ExpandEnv(projectPath)
			if len(projectPath) == 0 {
				continue
			}
			if exist, err := utils.PathExists(projectPath); err != nil || !exist {
				continue
			}
			project := &Project{
				Name: path.Base(projectPath),
				Dir:  projectPath,
			}
			FindAndSetProjectOpenTimestamp(ele, project)
			FindAndSetProjectActivationTimestamp(ele, project)
			projects = append(projects, project)
		}
	}
	return projects, nil
}

func FindAndSetProjectActivationTimestamp(ele *etree.Element, project *Project) {
	projectActivationEle := ele.FindElement("./value/RecentProjectMetaInfo/option[@name='activationTimestamp']")
	if projectActivationEle != nil {
		project.LastActivationTimestamp = conv.Int64Default(GetElementAttr(projectActivationEle.Attr, "value"), 0)
	}
}

func FindAndSetProjectOpenTimestamp(ele *etree.Element, project *Project) {
	projectOpenEle := ele.FindElement("./value/RecentProjectMetaInfo/option[@name='projectOpenTimestamp']")
	if projectOpenEle != nil {
		project.LastOpenTimestamp = conv.Int64Default(GetElementAttr(projectOpenEle.Attr, "value"), 0)
	}
}

func GetElementAttr(attr []etree.Attr, name string) string {
	for _, a := range attr {
		if a.Key == name {
			return a.Value
		}
	}
	return ""
}
