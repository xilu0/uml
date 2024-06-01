package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// ClusterStatus 定义顶层的JSON结构
type ClusterStatus struct {
	Alarms            []interface{}          `json:"alarms"`
	ClusterName       string                 `json:"cluster_name"`
	DiskNodes         []string               `json:"disk_nodes"`
	FeatureFlags      []FeatureFlag          `json:"feature_flags"`
	Listeners         map[string][]Listener  `json:"listeners"`
	MaintenanceStatus map[string]string      `json:"maintenance_status"`
	Partitions        map[string]interface{} `json:"partitions"`
	RamNodes          []interface{}          `json:"ram_nodes"`
	RunningNodes      []string               `json:"running_nodes"`
	Versions          map[string]VersionInfo `json:"versions"`
}

// FeatureFlag 定义feature_flags数组中的项
type FeatureFlag struct {
	Description string `json:"desc"`
	DocURL      string `json:"doc_url"`
	Name        string `json:"name"`
	ProvidedBy  string `json:"provided_by"`
	Stability   string `json:"stability"`
	State       string `json:"state"`
}

// Listener 定义listeners字典中的项
type Listener struct {
	Interface string `json:"interface"`
	Node      string `json:"node"`
	Port      int    `json:"port"`
	Protocol  string `json:"protocol"`
	Purpose   string `json:"purpose"`
}

// VersionInfo 定义versions字典中的版本信息
type VersionInfo struct {
	ErlangVersion   string `json:"erlang_version"`
	RabbitmqName    string `json:"rabbitmq_name"`
	RabbitmqVersion string `json:"rabbitmq_version"`
}

// parseJSON 解析JSON数据到结构体
func parseJSON(data string) *ClusterStatus {
	var status ClusterStatus
	err := json.Unmarshal([]byte(data), &status)
	if err != nil {
		log.Fatal(err)
	}
	return &status
}

func main() {
	jsonData := `{"alarms":[],"cluster_name":"rabbit@pam-rabbitmq","disk_nodes":["rabbit@pam-rabbitmq"],"feature_flags":[{"desc":"Count unroutable publishes to be dropped in stats","doc_url":"","name":"drop_unroutable_metric","provided_by":"rabbitmq_management_agent","stability":"stable","state":"enabled"}],"listeners":{"rabbit@pam-rabbitmq":[{"interface":"[::]","node":"rabbit@pam-rabbitmq","port":15692,"protocol":"http/prometheus","purpose":"Prometheus exporter API over HTTP"}]},"maintenance_status":{"rabbit@pam-rabbitmq":"not under maintenance"},"partitions":{},"ram_nodes":[],"running_nodes":["rabbit@pam-rabbitmq"],"versions":{"rabbit@pam-rabbitmq":{"erlang_version":"24.1.7","rabbitmq_name":"RabbitMQ","rabbitmq_version":"3.8.26"}}}`
	status := parseJSON(jsonData)
	fmt.Printf("Cluster Name: %s\n", status.ClusterName)
	fmt.Printf("Running Nodes: %v\n", status.RunningNodes)
	for _, flag := range status.FeatureFlags {
		fmt.Printf("Feature: %s, State: %s\n", flag.Name, flag.State)
	}
}
