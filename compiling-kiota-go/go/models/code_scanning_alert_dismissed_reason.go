package models
import (
    "errors"
)
// **Required when the state is dismissed.** The reason for dismissing or closing the alert.
type CodeScanningAlertDismissedReason int

const (
    FALSEPOSITIVE_CODESCANNINGALERTDISMISSEDREASON CodeScanningAlertDismissedReason = iota
    WONTFIX_CODESCANNINGALERTDISMISSEDREASON
    USEDINTESTS_CODESCANNINGALERTDISMISSEDREASON
)

func (i CodeScanningAlertDismissedReason) String() string {
    return []string{"falsePositive", "wonTFix", "usedInTests"}[i]
}
func ParseCodeScanningAlertDismissedReason(v string) (any, error) {
    result := FALSEPOSITIVE_CODESCANNINGALERTDISMISSEDREASON
    switch v {
        case "falsePositive":
            result = FALSEPOSITIVE_CODESCANNINGALERTDISMISSEDREASON
        case "wonTFix":
            result = WONTFIX_CODESCANNINGALERTDISMISSEDREASON
        case "usedInTests":
            result = USEDINTESTS_CODESCANNINGALERTDISMISSEDREASON
        default:
            return 0, errors.New("Unknown CodeScanningAlertDismissedReason value: " + v)
    }
    return &result, nil
}
func SerializeCodeScanningAlertDismissedReason(values []CodeScanningAlertDismissedReason) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
