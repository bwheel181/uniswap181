package models

type Asset struct {
	Data struct {
		Tokens []Token `json:"tokens"`
	} `json:"data"`
}

type Data struct {
	Tokens []Token `json:"tokens"`
}

type Token struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Volume    string `json:"volume"`
	VolumeUSD string `json:"volumeUSD"`
	PoolCount int64 `json:"poolCount"`
	Pools     []Pool `json:"whitelistPools"`
}

type Pool struct {
	ID           string `json:"id"`
	VolumeToken0 int64 `json:"volumeToken0"`
	VolumeToken1 int64 `json:"volumeToken1"`
}
