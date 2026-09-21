// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type ClusterJarPackage struct {
	Id        string              `json:"id"`
	Name      string              `json:"name"`
	Kind      ClusterJarKind      `json:"kind"`
	Filename  string              `json:"filename"`
	Sha256    string              `json:"sha256"`
	Size      int64               `json:"size"`
	Uri       string              `json:"uri"`
	Published bool                `json:"published"`
	Archived  bool                `json:"archived"`
	Status    string              `json:"status"`
	Synced    int64               `json:"synced"`
	Total     int64               `json:"total"`
	CreatedBy string              `json:"createdBy"`
	CreatedAt time.Time           `json:"createdAt"`
	Instances []ClusterJarReplica `json:"instances"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewClusterJarPackage instantiates a new ClusterJarPackage object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewClusterJarPackage(id string, name string, kind ClusterJarKind, filename string, sha256 string, size int64, uri string, published bool, archived bool, status string, synced int64, total int64, createdBy string, createdAt time.Time, instances []ClusterJarReplica) *ClusterJarPackage {
	this := ClusterJarPackage{}
	this.Id = id
	this.Name = name
	this.Kind = kind
	this.Filename = filename
	this.Sha256 = sha256
	this.Size = size
	this.Uri = uri
	this.Published = published
	this.Archived = archived
	this.Status = status
	this.Synced = synced
	this.Total = total
	this.CreatedBy = createdBy
	this.CreatedAt = createdAt
	this.Instances = instances
	return &this
}

// NewClusterJarPackageWithDefaults instantiates a new ClusterJarPackage object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewClusterJarPackageWithDefaults() *ClusterJarPackage {
	this := ClusterJarPackage{}
	return &this
}

// GetId returns the Id field value.
func (o *ClusterJarPackage) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ClusterJarPackage) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ClusterJarPackage) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value.
func (o *ClusterJarPackage) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ClusterJarPackage) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ClusterJarPackage) SetName(v string) {
	o.Name = v
}

// GetKind returns the Kind field value.
func (o *ClusterJarPackage) GetKind() ClusterJarKind {
	if o == nil {
		var ret ClusterJarKind
		return ret
	}
	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *ClusterJarPackage) GetKindOk() (*ClusterJarKind, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value.
func (o *ClusterJarPackage) SetKind(v ClusterJarKind) {
	o.Kind = v
}

// GetFilename returns the Filename field value.
func (o *ClusterJarPackage) GetFilename() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Filename
}

// GetFilenameOk returns a tuple with the Filename field value
// and a boolean to check if the value has been set.
func (o *ClusterJarPackage) GetFilenameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Filename, true
}

// SetFilename sets field value.
func (o *ClusterJarPackage) SetFilename(v string) {
	o.Filename = v
}

// GetSha256 returns the Sha256 field value.
func (o *ClusterJarPackage) GetSha256() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Sha256
}

// GetSha256Ok returns a tuple with the Sha256 field value
// and a boolean to check if the value has been set.
func (o *ClusterJarPackage) GetSha256Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sha256, true
}

// SetSha256 sets field value.
func (o *ClusterJarPackage) SetSha256(v string) {
	o.Sha256 = v
}

// GetSize returns the Size field value.
func (o *ClusterJarPackage) GetSize() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *ClusterJarPackage) GetSizeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value.
func (o *ClusterJarPackage) SetSize(v int64) {
	o.Size = v
}

// GetUri returns the Uri field value.
func (o *ClusterJarPackage) GetUri() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Uri
}

// GetUriOk returns a tuple with the Uri field value
// and a boolean to check if the value has been set.
func (o *ClusterJarPackage) GetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uri, true
}

// SetUri sets field value.
func (o *ClusterJarPackage) SetUri(v string) {
	o.Uri = v
}

// GetPublished returns the Published field value.
func (o *ClusterJarPackage) GetPublished() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Published
}

// GetPublishedOk returns a tuple with the Published field value
// and a boolean to check if the value has been set.
func (o *ClusterJarPackage) GetPublishedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Published, true
}

// SetPublished sets field value.
func (o *ClusterJarPackage) SetPublished(v bool) {
	o.Published = v
}

// GetArchived returns the Archived field value.
func (o *ClusterJarPackage) GetArchived() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Archived
}

// GetArchivedOk returns a tuple with the Archived field value
// and a boolean to check if the value has been set.
func (o *ClusterJarPackage) GetArchivedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Archived, true
}

// SetArchived sets field value.
func (o *ClusterJarPackage) SetArchived(v bool) {
	o.Archived = v
}

// GetStatus returns the Status field value.
func (o *ClusterJarPackage) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ClusterJarPackage) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *ClusterJarPackage) SetStatus(v string) {
	o.Status = v
}

// GetSynced returns the Synced field value.
func (o *ClusterJarPackage) GetSynced() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Synced
}

// GetSyncedOk returns a tuple with the Synced field value
// and a boolean to check if the value has been set.
func (o *ClusterJarPackage) GetSyncedOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Synced, true
}

// SetSynced sets field value.
func (o *ClusterJarPackage) SetSynced(v int64) {
	o.Synced = v
}

// GetTotal returns the Total field value.
func (o *ClusterJarPackage) GetTotal() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *ClusterJarPackage) GetTotalOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value.
func (o *ClusterJarPackage) SetTotal(v int64) {
	o.Total = v
}

// GetCreatedBy returns the CreatedBy field value.
func (o *ClusterJarPackage) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *ClusterJarPackage) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value.
func (o *ClusterJarPackage) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetCreatedAt returns the CreatedAt field value.
func (o *ClusterJarPackage) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ClusterJarPackage) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *ClusterJarPackage) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetInstances returns the Instances field value.
func (o *ClusterJarPackage) GetInstances() []ClusterJarReplica {
	if o == nil {
		var ret []ClusterJarReplica
		return ret
	}
	return o.Instances
}

// GetInstancesOk returns a tuple with the Instances field value
// and a boolean to check if the value has been set.
func (o *ClusterJarPackage) GetInstancesOk() (*[]ClusterJarReplica, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Instances, true
}

// SetInstances sets field value.
func (o *ClusterJarPackage) SetInstances(v []ClusterJarReplica) {
	o.Instances = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ClusterJarPackage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["kind"] = o.Kind
	toSerialize["filename"] = o.Filename
	toSerialize["sha256"] = o.Sha256
	toSerialize["size"] = o.Size
	toSerialize["uri"] = o.Uri
	toSerialize["published"] = o.Published
	toSerialize["archived"] = o.Archived
	toSerialize["status"] = o.Status
	toSerialize["synced"] = o.Synced
	toSerialize["total"] = o.Total
	toSerialize["createdBy"] = o.CreatedBy
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["instances"] = o.Instances

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ClusterJarPackage) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id        *string              `json:"id"`
		Name      *string              `json:"name"`
		Kind      *ClusterJarKind      `json:"kind"`
		Filename  *string              `json:"filename"`
		Sha256    *string              `json:"sha256"`
		Size      *int64               `json:"size"`
		Uri       *string              `json:"uri"`
		Published *bool                `json:"published"`
		Archived  *bool                `json:"archived"`
		Status    *string              `json:"status"`
		Synced    *int64               `json:"synced"`
		Total     *int64               `json:"total"`
		CreatedBy *string              `json:"createdBy"`
		CreatedAt *time.Time           `json:"createdAt"`
		Instances *[]ClusterJarReplica `json:"instances"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Kind == nil {
		return fmt.Errorf("required field kind missing")
	}
	if all.Filename == nil {
		return fmt.Errorf("required field filename missing")
	}
	if all.Sha256 == nil {
		return fmt.Errorf("required field sha256 missing")
	}
	if all.Size == nil {
		return fmt.Errorf("required field size missing")
	}
	if all.Uri == nil {
		return fmt.Errorf("required field uri missing")
	}
	if all.Published == nil {
		return fmt.Errorf("required field published missing")
	}
	if all.Archived == nil {
		return fmt.Errorf("required field archived missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	if all.Synced == nil {
		return fmt.Errorf("required field synced missing")
	}
	if all.Total == nil {
		return fmt.Errorf("required field total missing")
	}
	if all.CreatedBy == nil {
		return fmt.Errorf("required field createdBy missing")
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field createdAt missing")
	}
	if all.Instances == nil {
		return fmt.Errorf("required field instances missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"id", "name", "kind", "filename", "sha256", "size", "uri", "published", "archived", "status", "synced", "total", "createdBy", "createdAt", "instances"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Id = *all.Id
	o.Name = *all.Name
	if !all.Kind.IsValid() {
		hasInvalidField = true
	} else {
		o.Kind = *all.Kind
	}
	o.Filename = *all.Filename
	o.Sha256 = *all.Sha256
	o.Size = *all.Size
	o.Uri = *all.Uri
	o.Published = *all.Published
	o.Archived = *all.Archived
	o.Status = *all.Status
	o.Synced = *all.Synced
	o.Total = *all.Total
	o.CreatedBy = *all.CreatedBy
	o.CreatedAt = *all.CreatedAt
	o.Instances = *all.Instances

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
