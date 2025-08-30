package auth

import (
	"log"
	"net/http"
	"testing"
)

func TestGetApiKey(t *testing.T) {

	header := http.Header{}
	header.Add("Authorization", "ApiKey a9sud0dusas90dsd")

	actual, err := GetAPIKey(header)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	expected := "a9sud0dusas90dsd"
	if expected != actual {
		log.Fatalf("%s != %s", expected, actual)
	}

}
