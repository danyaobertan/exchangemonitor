package config

type Configuration struct {
	App  AppConfiguration
	DB   DBConfiguration
	SMTP SMTPConfig
}

type AppConfiguration struct {
	Port  int
	Name  string
	Debug bool
}

type DBConfiguration struct {
	ConnectionURL                string
	DefaultMaxConnections        int
	DefaultMaxConnectionLifetime int
	DefaultMaxConnectionIdleTime int
	DefaultHealthCheckPeriod     int
	DefaultHealthCheckTimeout    int
	MigrationSource              string
	MigrationQueryParams         string
}

type SMTPConfig struct {
	Host     string
	Port     string
	Username string
	Password string
}
