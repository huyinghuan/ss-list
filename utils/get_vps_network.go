package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type VPSInfo struct {
	PlanMonthlyData int   `json:"plan_monthly_data"`
	DataCounter     int   `json:"data_counter"`
	DataNextReset   int64 `json:"data_next_reset"`
}

func GetVPSNetwork(veid string, apikey string) (string, string) {
	apiURL := "https://api.64clouds.com/v1/getServiceInfo?veid=%s&api_key=%s"
	resp, err := http.Get(fmt.Sprintf(apiURL, veid, apikey))
	if err != nil {
		// handle err
		log.Fatal(err)
		return "", ""
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return "", ""
	}
	result := new(VPSInfo)
	data, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(data, &result)
	remaining := (result.PlanMonthlyData - result.DataCounter) / 1024 / 1024

	tm := time.Unix(result.DataNextReset, 0)

	if remaining/1024 <= 0 {
		return fmt.Sprintf("%d MB", remaining), tm.Format("2006/01/02 15:04:05")
	}

	return fmt.Sprintf("%f GB", float32(remaining)/1024), tm.Format("2006/01/02 15:04:05")
}
