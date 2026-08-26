package main

import (
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"os"
	"path"
	"ql-rename/backend"
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

func (a *App) setFiles(dirPath string) {
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
	a.fileIndex = 0
	for i := range a.files {
		filePath := a.files[i]
		onePreview := RenamePreview{
			OldDisPlayName: path.Base(filePath),
			OldName:        filePath,
			NewName:        a.getNewFileName(filePath),
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
