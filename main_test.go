package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"recipe-api/models"
	"reflect"
	"testing"
)

func TestHomeHandler(t *testing.T) {
	// テスト用のサーバーを立てる
	ts := httptest.NewServer(SetupServer())
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
	responseData, _ := ioutil.ReadAll(resp.Body)
	length := len(string(responseData))
	if length != 6 {
		t.Fatalf("Expected lengh 6, got %v", responseData)
	}

	// 期待した通りのレスポンスなのか?
	raw, err := ioutil.ReadFile("./testResponse.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var mr []models.Person
	var r []models.Person

	json.Unmarshal(raw, &mr)
	json.Unmarshal(responseData, &r)

	if reflect.DeepEqual(mr, r) {
		t.Fatalf("Expected hello world message, got %v", responseData)
	}
}
