package gorm_test

import (
	"errors"
	"testing"

	"github.com/jinzhu/gorm"
)

func TestErrorsCanBeUsedOutsideGorm(t *testing.T) {
	errs := []error{errors.New("First"), errors.New("Second")}

	gErrs := gorm.Errors(errs)
	gErrs = gErrs.Add(errors.New("Third"))
	gErrs = gErrs.Add(gErrs)

	if gErrs.Error() != "First; Second; Third" {
		t.Fatalf("Gave wrong error, got %s", gErrs.Error())
	}
}

func TestIsRecordNotFoundError(t *testing.T) {
	err := errors.New("record not found")
	if !gorm.IsRecordNotFoundError(err) {
		t.Fatal("The check for a single record not found error returned false while expecting true")
	}

	errs := gorm.Errors([]error{errors.New("record not found")})
	if !gorm.IsRecordNotFoundError(errs) {
		t.Fatal("The check for a record not found error returned false while expecting true")
	}
}
