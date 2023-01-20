package utils

import (
	"testing"
)

func TestIsPrime_11(t *testing.T) {
	t.Parallel()
	// arrange
	no := 11
	expectedResult := true

	// act
	actualResult := IsPrime(no)

	// assert
	if actualResult != expectedResult {

		// t.Logf("IsPrime() : arg - %d, expected : %t but go %t\n", no, expectedResult, actualResult)
		// t.Fail()

		t.Errorf("IsPrime() : arg - %d, expected : %t but got %t\n", no, expectedResult, actualResult)
		// t.Fatalf("IsPrime() : arg - %d, expected : %t but got %t\n", no, expectedResult, actualResult)
	}
}

func TestIsPrime_13(t *testing.T) {
	t.Parallel()
	// fmt.Println("Testing IsPrime : 11")
	// arrange
	no := 13
	expectedResult := true

	// act
	actualResult := IsPrime(no)

	// assert
	if actualResult != expectedResult {

		// t.Logf("IsPrime() : arg - %d, expected : %t but go %t\n", no, expectedResult, actualResult)
		// t.Fail()

		t.Errorf("IsPrime() : arg - %d, expected : %t but got %t\n", no, expectedResult, actualResult)
		// t.Fatalf("IsPrime() : arg - %d, expected : %t but got %t\n", no, expectedResult, actualResult)
	}
}

/*
type PrimeTestData struct {
	name           string
	no             int
	expectedResult bool
	actualResult   bool
}
*/

/*
func TestIsPrime(t *testing.T) {
	testData := []struct {
		name           string
		no             int
		expectedResult bool
		actualResult   bool
	}{
		{name: "IsPrime_11", no: 11, expectedResult: false},
		{name: "IsPrime_13", no: 13, expectedResult: true},
		{name: "IsPrime_15", no: 15, expectedResult: false},
		{name: "IsPrime_17", no: 17, expectedResult: true},
		{name: "IsPrime_20", no: 20, expectedResult: false},
	}
	for _, td := range testData {
		t.Run(td.name, func(t *testing.T) {
			td.actualResult = IsPrime(td.no)
			if td.actualResult != td.expectedResult {
				t.Errorf("IsPrime() : arg - %d, expected : %t but got %t\n", td.no, td.expectedResult, td.actualResult)
				// t.Fatalf("IsPrime() : arg - %d, expected : %t but got %t\n", no, expectedResult, actualResult)
			}
		})
	}
}
*/
