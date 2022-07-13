package app

import (
	"testing"
)

func TestGetUser(t *testing.T) {

	type testCase struct {
		id string
	}

	testCases := []testCase{{
		id: "success",
	}}

	for _, testCase := range testCases {
		t.Run(testCase.id, func(t *testing.T) {

		})
	}

}
