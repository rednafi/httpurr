package src

import (
	"slices"
	"testing"
)

// Test statusCodes slice
func TestStatusCodes(t *testing.T) {
	// Check length
	want := 68
	got := len(statusCodes)

	if got != want {
		t.Errorf("len(statusCodes) = %d, want %d", got, want)
	}

	var wantStatusCodes = []string{
		"------------------ 1xx ------------------",
		"100",
		"101",
		"102",
		"103",

		"------------------ 2xx ------------------",
		"200",
		"201",
		"202",
		"203",
		"204",
		"205",
		"206",
		"207",
		"208",
		"226",

		"------------------ 3xx ------------------",
		"300",
		"301",
		"302",
		"303",
		"304",
		"305",
		"306",
		"307",
		"308",

		"------------------ 4xx ------------------",
		"400",
		"401",
		"402",
		"403",
		"404",
		"405",
		"406",
		"407",
		"408",
		"409",
		"410",
		"411",
		"412",
		"413",
		"414",
		"415",
		"416",
		"417",
		"418",
		"421",
		"422",
		"423",
		"424",
		"425",
		"426",
		"428",
		"429",
		"431",
		"451",

		"------------------ 5xx ------------------",
		"500",
		"501",
		"502",
		"503",
		"504",
		"505",
		"506",
		"507",
		"508",
		"510",
		"511",
	}

	if !slices.Equal(statusCodes, wantStatusCodes) {
		t.Errorf("statusCodes = %v, want %v", statusCodes, wantStatusCodes)
	}
}

// Test statusCodeMap
func TestStatusCodeMap(t *testing.T) {

	// Check length
	want := 68 - 5 // 5 is the number section headers
	got := len(statusCodeMap)

	if got != want {
		t.Errorf("len(statusCodeMap) = %d, want %d", got, want)
	}
}
