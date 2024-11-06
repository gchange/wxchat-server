package httpserver

import "time"

type Config struct {
	Host         string        `yaml:"host" mapstructure:"host"`
	Port         string        `yaml:"port" mapstructure:"port"`
	ReadTimeout  time.Duration `yaml:"readTimeout" mapstructure:"readTimeout"`
	WriteTimeout time.Duration `yaml:"writeTimeout" mapstructure:"writeTimeout"`
	IdleTimeout  time.Duration `yaml:"idleTimeout" mapstructure:"idleTimeout"`
	CloseTimeout time.Duration `yaml:"closeTimeout" mapstructure:"closeTimeout"`
	PProf        bool          `yaml:"pprof" mapstructure:"pprof"`
	Verbose      bool          `yaml:"verbose" mapstructure:"verbose"`
}
