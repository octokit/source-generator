package models
import (
    "errors"
)
// 
type RepoRequiredWorkflow_state int

const (
    ACTIVE_REPOREQUIREDWORKFLOW_STATE RepoRequiredWorkflow_state = iota
    DELETED_REPOREQUIREDWORKFLOW_STATE
)

func (i RepoRequiredWorkflow_state) String() string {
    return []string{"active", "deleted"}[i]
}
func ParseRepoRequiredWorkflow_state(v string) (any, error) {
    result := ACTIVE_REPOREQUIREDWORKFLOW_STATE
    switch v {
        case "active":
            result = ACTIVE_REPOREQUIREDWORKFLOW_STATE
        case "deleted":
            result = DELETED_REPOREQUIREDWORKFLOW_STATE
        default:
            return 0, errors.New("Unknown RepoRequiredWorkflow_state value: " + v)
    }
    return &result, nil
}
func SerializeRepoRequiredWorkflow_state(values []RepoRequiredWorkflow_state) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
