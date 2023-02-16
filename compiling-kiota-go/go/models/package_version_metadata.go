package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PackageVersionMetadata 
type PackageVersionMetadata struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The container property
    container ContainerMetadataable
    // The docker property
    docker DockerMetadataable
    // The package_type property
    package_type *PackageVersionMetadata_package_type
}
// NewPackageVersionMetadata instantiates a new PackageVersionMetadata and sets the default values.
func NewPackageVersionMetadata()(*PackageVersionMetadata) {
    m := &PackageVersionMetadata{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePackageVersionMetadataFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePackageVersionMetadataFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPackageVersionMetadata(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PackageVersionMetadata) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetContainer gets the container property value. The container property
func (m *PackageVersionMetadata) GetContainer()(ContainerMetadataable) {
    return m.container
}
// GetDocker gets the docker property value. The docker property
func (m *PackageVersionMetadata) GetDocker()(DockerMetadataable) {
    return m.docker
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PackageVersionMetadata) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["container"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateContainerMetadataFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContainer(val.(ContainerMetadataable))
        }
        return nil
    }
    res["docker"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDockerMetadataFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDocker(val.(DockerMetadataable))
        }
        return nil
    }
    res["package_type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePackageVersionMetadata_package_type)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPackageType(val.(*PackageVersionMetadata_package_type))
        }
        return nil
    }
    return res
}
// GetPackageType gets the package_type property value. The package_type property
func (m *PackageVersionMetadata) GetPackageType()(*PackageVersionMetadata_package_type) {
    return m.package_type
}
// Serialize serializes information the current object
func (m *PackageVersionMetadata) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("container", m.GetContainer())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("docker", m.GetDocker())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PackageVersionMetadata) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetContainer sets the container property value. The container property
func (m *PackageVersionMetadata) SetContainer(value ContainerMetadataable)() {
    m.container = value
}
// SetDocker sets the docker property value. The docker property
func (m *PackageVersionMetadata) SetDocker(value DockerMetadataable)() {
    m.docker = value
}
// SetPackageType sets the package_type property value. The package_type property
func (m *PackageVersionMetadata) SetPackageType(value *PackageVersionMetadata_package_type)() {
    m.package_type = value
}
