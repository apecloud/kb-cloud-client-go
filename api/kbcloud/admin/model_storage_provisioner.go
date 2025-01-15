// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

// StorageProvisioner StorageProvisioner provides detailed information about the provisioner used by storage classes.
type StorageProvisioner struct {
	Provisioner   common.NullableString `json:"provisioner,omitempty"`
	Type          common.NullableString `json:"type,omitempty"`
	CloudProvider common.NullableString `json:"cloudProvider,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewStorageProvisioner instantiates a new StorageProvisioner object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewStorageProvisioner() *StorageProvisioner {
	this := StorageProvisioner{}
	return &this
}

// NewStorageProvisionerWithDefaults instantiates a new StorageProvisioner object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewStorageProvisionerWithDefaults() *StorageProvisioner {
	this := StorageProvisioner{}
	return &this
}

// GetProvisioner returns the Provisioner field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageProvisioner) GetProvisioner() string {
	if o == nil || o.Provisioner.Get() == nil {
		var ret string
		return ret
	}
	return *o.Provisioner.Get()
}

// GetProvisionerOk returns a tuple with the Provisioner field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *StorageProvisioner) GetProvisionerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Provisioner.Get(), o.Provisioner.IsSet()
}

// HasProvisioner returns a boolean if a field has been set.
func (o *StorageProvisioner) HasProvisioner() bool {
	return o != nil && o.Provisioner.IsSet()
}

// SetProvisioner gets a reference to the given common.NullableString and assigns it to the Provisioner field.
func (o *StorageProvisioner) SetProvisioner(v string) {
	o.Provisioner.Set(&v)
}

// SetProvisionerNil sets the value for Provisioner to be an explicit nil.
func (o *StorageProvisioner) SetProvisionerNil() {
	o.Provisioner.Set(nil)
}

// UnsetProvisioner ensures that no value is present for Provisioner, not even an explicit nil.
func (o *StorageProvisioner) UnsetProvisioner() {
	o.Provisioner.Unset()
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageProvisioner) GetType() string {
	if o == nil || o.Type.Get() == nil {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *StorageProvisioner) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *StorageProvisioner) HasType() bool {
	return o != nil && o.Type.IsSet()
}

// SetType gets a reference to the given common.NullableString and assigns it to the Type field.
func (o *StorageProvisioner) SetType(v string) {
	o.Type.Set(&v)
}

// SetTypeNil sets the value for Type to be an explicit nil.
func (o *StorageProvisioner) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil.
func (o *StorageProvisioner) UnsetType() {
	o.Type.Unset()
}

// GetCloudProvider returns the CloudProvider field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageProvisioner) GetCloudProvider() string {
	if o == nil || o.CloudProvider.Get() == nil {
		var ret string
		return ret
	}
	return *o.CloudProvider.Get()
}

// GetCloudProviderOk returns a tuple with the CloudProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *StorageProvisioner) GetCloudProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CloudProvider.Get(), o.CloudProvider.IsSet()
}

// HasCloudProvider returns a boolean if a field has been set.
func (o *StorageProvisioner) HasCloudProvider() bool {
	return o != nil && o.CloudProvider.IsSet()
}

// SetCloudProvider gets a reference to the given common.NullableString and assigns it to the CloudProvider field.
func (o *StorageProvisioner) SetCloudProvider(v string) {
	o.CloudProvider.Set(&v)
}

// SetCloudProviderNil sets the value for CloudProvider to be an explicit nil.
func (o *StorageProvisioner) SetCloudProviderNil() {
	o.CloudProvider.Set(nil)
}

// UnsetCloudProvider ensures that no value is present for CloudProvider, not even an explicit nil.
func (o *StorageProvisioner) UnsetCloudProvider() {
	o.CloudProvider.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o StorageProvisioner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Provisioner.IsSet() {
		toSerialize["provisioner"] = o.Provisioner.Get()
	}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	if o.CloudProvider.IsSet() {
		toSerialize["cloudProvider"] = o.CloudProvider.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *StorageProvisioner) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Provisioner   common.NullableString `json:"provisioner,omitempty"`
		Type          common.NullableString `json:"type,omitempty"`
		CloudProvider common.NullableString `json:"cloudProvider,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"provisioner", "type", "cloudProvider"})
	} else {
		return err
	}
	o.Provisioner = all.Provisioner
	o.Type = all.Type
	o.CloudProvider = all.CloudProvider

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
