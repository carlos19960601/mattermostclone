package model

const (
	SHOW_USERNAME = "username"
)

type Config struct {
	ServiceSettings ServiceSettings
	SqlSettings     SqlSettings
}

type ServiceSettings struct {
	ListenAddress          string
	SessionLengthWebInDays *int
	SessionCacheInMinutes  *int
}

type SqlSettings struct {
	DriverName        *string
	DataSource        *string
	DataSourceRelicas []string
}
