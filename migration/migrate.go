package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const awsIP = "100.26.164.166"
const doIP = "159.223.168.215"

func main() {
	// Loading from assets.json
	file, err := ioutil.ReadFile("assets.json")
	if err != nil {
		fmt.Printf("Couldn't read assets.json: %s\n", err)
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
		assetKey, _ := a["key"].(string)

		go readAndPost(assetKey)
		if n%200 == 0 {
			time.Sleep(10 * time.Second)
		} else {
			time.Sleep(200 * time.Millisecond)
		}
	}

}

func readAndPost(assetKey string) {
	// read asset
	readUrl := fmt.Sprintf("http://%s/api/query/readAsset", awsIP)

	assetType := strings.Split(assetKey, ":")[0]
	readBody, err := json.Marshal(map[string]interface{}{
		"key": map[string]interface{}{
			"@assetType": assetType,
			"@key":       assetKey,
		},
		"resolve": true,
	})
	if err != nil {
		fmt.Printf("FAIL: asset %s couldn't be read, err: %s\n", assetKey, err.Error())
		return
	}

	readRes, err := http.Post(readUrl, "application/json", bytes.NewBuffer(readBody))
	if err != nil {
		fmt.Printf("FAIL: asset %s couldn't be read, err: %s\n", assetKey, err.Error())
		return
	}

	if readRes.StatusCode != 200 {
		fmt.Printf("FAIL: asset %s couldn't be read\n", assetKey)
	}

	defer readRes.Body.Close()

	body, err := ioutil.ReadAll(readRes.Body)
	if err != nil {
		fmt.Printf("FAIL: asset %s couldn't be read\n", assetKey)
		return
	}

	resBody := make(map[string]interface{})
	err = json.Unmarshal(body, &resBody)
	if err != nil {
		fmt.Printf("FAIL: asset %s couldn't be read\n", assetKey)
		return
	}

	// post asset
	postUrl := fmt.Sprintf("http://%s/api/invoke/createAsset", doIP)
	postBody, err := json.Marshal(map[string]interface{}{
		"asset": []interface{}{resBody},
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
		fmt.Printf("asset %s couldn't be posted, status %d\n", assetKey, postRes.StatusCode)
		return
	}

	fmt.Printf("SUCCESS: asset %s\n", assetKey)

	defer postRes.Body.Close()
}
