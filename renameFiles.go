package main

import (
	"fmt"
	"os"
	"path"
	"regexp"
	"strings"
)

func (a *App) renameFiles() error {
	// 使用 channel 收集错误
	errChan := make(chan error, len(a.preview))

	for i := range a.preview {
		a.wg.Add(1)
		// 通过参数传递避免闭包变量捕获
		go func(idx int) {
			defer a.wg.Done()

			oldName := a.preview[idx].OldName
			newName := a.preview[idx].NewName

			if err := os.Rename(oldName, newName); err != nil {
				errChan <- fmt.Errorf("重命名文件 %s -> %s 失败: %w", oldName, newName, err)
			}
		}(i)
	}

	// 等待所有 goroutine 完成
	a.wg.Wait()
	close(errChan)

	// 收集所有错误
	var errors []error
	for err := range errChan {
		errors = append(errors, err)
	}

	// 如果有错误，返回组合错误
	if len(errors) > 0 {
		return fmt.Errorf("文件重命名过程中发生 %d 个错误: %v", len(errors), errors)
	}

	return nil
}

func (a *App) addFilePrefix(filePath string, prefix string) string {
	dir, file := path.Split(filePath)
	newPath := path.Join(dir, prefix+file)
	return newPath
}

func (a *App) addFileSuffix(filePath string, suffix string) string {
	dir, file := path.Split(filePath)
	ext := path.Ext(file)
	file = strings.ReplaceAll(file, ext, "")
	newPath := path.Join(dir, file+suffix+ext)
	return newPath
}

func (a *App) replaceFileStr(filePath string, obj *ReplaceObj) string {
	if obj == nil {
		return filePath
	}
	dir, file := path.Split(filePath)
	ext := path.Ext(file)
	file = strings.ReplaceAll(file, ext, "")
	file = strings.ReplaceAll(file, obj.OldStr, obj.NewStr)
	newPath := path.Join(dir, file+ext)
	return newPath
}

func (a *App) addFileNumber(filePath string, numberObj *NumberObj) string {
	dir, file := path.Split(filePath)
	ext := path.Ext(file)
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
	newPath := path.Join(dir, numberObj.NewName+suffix+ext)
	return newPath
}

func (a *App) changeUpperLower(filePath string) string {
	if a.rules.ToUpperCase {
		return a.fileToUpperCase(filePath)
	}
	return a.fileToLowerCase(filePath)
}

func (a *App) fileToUpperCase(filePath string) string {
	dir, file := path.Split(filePath)
	ext := path.Ext(file)
	file = strings.ReplaceAll(file, ext, "")
	file = strings.ToUpper(file)
	newPath := path.Join(dir, file+ext)
	return newPath
}

func (a *App) fileToLowerCase(filePath string) string {
	dir, file := path.Split(filePath)
	ext := path.Ext(file)
	file = strings.ReplaceAll(file, ext, "")
	file = strings.ToLower(file)
	newPath := path.Join(dir, file+ext)
	return newPath
}

func (a *App) deleteFileSpecialChars(filePath string) string {
	// 删除特殊字符 保留汉字数字字母下划线中英文括号短横线
	pattern := regexp.MustCompile(`[^0-9a-zA-Z_\x{4e00}-\x{9fa5}()（）\-]`)
	dir, file := path.Split(filePath)
	ext := path.Ext(file)
	file = strings.ReplaceAll(file, ext, "")
	file = pattern.ReplaceAllString(file, "")
	newPath := path.Join(dir, file+ext)
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
