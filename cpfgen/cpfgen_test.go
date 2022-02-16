package cpfgen_test

import (
	"gocpf/cpfgen"
	"testing"
)

func TestGenerateFirstDigit(t *testing.T) {
	cpf := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0}

	want := 1
	got := cpfgen.GenerateFirstDigit(cpf)

	if want != got {
		t.Errorf("got %d, want %d.", got, want)
	}
}

func TestGenerateSecondDigit(t *testing.T) {
	cpf := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0}

	want := 1
	got := cpfgen.GenerateSecondDigit(cpf)

	if want != got {
		t.Errorf("got %d, want %d.", got, want)
	}
}

func TestIsRepeatedTrue(t *testing.T) {
	cpf := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}

	want := true
	got := cpfgen.IsRepeated(cpf)

	if want != got {
		t.Errorf("got %t, want %t", got, want)
	}
}

func TestIsRepeatedFalse(t *testing.T) {
	cpf := []int{5, 2, 9, 9, 8, 2, 2, 4, 7, 2, 5}

	want := false
	got := cpfgen.IsRepeated(cpf)

	if want != got {
		t.Errorf("got %t, want %t", got, want)
	}
}

func TestValidateCPFInvalid(t *testing.T) {
	cpf := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2}

	want := false
	got := cpfgen.ValidateCPF(cpf)

	if want != got {
		t.Errorf("got %t, want %t.", got, want)
	}
}

func TestValidateCPFInvalidRepeated(t *testing.T) {
	cpf := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}

	want := false
	got := cpfgen.ValidateCPF(cpf)

	if want != got {
		t.Errorf("got %t, want %t.", got, want)
	}
}

func TestValidateCPFValid(t *testing.T) {
	cpf := []int{0, 3, 3, 3, 6, 8, 4, 9, 4, 0, 3}

	want := true
	got := cpfgen.ValidateCPF(cpf)

	if want != got {
		t.Errorf("got %t, want %t.", got, want)
	}
}

func TestCPFFormat(t *testing.T) {
	cpf := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2}

	want := "123.456.789-12"
	got := cpfgen.CPFFormat(cpf)

	if got != want {
		t.Errorf("got %s, want %s.", got, want)
	}
}
