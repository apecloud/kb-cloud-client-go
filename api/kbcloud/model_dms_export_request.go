// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type DmsExportRequest struct {
	// the database of the table or view
	Database *string `json:"database,omitempty"`
	// the sql string
	Query *string `json:"query,omitempty"`
	// return limited number of data
	Limit *int32 `json:"limit,omitempty"`
	// the file format for export data
	Format *string `json:"format,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDmsExportRequest instantiates a new DmsExportRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDmsExportRequest() *DmsExportRequest {
	this := DmsExportRequest{}
	var database string = ""
	this.Database = &database
	var limit int32 = 200
	this.Limit = &limit
	return &this
}

// GetDatabase returns the Database field value if set, zero value otherwise.
func (o *DmsExportRequest) GetDatabase() string {
	if o == nil || o.Database == nil {
		var ret string
		return ret
	}
	return *o.Database
}

// GetDatabaseOk returns a tuple with the Database field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsExportRequest) GetDatabaseOk() (*string, bool) {
	if o == nil || o.Database == nil {
		return nil, false
	}
	return o.Database, true
}

// HasDatabase returns a boolean if a field has been set.
func (o *DmsExportRequest) HasDatabase() bool {
	return o != nil && o.Database != nil
}

// SetDatabase gets a reference to the given string and assigns it to the Database field.
func (o *DmsExportRequest) SetDatabase(v string) {
	o.Database = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *DmsExportRequest) GetQuery() string {
	if o == nil || o.Query == nil {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsExportRequest) GetQueryOk() (*string, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *DmsExportRequest) HasQuery() bool {
	return o != nil && o.Query != nil
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *DmsExportRequest) SetQuery(v string) {
	o.Query = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *DmsExportRequest) GetLimit() int32 {
	if o == nil || o.Limit == nil {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsExportRequest) GetLimitOk() (*int32, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *DmsExportRequest) HasLimit() bool {
	return o != nil && o.Limit != nil
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *DmsExportRequest) SetLimit(v int32) {
	o.Limit = &v
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *DmsExportRequest) GetFormat() string {
	if o == nil || o.Format == nil {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsExportRequest) GetFormatOk() (*string, bool) {
	if o == nil || o.Format == nil {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *DmsExportRequest) HasFormat() bool {
	return o != nil && o.Format != nil
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *DmsExportRequest) SetFormat(v string) {
	o.Format = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DmsExportRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Database != nil {
		toSerialize["database"] = o.Database
	}
	if o.Query != nil {
		toSerialize["query"] = o.Query
	}
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	if o.Format != nil {
		toSerialize["format"] = o.Format
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DmsExportRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Database *string `json:"database,omitempty"`
		Query    *string `json:"query,omitempty"`
		Limit    *int32  `json:"limit,omitempty"`
		Format   *string `json:"format,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"database", "query", "limit", "format"})
	} else {
		return err
	}
	o.Database = all.Database
	o.Query = all.Query
	o.Limit = all.Limit
	o.Format = all.Format
	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
