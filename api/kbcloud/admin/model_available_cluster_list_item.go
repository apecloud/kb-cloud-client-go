// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

type AvailableClusterListItem struct {
	// ID of the cluster
	Id *string `json:"id,omitempty"`
	// Name of the cluster
	Name *string `json:"name,omitempty"`
	// Display name of the cluster
	DisplayName *string `json:"displayName,omitempty"`
	ClusterType *string `json:"clusterType,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAvailableClusterListItem instantiates a new AvailableClusterListItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAvailableClusterListItem() *AvailableClusterListItem {
	this := AvailableClusterListItem{}
	return &this
}

// NewAvailableClusterListItemWithDefaults instantiates a new AvailableClusterListItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAvailableClusterListItemWithDefaults() *AvailableClusterListItem {
	this := AvailableClusterListItem{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AvailableClusterListItem) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailableClusterListItem) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AvailableClusterListItem) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AvailableClusterListItem) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AvailableClusterListItem) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailableClusterListItem) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AvailableClusterListItem) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AvailableClusterListItem) SetName(v string) {
	o.Name = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *AvailableClusterListItem) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailableClusterListItem) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *AvailableClusterListItem) HasDisplayName() bool {
	return o != nil && o.DisplayName != nil
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *AvailableClusterListItem) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetClusterType returns the ClusterType field value if set, zero value otherwise.
func (o *AvailableClusterListItem) GetClusterType() string {
	if o == nil || o.ClusterType == nil {
		var ret string
		return ret
	}
	return *o.ClusterType
}

// GetClusterTypeOk returns a tuple with the ClusterType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailableClusterListItem) GetClusterTypeOk() (*string, bool) {
	if o == nil || o.ClusterType == nil {
		return nil, false
	}
	return o.ClusterType, true
}

// HasClusterType returns a boolean if a field has been set.
func (o *AvailableClusterListItem) HasClusterType() bool {
	return o != nil && o.ClusterType != nil
}

// SetClusterType gets a reference to the given string and assigns it to the ClusterType field.
func (o *AvailableClusterListItem) SetClusterType(v string) {
	o.ClusterType = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AvailableClusterListItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.ClusterType != nil {
		toSerialize["clusterType"] = o.ClusterType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AvailableClusterListItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id          *string `json:"id,omitempty"`
		Name        *string `json:"name,omitempty"`
		DisplayName *string `json:"displayName,omitempty"`
		ClusterType *string `json:"clusterType,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"id", "name", "displayName", "clusterType"})
	} else {
		return err
	}
	o.Id = all.Id
	o.Name = all.Name
	o.DisplayName = all.DisplayName
	o.ClusterType = all.ClusterType

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
