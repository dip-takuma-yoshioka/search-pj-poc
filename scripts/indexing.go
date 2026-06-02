package main

import (
	"fmt"
)

func indexData() {
	products := []map[string]interface{}{
		{
			"product_id":   "P001",
			"product_name": "ノートパソコン",
			"category":     "electronics",
			"price":        89800,
			"description":  "高性能な薄型ノートパソコン",
			"created_at":   "2024-01-15",
		},
		{
			"product_id":   "P002",
			"product_name": "ワイヤレスマウス",
			"category":     "electronics",
			"price":        3980,
			"description":  "Bluetooth接続の静音マウス",
			"created_at":   "2024-01-20",
		},
		{
			"product_id":   "P003",
			"product_name": "コーヒーメーカー",
			"category":     "kitchen",
			"price":        12800,
			"description":  "全自動のコーヒーメーカー",
			"created_at":   "2024-02-01",
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

	for i, product := range products {
		for _, service := range services {
			_, err := makeRequest("POST", service.url+"/products/_doc", product)
			if err != nil {
				fmt.Printf("❌ %s 商品%d: %v\n", service.name, i+1, err)
			} else {
				fmt.Printf("✅ %s 商品%d 投入成功\n", service.name, i+1)
			}
		}
	}

	// インデックスをリフレッシュ
	fmt.Println("\nインデックスをリフレッシュ中...")
	for _, service := range services {
		makeRequest("POST", service.url+"/products/_refresh", nil)
	}
}
