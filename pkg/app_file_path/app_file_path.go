package app_file_path

import (
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

var (
	ROOT_PATH, PUBLIC_PATH, CONFIG_PATH, ROUTE_PATH, STORAGE_PATH string
)

func Initialize() {
	GetRootPAth()
	PUBLIC_PATH = ROOT_PATH + "/public"
	CONFIG_PATH = ROOT_PATH + "/config"
	ROUTE_PATH = ROOT_PATH + "/routes"
	STORAGE_PATH = ROOT_PATH + "/storage"
}

func GetRootPAth() string {
	ROOT_PATH = getCurrentAbPathByExecutable()
	temp_dir := getTmpDir()
	if strings.Contains(ROOT_PATH, temp_dir) {
		ROOT_PATH = getCurrentAbPathByCaller()
		return ROOT_PATH
	}
	return ROOT_PATH
}

// 以执行环境获取绝对路径,适用于编译打包后在其他目录中运行。
func getCurrentAbPathByExecutable() string {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	res, _ := filepath.EvalSymlinks(filepath.Dir(exePath))
	return res
}

// 获取当前执行文件绝对路径（go run）。以go run方式运行，运行在临时缓存目录
func getCurrentAbPathByCaller() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = filepath.Dir(filepath.Dir(path.Dir(filename)))
	}
	return abPath
}

// 获取系统临时目录，兼容go run
func getTmpDir() string {
	dir := os.Getenv("TEMP")
	if dir == "" {
		dir = os.Getenv("TMP")
	}
	res, _ := filepath.EvalSymlinks(dir)
	return res
}
