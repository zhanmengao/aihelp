package config

type TGptConfig struct {
	V3ApiKey string `yaml:"v3_key" json:"v3_key"`
	V3Host   string `yaml:"v3_host" json:"v3_host"`
	V4ApiKey string `yaml:"v4_key" json:"v4_key"`
	V4Host   string `yaml:"v4_host" json:"v4_host"`
}

type TServerConfig struct {
}

type TTraceConfig struct {
	Addr string `yaml:"addr" json:"addr"`
}
