// Code generated . DO NOT EDIT.
// ################################## DO NOT EDIT THIS FILE ######################################
// Common Sense Coding (https://github.com/cscoding21/csgen)

// Generate Date: 2024-12-30 09:13:48.074341725 -0800 PST m=+0.000536574
// Implementation Name: csval
// Developer Note: The contents of this file will be recreated each time its generator is called

// -----------------------------------------------------------------------------------------------

package list

import "github.com/cscoding21/csval/validate"

// Validate checks the fields in the struct List to ensure it conforms to business rules
func (obj *List) Validate() validate.ValidationResult {
	result := validate.NewSuccessValidationResult()

	// ---Field: Name
	result.Append(validate.IsNotEmpty("Name", obj.Name))

	return result
}

// Validate checks the fields in the struct ListItem to ensure it conforms to business rules
func (obj *ListItem) Validate() validate.ValidationResult {
	result := validate.NewSuccessValidationResult()

	return result
}
