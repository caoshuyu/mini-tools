package cmd

import (
	"context"
	"fmt"
	"github.com/caoshuyu/mini-tools/mimi-tools/logic/build_shell_scripts"
	"github.com/spf13/cobra"
)

// 生成并发shell文件 输出目录
var bCSOutFilePath string

func init() {
	var buildDbCmd = &cobra.Command{
		Use:   "build-concurrency-shell",
		Short: "生成并发shell文件",
		Long: "生成并发shell文件\n " +
			"输入参数参考：\n " +
			`./mini-tools build-concurrency-shell --outFilePath "" `,
		Run: runBuildConcurrencyShell,
	}
	useCmd.AddCommand(buildDbCmd)
	buildDbCmd.Flags().StringVar(&bCSOutFilePath, "outFilePath", "./", "out put file path")
}

func runBuildConcurrencyShell(cmd *cobra.Command, args []string) {
	fmt.Println("start build-concurrency-shell function")
	bssLogic := build_shell_scripts.NewBuildShellScripts(context.Background())
	bssLogic.BuildConcurrencyShell(bCSOutFilePath)
}
