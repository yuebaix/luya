package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yuebaix/luya/internal/app/spider"
)

var m3u8Option = &spider.M3u8Option{
	Url: "",
}

func init() {
	m3u8Cmd := &cobra.Command{
		Use:   "m3u8",
		Short: "Find M3u8",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("m3u8 called")
			FindM3u8()
		},
	}
	rootCmd.AddCommand(m3u8Cmd)

	m3u8Cmd.Flags().StringVar(&m3u8Option.Url, "url", "", "")
}

func FindM3u8() {
	spider.FindM3u8(m3u8Option)
}
