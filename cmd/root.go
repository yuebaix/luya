package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "luya",
	Short: "常用命令集合",
}

func Execute() {
	rootCmd.AddCommand(adminCmd(), imCmd())
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
