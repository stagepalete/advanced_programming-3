package validator

import (
	"testing"
)

func TestEmailValidation_ValidEmail(t *testing.T) {
	v := New()
	email := "test@example.com"
	v.Check(Matches(email, EmailRX), "email", "Invalid email format")
	if !v.Valid() {
		t.Errorf("Expected no errors for valid email, but got errors: %v", v.Errors)
	}
}

func TestEmailValidation_InvalidEmail(t *testing.T) {
	v := New()
	email := "invalid-email"
	v.Check(Matches(email, EmailRX), "email", "Invalid email format")
	if v.Valid() {
		t.Errorf("Expected errors for invalid email, but got none")
	}
}

func TestUniqueValues_AllUnique(t *testing.T) {
	v := New()
	values := []string{"apple", "banana", "orange"}
	v.Check(Unique(values), "values", "Values must be unique")
	if !v.Valid() {
		t.Errorf("Expected no errors for unique values, but got errors: %v", v.Errors)
	}
}

func TestUniqueValues_NotUnique(t *testing.T) {
	v := New()
	values := []string{"apple", "banana", "apple"}
	v.Check(Unique(values), "values", "Values must be unique")
	if v.Valid() {
		t.Errorf("Expected errors for non-unique values, but got none")
	}
}

func TestInValue_InList(t *testing.T) {
	v := New()
	value := "apple"
	list := []string{"apple", "banana", "orange"}
	v.Check(In(value, list...), "value", "Value not found in list")
	if !v.Valid() {
		t.Errorf("Expected no errors for value in list, but got errors: %v", v.Errors)
	}
}

func TestInValue_NotInList(t *testing.T) {
	v := New()
	value := "grape"
	list := []string{"apple", "banana", "orange"}
	v.Check(In(value, list...), "value", "Value not found in list")
	if v.Valid() {
		t.Errorf("Expected errors for value not in list, but got none")
	}
}
