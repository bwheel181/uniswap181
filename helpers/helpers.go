package helpers

import (
	"encoding/json"
	"github.com/bwheel181/uniswap181/models"
)

type OrderDirection string

const (
	DESC OrderDirection = "desc"
	ASC OrderDirection = "asc"
)

func NewFetchAssetQuery (orderBy string, direction OrderDirection) []byte {
	jsonData := map[string]string{
		"query": `
             {
                tokens {
                    id
                    name
                    volume
                    volumeUSD
                    poolCount
                    whitelistPools {
                        id
                        volumeToken0
                        volumeToken1
                    }
                }
            }
        `,
	}
	
	jsonValue, _ := json.Marshal(jsonData)
	
	return jsonValue
}

func TranslateAssetResponse(assetID string, tokens []models.Token) []byte {
	for _, token := range tokens {
		if token.ID == assetID {
			token, _ := json.Marshal(token.Pools)
			return token
		}
	}
	
	return []byte{}
}