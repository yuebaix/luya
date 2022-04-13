package web

import "embed"

var (
	//go:embed static
	Static embed.FS
	//go:embed template
	Template embed.FS
)
