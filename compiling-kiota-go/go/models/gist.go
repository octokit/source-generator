package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Gist gist
type Gist struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The comments property
    comments *int32
    // The comments_url property
    comments_url *string
    // The commits_url property
    commits_url *string
    // The created_at property
    created_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The description property
    description *string
    // The files property
    files Gist_filesable
    // The forks_url property
    forks_url *string
    // The git_pull_url property
    git_pull_url *string
    // The git_push_url property
    git_push_url *string
    // The html_url property
    html_url *string
    // The id property
    id *string
    // The node_id property
    node_id *string
    // A GitHub user.
    owner NullableSimpleUserable
    // The public property
    public *bool
    // The truncated property
    truncated *bool
    // The updated_at property
    updated_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The url property
    url *string
    // A GitHub user.
    user NullableSimpleUserable
}
// NewGist instantiates a new Gist and sets the default values.
func NewGist()(*Gist) {
    m := &Gist{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateGistFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGistFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGist(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Gist) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetComments gets the comments property value. The comments property
func (m *Gist) GetComments()(*int32) {
    return m.comments
}
// GetCommentsUrl gets the comments_url property value. The comments_url property
func (m *Gist) GetCommentsUrl()(*string) {
    return m.comments_url
}
// GetCommitsUrl gets the commits_url property value. The commits_url property
func (m *Gist) GetCommitsUrl()(*string) {
    return m.commits_url
}
// GetCreatedAt gets the created_at property value. The created_at property
func (m *Gist) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.created_at
}
// GetDescription gets the description property value. The description property
func (m *Gist) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Gist) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["comments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComments(val)
        }
        return nil
    }
    res["comments_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCommentsUrl(val)
        }
        return nil
    }
    res["commits_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCommitsUrl(val)
        }
        return nil
    }
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
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["files"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGist_filesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFiles(val.(Gist_filesable))
        }
        return nil
    }
    res["forks_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetForksUrl(val)
        }
        return nil
    }
    res["git_pull_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGitPullUrl(val)
        }
        return nil
    }
    res["git_push_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGitPushUrl(val)
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
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["node_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNodeId(val)
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
    res["public"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublic(val)
        }
        return nil
    }
    res["truncated"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTruncated(val)
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
    res["user"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateNullableSimpleUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUser(val.(NullableSimpleUserable))
        }
        return nil
    }
    return res
}
// GetFiles gets the files property value. The files property
func (m *Gist) GetFiles()(Gist_filesable) {
    return m.files
}
// GetForksUrl gets the forks_url property value. The forks_url property
func (m *Gist) GetForksUrl()(*string) {
    return m.forks_url
}
// GetGitPullUrl gets the git_pull_url property value. The git_pull_url property
func (m *Gist) GetGitPullUrl()(*string) {
    return m.git_pull_url
}
// GetGitPushUrl gets the git_push_url property value. The git_push_url property
func (m *Gist) GetGitPushUrl()(*string) {
    return m.git_push_url
}
// GetHtmlUrl gets the html_url property value. The html_url property
func (m *Gist) GetHtmlUrl()(*string) {
    return m.html_url
}
// GetId gets the id property value. The id property
func (m *Gist) GetId()(*string) {
    return m.id
}
// GetNodeId gets the node_id property value. The node_id property
func (m *Gist) GetNodeId()(*string) {
    return m.node_id
}
// GetOwner gets the owner property value. A GitHub user.
func (m *Gist) GetOwner()(NullableSimpleUserable) {
    return m.owner
}
// GetPublic gets the public property value. The public property
func (m *Gist) GetPublic()(*bool) {
    return m.public
}
// GetTruncated gets the truncated property value. The truncated property
func (m *Gist) GetTruncated()(*bool) {
    return m.truncated
}
// GetUpdatedAt gets the updated_at property value. The updated_at property
func (m *Gist) GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.updated_at
}
// GetUrl gets the url property value. The url property
func (m *Gist) GetUrl()(*string) {
    return m.url
}
// GetUser gets the user property value. A GitHub user.
func (m *Gist) GetUser()(NullableSimpleUserable) {
    return m.user
}
// Serialize serializes information the current object
func (m *Gist) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("comments", m.GetComments())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("comments_url", m.GetCommentsUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("commits_url", m.GetCommitsUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("created_at", m.GetCreatedAt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("files", m.GetFiles())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("forks_url", m.GetForksUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("git_pull_url", m.GetGitPullUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("git_push_url", m.GetGitPushUrl())
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
        err := writer.WriteStringValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("node_id", m.GetNodeId())
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
    {
        err := writer.WriteBoolValue("public", m.GetPublic())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("truncated", m.GetTruncated())
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
        err := writer.WriteObjectValue("user", m.GetUser())
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
func (m *Gist) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetComments sets the comments property value. The comments property
func (m *Gist) SetComments(value *int32)() {
    m.comments = value
}
// SetCommentsUrl sets the comments_url property value. The comments_url property
func (m *Gist) SetCommentsUrl(value *string)() {
    m.comments_url = value
}
// SetCommitsUrl sets the commits_url property value. The commits_url property
func (m *Gist) SetCommitsUrl(value *string)() {
    m.commits_url = value
}
// SetCreatedAt sets the created_at property value. The created_at property
func (m *Gist) SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.created_at = value
}
// SetDescription sets the description property value. The description property
func (m *Gist) SetDescription(value *string)() {
    m.description = value
}
// SetFiles sets the files property value. The files property
func (m *Gist) SetFiles(value Gist_filesable)() {
    m.files = value
}
// SetForksUrl sets the forks_url property value. The forks_url property
func (m *Gist) SetForksUrl(value *string)() {
    m.forks_url = value
}
// SetGitPullUrl sets the git_pull_url property value. The git_pull_url property
func (m *Gist) SetGitPullUrl(value *string)() {
    m.git_pull_url = value
}
// SetGitPushUrl sets the git_push_url property value. The git_push_url property
func (m *Gist) SetGitPushUrl(value *string)() {
    m.git_push_url = value
}
// SetHtmlUrl sets the html_url property value. The html_url property
func (m *Gist) SetHtmlUrl(value *string)() {
    m.html_url = value
}
// SetId sets the id property value. The id property
func (m *Gist) SetId(value *string)() {
    m.id = value
}
// SetNodeId sets the node_id property value. The node_id property
func (m *Gist) SetNodeId(value *string)() {
    m.node_id = value
}
// SetOwner sets the owner property value. A GitHub user.
func (m *Gist) SetOwner(value NullableSimpleUserable)() {
    m.owner = value
}
// SetPublic sets the public property value. The public property
func (m *Gist) SetPublic(value *bool)() {
    m.public = value
}
// SetTruncated sets the truncated property value. The truncated property
func (m *Gist) SetTruncated(value *bool)() {
    m.truncated = value
}
// SetUpdatedAt sets the updated_at property value. The updated_at property
func (m *Gist) SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.updated_at = value
}
// SetUrl sets the url property value. The url property
func (m *Gist) SetUrl(value *string)() {
    m.url = value
}
// SetUser sets the user property value. A GitHub user.
func (m *Gist) SetUser(value NullableSimpleUserable)() {
    m.user = value
}
