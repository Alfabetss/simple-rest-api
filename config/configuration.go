package config

// Configuration exported
type Configuration struct {
	Server   Server
	Database Database
}

// Server exported
type Server struct {
	Port string
}

// Database exported
type Database struct {
	DBProtocol string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPass     string
	DBName     string
}
