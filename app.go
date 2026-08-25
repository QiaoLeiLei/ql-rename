package main

import (
	"context"
	"fmt"
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
	renameType RenameType
	prefix     string
	suffix     string
	replaceObj *ReplaceObj
	numberObj  *NumberObj
}
type RenamePreview struct {
	oldName string
	newName string
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
}

func NewApp() *App {
	return &App{
		rules: &Rules{
			renameType: ToUpperCase,
			prefix:     "",
			suffix:     "",
			replaceObj: &ReplaceObj{
				OldStr: "",
				NewStr: "",
			},
			numberObj: &NumberObj{
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
