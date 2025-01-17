// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

type DocumentList struct {
	Documents []map[string]interface{} `json:"documents,omitempty"`
	Total     *int32                   `json:"total,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDocumentList instantiates a new DocumentList object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDocumentList() *DocumentList {
	this := DocumentList{}
	return &this
}

// NewDocumentListWithDefaults instantiates a new DocumentList object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDocumentListWithDefaults() *DocumentList {
	this := DocumentList{}
	return &this
}

// GetDocuments returns the Documents field value if set, zero value otherwise.
func (o *DocumentList) GetDocuments() []map[string]interface{} {
	if o == nil || o.Documents == nil {
		var ret []map[string]interface{}
		return ret
	}
	return o.Documents
}

// GetDocumentsOk returns a tuple with the Documents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentList) GetDocumentsOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.Documents == nil {
		return nil, false
	}
	return &o.Documents, true
}

// HasDocuments returns a boolean if a field has been set.
func (o *DocumentList) HasDocuments() bool {
	return o != nil && o.Documents != nil
}

// SetDocuments gets a reference to the given []map[string]interface{} and assigns it to the Documents field.
func (o *DocumentList) SetDocuments(v []map[string]interface{}) {
	o.Documents = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *DocumentList) GetTotal() int32 {
	if o == nil || o.Total == nil {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentList) GetTotalOk() (*int32, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *DocumentList) HasTotal() bool {
	return o != nil && o.Total != nil
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *DocumentList) SetTotal(v int32) {
	o.Total = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DocumentList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Documents != nil {
		toSerialize["documents"] = o.Documents
	}
	if o.Total != nil {
		toSerialize["total"] = o.Total
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DocumentList) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Documents []map[string]interface{} `json:"documents,omitempty"`
		Total     *int32                   `json:"total,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"documents", "total"})
	} else {
		return err
	}
	o.Documents = all.Documents
	o.Total = all.Total

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
