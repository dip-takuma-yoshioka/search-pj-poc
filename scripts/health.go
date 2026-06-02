package main

import (
	"encoding/json"
	"fmt"
)

func checkHealth() {
	// Elasticsearch・OpenSearchに同じ検証を行うため、サービス情報を構造体で管理
	services := []struct {
		name string
		url  string
	}{
		{"Elasticsearch", esURL},
		{"OpenSearch", osURL},
	}

	for _, service := range services {
		resp, err := makeRequest("GET", service.url+"/_cluster/health", nil)
		if err != nil {
			fmt.Printf("❌ %s: %v\n", service.name, err)
		} else {
			var result map[string]interface{}
			json.Unmarshal(resp, &result)
			fmt.Printf("✅ %s: status=%s\n", service.name, result["status"])
		}
	}
}
