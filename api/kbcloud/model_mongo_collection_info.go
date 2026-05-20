// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type MongoCollectionInfo struct {
	// collection or view name
	Name *string `json:"name,omitempty"`
	// collection type (collection or view)
	Type *string `json:"type,omitempty"`
	// whether this is a MongoDB system collection
	IsSystem     *bool                        `json:"isSystem,omitempty"`
	Capabilities *MongoCollectionCapabilities `json:"capabilities,omitempty"`
	// collection options, including viewOn/pipeline when type is view
	Options interface{} `json:"options,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMongoCollectionInfo instantiates a new MongoCollectionInfo object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMongoCollectionInfo() *MongoCollectionInfo {
	this := MongoCollectionInfo{}
	return &this
}

// NewMongoCollectionInfoWithDefaults instantiates a new MongoCollectionInfo object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMongoCollectionInfoWithDefaults() *MongoCollectionInfo {
	this := MongoCollectionInfo{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *MongoCollectionInfo) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoCollectionInfo) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MongoCollectionInfo) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MongoCollectionInfo) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *MongoCollectionInfo) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoCollectionInfo) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *MongoCollectionInfo) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *MongoCollectionInfo) SetType(v string) {
	o.Type = &v
}

// GetIsSystem returns the IsSystem field value if set, zero value otherwise.
func (o *MongoCollectionInfo) GetIsSystem() bool {
	if o == nil || o.IsSystem == nil {
		var ret bool
		return ret
	}
	return *o.IsSystem
}

// GetIsSystemOk returns a tuple with the IsSystem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoCollectionInfo) GetIsSystemOk() (*bool, bool) {
	if o == nil || o.IsSystem == nil {
		return nil, false
	}
	return o.IsSystem, true
}

// HasIsSystem returns a boolean if a field has been set.
func (o *MongoCollectionInfo) HasIsSystem() bool {
	return o != nil && o.IsSystem != nil
}

// SetIsSystem gets a reference to the given bool and assigns it to the IsSystem field.
func (o *MongoCollectionInfo) SetIsSystem(v bool) {
	o.IsSystem = &v
}

// GetCapabilities returns the Capabilities field value if set, zero value otherwise.
func (o *MongoCollectionInfo) GetCapabilities() MongoCollectionCapabilities {
	if o == nil || o.Capabilities == nil {
		var ret MongoCollectionCapabilities
		return ret
	}
	return *o.Capabilities
}

// GetCapabilitiesOk returns a tuple with the Capabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoCollectionInfo) GetCapabilitiesOk() (*MongoCollectionCapabilities, bool) {
	if o == nil || o.Capabilities == nil {
		return nil, false
	}
	return o.Capabilities, true
}

// HasCapabilities returns a boolean if a field has been set.
func (o *MongoCollectionInfo) HasCapabilities() bool {
	return o != nil && o.Capabilities != nil
}

// SetCapabilities gets a reference to the given MongoCollectionCapabilities and assigns it to the Capabilities field.
func (o *MongoCollectionInfo) SetCapabilities(v MongoCollectionCapabilities) {
	o.Capabilities = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *MongoCollectionInfo) GetOptions() interface{} {
	if o == nil || o.Options == nil {
		var ret interface{}
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoCollectionInfo) GetOptionsOk() (*interface{}, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return &o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *MongoCollectionInfo) HasOptions() bool {
	return o != nil && o.Options != nil
}

// SetOptions gets a reference to the given interface{} and assigns it to the Options field.
func (o *MongoCollectionInfo) SetOptions(v interface{}) {
	o.Options = v
}

// MarshalJSON serializes the struct using spec logic.
func (o MongoCollectionInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.IsSystem != nil {
		toSerialize["isSystem"] = o.IsSystem
	}
	if o.Capabilities != nil {
		toSerialize["capabilities"] = o.Capabilities
	}
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MongoCollectionInfo) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name         *string                      `json:"name,omitempty"`
		Type         *string                      `json:"type,omitempty"`
		IsSystem     *bool                        `json:"isSystem,omitempty"`
		Capabilities *MongoCollectionCapabilities `json:"capabilities,omitempty"`
		Options      interface{}                  `json:"options,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "type", "isSystem", "capabilities", "options"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Name = all.Name
	o.Type = all.Type
	o.IsSystem = all.IsSystem
	if all.Capabilities != nil && all.Capabilities.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Capabilities = all.Capabilities
	o.Options = all.Options

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
