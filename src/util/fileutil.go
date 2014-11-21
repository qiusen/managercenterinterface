package util

import (
	"container/list"
	"fmt"
	"os"
	"path/filepath"
)

func GetAllFileListByFilePath(filePath string) *list.List {
	fmt.Println(filePath)

	l := list.New()

	filepath.Walk(filePath, func(path string, info os.FileInfo, err error) error {
		if nil == info {
			fmt.Println("nil")
			return err
		}

		if info.IsDir() {
			return nil
		}

		l.PushBack(path)

		return nil
	})

	return l
}
