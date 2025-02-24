// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type DmsExplainRequest struct {
	// the sql string
	Query *string `json:"query,omitempty"`
	// the database for explaining the SQL
	Database *string `json:"database,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDmsExplainRequest instantiates a new DmsExplainRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDmsExplainRequest() *DmsExplainRequest {
	this := DmsExplainRequest{}
	var database string = ""
	this.Database = &database
	return &this
}

// NewDmsExplainRequestWithDefaults instantiates a new DmsExplainRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDmsExplainRequestWithDefaults() *DmsExplainRequest {
	this := DmsExplainRequest{}
	var database string = ""
	this.Database = &database
	return &this
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *DmsExplainRequest) GetQuery() string {
	if o == nil || o.Query == nil {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsExplainRequest) GetQueryOk() (*string, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *DmsExplainRequest) HasQuery() bool {
	return o != nil && o.Query != nil
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *DmsExplainRequest) SetQuery(v string) {
	o.Query = &v
}

// GetDatabase returns the Database field value if set, zero value otherwise.
func (o *DmsExplainRequest) GetDatabase() string {
	if o == nil || o.Database == nil {
		var ret string
		return ret
	}
	return *o.Database
}

// GetDatabaseOk returns a tuple with the Database field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsExplainRequest) GetDatabaseOk() (*string, bool) {
	if o == nil || o.Database == nil {
		return nil, false
	}
	return o.Database, true
}

// HasDatabase returns a boolean if a field has been set.
func (o *DmsExplainRequest) HasDatabase() bool {
	return o != nil && o.Database != nil
}

// SetDatabase gets a reference to the given string and assigns it to the Database field.
func (o *DmsExplainRequest) SetDatabase(v string) {
	o.Database = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DmsExplainRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Query != nil {
		toSerialize["query"] = o.Query
	}
	if o.Database != nil {
		toSerialize["database"] = o.Database
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DmsExplainRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Query    *string `json:"query,omitempty"`
		Database *string `json:"database,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"query", "database"})
	} else {
		return err
	}
	o.Query = all.Query
	o.Database = all.Database

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
