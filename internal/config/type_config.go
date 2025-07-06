package config

type Config struct {
	DbUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (c *Config) SetUser(userName string) error {
	c.CurrentUserName = userName

	err := Write(*c)
	if err != nil {
		return err
	}

	return nil
}
