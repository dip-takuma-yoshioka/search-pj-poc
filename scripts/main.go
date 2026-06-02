package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	esURL = "http://localhost:9200"
	osURL = "http://localhost:9201"
)

func main() {
	fmt.Println("=== Elasticsearch & OpenSearch 互換性テスト ===")
	fmt.Println()

	// 1. ヘルスチェック
	fmt.Println("1. ヘルスチェック")
	checkHealth()
	time.Sleep(1 * time.Second)

	// 2. マッピング作成
	fmt.Println("\n2. マッピング作成")
	createMapping()
	time.Sleep(1 * time.Second)

	// 3. データ投入
	fmt.Println("\n3. データ投入")
	indexData()
	time.Sleep(1 * time.Second)

	// 4. 検索テスト
	fmt.Println("\n4. 検索テスト")
	searchTests()

	fmt.Println("\n=== テスト完了 ===")
}

func makeRequest(method, url string, body interface{}) ([]byte, error) {
	var bodyReader io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		bodyReader = bytes.NewBuffer(jsonBody)
	}

	req, err := http.NewRequest(method, url, bodyReader)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}