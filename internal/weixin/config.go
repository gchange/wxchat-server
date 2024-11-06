package weixin

type Config struct {
	AppId  string `json:"app-id" yaml:"app-id" mapstructure:"app-id"`
	Secret string `json:"secret" yaml:"secret" mapstructure:"secret"`
}
