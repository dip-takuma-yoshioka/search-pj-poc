package main

import (
	"fmt"
)

func createMapping() {
	mapping := map[string]interface{}{
		"mappings": map[string]interface{}{
			"properties": map[string]interface{}{
				"product_id": map[string]string{
					"type": "keyword",
				},
				"product_name": map[string]string{
					"type": "text",
				},
				"category": map[string]string{
					"type": "keyword",
				},
				"price": map[string]string{
					"type": "float",
				},
				"description": map[string]string{
					"type": "text",
				},
				"created_at": map[string]string{
					"type": "date",
				},
			},
		},
	}

	// Elasticsearch・OpenSearchに同じ検証を行うため、サービス情報を構造体で管理
	services := []struct {
		name string
		url  string
	}{
		{"Elasticsearch", esURL},
		{"OpenSearch", osURL},
	}

	for _, service := range services {
		resp, err := makeRequest("PUT", service.url+"/products", mapping)
		if err != nil {
			fmt.Printf("❌ %s: %v\n", service.name, err)
		} else {
			fmt.Printf("✅ %s: %s\n", service.name, string(resp))
		}
	}
}
