package models
import (
    "errors"
)
// The level of permission to grant the access token to view and manage announcement banners for a repository.
type AppPermissions_repository_announcement_banners int

const (
    READ_APPPERMISSIONS_REPOSITORY_ANNOUNCEMENT_BANNERS AppPermissions_repository_announcement_banners = iota
    WRITE_APPPERMISSIONS_REPOSITORY_ANNOUNCEMENT_BANNERS
)

func (i AppPermissions_repository_announcement_banners) String() string {
    return []string{"read", "write"}[i]
}
func ParseAppPermissions_repository_announcement_banners(v string) (any, error) {
    result := READ_APPPERMISSIONS_REPOSITORY_ANNOUNCEMENT_BANNERS
    switch v {
        case "read":
            result = READ_APPPERMISSIONS_REPOSITORY_ANNOUNCEMENT_BANNERS
        case "write":
            result = WRITE_APPPERMISSIONS_REPOSITORY_ANNOUNCEMENT_BANNERS
        default:
            return 0, errors.New("Unknown AppPermissions_repository_announcement_banners value: " + v)
    }
    return &result, nil
}
func SerializeAppPermissions_repository_announcement_banners(values []AppPermissions_repository_announcement_banners) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
