// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

type ComponentInfoComponentsItem struct {
	// Component name
	Name *string `json:"name,omitempty"`
	// Component version
	Version *string `json:"version,omitempty"`
	// Component status (running, stopped, error, not_installed)
	Status *string `json:"status,omitempty"`
	// Management type (unmanageable, manageable, kb-cluster)
	ManagementType *string `json:"management_type,omitempty"`
	// Number of replicas
	Replicas *int32 `json:"replicas,omitempty"`
	// Deployment location
	Location *string `json:"location,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewComponentInfoComponentsItem instantiates a new ComponentInfoComponentsItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewComponentInfoComponentsItem() *ComponentInfoComponentsItem {
	this := ComponentInfoComponentsItem{}
	return &this
}

// NewComponentInfoComponentsItemWithDefaults instantiates a new ComponentInfoComponentsItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewComponentInfoComponentsItemWithDefaults() *ComponentInfoComponentsItem {
	this := ComponentInfoComponentsItem{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ComponentInfoComponentsItem) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentInfoComponentsItem) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ComponentInfoComponentsItem) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ComponentInfoComponentsItem) SetName(v string) {
	o.Name = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ComponentInfoComponentsItem) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentInfoComponentsItem) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ComponentInfoComponentsItem) HasVersion() bool {
	return o != nil && o.Version != nil
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *ComponentInfoComponentsItem) SetVersion(v string) {
	o.Version = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ComponentInfoComponentsItem) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentInfoComponentsItem) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ComponentInfoComponentsItem) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ComponentInfoComponentsItem) SetStatus(v string) {
	o.Status = &v
}

// GetManagementType returns the ManagementType field value if set, zero value otherwise.
func (o *ComponentInfoComponentsItem) GetManagementType() string {
	if o == nil || o.ManagementType == nil {
		var ret string
		return ret
	}
	return *o.ManagementType
}

// GetManagementTypeOk returns a tuple with the ManagementType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentInfoComponentsItem) GetManagementTypeOk() (*string, bool) {
	if o == nil || o.ManagementType == nil {
		return nil, false
	}
	return o.ManagementType, true
}

// HasManagementType returns a boolean if a field has been set.
func (o *ComponentInfoComponentsItem) HasManagementType() bool {
	return o != nil && o.ManagementType != nil
}

// SetManagementType gets a reference to the given string and assigns it to the ManagementType field.
func (o *ComponentInfoComponentsItem) SetManagementType(v string) {
	o.ManagementType = &v
}

// GetReplicas returns the Replicas field value if set, zero value otherwise.
func (o *ComponentInfoComponentsItem) GetReplicas() int32 {
	if o == nil || o.Replicas == nil {
		var ret int32
		return ret
	}
	return *o.Replicas
}

// GetReplicasOk returns a tuple with the Replicas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentInfoComponentsItem) GetReplicasOk() (*int32, bool) {
	if o == nil || o.Replicas == nil {
		return nil, false
	}
	return o.Replicas, true
}

// HasReplicas returns a boolean if a field has been set.
func (o *ComponentInfoComponentsItem) HasReplicas() bool {
	return o != nil && o.Replicas != nil
}

// SetReplicas gets a reference to the given int32 and assigns it to the Replicas field.
func (o *ComponentInfoComponentsItem) SetReplicas(v int32) {
	o.Replicas = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *ComponentInfoComponentsItem) GetLocation() string {
	if o == nil || o.Location == nil {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentInfoComponentsItem) GetLocationOk() (*string, bool) {
	if o == nil || o.Location == nil {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *ComponentInfoComponentsItem) HasLocation() bool {
	return o != nil && o.Location != nil
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *ComponentInfoComponentsItem) SetLocation(v string) {
	o.Location = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ComponentInfoComponentsItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.ManagementType != nil {
		toSerialize["management_type"] = o.ManagementType
	}
	if o.Replicas != nil {
		toSerialize["replicas"] = o.Replicas
	}
	if o.Location != nil {
		toSerialize["location"] = o.Location
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ComponentInfoComponentsItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name           *string `json:"name,omitempty"`
		Version        *string `json:"version,omitempty"`
		Status         *string `json:"status,omitempty"`
		ManagementType *string `json:"management_type,omitempty"`
		Replicas       *int32  `json:"replicas,omitempty"`
		Location       *string `json:"location,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "version", "status", "management_type", "replicas", "location"})
	} else {
		return err
	}
	o.Name = all.Name
	o.Version = all.Version
	o.Status = all.Status
	o.ManagementType = all.ManagementType
	o.Replicas = all.Replicas
	o.Location = all.Location

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
