package httpserver

import "time"

type Config struct {
	Host         string        `json:"host" yaml:"host" mapstructure:"host"`
	Port         string        `json:"port" yaml:"port" mapstructure:"port"`
	ReadTimeout  time.Duration `json:"read-timeout" yaml:"read-timeout" mapstructure:"read-timeout"`
	WriteTimeout time.Duration `json:"write-timeout" yaml:"write-timeout" mapstructure:"write-timeout"`
	IdleTimeout  time.Duration `json:"idle-timeout" yaml:"idle-timeout" mapstructure:"idle-timeout"`
	CloseTimeout time.Duration `json:"close-timeout" yaml:"close-timeout" mapstructure:"close-timeout"`
	PProf        bool          `json:"pprof" yaml:"pprof" mapstructure:"pprof"`
	Verbose      bool          `json:"verbose" yaml:"verbose" mapstructure:"verbose"`
}
