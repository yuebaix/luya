package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yuebaix/luya/internal/app/im"
)

func init() {
	imCmd := &cobra.Command{
		Use:   "startIm",
		Short: "Start Im Server",
		Run: func(cmd *cobra.Command, args []string) {
			StartImGateway()
			StartImWorker()
		},
	}
	rootCmd.AddCommand(imCmd)
}

func StartImGateway() {
	im.StartGatewayServer()
}

func StartImWorker() {
	im.StarWorkerServer()
}
