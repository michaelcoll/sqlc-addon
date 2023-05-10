package addon

type SqlcConf struct {
	Sql []Sql `yaml:"sql"`
}

type Sql struct {
	Schema string `yaml:"schema"`
}

type SqlcAddonConf struct {
	AddonOut     string `yaml:"addon_out"`
	DatabaseName string `yaml:"database_name"`
}

type TemplateValues struct {
	DatabaseName     string
	MigrationFiles   []MigrationFile
	SqlcAddonVersion string
}

type MigrationFile struct {
	Name    string
	Version int
	Content string
	Last    bool
}
