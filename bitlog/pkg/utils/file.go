package utils

import (
	"fmt"
	"io"
	"os"
)

func ReadLastLine(filePath string) (string, error) {
	fileHandle, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer fileHandle.Close()

	line := ""
	var cursor int64 = 0
	stat, _ := fileHandle.Stat()
	filesize := stat.Size()
	for {
		cursor -= 1
		fileHandle.Seek(cursor, io.SeekEnd)

		char := make([]byte, 1)
		fileHandle.Read(char)
		// stop if we find a line
		if cursor != -1 && (char[0] == 10 || char[0] == 13) {
			break
		}
		// there is more efficient way
		line = fmt.Sprintf("%s%s", string(char), line)
		// stop if we are at the begining
		if cursor == -filesize {
			break
		}
	}
	return line, nil
}

func IsFileExisted(filePath string) bool {
	info, err := os.Stat(filePath)
	return (err == nil || os.IsExist(err)) && !info.IsDir()
}

func IsDirExisted(dirPath string) bool {
	info, err := os.Stat(dirPath)
	return (err == nil || os.IsExist(err)) && info.IsDir()
}

func DirFileCount(dirPath string) int {
	files, _ := os.ReadDir(dirPath)
	return len(files)
}
