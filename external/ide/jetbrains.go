package ide

import (
	"fmt"
	"github.com/beevik/etree"
	"github.com/icankeep/simplego/conv"
	"github.com/icankeep/simplego/setx"
	"github.com/icankeep/simplego/utils"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

var (
	XMLFiles = []string{"recentProjects.xml", "recentSolutions.xml"}
)

// GetJetBrainsBasePath https://www.jetbrains.com/help/idea/directories-used-by-the-ide-to-store-settings-caches-plugins-and-logs.html#config-directory
func GetJetBrainsBasePath() string {
	switch runtime.GOOS {
	case "darwin":
		return os.ExpandEnv("$HOME/Library/Application Support/JetBrains")
	case "windows":
		return os.ExpandEnv("$APPDATA\\JetBrains")
	case "linux":
		return os.ExpandEnv("$HOME/.config/JetBrains")
	default:
		return ""
	}
}

func GetJetBrainsRecentProjects(jetBrainsType IDEType) ([]*Project, error) {
	// 1. Get path
	basePath := GetJetBrainsBasePath()
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

func GetAllXMLFiles(jetBrainsType IDEType, dirs []os.DirEntry) []string {
	walkFiles := make([]string, 0)
	for _, dir := range dirs {
		if !dir.IsDir() || !strings.HasPrefix(dir.Name(), string(jetBrainsType)) {
			continue
		}

		for _, file := range XMLFiles {
			recentProjectXmlFilePath := filepath.Join(GetJetBrainsBasePath(), dir.Name(), "options", file)
			if exist, err := utils.PathExists(recentProjectXmlFilePath); err != nil || !exist {
				continue
			}
			walkFiles = append(walkFiles, recentProjectXmlFilePath)
		}
	}
	return walkFiles
}

func GetProjectsFromXML(walkFiles []string) ([]*Project, error) {
	projectDirs := setx.NewSet[string]()
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
		projectsEle := rootEle.FindElement("./component/option[@name='additionalInfo']/map")
		if projectsEle == nil {
			continue
		}
		configDir, _, _ := strings.Cut(xmlFile, "/options/")
		for _, ele := range projectsEle.ChildElements() {
			// TODO: check for windows
			projectPath := strings.Replace(GetElementAttr(ele.Attr, "key"), "$USER_HOME$", "$HOME", 1)
			projectPath = strings.Replace(projectPath, "$APPLICATION_CONFIG_DIR$", configDir, 1)
			projectPath = os.ExpandEnv(projectPath)
			if len(projectPath) == 0 || projectDirs.Contains(projectPath) {
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
			projectDirs.Add(projectPath)
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
