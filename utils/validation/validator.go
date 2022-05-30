package validation

import (
	"regexp"
	"strings"
	"yamlrest/utils/model"
)

func validateEmail(email string) bool {
	if !regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9-]+(?:\\.[a-zA-Z0-9-]+)*$").MatchString(email) {
		return false
	}
	return true
}

func fieldValidation(body *model.Metadata) []string {
	var individualFields []string

	if len(strings.TrimSpace(body.Version)) < 1 {
		individualFields = append(individualFields, "Version field received empty")
	}
	if len(strings.TrimSpace(body.Company)) < 1 {
		individualFields = append(individualFields, "Company field received empty")
	}
	if len(strings.TrimSpace(body.Description)) < 1 {
		individualFields = append(individualFields, "Description field received empty")
	}
	if len(strings.TrimSpace(body.License)) < 1 {
		individualFields = append(individualFields, "License field received empty")
	}
	if len(strings.TrimSpace(body.Source)) < 1 {
		individualFields = append(individualFields, "Source field received empty")
	}
	if len(strings.TrimSpace(body.Title)) < 1 {
		individualFields = append(individualFields, "Title field received empty")
	}
	if len(strings.TrimSpace(body.Website)) < 1 {
		individualFields = append(individualFields, "Website field received empty")
	}

	if len(body.Maintainers) < 1 {
		individualFields = append(individualFields, "Maintainers field received empty")
	} else {
		for _, person := range body.Maintainers {
			if person.Name == "" {
				individualFields = append(individualFields, "Maintainer name field received empty")
			}
			if person.Email == "" {
				individualFields = append(individualFields, "Maintainer email field received empty")
			}
			if !validateEmail(person.Email) {
				individualFields = append(individualFields, "Maintainer email address not correct")
			}
		}
	}
	return individualFields
}

func ValidateMetadata(input model.Metadata) (bool, string) {

	if result := fieldValidation(&input); len(result) > 0 {
		return false, strings.Join(result, "\n")
	}
	return true, ""
}
