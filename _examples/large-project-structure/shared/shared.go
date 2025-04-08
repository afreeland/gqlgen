package shared

import (
	"embed"
)

//go:embed "graph/schema.graphqls"
var sourcesFS embed.FS
