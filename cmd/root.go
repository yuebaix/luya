package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

/*
______
___  / ____  ______  _______ _
__  /  _  / / /_  / / /  __ `/
_  /___/ /_/ /_  /_/ // /_/ /
/_____/\__,_/ _\__, / \__,_/
              /____/           :~) luya 常用命令集合
*/

var rootCmd = &cobra.Command{
	Use:   "luya",
	Short: "常用命令集合",
	Long:  "______\n___  / ____  ______  _______ _\n__  /  _  / / /_  / / /  __ `/\n_  /___/ /_/ /_  /_/ // /_/ /\n/_____/\\__,_/ _\\__, / \\__,_/\n              /____/           :~) luya 常用命令集合",
}

func init() {

}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
