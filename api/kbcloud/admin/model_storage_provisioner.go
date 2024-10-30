// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd


package admin

import (
	"github.com/google/uuid"
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api"

)



// StorageProvisioner StorageProvisioner provides detailed information about the provisioner used by storage classes. 
type StorageProvisioner struct {
	Provisioner string `json:"provisioner"`
	Type string `json:"type"`
	CloudProvider string `json:"cloudProvider"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}


// NewStorageProvisioner instantiates a new StorageProvisioner object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewStorageProvisioner(provisioner string, typeVar string, cloudProvider string) *StorageProvisioner {
	this := StorageProvisioner{}
	this.Provisioner = provisioner
	this.Type = typeVar
	this.CloudProvider = cloudProvider
	return &this
}

// NewStorageProvisionerWithDefaults instantiates a new StorageProvisioner object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewStorageProvisionerWithDefaults() *StorageProvisioner {
	this := StorageProvisioner{}
	return &this
}
// GetProvisioner returns the Provisioner field value.
func (o *StorageProvisioner) GetProvisioner() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Provisioner
}

// GetProvisionerOk returns a tuple with the Provisioner field value
// and a boolean to check if the value has been set.
func (o *StorageProvisioner) GetProvisionerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Provisioner, true
}

// SetProvisioner sets field value.
func (o *StorageProvisioner) SetProvisioner(v string) {
	o.Provisioner = v
}


// GetType returns the Type field value.
func (o *StorageProvisioner) GetType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *StorageProvisioner) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *StorageProvisioner) SetType(v string) {
	o.Type = v
}


// GetCloudProvider returns the CloudProvider field value.
func (o *StorageProvisioner) GetCloudProvider() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CloudProvider
}

// GetCloudProviderOk returns a tuple with the CloudProvider field value
// and a boolean to check if the value has been set.
func (o *StorageProvisioner) GetCloudProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudProvider, true
}

// SetCloudProvider sets field value.
func (o *StorageProvisioner) SetCloudProvider(v string) {
	o.CloudProvider = v
}



// MarshalJSON serializes the struct using spec logic.
func (o StorageProvisioner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["provisioner"] = o.Provisioner
	toSerialize["type"] = o.Type
	toSerialize["cloudProvider"] = o.CloudProvider

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *StorageProvisioner) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Provisioner *string `json:"provisioner"`
		Type *string `json:"type"`
		CloudProvider *string `json:"cloudProvider"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Provisioner == nil {
		return fmt.Errorf("required field provisioner missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	if all.CloudProvider == nil {
		return fmt.Errorf("required field cloudProvider missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{ "provisioner", "type", "cloudProvider",  })
	} else {
		return err
	}
	o.Provisioner = *all.Provisioner
	o.Type = *all.Type
	o.CloudProvider = *all.CloudProvider

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
