package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"recipe-api/models"
	"recipe-api/routes"
	"reflect"
	"testing"
)

func TestHomeHandler(t *testing.T) {
	// テスト用のサーバーを立てる
	ts := httptest.NewServer(routes.NewRoutes())
	defer ts.Close()

	// リクエストを送れるか?
	resp, err := http.Get(fmt.Sprintf("%s/persons", ts.URL))
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	defer resp.Body.Close()

	// Statusコードは200か?
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
	}

	// 期待した通りのレコード数なのか?
	respArr, _ := ioutil.ReadAll(resp.Body)

	var r []models.Person
	json.Unmarshal(respArr, &r)
	length := len(r)

	if length != 6 {
		t.Fatalf("Expected lengh 6, got %v", respArr)
	}

	// 期待した通りのレスポンスなのか?
	raw, err := ioutil.ReadFile("./testResponse.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var mr []models.Person
	json.Unmarshal(raw, &mr)

	if reflect.DeepEqual(mr, r) {
		t.Fatalf("Expected hello world message, got %v", respArr)
	}
}
