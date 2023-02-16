package pages
import (
    "errors"
)
// The repository directory that includes the source files for the Pages site. Allowed paths are `/` or `/docs`.
type PagesMember1_path int

const (
    DOCS_PAGESMEMBER1_PATH PagesMember1_path = iota
)

func (i PagesMember1_path) String() string {
    return []string{"Docs"}[i]
}
func ParsePagesMember1_path(v string) (any, error) {
    result := DOCS_PAGESMEMBER1_PATH
    switch v {
        case "Docs":
            result = DOCS_PAGESMEMBER1_PATH
        default:
            return 0, errors.New("Unknown PagesMember1_path value: " + v)
    }
    return &result, nil
}
func SerializePagesMember1_path(values []PagesMember1_path) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
