package errors

import "strings"

func removeOrgInError(fields, values string) (string, string) {
	if strings.Contains(fields, "organization_id") {
		tmpFields := strings.Split(strings.ReplaceAll(fields, " ", ""), ",")
		tmpValues := strings.Split(strings.ReplaceAll(values, " ", ""), ",")

		resultFields := make([]string, 0, len(tmpFields)-1)
		resultValues := make([]string, 0, len(tmpValues)-1)

		for index, value := range tmpFields {
			if value != "organization_id" {
				resultFields = append(resultFields, value)
				resultValues = append(resultValues, tmpValues[index])
			}
		}
		fields = strings.Join(resultFields, ",")
		values = strings.Join(resultValues, ",")
	}
	return fields, values
}
