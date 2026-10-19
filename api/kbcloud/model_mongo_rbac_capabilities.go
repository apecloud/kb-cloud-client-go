// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type MongoRBACCapabilities struct {
	EngineVersion *string                       `json:"engineVersion,omitempty"`
	Resources     []MongoRBACResourceCapability `json:"resources,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMongoRBACCapabilities instantiates a new MongoRBACCapabilities object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMongoRBACCapabilities() *MongoRBACCapabilities {
	this := MongoRBACCapabilities{}
	return &this
}

// NewMongoRBACCapabilitiesWithDefaults instantiates a new MongoRBACCapabilities object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMongoRBACCapabilitiesWithDefaults() *MongoRBACCapabilities {
	this := MongoRBACCapabilities{}
	return &this
}

// GetEngineVersion returns the EngineVersion field value if set, zero value otherwise.
func (o *MongoRBACCapabilities) GetEngineVersion() string {
	if o == nil || o.EngineVersion == nil {
		var ret string
		return ret
	}
	return *o.EngineVersion
}

// GetEngineVersionOk returns a tuple with the EngineVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoRBACCapabilities) GetEngineVersionOk() (*string, bool) {
	if o == nil || o.EngineVersion == nil {
		return nil, false
	}
	return o.EngineVersion, true
}

// HasEngineVersion returns a boolean if a field has been set.
func (o *MongoRBACCapabilities) HasEngineVersion() bool {
	return o != nil && o.EngineVersion != nil
}

// SetEngineVersion gets a reference to the given string and assigns it to the EngineVersion field.
func (o *MongoRBACCapabilities) SetEngineVersion(v string) {
	o.EngineVersion = &v
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *MongoRBACCapabilities) GetResources() []MongoRBACResourceCapability {
	if o == nil || o.Resources == nil {
		var ret []MongoRBACResourceCapability
		return ret
	}
	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoRBACCapabilities) GetResourcesOk() (*[]MongoRBACResourceCapability, bool) {
	if o == nil || o.Resources == nil {
		return nil, false
	}
	return &o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *MongoRBACCapabilities) HasResources() bool {
	return o != nil && o.Resources != nil
}

// SetResources gets a reference to the given []MongoRBACResourceCapability and assigns it to the Resources field.
func (o *MongoRBACCapabilities) SetResources(v []MongoRBACResourceCapability) {
	o.Resources = v
}

// MarshalJSON serializes the struct using spec logic.
func (o MongoRBACCapabilities) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.EngineVersion != nil {
		toSerialize["engineVersion"] = o.EngineVersion
	}
	if o.Resources != nil {
		toSerialize["resources"] = o.Resources
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MongoRBACCapabilities) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		EngineVersion *string                       `json:"engineVersion,omitempty"`
		Resources     []MongoRBACResourceCapability `json:"resources,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"engineVersion", "resources"})
	} else {
		return err
	}
	o.EngineVersion = all.EngineVersion
	o.Resources = all.Resources

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
