package parser

import (
	"crawler/fetcher"
	"fmt"
	"testing"
)

func TestParseCityList(t *testing.T) {
	contents, err := fetcher.Fetch("http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", contents)

	result := ParseCityList(contents)

	const resultSize = 470
	expectedUrls := []string{
		"", "", "",
	}
	expectedCities := []string{
		"", "", "",
	}
	if len(result.Requests) != resultSize {
		t.Errorf("result should have %d requests, but had", resultSize, len(result.Requests))
	}
	if len(result.Items) != resultSize {
		t.Errorf("result should have %d requests, but had", resultSize, len(result.Items))
	}
	//ParseCityList(contents)
}
