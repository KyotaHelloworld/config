package config

import "github.com/KyotaHelloworld/config/input"

type Config interface {
	GET(key string) (value string)
	UPSERT(key, value string)
}

func NewConf(path string, overwrite ...map[string]string) Config {
	v, err := input.ReadConfigFile(path)
	if err != nil {
		return nil
	}

	for _, m := range overwrite {
		for k, overwriteValue := range m {
			v[k] = overwriteValue
		}
	}

	return &conf{
		path:   path,
		values: v,
	}
}

type conf struct {
	path   string
	values map[string]string
}

func (c *conf) GET(key string) string {
	v, ok := c.values[key]
	if !ok {
		return ""
	}
	return v
}

func (c *conf) UPSERT(key, value string) {
	c.values[key] = value
}
