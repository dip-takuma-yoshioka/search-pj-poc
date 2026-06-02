package main

import (
	"encoding/json"
	"fmt"
)

func searchTests() {
	tests := []struct {
		name  string
		query map[string]interface{}
	}{
		{
			name: "全件検索",
			query: map[string]interface{}{
				"query": map[string]interface{}{
					"match_all": map[string]interface{}{},
				},
			},
		},
		{
			name: "キーワード検索（マウス）",
			query: map[string]interface{}{
				"query": map[string]interface{}{
					"match": map[string]interface{}{
						"product_name": "マウス",
					},
				},
			},
		},
		{
			name: "範囲検索（価格10000円以下）",
			query: map[string]interface{}{
				"query": map[string]interface{}{
					"range": map[string]interface{}{
						"price": map[string]interface{}{
							"lte": 10000,
						},
					},
				},
			},
		},
		{
			name: "複合検索（electronics カテゴリで50000円以下）",
			query: map[string]interface{}{
				"query": map[string]interface{}{
					"bool": map[string]interface{}{
						"must": []interface{}{
							map[string]interface{}{
								"term": map[string]interface{}{
									"category": "electronics",
								},
							},
							map[string]interface{}{
								"range": map[string]interface{}{
									"price": map[string]interface{}{
										"lte": 50000,
									},
								},
							},
						},
					},
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

	for _, test := range tests {
		fmt.Printf("\n--- %s ---\n", test.name)

		for _, service := range services {
			resp, err := makeRequest("GET", service.url+"/products/_search", test.query)
			if err != nil {
				fmt.Printf("❌ %s: %v\n", service.name, err)
			} else {
				var result map[string]interface{}
				json.Unmarshal(resp, &result)
				hits := result["hits"].(map[string]interface{})
				total := hits["total"].(map[string]interface{})
				fmt.Printf("✅ %s: %v件ヒット\n", service.name, total["value"])
			}
		}
	}
}
