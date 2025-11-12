// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// ImportPreflightRequest Data source preflight request
type ImportPreflightRequest struct {
	// Data source type
	SourceType string `json:"sourceType"`
	// Connection configuration parameters that vary by data source type
	ConnectionConfig interface{}        `json:"connectionConfig"`
	ObjectSelection  *DataChannelObject `json:"objectSelection,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewImportPreflightRequest instantiates a new ImportPreflightRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewImportPreflightRequest(sourceType string, connectionConfig interface{}) *ImportPreflightRequest {
	this := ImportPreflightRequest{}
	this.SourceType = sourceType
	this.ConnectionConfig = connectionConfig
	return &this
}

// NewImportPreflightRequestWithDefaults instantiates a new ImportPreflightRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewImportPreflightRequestWithDefaults() *ImportPreflightRequest {
	this := ImportPreflightRequest{}
	return &this
}

// GetSourceType returns the SourceType field value.
func (o *ImportPreflightRequest) GetSourceType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SourceType
}

// GetSourceTypeOk returns a tuple with the SourceType field value
// and a boolean to check if the value has been set.
func (o *ImportPreflightRequest) GetSourceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceType, true
}

// SetSourceType sets field value.
func (o *ImportPreflightRequest) SetSourceType(v string) {
	o.SourceType = v
}

// GetConnectionConfig returns the ConnectionConfig field value.
func (o *ImportPreflightRequest) GetConnectionConfig() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ConnectionConfig
}

// GetConnectionConfigOk returns a tuple with the ConnectionConfig field value
// and a boolean to check if the value has been set.
func (o *ImportPreflightRequest) GetConnectionConfigOk() (*interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectionConfig, true
}

// SetConnectionConfig sets field value.
func (o *ImportPreflightRequest) SetConnectionConfig(v interface{}) {
	o.ConnectionConfig = v
}

// GetObjectSelection returns the ObjectSelection field value if set, zero value otherwise.
func (o *ImportPreflightRequest) GetObjectSelection() DataChannelObject {
	if o == nil || o.ObjectSelection == nil {
		var ret DataChannelObject
		return ret
	}
	return *o.ObjectSelection
}

// GetObjectSelectionOk returns a tuple with the ObjectSelection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportPreflightRequest) GetObjectSelectionOk() (*DataChannelObject, bool) {
	if o == nil || o.ObjectSelection == nil {
		return nil, false
	}
	return o.ObjectSelection, true
}

// HasObjectSelection returns a boolean if a field has been set.
func (o *ImportPreflightRequest) HasObjectSelection() bool {
	return o != nil && o.ObjectSelection != nil
}

// SetObjectSelection gets a reference to the given DataChannelObject and assigns it to the ObjectSelection field.
func (o *ImportPreflightRequest) SetObjectSelection(v DataChannelObject) {
	o.ObjectSelection = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ImportPreflightRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
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
func (o *ImportPreflightRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		SourceType       *string            `json:"sourceType"`
		ConnectionConfig *interface{}       `json:"connectionConfig"`
		ObjectSelection  *DataChannelObject `json:"objectSelection,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.SourceType == nil {
		return fmt.Errorf("required field sourceType missing")
	}
	if all.ConnectionConfig == nil {
		return fmt.Errorf("required field connectionConfig missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"sourceType", "connectionConfig", "objectSelection"})
	} else {
		return err
	}

	hasInvalidField := false
	o.SourceType = *all.SourceType
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
