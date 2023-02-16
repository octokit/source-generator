package billing
import (
    "errors"
)
// Which users can access codespaces in the organization. `disabled` means that no users can access codespaces in the organization.
type BillingPutRequestBody_visibility int

const (
    DISABLED_BILLINGPUTREQUESTBODY_VISIBILITY BillingPutRequestBody_visibility = iota
    SELECTED_MEMBERS_BILLINGPUTREQUESTBODY_VISIBILITY
    ALL_MEMBERS_BILLINGPUTREQUESTBODY_VISIBILITY
    ALL_MEMBERS_AND_OUTSIDE_COLLABORATORS_BILLINGPUTREQUESTBODY_VISIBILITY
)

func (i BillingPutRequestBody_visibility) String() string {
    return []string{"disabled", "selected_members", "all_members", "all_members_and_outside_collaborators"}[i]
}
func ParseBillingPutRequestBody_visibility(v string) (any, error) {
    result := DISABLED_BILLINGPUTREQUESTBODY_VISIBILITY
    switch v {
        case "disabled":
            result = DISABLED_BILLINGPUTREQUESTBODY_VISIBILITY
        case "selected_members":
            result = SELECTED_MEMBERS_BILLINGPUTREQUESTBODY_VISIBILITY
        case "all_members":
            result = ALL_MEMBERS_BILLINGPUTREQUESTBODY_VISIBILITY
        case "all_members_and_outside_collaborators":
            result = ALL_MEMBERS_AND_OUTSIDE_COLLABORATORS_BILLINGPUTREQUESTBODY_VISIBILITY
        default:
            return 0, errors.New("Unknown BillingPutRequestBody_visibility value: " + v)
    }
    return &result, nil
}
func SerializeBillingPutRequestBody_visibility(values []BillingPutRequestBody_visibility) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
