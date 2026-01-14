// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

// ClusterBatchUpdateData ClusterBatchUpdateData contains the update information to be applied to all clusters in the batch
type ClusterBatchUpdateData struct {
	// Display name of cluster. This will be applied to all clusters in the batch.
	DisplayName common.NullableString `json:"displayName,omitempty"`
	// the maintenance window for a cluster
	MaintainceWindow *ClusterMaintainceWindow `json:"maintainceWindow,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewClusterBatchUpdateData instantiates a new ClusterBatchUpdateData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewClusterBatchUpdateData() *ClusterBatchUpdateData {
	this := ClusterBatchUpdateData{}
	return &this
}

// NewClusterBatchUpdateDataWithDefaults instantiates a new ClusterBatchUpdateData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewClusterBatchUpdateDataWithDefaults() *ClusterBatchUpdateData {
	this := ClusterBatchUpdateData{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterBatchUpdateData) GetDisplayName() string {
	if o == nil || o.DisplayName.Get() == nil {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ClusterBatchUpdateData) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *ClusterBatchUpdateData) HasDisplayName() bool {
	return o != nil && o.DisplayName.IsSet()
}

// SetDisplayName gets a reference to the given common.NullableString and assigns it to the DisplayName field.
func (o *ClusterBatchUpdateData) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}

// SetDisplayNameNil sets the value for DisplayName to be an explicit nil.
func (o *ClusterBatchUpdateData) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil.
func (o *ClusterBatchUpdateData) UnsetDisplayName() {
	o.DisplayName.Unset()
}

// GetMaintainceWindow returns the MaintainceWindow field value if set, zero value otherwise.
func (o *ClusterBatchUpdateData) GetMaintainceWindow() ClusterMaintainceWindow {
	if o == nil || o.MaintainceWindow == nil {
		var ret ClusterMaintainceWindow
		return ret
	}
	return *o.MaintainceWindow
}

// GetMaintainceWindowOk returns a tuple with the MaintainceWindow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterBatchUpdateData) GetMaintainceWindowOk() (*ClusterMaintainceWindow, bool) {
	if o == nil || o.MaintainceWindow == nil {
		return nil, false
	}
	return o.MaintainceWindow, true
}

// HasMaintainceWindow returns a boolean if a field has been set.
func (o *ClusterBatchUpdateData) HasMaintainceWindow() bool {
	return o != nil && o.MaintainceWindow != nil
}

// SetMaintainceWindow gets a reference to the given ClusterMaintainceWindow and assigns it to the MaintainceWindow field.
func (o *ClusterBatchUpdateData) SetMaintainceWindow(v ClusterMaintainceWindow) {
	o.MaintainceWindow = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ClusterBatchUpdateData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.DisplayName.IsSet() {
		toSerialize["displayName"] = o.DisplayName.Get()
	}
	if o.MaintainceWindow != nil {
		toSerialize["maintainceWindow"] = o.MaintainceWindow
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ClusterBatchUpdateData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DisplayName      common.NullableString    `json:"displayName,omitempty"`
		MaintainceWindow *ClusterMaintainceWindow `json:"maintainceWindow,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"displayName", "maintainceWindow"})
	} else {
		return err
	}

	hasInvalidField := false
	o.DisplayName = all.DisplayName
	if all.MaintainceWindow != nil && all.MaintainceWindow.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.MaintainceWindow = all.MaintainceWindow

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
