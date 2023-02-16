package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Package a software package
type Package struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The created_at property
    created_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The html_url property
    html_url *string
    // Unique identifier of the package.
    id *int32
    // The name of the package.
    name *string
    // A GitHub user.
    owner NullableSimpleUserable
    // The package_type property
    package_type *Package_package_type
    // Minimal Repository
    repository NullableMinimalRepositoryable
    // The updated_at property
    updated_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The url property
    url *string
    // The number of versions of the package.
    version_count *int32
    // The visibility property
    visibility *Package_visibility
}
// NewPackage instantiates a new Package and sets the default values.
func NewPackage()(*Package) {
    m := &Package{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePackageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePackageFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPackage(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Package) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCreatedAt gets the created_at property value. The created_at property
func (m *Package) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.created_at
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Package) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["created_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedAt(val)
        }
        return nil
    }
    res["html_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHtmlUrl(val)
        }
        return nil
    }
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["owner"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateNullableSimpleUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwner(val.(NullableSimpleUserable))
        }
        return nil
    }
    res["package_type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePackage_package_type)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPackageType(val.(*Package_package_type))
        }
        return nil
    }
    res["repository"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateNullableMinimalRepositoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRepository(val.(NullableMinimalRepositoryable))
        }
        return nil
    }
    res["updated_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdatedAt(val)
        }
        return nil
    }
    res["url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUrl(val)
        }
        return nil
    }
    res["version_count"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersionCount(val)
        }
        return nil
    }
    res["visibility"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePackage_visibility)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVisibility(val.(*Package_visibility))
        }
        return nil
    }
    return res
}
// GetHtmlUrl gets the html_url property value. The html_url property
func (m *Package) GetHtmlUrl()(*string) {
    return m.html_url
}
// GetId gets the id property value. Unique identifier of the package.
func (m *Package) GetId()(*int32) {
    return m.id
}
// GetName gets the name property value. The name of the package.
func (m *Package) GetName()(*string) {
    return m.name
}
// GetOwner gets the owner property value. A GitHub user.
func (m *Package) GetOwner()(NullableSimpleUserable) {
    return m.owner
}
// GetPackageType gets the package_type property value. The package_type property
func (m *Package) GetPackageType()(*Package_package_type) {
    return m.package_type
}
// GetRepository gets the repository property value. Minimal Repository
func (m *Package) GetRepository()(NullableMinimalRepositoryable) {
    return m.repository
}
// GetUpdatedAt gets the updated_at property value. The updated_at property
func (m *Package) GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.updated_at
}
// GetUrl gets the url property value. The url property
func (m *Package) GetUrl()(*string) {
    return m.url
}
// GetVersionCount gets the version_count property value. The number of versions of the package.
func (m *Package) GetVersionCount()(*int32) {
    return m.version_count
}
// GetVisibility gets the visibility property value. The visibility property
func (m *Package) GetVisibility()(*Package_visibility) {
    return m.visibility
}
// Serialize serializes information the current object
func (m *Package) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteTimeValue("created_at", m.GetCreatedAt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("html_url", m.GetHtmlUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("owner", m.GetOwner())
        if err != nil {
            return err
        }
    }
    if m.GetPackageType() != nil {
        cast := (*m.GetPackageType()).String()
        err := writer.WriteStringValue("package_type", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("repository", m.GetRepository())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("updated_at", m.GetUpdatedAt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("url", m.GetUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("version_count", m.GetVersionCount())
        if err != nil {
            return err
        }
    }
    if m.GetVisibility() != nil {
        cast := (*m.GetVisibility()).String()
        err := writer.WriteStringValue("visibility", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Package) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCreatedAt sets the created_at property value. The created_at property
func (m *Package) SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.created_at = value
}
// SetHtmlUrl sets the html_url property value. The html_url property
func (m *Package) SetHtmlUrl(value *string)() {
    m.html_url = value
}
// SetId sets the id property value. Unique identifier of the package.
func (m *Package) SetId(value *int32)() {
    m.id = value
}
// SetName sets the name property value. The name of the package.
func (m *Package) SetName(value *string)() {
    m.name = value
}
// SetOwner sets the owner property value. A GitHub user.
func (m *Package) SetOwner(value NullableSimpleUserable)() {
    m.owner = value
}
// SetPackageType sets the package_type property value. The package_type property
func (m *Package) SetPackageType(value *Package_package_type)() {
    m.package_type = value
}
// SetRepository sets the repository property value. Minimal Repository
func (m *Package) SetRepository(value NullableMinimalRepositoryable)() {
    m.repository = value
}
// SetUpdatedAt sets the updated_at property value. The updated_at property
func (m *Package) SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.updated_at = value
}
// SetUrl sets the url property value. The url property
func (m *Package) SetUrl(value *string)() {
    m.url = value
}
// SetVersionCount sets the version_count property value. The number of versions of the package.
func (m *Package) SetVersionCount(value *int32)() {
    m.version_count = value
}
// SetVisibility sets the visibility property value. The visibility property
func (m *Package) SetVisibility(value *Package_visibility)() {
    m.visibility = value
}
