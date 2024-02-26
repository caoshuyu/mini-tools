package build_shell_scripts

import "context"

// IBuildShellScripts 制作shell脚本
type IBuildShellScripts interface {
	// BuildConcurrencyShell 制作并发shell脚本
	BuildConcurrencyShell(outFilePath string)
}

// NewBuildShellScripts 新建制作shell脚本
func NewBuildShellScripts(ctx context.Context) IBuildShellScripts {
	return &buildShellScriptsImpl{
		ctx: ctx,
	}
}
