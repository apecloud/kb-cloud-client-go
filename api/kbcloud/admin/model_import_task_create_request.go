// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// ImportTaskCreateRequest Create import task request
type ImportTaskCreateRequest struct {
	// Import name
	Name string `json:"name"`
	// Import source type
	SourceType ImportSourceType `json:"sourceType"`
	// Connection configuration
	ConnectionConfig interface{}        `json:"connectionConfig"`
	ObjectSelection  *DataChannelObject `json:"objectSelection,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewImportTaskCreateRequest instantiates a new ImportTaskCreateRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewImportTaskCreateRequest(name string, sourceType ImportSourceType, connectionConfig interface{}) *ImportTaskCreateRequest {
	this := ImportTaskCreateRequest{}
	this.Name = name
	this.SourceType = sourceType
	this.ConnectionConfig = connectionConfig
	return &this
}

// NewImportTaskCreateRequestWithDefaults instantiates a new ImportTaskCreateRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewImportTaskCreateRequestWithDefaults() *ImportTaskCreateRequest {
	this := ImportTaskCreateRequest{}
	return &this
}

// GetName returns the Name field value.
func (o *ImportTaskCreateRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ImportTaskCreateRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ImportTaskCreateRequest) SetName(v string) {
	o.Name = v
}

// GetSourceType returns the SourceType field value.
func (o *ImportTaskCreateRequest) GetSourceType() ImportSourceType {
	if o == nil {
		var ret ImportSourceType
		return ret
	}
	return o.SourceType
}

// GetSourceTypeOk returns a tuple with the SourceType field value
// and a boolean to check if the value has been set.
func (o *ImportTaskCreateRequest) GetSourceTypeOk() (*ImportSourceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceType, true
}

// SetSourceType sets field value.
func (o *ImportTaskCreateRequest) SetSourceType(v ImportSourceType) {
	o.SourceType = v
}

// GetConnectionConfig returns the ConnectionConfig field value.
func (o *ImportTaskCreateRequest) GetConnectionConfig() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ConnectionConfig
}

// GetConnectionConfigOk returns a tuple with the ConnectionConfig field value
// and a boolean to check if the value has been set.
func (o *ImportTaskCreateRequest) GetConnectionConfigOk() (*interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectionConfig, true
}

// SetConnectionConfig sets field value.
func (o *ImportTaskCreateRequest) SetConnectionConfig(v interface{}) {
	o.ConnectionConfig = v
}

// GetObjectSelection returns the ObjectSelection field value if set, zero value otherwise.
func (o *ImportTaskCreateRequest) GetObjectSelection() DataChannelObject {
	if o == nil || o.ObjectSelection == nil {
		var ret DataChannelObject
		return ret
	}
	return *o.ObjectSelection
}

// GetObjectSelectionOk returns a tuple with the ObjectSelection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportTaskCreateRequest) GetObjectSelectionOk() (*DataChannelObject, bool) {
	if o == nil || o.ObjectSelection == nil {
		return nil, false
	}
	return o.ObjectSelection, true
}

// HasObjectSelection returns a boolean if a field has been set.
func (o *ImportTaskCreateRequest) HasObjectSelection() bool {
	return o != nil && o.ObjectSelection != nil
}

// SetObjectSelection gets a reference to the given DataChannelObject and assigns it to the ObjectSelection field.
func (o *ImportTaskCreateRequest) SetObjectSelection(v DataChannelObject) {
	o.ObjectSelection = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ImportTaskCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	toSerialize["sourceType"] = o.SourceType
	toSerialize["connectionConfig"] = o.ConnectionConfig
	if o.ObjectSelection != nil {
		toSerialize["objectSelection"] = o.ObjectSelection
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ImportTaskCreateRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name             *string            `json:"name"`
		SourceType       *ImportSourceType  `json:"sourceType"`
		ConnectionConfig *interface{}       `json:"connectionConfig"`
		ObjectSelection  *DataChannelObject `json:"objectSelection,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.SourceType == nil {
		return fmt.Errorf("required field sourceType missing")
	}
	if all.ConnectionConfig == nil {
		return fmt.Errorf("required field connectionConfig missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "sourceType", "connectionConfig", "objectSelection"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Name = *all.Name
	if !all.SourceType.IsValid() {
		hasInvalidField = true
	} else {
		o.SourceType = *all.SourceType
	}
	o.ConnectionConfig = *all.ConnectionConfig
	if all.ObjectSelection != nil && all.ObjectSelection.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ObjectSelection = all.ObjectSelection

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
