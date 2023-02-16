package required_workflows
import (
    "errors"
)
// Enable the required workflow for all repositories or selected repositories in the organization.
type Required_workflowsPostRequestBody_scope int

const (
    SELECTED_REQUIRED_WORKFLOWSPOSTREQUESTBODY_SCOPE Required_workflowsPostRequestBody_scope = iota
    ALL_REQUIRED_WORKFLOWSPOSTREQUESTBODY_SCOPE
)

func (i Required_workflowsPostRequestBody_scope) String() string {
    return []string{"selected", "all"}[i]
}
func ParseRequired_workflowsPostRequestBody_scope(v string) (any, error) {
    result := SELECTED_REQUIRED_WORKFLOWSPOSTREQUESTBODY_SCOPE
    switch v {
        case "selected":
            result = SELECTED_REQUIRED_WORKFLOWSPOSTREQUESTBODY_SCOPE
        case "all":
            result = ALL_REQUIRED_WORKFLOWSPOSTREQUESTBODY_SCOPE
        default:
            return 0, errors.New("Unknown Required_workflowsPostRequestBody_scope value: " + v)
    }
    return &result, nil
}
func SerializeRequired_workflowsPostRequestBody_scope(values []Required_workflowsPostRequestBody_scope) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
