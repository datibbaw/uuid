package uuid_test

import (
	"errors"
	"github.com/datibbaw/uuid"
	"testing"
)

type randomSource struct {
	data []byte
	err error
}

func (source randomSource) Read(target []byte) (int, error) {
	copy(target[:], source.data)
	return len(source.data), source.err
}

func TestCreateEmptyV4(t *testing.T)  {
	data := make([]byte, 16)
	expected := "00000000-0000-4000-8000-000000000000"
	actual, _ := uuid.CreateV4(randomSource{data, nil})
	if actual != expected {
		t.Errorf("Expected app of %v to be %s, but got %s", data, expected, actual)
	}
}

func TestCreateRandomV4(t *testing.T) {
	data := []byte{104, 11, 159, 228, 50, 48, 100, 161, 11, 76, 155, 145, 49, 121, 236, 37}
	expected := "680b9fe4-3230-44a1-8b4c-9b913179ec25"
	actual, _ := uuid.CreateV4(randomSource{data, nil})
	if expected != actual {
		t.Errorf("Expected app of %v to be %s, but got %s", data, expected, actual)
	}
}

func TestFailureV4(t *testing.T) {
	expectedError := errors.New("read error")
	_, err := uuid.CreateV4(randomSource{make([]byte, 16), expectedError})
	if err == nil {
		t.Error("Expected an error to be returned")
	}
}