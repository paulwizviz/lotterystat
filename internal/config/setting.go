package config

// Setting data type representing app configuration
type Setting struct {
	SQLiteSetting `yaml:"sqlite"`
}

// SQLiteSetting for sqlite db
type SQLiteSetting struct {
	Path string `yaml:"path"`
	File string `yaml:"file"`
}

func defaultSetting(p string) Setting {

	return Setting{
		SQLiteSetting: SQLiteSetting{
			Path: p,
			File: "data.db",
		},
	}
}
