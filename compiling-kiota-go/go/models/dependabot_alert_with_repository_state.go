package models
import (
    "errors"
)
// The state of the Dependabot alert.
type DependabotAlertWithRepository_state int

const (
    DISMISSED_DEPENDABOTALERTWITHREPOSITORY_STATE DependabotAlertWithRepository_state = iota
    FIXED_DEPENDABOTALERTWITHREPOSITORY_STATE
    OPEN_DEPENDABOTALERTWITHREPOSITORY_STATE
)

func (i DependabotAlertWithRepository_state) String() string {
    return []string{"dismissed", "fixed", "open"}[i]
}
func ParseDependabotAlertWithRepository_state(v string) (any, error) {
    result := DISMISSED_DEPENDABOTALERTWITHREPOSITORY_STATE
    switch v {
        case "dismissed":
            result = DISMISSED_DEPENDABOTALERTWITHREPOSITORY_STATE
        case "fixed":
            result = FIXED_DEPENDABOTALERTWITHREPOSITORY_STATE
        case "open":
            result = OPEN_DEPENDABOTALERTWITHREPOSITORY_STATE
        default:
            return 0, errors.New("Unknown DependabotAlertWithRepository_state value: " + v)
    }
    return &result, nil
}
func SerializeDependabotAlertWithRepository_state(values []DependabotAlertWithRepository_state) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
