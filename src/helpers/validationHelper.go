package helpers

import (
	"log"
	"strings"

	"github.com/dongri/phonenumber"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func InitializeValidator() {
	validate = validator.New()

	validate.RegisterValidation("phone", ValidatePhone)
}

func StructValidator(theStruct interface{}) error {
	err := validate.Struct(theStruct)
	return err
}

func ValidatePhone(fld validator.FieldLevel) bool {
    // Log the field value for debugging
    log.Printf("Field Value: %s", fld.Field().String())

    contplus := strings.Contains(fld.Field().String(), "+")
    if !contplus {
        return false
    }

    remplus := strings.Trim(fld.Field().String(), "+")
    
    // Log the trimmed value for debugging
    log.Printf("Trimmed Value: %s", remplus)

    asd := phonenumber.GetISO3166ByNumber(remplus, true)

    // Log the result for debugging
    log.Printf("ISO3166 Result: %+v", asd)

    return asd.CountryName != ""
}
