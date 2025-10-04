package divide

import "testing"

func TestDivide(t *testing.T) {
	_, err := Divide(10.0, 1.0)
	if err != nil {
		t.Error("Got an error for valid division:", err)
	}
}

func TestDivideWithValue(t *testing.T) {
	result, err := Divide(10, 2)
	if err != nil {
		t.Errorf("Did not expect an error for valid division, but got: %v", err)
	}
	if result != 5 {
		t.Errorf("Expected result to be 5, but got: %v", result)
	}
}
func TestDivideWithZero(t *testing.T) {
	_, err := Divide(10.0, 0)
	if err == nil {
		t.Error("Expected an error for division by zero, but got none")
	}
	if err.Error() != "division by zero" {
		t.Errorf("Expected 'division by zero' error, but got: %v", err)
	}
}
