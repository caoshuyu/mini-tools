package build_shell_scripts

import (
	"context"
	"fmt"
	"github.com/sframe-go/sframe-kit/filekit"
	"path/filepath"
	"strings"
)

type buildShellScriptsImpl struct {
	ctx context.Context
}

// BuildConcurrencyShell 制作并发shell脚本
func (b buildShellScriptsImpl) BuildConcurrencyShell(outFilePath string) {
	fileDir, fileName, demoFileName := "", "", ""
	if vIndex := strings.Index(outFilePath, "."); vIndex == -1 {
		// 目录
		fileDir = outFilePath
		fileName = "concurrency.sh"
		demoFileName = "concurrency_demo.sh"
	} else {
		fileDir, fileName = filepath.Split(outFilePath)
		if nameIndex := strings.Index(fileName, "."); nameIndex == -1 {
			fileName += ".sh"
			demoFileName = fileName + "_demo.sh"
		} else {
			val := fileName[:nameIndex]
			demoFileName = val + "_demo.sh"
		}
	}

	fmt.Println("==========fileDir==========:", fileDir)
	fmt.Println("==========fileName=========:", fileName)
	fmt.Println("==========demoFileName=====:", demoFileName)

	conShellFile := b.buildConcurrencySh()
	conShellDemoFile := b.buildConcurrencyShDemo()
	if err := filekit.WriteFileByteWithDir(fileDir, fileName, []byte(conShellFile)); err != nil {
		panic(err)
	}
	if err := filekit.WriteFileByteWithDir(fileDir, demoFileName, []byte(conShellDemoFile)); err != nil {
		panic(err)
	}
}

// 制作并发脚本
func (b buildShellScriptsImpl) buildConcurrencySh() string {
	dataList := make([]string, 0)
	dataList = append(dataList, "#!/bin/bash", "")
	dataList = append(dataList, "# shell并发方法", "# 必要参数range_arr，thread_num可选参数，默认为10")
	dataList = append(dataList, "# range_arr执行无序", "function concurrency() {",
		"    # check use params",
		"    if [[ ${#range_arr[*]} -eq 0 ]];then", `        echo "params range_arr not have or empty!!!"`, "    fi",
		"    if [[ ${thread_num} -eq 0 ]];then",
		`        echo "params thread_num not have or eq 0,use defalt value 10 !!!"`,
		"        thread_num=10", "    fi", "",
		"    # mkfifo 创建命名管道", `    tempfifo="my_temp_fifo"`, "    mkfifo ${tempfifo}",
		"    #关联fifo文件和fd6，使文件描述符为非阻塞式", "    exec 6<>${tempfifo}", "    rm -f ${tempfifo}",
		"    # 为文件描述符创建占位信息", "    for ((i=1;i<=${thread_num};i++))", "    do", "    {", "        echo",
		"    }", "    done >&6", "",
		"    # 遍历应用程序", "    for i in ${!range_arr[@]}", "    do", "    {",
		"        read -u6 ##read -u6命令执行一次，相当于尝试从fd6中获取一行，如果获取不到，则阻塞获取到了一行后，"+
			"fd6就少了一行了，开始处理子进程，子进程放在后台执行", "        {", "            ${range_arr[${i}]}",
		"            echo >&6 #完成后再补充一个空值到fd6中，释放一个锁", "        } &", "    }", "    done", "    wait", "",
		"    # 关闭fd6管道", "    exec 6>&-", "}", "",
	)
	dataList = append(dataList, "# declare -A range_arr", `# range_arr[0]="echo info 0"`,
		`# range_arr[1]="echo info 1"`, "", "# 使用并发方法，需要存在range_arr数组和thread_num并发数",
		"# concurrency ${range_arr[*]}")

	return strings.Join(dataList, "\n")
}

// 制作并发脚本demo
func (b buildShellScriptsImpl) buildConcurrencyShDemo() string {
	dataList := make([]string, 0)
	dataList = append(dataList, "#!/bin/bash", "")
	dataList = append(dataList, "# 引入并发函数文件", ". ./concurrency.sh", "")
	dataList = append(dataList, "# 设置并发使用数组", "declare -A range_arr", `range_arr[0]="echo info 0"`,
		`range_arr[1]="echo info 1"`, `range_arr[2]="echo info 2"`, `range_arr[3]="echo info 3"`,
		`range_arr[4]="echo info 4"`, `range_arr[5]="echo info 5"`, `range_arr[6]="echo info 6"`,
		`range_arr[7]="echo info 7"`, `range_arr[8]="echo info 8"`, `range_arr[9]="echo info 9"`,
		`range_arr[10]="echo info 10"`, `range_arr[11]="echo info 11"`, `range_arr[12]="echo info 12"`,
		`range_arr[13]="echo info 13"`, `range_arr[14]="echo info 14"`, "",
	)
	dataList = append(dataList, "# 并发数", "thread_num=5", "")
	dataList = append(dataList, "# 调用并发方法", "concurrency")

	return strings.Join(dataList, "\n")
}
