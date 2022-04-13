package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yuebaix/luya/internal/app/admin"
)

func init() {

}

func adminCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "startAdmin",
		Short: "Start Admin Server",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("startAdmin called")
			startAdmin()
		},
	}
}

func startAdmin() {
	admin.StartAdminServer()
}
