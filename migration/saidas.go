package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const awsIP = "100.26.164.166"
const doIP = "159.223.168.215"

func main() {
	// Loading from assets.json
	file, err := ioutil.ReadFile("saidas.json")
	if err != nil {
		fmt.Printf("Couldn't read saidas.json: %s\n", err)
		return
	}

	var assetsMap map[string]interface{}
	err = json.Unmarshal(file, &assetsMap)
	if err != nil {
		fmt.Println("Could not unmarshal into assets map")
		return
	}

	assets, _ := assetsMap["rows"].([]interface{})

	for n, asset := range assets {
		// read asset from AWS network
		// post asset to Digital Ocean network
		a, _ := asset.(map[string]interface{})

		go readAndPost(a)
		if n%200 == 0 {
			time.Sleep(10 * time.Second)
		} else {
			time.Sleep(200 * time.Millisecond)
		}
	}

}

func readAndPost(asset interface{}) {
	a, _ := asset.(map[string]interface{})
	assetKey := a["@key"]

	// post asset
	postUrl := fmt.Sprintf("http://%s/api/invoke/createAsset", doIP)
	postBody, err := json.Marshal(map[string]interface{}{
		"asset": []interface{}{asset},
	})
	if err != nil {
		return
	}

	postRes, err := http.Post(postUrl, "application/json", bytes.NewBuffer(postBody))
	if err != nil {
		fmt.Printf("asset %s couldn't be posted, status %s\n", assetKey, err.Error())
		return
	}

	if postRes.StatusCode != 200 {
		b, _ := ioutil.ReadAll(postRes.Body)
		fmt.Printf("asset %s couldn't be posted, status %d error %s \n", assetKey, postRes.StatusCode, string(b))
		return
	}

	fmt.Printf("SUCCESS: asset %s\n", assetKey)

	defer postRes.Body.Close()
}
