package ide

import (
	"fmt"
	"github.com/icankeep/simplego/jsonx"
	"github.com/icankeep/simplego/logx"
	"github.com/icankeep/simplego/setx"
	"github.com/icankeep/simplego/utils"
	"net/url"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

type WorkspaceInfo struct {
	Folder string `json:"folder"`
}

// GetVsCodeBasePath https://code.visualstudio.com/docs/getstarted/settings#_settings-file-locations
func GetVsCodeBasePath() string {
	switch runtime.GOOS {
	case "darwin":
		return os.ExpandEnv("$HOME/Library/Application Support/Code/User/workspaceStorage")
	case "windows":
		return os.ExpandEnv("$APPDATA\\Code\\User\\workspaceStorage")
	case "linux":
		return os.ExpandEnv("$HOME/.config/Code/User/workspaceStorage")
	default:
		return ""
	}
}

func GetVsCodeRecentProjects() ([]*Project, error) {
	// 1. 获取workspace目录
	basePath := GetVsCodeBasePath()
	if exist, err := utils.PathExists(basePath); err != nil {
		return nil, err
	} else if !exist {
		return make([]*Project, 0), nil
	}

	dirs, err := os.ReadDir(basePath)
	if err != nil {
		return nil, fmt.Errorf("list dirs failed: %v", err)
	}

	// 2. 遍历base path下所有workspace, 解析出json
	projects := make([]*Project, 0)
	projectDirs := setx.NewSet[string]()
	for _, dir := range dirs {
		workspaceFilePath := filepath.Join(basePath, dir.Name(), "workspace.json")
		if exist, err := utils.PathExists(workspaceFilePath); err != nil || !exist {
			logx.Debug("path err or not exist, path: %v, err: %v", workspaceFilePath, err)
			continue
		}
		content, err := utils.ReadStringFromFile(workspaceFilePath)
		if err != nil {
			logx.Debug("failed to read workspace, path: %v, err: %v", workspaceFilePath, err)
			continue
		}

		workspace := &WorkspaceInfo{}
		err = jsonx.JSONToAny(content, workspace)
		if err != nil {
			logx.Debug("failed to read workspace, path: %v, err: %v", workspaceFilePath, err)
			continue
		}
		if !strings.HasPrefix(workspace.Folder, "file://") {
			continue
		}
		workspace.Folder = strings.TrimPrefix(workspace.Folder, "file://")
		if exist, err := utils.PathExists(workspace.Folder); err != nil || !exist {
			logx.Debug("project path err or not exist, path: %v, err: %v", workspaceFilePath, err)
		}

		unEscapeFolder, err := url.QueryUnescape(workspace.Folder)
		if err != nil {
			logx.Debug("failed to unescape workspace, path: %v, err: %v", workspaceFilePath, err)
		} else {
			workspace.Folder = unEscapeFolder
		}
		if projectDirs.Contains(workspace.Folder) {
			continue
		}
		projectDirs.Add(workspace.Folder)
		projects = append(projects, &Project{Dir: workspace.Folder, Name: filepath.Base(workspace.Folder)})
	}

	return projects, nil
}
