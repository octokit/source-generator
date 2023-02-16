package models
import (
    "errors"
)
// 
type PackageVersionMetadata_package_type int

const (
    NPM_PACKAGEVERSIONMETADATA_PACKAGE_TYPE PackageVersionMetadata_package_type = iota
    MAVEN_PACKAGEVERSIONMETADATA_PACKAGE_TYPE
    RUBYGEMS_PACKAGEVERSIONMETADATA_PACKAGE_TYPE
    DOCKER_PACKAGEVERSIONMETADATA_PACKAGE_TYPE
    NUGET_PACKAGEVERSIONMETADATA_PACKAGE_TYPE
    CONTAINER_PACKAGEVERSIONMETADATA_PACKAGE_TYPE
)

func (i PackageVersionMetadata_package_type) String() string {
    return []string{"npm", "maven", "rubygems", "docker", "nuget", "container"}[i]
}
func ParsePackageVersionMetadata_package_type(v string) (any, error) {
    result := NPM_PACKAGEVERSIONMETADATA_PACKAGE_TYPE
    switch v {
        case "npm":
            result = NPM_PACKAGEVERSIONMETADATA_PACKAGE_TYPE
        case "maven":
            result = MAVEN_PACKAGEVERSIONMETADATA_PACKAGE_TYPE
        case "rubygems":
            result = RUBYGEMS_PACKAGEVERSIONMETADATA_PACKAGE_TYPE
        case "docker":
            result = DOCKER_PACKAGEVERSIONMETADATA_PACKAGE_TYPE
        case "nuget":
            result = NUGET_PACKAGEVERSIONMETADATA_PACKAGE_TYPE
        case "container":
            result = CONTAINER_PACKAGEVERSIONMETADATA_PACKAGE_TYPE
        default:
            return 0, errors.New("Unknown PackageVersionMetadata_package_type value: " + v)
    }
    return &result, nil
}
func SerializePackageVersionMetadata_package_type(values []PackageVersionMetadata_package_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
