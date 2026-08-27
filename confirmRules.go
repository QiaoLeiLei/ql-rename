package main

import "path"

// SetRules 设置重命名规则
func (a *App) SetRules(rules *Rules) {
	if rules == nil {
		return
	}
	a.rules = rules
	a.updatePreview()
}

func (a *App) GetRules() *Rules {
	return a.rules
}

// ExecsRename 执行重命名
func (a *App) ExecsRename() {
	err := a.renameFiles()
	if err != nil {
		return
	}
}

// resetPreview 更新预览
func (a *App) updatePreview() {
	if a.preview == nil {
		return
	}
	a.fileIndex = 0
	for i := range a.preview {
		a.preview[i].NewName = a.getNewFileName(a.preview[i].OldName)
		a.preview[i].NewDisPlayName = path.Base(a.preview[i].NewName)
	}
	a.fileIndex = 0
	a.notifyPreviewUpdate()
}
