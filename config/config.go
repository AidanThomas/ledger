package config

type Configuration struct {
	SupportedDBs map[string]string
}

func Load() (*Configuration, error) {
	var conf Configuration

	conf.SupportedDBs = map[string]string{
		"psql": "postgres",
	}

	if err := conf.Validate(); err != nil {
		return nil, err
	}

	return &conf, nil
}

func (c *Configuration) Validate() error {
	return nil
}
