package models
import (
    "errors"
)
// State of the required workflow
type RequiredWorkflow_state int

const (
    ACTIVE_REQUIREDWORKFLOW_STATE RequiredWorkflow_state = iota
    DELETED_REQUIREDWORKFLOW_STATE
)

func (i RequiredWorkflow_state) String() string {
    return []string{"active", "deleted"}[i]
}
func ParseRequiredWorkflow_state(v string) (any, error) {
    result := ACTIVE_REQUIREDWORKFLOW_STATE
    switch v {
        case "active":
            result = ACTIVE_REQUIREDWORKFLOW_STATE
        case "deleted":
            result = DELETED_REQUIREDWORKFLOW_STATE
        default:
            return 0, errors.New("Unknown RequiredWorkflow_state value: " + v)
    }
    return &result, nil
}
func SerializeRequiredWorkflow_state(values []RequiredWorkflow_state) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
