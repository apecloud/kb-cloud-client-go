// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

type DmsVisualExplainRequest struct {
	// the sql string
	Query *string `json:"query,omitempty"`
	// the database for explaining the SQL
	Database *string `json:"database,omitempty"`
	// whether to run an actual/analyze plan when the engine supports it
	Analyze *bool `json:"analyze,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDmsVisualExplainRequest instantiates a new DmsVisualExplainRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDmsVisualExplainRequest() *DmsVisualExplainRequest {
	this := DmsVisualExplainRequest{}
	var database string = ""
	this.Database = &database
	var analyze bool = false
	this.Analyze = &analyze
	return &this
}

// NewDmsVisualExplainRequestWithDefaults instantiates a new DmsVisualExplainRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDmsVisualExplainRequestWithDefaults() *DmsVisualExplainRequest {
	this := DmsVisualExplainRequest{}
	var database string = ""
	this.Database = &database
	var analyze bool = false
	this.Analyze = &analyze
	return &this
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *DmsVisualExplainRequest) GetQuery() string {
	if o == nil || o.Query == nil {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsVisualExplainRequest) GetQueryOk() (*string, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *DmsVisualExplainRequest) HasQuery() bool {
	return o != nil && o.Query != nil
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *DmsVisualExplainRequest) SetQuery(v string) {
	o.Query = &v
}

// GetDatabase returns the Database field value if set, zero value otherwise.
func (o *DmsVisualExplainRequest) GetDatabase() string {
	if o == nil || o.Database == nil {
		var ret string
		return ret
	}
	return *o.Database
}

// GetDatabaseOk returns a tuple with the Database field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsVisualExplainRequest) GetDatabaseOk() (*string, bool) {
	if o == nil || o.Database == nil {
		return nil, false
	}
	return o.Database, true
}

// HasDatabase returns a boolean if a field has been set.
func (o *DmsVisualExplainRequest) HasDatabase() bool {
	return o != nil && o.Database != nil
}

// SetDatabase gets a reference to the given string and assigns it to the Database field.
func (o *DmsVisualExplainRequest) SetDatabase(v string) {
	o.Database = &v
}

// GetAnalyze returns the Analyze field value if set, zero value otherwise.
func (o *DmsVisualExplainRequest) GetAnalyze() bool {
	if o == nil || o.Analyze == nil {
		var ret bool
		return ret
	}
	return *o.Analyze
}

// GetAnalyzeOk returns a tuple with the Analyze field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsVisualExplainRequest) GetAnalyzeOk() (*bool, bool) {
	if o == nil || o.Analyze == nil {
		return nil, false
	}
	return o.Analyze, true
}

// HasAnalyze returns a boolean if a field has been set.
func (o *DmsVisualExplainRequest) HasAnalyze() bool {
	return o != nil && o.Analyze != nil
}

// SetAnalyze gets a reference to the given bool and assigns it to the Analyze field.
func (o *DmsVisualExplainRequest) SetAnalyze(v bool) {
	o.Analyze = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DmsVisualExplainRequest) MarshalJSON() ([]byte, error) {
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
	if o.Analyze != nil {
		toSerialize["analyze"] = o.Analyze
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DmsVisualExplainRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Query    *string `json:"query,omitempty"`
		Database *string `json:"database,omitempty"`
		Analyze  *bool   `json:"analyze,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"query", "database", "analyze"})
	} else {
		return err
	}
	o.Query = all.Query
	o.Database = all.Database
	o.Analyze = all.Analyze

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
