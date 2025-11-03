// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// ImportQueryObjectsRequest Request payload for querying import data source objects.
type ImportQueryObjectsRequest struct {
	// Data source preflight request
	Preflight ImportPreflightRequest `json:"preflight"`
	// Current node type in the replication object tree.
	NodeType *string `json:"nodeType,omitempty"`
	// Encoded node value of the current tree node.
	NodeValue *string `json:"nodeValue,omitempty"`
	// Target node type to enumerate for the next level.
	QueryNodeType *string `json:"queryNodeType,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewImportQueryObjectsRequest instantiates a new ImportQueryObjectsRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewImportQueryObjectsRequest(preflight ImportPreflightRequest) *ImportQueryObjectsRequest {
	this := ImportQueryObjectsRequest{}
	this.Preflight = preflight
	return &this
}

// NewImportQueryObjectsRequestWithDefaults instantiates a new ImportQueryObjectsRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewImportQueryObjectsRequestWithDefaults() *ImportQueryObjectsRequest {
	this := ImportQueryObjectsRequest{}
	return &this
}

// GetPreflight returns the Preflight field value.
func (o *ImportQueryObjectsRequest) GetPreflight() ImportPreflightRequest {
	if o == nil {
		var ret ImportPreflightRequest
		return ret
	}
	return o.Preflight
}

// GetPreflightOk returns a tuple with the Preflight field value
// and a boolean to check if the value has been set.
func (o *ImportQueryObjectsRequest) GetPreflightOk() (*ImportPreflightRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Preflight, true
}

// SetPreflight sets field value.
func (o *ImportQueryObjectsRequest) SetPreflight(v ImportPreflightRequest) {
	o.Preflight = v
}

// GetNodeType returns the NodeType field value if set, zero value otherwise.
func (o *ImportQueryObjectsRequest) GetNodeType() string {
	if o == nil || o.NodeType == nil {
		var ret string
		return ret
	}
	return *o.NodeType
}

// GetNodeTypeOk returns a tuple with the NodeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportQueryObjectsRequest) GetNodeTypeOk() (*string, bool) {
	if o == nil || o.NodeType == nil {
		return nil, false
	}
	return o.NodeType, true
}

// HasNodeType returns a boolean if a field has been set.
func (o *ImportQueryObjectsRequest) HasNodeType() bool {
	return o != nil && o.NodeType != nil
}

// SetNodeType gets a reference to the given string and assigns it to the NodeType field.
func (o *ImportQueryObjectsRequest) SetNodeType(v string) {
	o.NodeType = &v
}

// GetNodeValue returns the NodeValue field value if set, zero value otherwise.
func (o *ImportQueryObjectsRequest) GetNodeValue() string {
	if o == nil || o.NodeValue == nil {
		var ret string
		return ret
	}
	return *o.NodeValue
}

// GetNodeValueOk returns a tuple with the NodeValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportQueryObjectsRequest) GetNodeValueOk() (*string, bool) {
	if o == nil || o.NodeValue == nil {
		return nil, false
	}
	return o.NodeValue, true
}

// HasNodeValue returns a boolean if a field has been set.
func (o *ImportQueryObjectsRequest) HasNodeValue() bool {
	return o != nil && o.NodeValue != nil
}

// SetNodeValue gets a reference to the given string and assigns it to the NodeValue field.
func (o *ImportQueryObjectsRequest) SetNodeValue(v string) {
	o.NodeValue = &v
}

// GetQueryNodeType returns the QueryNodeType field value if set, zero value otherwise.
func (o *ImportQueryObjectsRequest) GetQueryNodeType() string {
	if o == nil || o.QueryNodeType == nil {
		var ret string
		return ret
	}
	return *o.QueryNodeType
}

// GetQueryNodeTypeOk returns a tuple with the QueryNodeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportQueryObjectsRequest) GetQueryNodeTypeOk() (*string, bool) {
	if o == nil || o.QueryNodeType == nil {
		return nil, false
	}
	return o.QueryNodeType, true
}

// HasQueryNodeType returns a boolean if a field has been set.
func (o *ImportQueryObjectsRequest) HasQueryNodeType() bool {
	return o != nil && o.QueryNodeType != nil
}

// SetQueryNodeType gets a reference to the given string and assigns it to the QueryNodeType field.
func (o *ImportQueryObjectsRequest) SetQueryNodeType(v string) {
	o.QueryNodeType = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ImportQueryObjectsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["preflight"] = o.Preflight
	if o.NodeType != nil {
		toSerialize["nodeType"] = o.NodeType
	}
	if o.NodeValue != nil {
		toSerialize["nodeValue"] = o.NodeValue
	}
	if o.QueryNodeType != nil {
		toSerialize["queryNodeType"] = o.QueryNodeType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ImportQueryObjectsRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Preflight     *ImportPreflightRequest `json:"preflight"`
		NodeType      *string                 `json:"nodeType,omitempty"`
		NodeValue     *string                 `json:"nodeValue,omitempty"`
		QueryNodeType *string                 `json:"queryNodeType,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Preflight == nil {
		return fmt.Errorf("required field preflight missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"preflight", "nodeType", "nodeValue", "queryNodeType"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Preflight.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Preflight = *all.Preflight
	o.NodeType = all.NodeType
	o.NodeValue = all.NodeValue
	o.QueryNodeType = all.QueryNodeType

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
