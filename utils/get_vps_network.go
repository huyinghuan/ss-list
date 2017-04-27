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
	PlanMonthlyData int    `json:"plan_monthly_data"`
	DataCounter     int    `json:"data_counter"`
	DataNextReset   int64  `json:"data_next_reset"`
	NodeIP          string `json:"node_ip"`
	Remaining       string
	DataReset       string
	IPAddresses     []string `json:"ip_addresses"`
}

func GetVPSNetwork(veid string, apikey string) *VPSInfo {
	apiURL := "https://api.64clouds.com/v1/getServiceInfo?veid=%s&api_key=%s"
	resp, err := http.Get(fmt.Sprintf(apiURL, veid, apikey))
	result := new(VPSInfo)
	if err != nil {
		// handle err
		log.Fatal(err)
		return result
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return result
	}
	data, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(data, &result)
	remaining := (result.PlanMonthlyData - result.DataCounter) / 1024 / 1024

	tm := time.Unix(result.DataNextReset, 0)
	result.DataReset = tm.Format("2006/01/02 15:04:05")
	if remaining/1024 <= 0 {
		result.Remaining = fmt.Sprintf("%d MB", remaining)
	} else {
		result.Remaining = fmt.Sprintf("%f GB", float32(remaining)/1024)
	}
	return result

}
