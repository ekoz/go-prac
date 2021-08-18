package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

// 根据指定目录和文件类型，返回目录下所有文件，包含子目录
// dirPath 目标路径
// filetypes 目标文件类型map，如 .jpg .png .mp4 等
func GetAllFiles(dirPath string, filetypes map[string]string) (files []string, err error) {
	fis, err := ioutil.ReadDir(filepath.Clean(filepath.ToSlash(dirPath)))
	if err != nil {
		return nil, err
	}

	for _, f := range fis {
		_path := filepath.Join(dirPath, f.Name())

		if f.IsDir() {
			fs, _ := GetAllFiles(_path, filetypes)
			files = append(files, fs...)
			continue
		}

		// 指定格式
		extention := filepath.Ext(f.Name())
		if filetypes == nil {
			files = append(files, _path)
		} else if filetypes[extention] != "" {
			files = append(files, _path)
		}
		// switch filepath.Ext(f.Name()) {
		// case ".png", ".jpg", ".mp4":
		// 	files = append(files, _path)
		// }
	}

	return files, nil
}

func main() {

	// 找到目标文件夹下，所有以mp4结尾的文件，并打印路径
	targetDir := "D:/Images/Video"
	filetypes := make(map[string]string)
	filetypes[".png"] = ".png"
	filetypes[".jpg"] = ".jpg"
	filetypes[".mp4"] = ".mp4"
	fis, err := GetAllFiles(targetDir, filetypes)
	if err != nil {
		fmt.Printf("read dir error: %v\n", err)
		return
	}

	for _, f := range fis {
		fmt.Printf("%s\n", strings.Replace(f, "\\", "/", -1))
	}

}
