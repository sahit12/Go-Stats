package etlutils

import (
	"testing"
)

// Testing for MinMaxNormalization correctness
func TestMinMaxNormalization(t *testing.T) {
	input := []float64{56, 1}
	output := MinMaxNormalization(input)
	check := []float64{1, 0}
	for i, v := range output {
		t.Logf("Testing for input:%v", input[i])
		t.Logf("Input:[%v], Expected Output:[%v]", input[i], check[i])
		if v != check[i] {
			t.Errorf("Implementation Error")
		} else {
			t.Logf("Test Pass - Successful")
		}
	}
}

// Testing for Mean correctness
func TestMean(t *testing.T) {
	input := []float64{23, 12, 34}
	actualOutput := Mean(input)
	expectedOutput := 23.0
	if actualOutput == expectedOutput {
		t.Logf("Correct. Successfull run.")
	} else {
		t.Fatalf("Error while running Test.")
	}
}

// Testing performance benchmark of MinMaxNormalization
func BenchmarkMinMaxNormalization(b *testing.B) {
	input := []float64{12, 1, 45, 56, 34}
	for i := 0; i < b.N; i++ {
		MinMaxNormalization(input)
	}
}

// Testing for performance benchmark of Mean
func BenchmarkMean(b *testing.B) {
	input := []float64{12, 1, 45, 56, 34}
	for i := 0; i < b.N; i++ {
		MinMaxNormalization(input)
	}
}
