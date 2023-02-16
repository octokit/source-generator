package pages
import (
    "errors"
)
// The repository directory that includes the source files for the Pages site. Allowed paths are `/` or `/docs`. Default: `/`
type PagesPostRequestBody_source_path int

const (
    DOCS_PAGESPOSTREQUESTBODY_SOURCE_PATH PagesPostRequestBody_source_path = iota
)

func (i PagesPostRequestBody_source_path) String() string {
    return []string{"Docs"}[i]
}
func ParsePagesPostRequestBody_source_path(v string) (any, error) {
    result := DOCS_PAGESPOSTREQUESTBODY_SOURCE_PATH
    switch v {
        case "Docs":
            result = DOCS_PAGESPOSTREQUESTBODY_SOURCE_PATH
        default:
            return 0, errors.New("Unknown PagesPostRequestBody_source_path value: " + v)
    }
    return &result, nil
}
func SerializePagesPostRequestBody_source_path(values []PagesPostRequestBody_source_path) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
