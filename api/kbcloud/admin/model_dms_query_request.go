// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

type DmsQueryRequest struct {
	// the database of the table or view
	Database *string `json:"database,omitempty"`
	// the target table or view name
	Table *string `json:"table,omitempty"`
	// return limited number of data
	Limit *int32 `json:"limit,omitempty"`
	// if noRecord is true, the query(s) will not be added to query history
	NoRecord common.NullableBool `json:"noRecord,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDmsQueryRequest instantiates a new DmsQueryRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDmsQueryRequest() *DmsQueryRequest {
	this := DmsQueryRequest{}
	var database string = ""
	this.Database = &database
	var limit int32 = 200
	this.Limit = &limit
	return &this
}

// GetDatabase returns the Database field value if set, zero value otherwise.
func (o *DmsQueryRequest) GetDatabase() string {
	if o == nil || o.Database == nil {
		var ret string
		return ret
	}
	return *o.Database
}

// GetDatabaseOk returns a tuple with the Database field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsQueryRequest) GetDatabaseOk() (*string, bool) {
	if o == nil || o.Database == nil {
		return nil, false
	}
	return o.Database, true
}

// HasDatabase returns a boolean if a field has been set.
func (o *DmsQueryRequest) HasDatabase() bool {
	return o != nil && o.Database != nil
}

// SetDatabase gets a reference to the given string and assigns it to the Database field.
func (o *DmsQueryRequest) SetDatabase(v string) {
	o.Database = &v
}

// GetTable returns the Table field value if set, zero value otherwise.
func (o *DmsQueryRequest) GetTable() string {
	if o == nil || o.Table == nil {
		var ret string
		return ret
	}
	return *o.Table
}

// GetTableOk returns a tuple with the Table field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsQueryRequest) GetTableOk() (*string, bool) {
	if o == nil || o.Table == nil {
		return nil, false
	}
	return o.Table, true
}

// HasTable returns a boolean if a field has been set.
func (o *DmsQueryRequest) HasTable() bool {
	return o != nil && o.Table != nil
}

// SetTable gets a reference to the given string and assigns it to the Table field.
func (o *DmsQueryRequest) SetTable(v string) {
	o.Table = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *DmsQueryRequest) GetLimit() int32 {
	if o == nil || o.Limit == nil {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsQueryRequest) GetLimitOk() (*int32, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *DmsQueryRequest) HasLimit() bool {
	return o != nil && o.Limit != nil
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *DmsQueryRequest) SetLimit(v int32) {
	o.Limit = &v
}

// GetNoRecord returns the NoRecord field value if set, zero value otherwise.
func (o *DmsQueryRequest) GetNoRecord() bool {
	if o == nil || o.NoRecord.Get() == nil {
		var ret bool
		return ret
	}
	return *o.NoRecord.Get()
}

// GetNoRecordOk returns a tuple with the NoRecord field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsQueryRequest) GetNoRecordOk() (*bool, bool) {
	if o == nil || o.NoRecord.Get() == nil {
		return nil, false
	}
	return o.NoRecord.Get(), o.NoRecord.IsSet()
}

// HasNoRecord returns a boolean if a field has been set.
func (o *DmsQueryRequest) HasNoRecord() bool {
	return o != nil && o.NoRecord.IsSet()
}

// SetNoRecord gets a reference to the given common.NullableBool and assigns it to the NoRecord field.
func (o *DmsQueryRequest) SetNoRecord(v bool) {
	o.NoRecord.Set(&v)
}

// SetNoRecordNil sets the value for NoRecord to be an explicit nil.
func (o *DmsQueryRequest) SetNoRecordNil() {
	o.NoRecord.Set(nil)
}

// UnsetNoRecord ensures that no value is present for NoRecord, not even an explicit nil.
func (o *DmsQueryRequest) UnsetNoRecord() {
	o.NoRecord.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o DmsQueryRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Database != nil {
		toSerialize["database"] = o.Database
	}
	if o.Table != nil {
		toSerialize["table"] = o.Table
	}
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	if o.NoRecord.IsSet() {
		toSerialize["noRecord"] = o.NoRecord.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DmsQueryRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Database *string             `json:"database,omitempty"`
		Table    *string             `json:"table,omitempty"`
		Limit    *int32              `json:"limit,omitempty"`
		NoRecord common.NullableBool `json:"noRecord,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"database", "table", "limit", "noRecord"})
	} else {
		return err
	}
	o.Database = all.Database
	o.Table = all.Table
	o.Limit = all.Limit
	o.NoRecord = all.NoRecord
	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
