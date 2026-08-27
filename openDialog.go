package main

import (
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"ql-rename/backend"
	goruntime "runtime"
)

// OpenMultipleFilesDialog 显示打开文件对话框，返回选择的文件路径
func (a *App) OpenMultipleFilesDialog() {
	filesPath, err := runtime.OpenMultipleFilesDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "导入文件",
	})
	if err != nil {
		runtime.LogError(a.ctx, err.Error())
		return
	}
	runtime.LogInfo(a.ctx, fmt.Sprintf("OpenMultipleFilesDialog: %v", filesPath))
	a.files = []string{}
	a.files = append(a.files, filesPath...)
	a.setPreview()
}

// OpenDirectoryDialog 显示打开目录对话框，返回选择的目录路径
func (a *App) OpenDirectoryDialog() {
	dirPath, err := runtime.OpenDirectoryDialog(a.ctx, runtime.OpenDialogOptions{})
	if err != nil {
		runtime.LogError(a.ctx, err.Error())
		return
	}
	runtime.LogInfo(a.ctx, fmt.Sprintf("OpenDirectoryDialog: %v", dirPath))
	a.setFiles(dirPath)
	a.setPreview()
}

func (a *App) showDialog(s string, errors []error) {
	selected, _ := runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
		Title:         s,
		Message:       fmt.Sprintf("共重命名 %d 个文件", len(a.files)),
		Buttons:       []string{"取消", "打开文件夹"},
		DefaultButton: "打开文件夹",
	})
	if selected == "打开文件夹" {
		a.OpenInFinder(a.files[0])
		return
	}
}

// OpenInFinder 在系统文件管理器中打开并选中指定文件
// 支持 macOS（Finder）、Windows（资源管理器）、Linux（xdg-open 打开所在目录）
func (a *App) OpenInFinder(filePath string) {
	runtime.LogInfo(a.ctx, fmt.Sprintf("Opening in file manager: %s", filePath))

	var cmd *exec.Cmd
	switch goruntime.GOOS {
	case "darwin":
		cmd = exec.Command("open", "-R", filePath)
	case "windows":
		cmd = exec.Command("explorer", "/select,", filePath)
	default:
		dir := filepath.Dir(filePath)
		cmd = exec.Command("xdg-open", dir)
	}

	if err := cmd.Run(); err != nil {
		runtime.LogError(a.ctx, fmt.Sprintf("Failed to open file manager: %v", err))
	}
}

func (a *App) setFiles(dirPath string) {
	a.files = []string{}
	files, err := os.ReadDir(dirPath)
	if err != nil {
		runtime.LogError(a.ctx, err.Error())
		return
	}
	for _, f := range files {
		if f.IsDir() {
			continue
		}

		fileName := path.Join(dirPath, f.Name())
		a.files = append(a.files, fileName)
	}
}

func (a *App) setPreview() {
	a.preview = []RenamePreview{}
	a.fileIndex = 0
	for i := range a.files {
		filePath := a.files[i]
		onePreview := RenamePreview{
			OldDisPlayName: path.Base(filePath),
			OldName:        filePath,
			NewName:        a.getNewFileName(filePath),
			Selected:       true,
		}
		onePreview.NewDisPlayName = path.Base(onePreview.NewName)
		a.preview = append(a.preview, onePreview)
	}
	a.fileIndex = 0
	a.notifyPreviewUpdate()
}

func (a *App) notifyPreviewUpdate() {
	runtime.EventsEmit(a.ctx, backend.EventDataUpdate)
}
