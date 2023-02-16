package statuses
import (
    "errors"
)
// Name for the target deployment environment, which can be changed when setting a deploy status. For example, `production`, `staging`, or `qa`.
type StatusesPostRequestBody_environment int

const (
    PRODUCTION_STATUSESPOSTREQUESTBODY_ENVIRONMENT StatusesPostRequestBody_environment = iota
    STAGING_STATUSESPOSTREQUESTBODY_ENVIRONMENT
    QA_STATUSESPOSTREQUESTBODY_ENVIRONMENT
)

func (i StatusesPostRequestBody_environment) String() string {
    return []string{"production", "staging", "qa"}[i]
}
func ParseStatusesPostRequestBody_environment(v string) (any, error) {
    result := PRODUCTION_STATUSESPOSTREQUESTBODY_ENVIRONMENT
    switch v {
        case "production":
            result = PRODUCTION_STATUSESPOSTREQUESTBODY_ENVIRONMENT
        case "staging":
            result = STAGING_STATUSESPOSTREQUESTBODY_ENVIRONMENT
        case "qa":
            result = QA_STATUSESPOSTREQUESTBODY_ENVIRONMENT
        default:
            return 0, errors.New("Unknown StatusesPostRequestBody_environment value: " + v)
    }
    return &result, nil
}
func SerializeStatusesPostRequestBody_environment(values []StatusesPostRequestBody_environment) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
