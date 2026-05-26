// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type MongoFindResponse struct {
	// column names (flattened document key union)
	Columns []string `json:"columns,omitempty"`
	// row data (flattened document values)
	Rows [][]interface{} `json:"rows,omitempty"`
	// raw documents (when raw mode)
	Documents []interface{} `json:"documents,omitempty"`
	// total matching documents
	TotalCount *int64 `json:"totalCount,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMongoFindResponse instantiates a new MongoFindResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMongoFindResponse() *MongoFindResponse {
	this := MongoFindResponse{}
	return &this
}

// NewMongoFindResponseWithDefaults instantiates a new MongoFindResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMongoFindResponseWithDefaults() *MongoFindResponse {
	this := MongoFindResponse{}
	return &this
}

// GetColumns returns the Columns field value if set, zero value otherwise.
func (o *MongoFindResponse) GetColumns() []string {
	if o == nil || o.Columns == nil {
		var ret []string
		return ret
	}
	return o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoFindResponse) GetColumnsOk() (*[]string, bool) {
	if o == nil || o.Columns == nil {
		return nil, false
	}
	return &o.Columns, true
}

// HasColumns returns a boolean if a field has been set.
func (o *MongoFindResponse) HasColumns() bool {
	return o != nil && o.Columns != nil
}

// SetColumns gets a reference to the given []string and assigns it to the Columns field.
func (o *MongoFindResponse) SetColumns(v []string) {
	o.Columns = v
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *MongoFindResponse) GetRows() [][]interface{} {
	if o == nil || o.Rows == nil {
		var ret [][]interface{}
		return ret
	}
	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoFindResponse) GetRowsOk() (*[][]interface{}, bool) {
	if o == nil || o.Rows == nil {
		return nil, false
	}
	return &o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *MongoFindResponse) HasRows() bool {
	return o != nil && o.Rows != nil
}

// SetRows gets a reference to the given [][]interface{} and assigns it to the Rows field.
func (o *MongoFindResponse) SetRows(v [][]interface{}) {
	o.Rows = v
}

// GetDocuments returns the Documents field value if set, zero value otherwise.
func (o *MongoFindResponse) GetDocuments() []interface{} {
	if o == nil || o.Documents == nil {
		var ret []interface{}
		return ret
	}
	return o.Documents
}

// GetDocumentsOk returns a tuple with the Documents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoFindResponse) GetDocumentsOk() (*[]interface{}, bool) {
	if o == nil || o.Documents == nil {
		return nil, false
	}
	return &o.Documents, true
}

// HasDocuments returns a boolean if a field has been set.
func (o *MongoFindResponse) HasDocuments() bool {
	return o != nil && o.Documents != nil
}

// SetDocuments gets a reference to the given []interface{} and assigns it to the Documents field.
func (o *MongoFindResponse) SetDocuments(v []interface{}) {
	o.Documents = v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *MongoFindResponse) GetTotalCount() int64 {
	if o == nil || o.TotalCount == nil {
		var ret int64
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoFindResponse) GetTotalCountOk() (*int64, bool) {
	if o == nil || o.TotalCount == nil {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *MongoFindResponse) HasTotalCount() bool {
	return o != nil && o.TotalCount != nil
}

// SetTotalCount gets a reference to the given int64 and assigns it to the TotalCount field.
func (o *MongoFindResponse) SetTotalCount(v int64) {
	o.TotalCount = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o MongoFindResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Columns != nil {
		toSerialize["columns"] = o.Columns
	}
	if o.Rows != nil {
		toSerialize["rows"] = o.Rows
	}
	if o.Documents != nil {
		toSerialize["documents"] = o.Documents
	}
	if o.TotalCount != nil {
		toSerialize["totalCount"] = o.TotalCount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MongoFindResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Columns    []string        `json:"columns,omitempty"`
		Rows       [][]interface{} `json:"rows,omitempty"`
		Documents  []interface{}   `json:"documents,omitempty"`
		TotalCount *int64          `json:"totalCount,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"columns", "rows", "documents", "totalCount"})
	} else {
		return err
	}
	o.Columns = all.Columns
	o.Rows = all.Rows
	o.Documents = all.Documents
	o.TotalCount = all.TotalCount

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
