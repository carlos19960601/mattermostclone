package model

type ServiceSettings struct {
	ListenAddress string
}

type Config struct {
	ServiceSettings ServiceSettings
}
