package addon

import "embed"

//go:embed templates/connect.go.gotmpl
//go:embed templates/migration.go.gotmpl
var templates embed.FS
