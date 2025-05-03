package firecracker

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"time"
)

type VMInfo struct {
	State     string `json:"state"`
	CPUCount  int    `json:"vcpu_count"`
	MemSizeMB int    `json:"mem_size_mib"`
}

type VMMetrics struct {
	CPUUsageUs      int `json:"cpu_usage_us"`
	LoadAvgOneMin   int `json:"loadavg_1min"`
	LoadAvgFiveMin  int `json:"loadavg_5min"`
	LoadAvgFifteen  int `json:"loadavg_15min"`
	MemoryRSS       int `json:"rss"`
	NetworkRxBytes  int `json:"net_rx_bytes"`
	NetworkTxBytes  int `json:"net_tx_bytes"`
}

// GetVMInfo fetches /vm from Firecracker API socket
func GetVMInfo(socketPath string) (*VMInfo, error) {
	dialer := func(proto, addr string) (net.Conn, error) {
		return net.Dial("unix", socketPath)
	}
	client := &http.Client{
		Transport: &http.Transport{
			Dial: dialer,
		},
		Timeout: 2 * time.Second,
	}

	req, err := http.NewRequest("GET", "http://unix/machine-config", nil)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		dump, _ := httputil.DumpResponse(resp, true)
		return nil, fmt.Errorf("unexpected response: %s", dump)
	}

	var info VMInfo
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	err = json.Unmarshal(buf.Bytes(), &info)
	if err != nil {
		return nil, err
	}
	return &info, nil
}

// GetVMMetrics fetches /metrics from Firecracker API socket
func GetVMMetrics(socketPath string) (*VMMetrics, error) {
	dialer := func(proto, addr string) (net.Conn, error) {
		return net.Dial("unix", socketPath)
	}
	client := &http.Client{
		Transport: &http.Transport{
			Dial: dialer,
		},
		Timeout: 2 * time.Second,
	}

	req, err := http.NewRequest("GET", "http://unix/metrics", nil)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		dump, _ := httputil.DumpResponse(resp, true)
		return nil, fmt.Errorf("unexpected response or metrics are not enabled in vm: %s", dump)
	}

	var metrics VMMetrics
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	err = json.Unmarshal(buf.Bytes(), &metrics)
	if err != nil {
		return nil, err
	}
	return &metrics, nil
}
