package model

const (
	DatabaseDriverMysql    = "mysql"
	DatabaseDriverPostgres = "postgres"

	SqlSettingsDefaultDataSource = "postgres://mmuser:mostest@localhost/mattermost_test?sslmode=disable&connect_timeout=10"

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

func (s *SqlSettings) SetDefaults(isUpdate bool) {
	if s.DriverName == nil {
		s.DriverName = NewString(DatabaseDriverMysql)
	}

	if s.DataSource == nil {
		s.DataSource = NewString(SqlSettingsDefaultDataSource)
	}
}
