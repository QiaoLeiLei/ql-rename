package main

import (
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
	goruntime "runtime"
)

func (a *App) getMenu() *menu.Menu {
	AppMenu := menu.NewMenu()
	if goruntime.GOOS == "darwin" {
		AppMenu.Append(menu.AppMenu())
	}
	FileMenu := AppMenu.AddSubmenu("文件")
	FileMenu.AddText("打开...", keys.CmdOrCtrl("o"), func(_ *menu.CallbackData) {
		a.OpenMultipleFilesDialog()
	})
	FileMenu.AddSeparator()
	FileMenu.AddText("打开文件夹", keys.CmdOrCtrl("f"), func(_ *menu.CallbackData) {
		a.OpenDirectoryDialog()
	})
	//FileMenu.AddText("退出", keys.CmdOrCtrl("q"), func(_ *menu.CallbackData) {
	//	runtime.Quit(a.ctx)
	//})
	return AppMenu
}
