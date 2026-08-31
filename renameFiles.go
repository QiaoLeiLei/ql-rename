package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func (a *App) renameFiles() error {
	var errors []error
	for _, p := range a.preview {
		if err := os.Rename(p.OldName, p.NewName); err != nil {
			errors = append(errors, fmt.Errorf("重命名文件 %s -> %s 失败: %w", p.OldName, p.NewName, err))
		}
	}
	a.resetData()
	a.showDialog("重命名完成", errors)
	if len(errors) > 0 {
		return fmt.Errorf("文件重命名过程中发生 %d 个错误: %v", len(errors), errors)
	}
	return nil
}

func (a *App) resetData() {
	a.files = []string{}
	for _, preview := range a.preview {
		a.files = append(a.files, preview.NewName)
	}
	a.setPreview()
}

func (a *App) addFilePrefix(filePath string, prefix string) string {
	dir, file := filepath.Split(filePath)
	newPath := filepath.Join(dir, prefix+file)
	return newPath
}

func (a *App) addFileSuffix(filePath string, suffix string) string {
	dir, file := filepath.Split(filePath)
	ext := filepath.Ext(file)
	file = strings.TrimSuffix(file, ext)
	newPath := filepath.Join(dir, file+suffix+ext)
	return newPath
}

func (a *App) replaceFileStr(filePath string, obj *ReplaceObj) string {
	if obj == nil {
		return filePath
	}
	dir, file := filepath.Split(filePath)
	ext := filepath.Ext(file)
	file = strings.TrimSuffix(file, ext)
	file = strings.ReplaceAll(file, obj.OldStr, obj.NewStr)
	newPath := filepath.Join(dir, file+ext)
	return newPath
}

func (a *App) addFileNumber(filePath string, numberObj *NumberObj) string {
	dir, file := filepath.Split(filePath)
	ext := filepath.Ext(file)
	suffix := ""
	if a.fileIndex != 0 {
		if numberObj.Suffix == bracket {
			suffix = fmt.Sprintf("%s%d%s", "(", a.fileIndex, ")")
		}
		if numberObj.Suffix == underLine {
			suffix = fmt.Sprintf("%s%d", "_", a.fileIndex)
		}
	}
	a.fileIndex++
	newPath := filepath.Join(dir, numberObj.NewName+suffix+ext)
	return newPath
}

func (a *App) changeUpperLower(filePath string) string {
	if a.rules.ToUpperCase {
		return a.fileToUpperCase(filePath)
	}
	return a.fileToLowerCase(filePath)
}

func (a *App) fileToUpperCase(filePath string) string {
	dir, file := filepath.Split(filePath)
	ext := filepath.Ext(file)
	file = strings.TrimSuffix(file, ext)
	file = strings.ToUpper(file)
	newPath := filepath.Join(dir, file+ext)
	return newPath
}

func (a *App) fileToLowerCase(filePath string) string {
	dir, file := filepath.Split(filePath)
	ext := filepath.Ext(file)
	file = strings.TrimSuffix(file, ext)
	file = strings.ToLower(file)
	newPath := filepath.Join(dir, file+ext)
	return newPath
}

func (a *App) deleteFileSpecialChars(filePath string) string {
	// 删除特殊字符 保留汉字数字字母下划线中英文括号短横线
	pattern := regexp.MustCompile(`[^0-9a-zA-Z_\x{4e00}-\x{9fa5}()（）\-]`)
	dir, file := filepath.Split(filePath)
	ext := filepath.Ext(file)
	file = strings.TrimSuffix(file, ext)
	file = pattern.ReplaceAllString(file, "")
	newPath := filepath.Join(dir, file+ext)
	return newPath
}

// getNewFileName 获取新的文件名
func (a *App) getNewFileName(filePath string) string {
	rules := a.rules
	switch rules.RenameType {
	case AddPrefix:
		return a.addFilePrefix(filePath, rules.Prefix)
	case AddSuffix:
		return a.addFileSuffix(filePath, rules.Suffix)
	case ReplaceStr:
		return a.replaceFileStr(filePath, rules.ReplaceObj)
	case AddNumber:
		return a.addFileNumber(filePath, rules.NumberObj)
	case ChangeUpperLower:
		return a.changeUpperLower(filePath)
	case DeleteSpecialChars:
		return a.deleteFileSpecialChars(filePath)
	default:
		return filePath
	}
}
