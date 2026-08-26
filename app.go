package main

import (
	"context"
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"os"
	"sync"
)

type (
	RenameType int //重命名类型
	NumberType int //数字序号类型
)

const (
	AddPrefix          RenameType = iota + 1 // 添加前缀
	AddSuffix                                // 添加后缀
	ReplaceStr                               // 替换字符串
	AddNumber                                // 添加数字序号
	ToUpperCase                              // 转换为大写
	ToLowerCase                              // 转换为小写
	DeleteSpecialChars                       // 删除特殊字符
)

const (
	bracket   NumberType = iota + 1 // 括号
	underLine                       // 下划线
)

type ReplaceObj struct {
	OldStr string
	NewStr string
}

type NumberObj struct {
	NewName string
	Suffix  NumberType
}

type Rules struct {
	RenameType RenameType
	Prefix     string
	Suffix     string
	ReplaceObj *ReplaceObj
	NumberObj  *NumberObj
}
type RenamePreview struct {
	OldDisPlayName string
	NewDisPlayName string
	OldName        string
	NewName        string
	Selected       bool
}

// App struct
type App struct {
	ctx       context.Context
	files     []string        // 原始文件路径
	rules     *Rules          // 重命名规则
	fileIndex int             // 文件序号
	preview   []RenamePreview // 重命名预览
	wg        *sync.WaitGroup // 并发重命名等待组
}

// NewApp creates a new App application struct

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	runtime.LogInfo(ctx, "App startup")
	runtime.OnFileDrop(ctx, func(x, y int, paths []string) {
		for _, p := range paths {
			runtime.LogInfo(ctx, p)
			stat, err := os.Stat(p)
			if err != nil {
				continue
			}
			if stat.IsDir() {
				a.setFiles(p)
			} else {
				a.files = append(a.files, p)
			}
		}
		a.setPreview()
	})
}

func (a *App) shutdown(ctx context.Context) {
	runtime.OnFileDropOff(ctx)
}

func NewApp() *App {
	return &App{
		rules: &Rules{
			RenameType: ToUpperCase,
			Prefix:     "",
			Suffix:     "",
			ReplaceObj: &ReplaceObj{
				OldStr: "",
				NewStr: "",
			},
			NumberObj: &NumberObj{
				NewName: "新名字",
				Suffix:  bracket,
			},
		},
		wg: &sync.WaitGroup{},
	}
}

func (a *App) GetPreview() []RenamePreview {
	return a.preview
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}
