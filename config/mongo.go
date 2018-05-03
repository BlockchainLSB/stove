package config

// MongoDB is configurations for connecting Mongo DB.
type MongoDB struct {
	Host   string `yaml:"host"`
	Port   int    `yaml:"port"`
	DBName string `yaml:"db_name"`
}
