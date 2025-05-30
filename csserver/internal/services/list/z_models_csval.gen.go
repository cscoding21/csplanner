// Code generated . DO NOT EDIT.
// ################################## DO NOT EDIT THIS FILE ######################################
// Common Sense Coding (https://github.com/cscoding21/csgen)

// Generate Date: 2025-05-28 15:10:25.758662272 -0700 PDT m=+0.000522437
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

	// ---Field: Description
	result.Append(validate.IsNotEmpty("Description", obj.Description))

	return result
}

// Validate checks the fields in the struct ListItem to ensure it conforms to business rules
func (obj *ListItem) Validate() validate.ValidationResult {
	result := validate.NewSuccessValidationResult()

	return result
}
