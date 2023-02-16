package item
import (
    "errors"
)
// Enable the required workflow for all repositories or selected repositories in the organization.
type WithRequired_workflow_PatchRequestBody_scope int

const (
    SELECTED_WITHREQUIRED_WORKFLOW_PATCHREQUESTBODY_SCOPE WithRequired_workflow_PatchRequestBody_scope = iota
    ALL_WITHREQUIRED_WORKFLOW_PATCHREQUESTBODY_SCOPE
)

func (i WithRequired_workflow_PatchRequestBody_scope) String() string {
    return []string{"selected", "all"}[i]
}
func ParseWithRequired_workflow_PatchRequestBody_scope(v string) (any, error) {
    result := SELECTED_WITHREQUIRED_WORKFLOW_PATCHREQUESTBODY_SCOPE
    switch v {
        case "selected":
            result = SELECTED_WITHREQUIRED_WORKFLOW_PATCHREQUESTBODY_SCOPE
        case "all":
            result = ALL_WITHREQUIRED_WORKFLOW_PATCHREQUESTBODY_SCOPE
        default:
            return 0, errors.New("Unknown WithRequired_workflow_PatchRequestBody_scope value: " + v)
    }
    return &result, nil
}
func SerializeWithRequired_workflow_PatchRequestBody_scope(values []WithRequired_workflow_PatchRequestBody_scope) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
