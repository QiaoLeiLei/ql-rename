package main

// SetRules 设置重命名规则
func (a *App) SetRules(rules *Rules) {
	a.rules = rules
	a.updatePreview()
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
		a.preview[i].newName = a.getNewFileName(a.preview[i].oldName)
	}
	a.fileIndex = 0
}
