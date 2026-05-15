// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type MongoInsertRequest struct {
	// single document to insert
	Document interface{} `json:"document,omitempty"`
	// low-batch documents to insert
	Documents []interface{} `json:"documents,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMongoInsertRequest instantiates a new MongoInsertRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMongoInsertRequest() *MongoInsertRequest {
	this := MongoInsertRequest{}
	return &this
}

// NewMongoInsertRequestWithDefaults instantiates a new MongoInsertRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMongoInsertRequestWithDefaults() *MongoInsertRequest {
	this := MongoInsertRequest{}
	return &this
}

// GetDocument returns the Document field value if set, zero value otherwise.
func (o *MongoInsertRequest) GetDocument() interface{} {
	if o == nil || o.Document == nil {
		var ret interface{}
		return ret
	}
	return o.Document
}

// GetDocumentOk returns a tuple with the Document field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoInsertRequest) GetDocumentOk() (*interface{}, bool) {
	if o == nil || o.Document == nil {
		return nil, false
	}
	return &o.Document, true
}

// HasDocument returns a boolean if a field has been set.
func (o *MongoInsertRequest) HasDocument() bool {
	return o != nil && o.Document != nil
}

// SetDocument gets a reference to the given interface{} and assigns it to the Document field.
func (o *MongoInsertRequest) SetDocument(v interface{}) {
	o.Document = v
}

// GetDocuments returns the Documents field value if set, zero value otherwise.
func (o *MongoInsertRequest) GetDocuments() []interface{} {
	if o == nil || o.Documents == nil {
		var ret []interface{}
		return ret
	}
	return o.Documents
}

// GetDocumentsOk returns a tuple with the Documents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoInsertRequest) GetDocumentsOk() (*[]interface{}, bool) {
	if o == nil || o.Documents == nil {
		return nil, false
	}
	return &o.Documents, true
}

// HasDocuments returns a boolean if a field has been set.
func (o *MongoInsertRequest) HasDocuments() bool {
	return o != nil && o.Documents != nil
}

// SetDocuments gets a reference to the given []interface{} and assigns it to the Documents field.
func (o *MongoInsertRequest) SetDocuments(v []interface{}) {
	o.Documents = v
}

// MarshalJSON serializes the struct using spec logic.
func (o MongoInsertRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Document != nil {
		toSerialize["document"] = o.Document
	}
	if o.Documents != nil {
		toSerialize["documents"] = o.Documents
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MongoInsertRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Document  interface{}   `json:"document,omitempty"`
		Documents []interface{} `json:"documents,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"document", "documents"})
	} else {
		return err
	}
	o.Document = all.Document
	o.Documents = all.Documents

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
