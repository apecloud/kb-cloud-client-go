// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

type ExtraConfig struct {
	Name  *string            `json:"name,omitempty"`
	Title *InternationalDesc `json:"title,omitempty"`
	// the type of config
	Type              *ConfigType        `json:"type,omitempty"`
	IsRequired        *bool              `json:"isRequired,omitempty"`
	ConnectionCfgType *ConnectionCfgType `json:"connectionCfgType,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExtraConfig instantiates a new ExtraConfig object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExtraConfig() *ExtraConfig {
	this := ExtraConfig{}
	return &this
}

// NewExtraConfigWithDefaults instantiates a new ExtraConfig object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExtraConfigWithDefaults() *ExtraConfig {
	this := ExtraConfig{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ExtraConfig) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtraConfig) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ExtraConfig) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ExtraConfig) SetName(v string) {
	o.Name = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *ExtraConfig) GetTitle() InternationalDesc {
	if o == nil || o.Title == nil {
		var ret InternationalDesc
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtraConfig) GetTitleOk() (*InternationalDesc, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ExtraConfig) HasTitle() bool {
	return o != nil && o.Title != nil
}

// SetTitle gets a reference to the given InternationalDesc and assigns it to the Title field.
func (o *ExtraConfig) SetTitle(v InternationalDesc) {
	o.Title = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ExtraConfig) GetType() ConfigType {
	if o == nil || o.Type == nil {
		var ret ConfigType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtraConfig) GetTypeOk() (*ConfigType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ExtraConfig) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given ConfigType and assigns it to the Type field.
func (o *ExtraConfig) SetType(v ConfigType) {
	o.Type = &v
}

// GetIsRequired returns the IsRequired field value if set, zero value otherwise.
func (o *ExtraConfig) GetIsRequired() bool {
	if o == nil || o.IsRequired == nil {
		var ret bool
		return ret
	}
	return *o.IsRequired
}

// GetIsRequiredOk returns a tuple with the IsRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtraConfig) GetIsRequiredOk() (*bool, bool) {
	if o == nil || o.IsRequired == nil {
		return nil, false
	}
	return o.IsRequired, true
}

// HasIsRequired returns a boolean if a field has been set.
func (o *ExtraConfig) HasIsRequired() bool {
	return o != nil && o.IsRequired != nil
}

// SetIsRequired gets a reference to the given bool and assigns it to the IsRequired field.
func (o *ExtraConfig) SetIsRequired(v bool) {
	o.IsRequired = &v
}

// GetConnectionCfgType returns the ConnectionCfgType field value if set, zero value otherwise.
func (o *ExtraConfig) GetConnectionCfgType() ConnectionCfgType {
	if o == nil || o.ConnectionCfgType == nil {
		var ret ConnectionCfgType
		return ret
	}
	return *o.ConnectionCfgType
}

// GetConnectionCfgTypeOk returns a tuple with the ConnectionCfgType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtraConfig) GetConnectionCfgTypeOk() (*ConnectionCfgType, bool) {
	if o == nil || o.ConnectionCfgType == nil {
		return nil, false
	}
	return o.ConnectionCfgType, true
}

// HasConnectionCfgType returns a boolean if a field has been set.
func (o *ExtraConfig) HasConnectionCfgType() bool {
	return o != nil && o.ConnectionCfgType != nil
}

// SetConnectionCfgType gets a reference to the given ConnectionCfgType and assigns it to the ConnectionCfgType field.
func (o *ExtraConfig) SetConnectionCfgType(v ConnectionCfgType) {
	o.ConnectionCfgType = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExtraConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.IsRequired != nil {
		toSerialize["isRequired"] = o.IsRequired
	}
	if o.ConnectionCfgType != nil {
		toSerialize["connectionCfgType"] = o.ConnectionCfgType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExtraConfig) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name              *string            `json:"name,omitempty"`
		Title             *InternationalDesc `json:"title,omitempty"`
		Type              *ConfigType        `json:"type,omitempty"`
		IsRequired        *bool              `json:"isRequired,omitempty"`
		ConnectionCfgType *ConnectionCfgType `json:"connectionCfgType,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "title", "type", "isRequired", "connectionCfgType"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Name = all.Name
	if all.Title != nil && all.Title.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Title = all.Title
	if all.Type != nil && !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = all.Type
	}
	o.IsRequired = all.IsRequired
	if all.ConnectionCfgType != nil && !all.ConnectionCfgType.IsValid() {
		hasInvalidField = true
	} else {
		o.ConnectionCfgType = all.ConnectionCfgType
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
