package models
import (
    "errors"
)
// Scope of the required workflow
type RequiredWorkflow_scope int

const (
    ALL_REQUIREDWORKFLOW_SCOPE RequiredWorkflow_scope = iota
    SELECTED_REQUIREDWORKFLOW_SCOPE
)

func (i RequiredWorkflow_scope) String() string {
    return []string{"all", "selected"}[i]
}
func ParseRequiredWorkflow_scope(v string) (any, error) {
    result := ALL_REQUIREDWORKFLOW_SCOPE
    switch v {
        case "all":
            result = ALL_REQUIREDWORKFLOW_SCOPE
        case "selected":
            result = SELECTED_REQUIREDWORKFLOW_SCOPE
        default:
            return 0, errors.New("Unknown RequiredWorkflow_scope value: " + v)
    }
    return &result, nil
}
func SerializeRequiredWorkflow_scope(values []RequiredWorkflow_scope) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
