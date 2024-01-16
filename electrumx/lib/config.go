package lib

type Config struct {
	coin Coin
	// server data and database dir
	Datadir string
	//
	PreFetchPollingDelaySecs int
}

func GetDefaultConfig() *Config {
	return &Config{
		Datadir:                  "datadir",
		PreFetchPollingDelaySecs: 10,
	}
}

func (c *Config) SetCoin(coin Coin) error {
	// TODO: error checking
	c.coin = coin
	return nil
}

func (c *Config) GetCoin() Coin {
	return c.coin
}
