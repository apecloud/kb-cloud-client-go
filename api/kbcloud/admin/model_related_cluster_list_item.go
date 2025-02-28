// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// RelatedClusterListItem the simplified cluster info for relation
type RelatedClusterListItem struct {
	// ID of the cluster
	Id string `json:"id"`
	// Name of the cluster
	Name string `json:"name"`
	// Type of the cluster
	ClusterDefinition string `json:"clusterDefinition"`
	Version           string `json:"version"`
	Mode              string `json:"mode"`
	Status            string `json:"status"`
	// Name of the environment
	EnvironmentName string `json:"environmentName"`
	// Resource of the cluster
	Resource string `json:"resource"`
	Storage  string `json:"storage"`
	// Time when the cluster was created
	CreatedAt time.Time `json:"createdAt"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRelatedClusterListItem instantiates a new RelatedClusterListItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRelatedClusterListItem(id string, name string, clusterDefinition string, version string, mode string, status string, environmentName string, resource string, storage string, createdAt time.Time) *RelatedClusterListItem {
	this := RelatedClusterListItem{}
	this.Id = id
	this.Name = name
	this.ClusterDefinition = clusterDefinition
	this.Version = version
	this.Mode = mode
	this.Status = status
	this.EnvironmentName = environmentName
	this.Resource = resource
	this.Storage = storage
	this.CreatedAt = createdAt
	return &this
}

// NewRelatedClusterListItemWithDefaults instantiates a new RelatedClusterListItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRelatedClusterListItemWithDefaults() *RelatedClusterListItem {
	this := RelatedClusterListItem{}
	return &this
}

// GetId returns the Id field value.
func (o *RelatedClusterListItem) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RelatedClusterListItem) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *RelatedClusterListItem) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value.
func (o *RelatedClusterListItem) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RelatedClusterListItem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *RelatedClusterListItem) SetName(v string) {
	o.Name = v
}

// GetClusterDefinition returns the ClusterDefinition field value.
func (o *RelatedClusterListItem) GetClusterDefinition() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ClusterDefinition
}

// GetClusterDefinitionOk returns a tuple with the ClusterDefinition field value
// and a boolean to check if the value has been set.
func (o *RelatedClusterListItem) GetClusterDefinitionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClusterDefinition, true
}

// SetClusterDefinition sets field value.
func (o *RelatedClusterListItem) SetClusterDefinition(v string) {
	o.ClusterDefinition = v
}

// GetVersion returns the Version field value.
func (o *RelatedClusterListItem) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *RelatedClusterListItem) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value.
func (o *RelatedClusterListItem) SetVersion(v string) {
	o.Version = v
}

// GetMode returns the Mode field value.
func (o *RelatedClusterListItem) GetMode() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *RelatedClusterListItem) GetModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value.
func (o *RelatedClusterListItem) SetMode(v string) {
	o.Mode = v
}

// GetStatus returns the Status field value.
func (o *RelatedClusterListItem) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *RelatedClusterListItem) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *RelatedClusterListItem) SetStatus(v string) {
	o.Status = v
}

// GetEnvironmentName returns the EnvironmentName field value.
func (o *RelatedClusterListItem) GetEnvironmentName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.EnvironmentName
}

// GetEnvironmentNameOk returns a tuple with the EnvironmentName field value
// and a boolean to check if the value has been set.
func (o *RelatedClusterListItem) GetEnvironmentNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnvironmentName, true
}

// SetEnvironmentName sets field value.
func (o *RelatedClusterListItem) SetEnvironmentName(v string) {
	o.EnvironmentName = v
}

// GetResource returns the Resource field value.
func (o *RelatedClusterListItem) GetResource() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Resource
}

// GetResourceOk returns a tuple with the Resource field value
// and a boolean to check if the value has been set.
func (o *RelatedClusterListItem) GetResourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Resource, true
}

// SetResource sets field value.
func (o *RelatedClusterListItem) SetResource(v string) {
	o.Resource = v
}

// GetStorage returns the Storage field value.
func (o *RelatedClusterListItem) GetStorage() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Storage
}

// GetStorageOk returns a tuple with the Storage field value
// and a boolean to check if the value has been set.
func (o *RelatedClusterListItem) GetStorageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Storage, true
}

// SetStorage sets field value.
func (o *RelatedClusterListItem) SetStorage(v string) {
	o.Storage = v
}

// GetCreatedAt returns the CreatedAt field value.
func (o *RelatedClusterListItem) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *RelatedClusterListItem) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *RelatedClusterListItem) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RelatedClusterListItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["clusterDefinition"] = o.ClusterDefinition
	toSerialize["version"] = o.Version
	toSerialize["mode"] = o.Mode
	toSerialize["status"] = o.Status
	toSerialize["environmentName"] = o.EnvironmentName
	toSerialize["resource"] = o.Resource
	toSerialize["storage"] = o.Storage
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RelatedClusterListItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id                *string    `json:"id"`
		Name              *string    `json:"name"`
		ClusterDefinition *string    `json:"clusterDefinition"`
		Version           *string    `json:"version"`
		Mode              *string    `json:"mode"`
		Status            *string    `json:"status"`
		EnvironmentName   *string    `json:"environmentName"`
		Resource          *string    `json:"resource"`
		Storage           *string    `json:"storage"`
		CreatedAt         *time.Time `json:"createdAt"`
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
	if all.ClusterDefinition == nil {
		return fmt.Errorf("required field clusterDefinition missing")
	}
	if all.Version == nil {
		return fmt.Errorf("required field version missing")
	}
	if all.Mode == nil {
		return fmt.Errorf("required field mode missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	if all.EnvironmentName == nil {
		return fmt.Errorf("required field environmentName missing")
	}
	if all.Resource == nil {
		return fmt.Errorf("required field resource missing")
	}
	if all.Storage == nil {
		return fmt.Errorf("required field storage missing")
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field createdAt missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"id", "name", "clusterDefinition", "version", "mode", "status", "environmentName", "resource", "storage", "createdAt"})
	} else {
		return err
	}
	o.Id = *all.Id
	o.Name = *all.Name
	o.ClusterDefinition = *all.ClusterDefinition
	o.Version = *all.Version
	o.Mode = *all.Mode
	o.Status = *all.Status
	o.EnvironmentName = *all.EnvironmentName
	o.Resource = *all.Resource
	o.Storage = *all.Storage
	o.CreatedAt = *all.CreatedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
