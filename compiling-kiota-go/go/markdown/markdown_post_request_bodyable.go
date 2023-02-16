package markdown

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MarkdownPostRequestBodyable 
type MarkdownPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetContext()(*string)
    GetMode()(*MarkdownPostRequestBody_mode)
    GetText()(*string)
    SetContext(value *string)()
    SetMode(value *MarkdownPostRequestBody_mode)()
    SetText(value *string)()
}
