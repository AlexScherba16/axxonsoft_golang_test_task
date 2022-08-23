package test_test

import (
	_map "axxonsoft_golang_test_task/pkg/map"
	proxy2 "axxonsoft_golang_test_task/pkg/services/proxy"
	"errors"
	"testing"
)

func TestMapGetFromEmptyMap(t *testing.T) {
	testMap := _map.NewRequestResponseMap()
	err, emptyResponse := testMap.GetResponse(proxy2.ClientRequest{})

	if err == nil {
		t.Fatalf("expected : %v, current result : %v", errors.New("Empty map"), err)
	}

	if proxy2.CompareResponseStructs(&proxy2.ClientResponse{}, &emptyResponse) == false {
		t.Fatalf("expected : %t, current result : %t", false, true)
	}
}

func TestMapUpdateGetOk(t *testing.T) {
	testMap := _map.NewRequestResponseMap()
	emptyRequest := proxy2.ClientRequest{}
	emptyResponse := proxy2.ClientResponse{}

	err := testMap.Update(emptyRequest, emptyResponse)
	if err != nil {
		t.Fatal(err)
	}

	err, getResp := testMap.GetResponse(emptyRequest)
	if err != nil {
		t.Fatal(err)
	}

	if proxy2.CompareResponseStructs(&emptyResponse, &getResp) == false {
		t.Fatal("Failed response compare")
	}
}

func TestMapUpdateDoubleGetOk(t *testing.T) {
	testMap := _map.NewRequestResponseMap()
	emptyRequest := proxy2.ClientRequest{}
	emptyResponse := proxy2.ClientResponse{}

	err := testMap.Update(emptyRequest, emptyResponse)
	if err != nil {
		t.Fatal(err)
	}

	err, getResp := testMap.GetResponse(emptyRequest)
	if err != nil {
		t.Fatal(err)
	}
	err, getResp1 := testMap.GetResponse(emptyRequest)
	if err != nil {
		t.Fatal(err)
	}

	first := proxy2.CompareResponseStructs(&emptyResponse, &getResp)
	second := proxy2.CompareResponseStructs(&emptyResponse, &getResp1)

	if (first != true) || (second != true) {
		t.Fatal("Failed response compare")
	}
}

func TestMapSeveralUpdateGetOk(t *testing.T) {
	testMap := _map.NewRequestResponseMap()
	firstRequest := proxy2.ClientRequest{}
	firstResponse := proxy2.ClientResponse{}

	err := testMap.Update(firstRequest, firstResponse)
	if err != nil {
		t.Fatal(err)
	}

	secondRequest := proxy2.ClientRequest{Url: "secondUrl"}
	secondResponse := proxy2.ClientResponse{ID: "secondId"}
	err = testMap.Update(secondRequest, secondResponse)
	if err != nil {
		t.Fatal(err)
	}

	err, getFirstResp := testMap.GetResponse(firstRequest)
	if err != nil {
		t.Fatal(err)
	}
	err, getSecondResp := testMap.GetResponse(secondRequest)
	if err != nil {
		t.Fatal(err)
	}

	first := proxy2.CompareResponseStructs(&firstResponse, &getFirstResp)
	second := proxy2.CompareResponseStructs(&secondResponse, &getSecondResp)
	third := proxy2.CompareResponseStructs(&firstResponse, &getSecondResp)

	if (first != true) || (second != true) || (third != false) {
		t.Fatal("Failed response compare")
	}
}
