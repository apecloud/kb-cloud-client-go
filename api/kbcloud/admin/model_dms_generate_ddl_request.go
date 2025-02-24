// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type DmsGenerateDDLRequest struct {
	// operation type
	Type          DmsGenerateDdlOperationType `json:"type"`
	TableMetadata *DmsTableMetadata           `json:"tableMetadata,omitempty"`
	ViewMetadata  *DmsViewMetadata            `json:"viewMetadata,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDmsGenerateDDLRequest instantiates a new DmsGenerateDDLRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDmsGenerateDDLRequest(typeVar DmsGenerateDdlOperationType) *DmsGenerateDDLRequest {
	this := DmsGenerateDDLRequest{}
	this.Type = typeVar
	return &this
}

// NewDmsGenerateDDLRequestWithDefaults instantiates a new DmsGenerateDDLRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDmsGenerateDDLRequestWithDefaults() *DmsGenerateDDLRequest {
	this := DmsGenerateDDLRequest{}
	return &this
}

// GetType returns the Type field value.
func (o *DmsGenerateDDLRequest) GetType() DmsGenerateDdlOperationType {
	if o == nil {
		var ret DmsGenerateDdlOperationType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DmsGenerateDDLRequest) GetTypeOk() (*DmsGenerateDdlOperationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *DmsGenerateDDLRequest) SetType(v DmsGenerateDdlOperationType) {
	o.Type = v
}

// GetTableMetadata returns the TableMetadata field value if set, zero value otherwise.
func (o *DmsGenerateDDLRequest) GetTableMetadata() DmsTableMetadata {
	if o == nil || o.TableMetadata == nil {
		var ret DmsTableMetadata
		return ret
	}
	return *o.TableMetadata
}

// GetTableMetadataOk returns a tuple with the TableMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsGenerateDDLRequest) GetTableMetadataOk() (*DmsTableMetadata, bool) {
	if o == nil || o.TableMetadata == nil {
		return nil, false
	}
	return o.TableMetadata, true
}

// HasTableMetadata returns a boolean if a field has been set.
func (o *DmsGenerateDDLRequest) HasTableMetadata() bool {
	return o != nil && o.TableMetadata != nil
}

// SetTableMetadata gets a reference to the given DmsTableMetadata and assigns it to the TableMetadata field.
func (o *DmsGenerateDDLRequest) SetTableMetadata(v DmsTableMetadata) {
	o.TableMetadata = &v
}

// GetViewMetadata returns the ViewMetadata field value if set, zero value otherwise.
func (o *DmsGenerateDDLRequest) GetViewMetadata() DmsViewMetadata {
	if o == nil || o.ViewMetadata == nil {
		var ret DmsViewMetadata
		return ret
	}
	return *o.ViewMetadata
}

// GetViewMetadataOk returns a tuple with the ViewMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsGenerateDDLRequest) GetViewMetadataOk() (*DmsViewMetadata, bool) {
	if o == nil || o.ViewMetadata == nil {
		return nil, false
	}
	return o.ViewMetadata, true
}

// HasViewMetadata returns a boolean if a field has been set.
func (o *DmsGenerateDDLRequest) HasViewMetadata() bool {
	return o != nil && o.ViewMetadata != nil
}

// SetViewMetadata gets a reference to the given DmsViewMetadata and assigns it to the ViewMetadata field.
func (o *DmsGenerateDDLRequest) SetViewMetadata(v DmsViewMetadata) {
	o.ViewMetadata = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DmsGenerateDDLRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["type"] = o.Type
	if o.TableMetadata != nil {
		toSerialize["tableMetadata"] = o.TableMetadata
	}
	if o.ViewMetadata != nil {
		toSerialize["viewMetadata"] = o.ViewMetadata
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DmsGenerateDDLRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Type          *DmsGenerateDdlOperationType `json:"type"`
		TableMetadata *DmsTableMetadata            `json:"tableMetadata,omitempty"`
		ViewMetadata  *DmsViewMetadata             `json:"viewMetadata,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"type", "tableMetadata", "viewMetadata"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}
	if all.TableMetadata != nil && all.TableMetadata.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.TableMetadata = all.TableMetadata
	if all.ViewMetadata != nil && all.ViewMetadata.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ViewMetadata = all.ViewMetadata

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
