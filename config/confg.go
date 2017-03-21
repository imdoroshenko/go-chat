package config

type Config struct {
	Settings map[string] interface{}
}

var Settings *Config = new(Config)

func Get(key string) interface{} {
	return Settings[key];
}

func Put(key string, value interface{}) {
	Settings[key] = value;
}

func Override(m map[string] interface{}) {

}

func init() {
	Settings["AAAAA"] = 100;

}
