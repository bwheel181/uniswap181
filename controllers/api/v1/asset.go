package v1

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/bwheel181/uniswap181/models"
	"io/ioutil"
	"net/http"
	
	"github.com/bwheel181/uniswap181/helpers"
)

type GraphAssetController interface {
	GetAssetByID(w http.ResponseWriter, r *http.Request)
}

type graphAssetController struct {
	queryURL string
}

func NewGraphAssetController(queryUrl string) GraphAssetController {
	return &graphAssetController{queryURL: queryUrl}
}

func (a *graphAssetController) GetAssetByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	
	req, err := http.NewRequest("POST", a.queryURL, bytes.NewBuffer(helpers.NewFetchAssetQuery("", helpers.DESC)))
	
	client := &http.Client{}
	resp, err := client.Do(req)
	
	defer resp.Body.Close()
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}
	
	data, err := ioutil.ReadAll(resp.Body)
	var asset models.Asset
	err = json.Unmarshal(data, &asset)
	if err != nil {
		fmt.Printf("The response could not be unmarshalled %s", err)
	}
	
	respData := helpers.TranslateAssetResponse(id, asset.Data.Tokens)
	w.Write(respData)
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}
