package config

type Config struct {
	Run Run

	Output Output

	LintersSettings LintersSettings `mapstructure:"linters-settings"`
	Linters         Linters
	Issues          Issues
	Severity        Severity
	Version         Version

	InternalCmdTest bool `mapstructure:"internal-cmd-test"` // Option is used only for testing go-readability command, don't use it
	InternalTest    bool // Option is used only for testing go-readability code, don't use it
}

func NewDefault() *Config {
	return &Config{
		LintersSettings: defaultLintersSettings,
	}
}

type Version struct {
	Format string `mapstructure:"format"`
}
