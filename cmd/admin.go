package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yuebaix/luya/internal/app/admin"
	"github.com/yuebaix/luya/internal/pkg/util"
)

func init() {
	util.LoadConfig("admin-server")

	adminCmd := &cobra.Command{
		Use:   "startAdmin",
		Short: "Start Admin Server",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("startAdmin called")
			startAdmin()
		},
	}
	rootCmd.AddCommand(adminCmd)
}

func startAdmin() {
	admin.StartAdminServer()
}
