// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd


package admin

import (
	"github.com/google/uuid"
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api"

)



// MonitorDataSink MonitorDataSink 
type MonitorDataSink struct {
	// MonitorDataSink ID
	Id string `json:"id"`
	// url for exporter(eg. url for ElasticSearch)
	ExporterUrl string `json:"exporterUrl"`
	// ID for relative Environment
	EnvironmentId string `json:"environmentID"`
	// indexName for ElasticSearch
	IndexName common.NullableString `json:"indexName,omitempty"`
	// type of monitor data sink(logs or metrics)
	MonitorDataSinkType string `json:"monitorDataSinkType"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}


// NewMonitorDataSink instantiates a new MonitorDataSink object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMonitorDataSink(id string, exporterUrl string, environmentId string, monitorDataSinkType string) *MonitorDataSink {
	this := MonitorDataSink{}
	this.Id = id
	this.ExporterUrl = exporterUrl
	this.EnvironmentId = environmentId
	this.MonitorDataSinkType = monitorDataSinkType
	return &this
}

// NewMonitorDataSinkWithDefaults instantiates a new MonitorDataSink object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMonitorDataSinkWithDefaults() *MonitorDataSink {
	this := MonitorDataSink{}
	return &this
}
// GetId returns the Id field value.
func (o *MonitorDataSink) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *MonitorDataSink) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *MonitorDataSink) SetId(v string) {
	o.Id = v
}


// GetExporterUrl returns the ExporterUrl field value.
func (o *MonitorDataSink) GetExporterUrl() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ExporterUrl
}

// GetExporterUrlOk returns a tuple with the ExporterUrl field value
// and a boolean to check if the value has been set.
func (o *MonitorDataSink) GetExporterUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExporterUrl, true
}

// SetExporterUrl sets field value.
func (o *MonitorDataSink) SetExporterUrl(v string) {
	o.ExporterUrl = v
}


// GetEnvironmentId returns the EnvironmentId field value.
func (o *MonitorDataSink) GetEnvironmentId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.EnvironmentId
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value
// and a boolean to check if the value has been set.
func (o *MonitorDataSink) GetEnvironmentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnvironmentId, true
}

// SetEnvironmentId sets field value.
func (o *MonitorDataSink) SetEnvironmentId(v string) {
	o.EnvironmentId = v
}


// GetIndexName returns the IndexName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonitorDataSink) GetIndexName() string {
	if o == nil || o.IndexName.Get() == nil {
		var ret string
		return ret
	}
	return *o.IndexName.Get()
}

// GetIndexNameOk returns a tuple with the IndexName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MonitorDataSink) GetIndexNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IndexName.Get(), o.IndexName.IsSet()
}

// HasIndexName returns a boolean if a field has been set.
func (o *MonitorDataSink) HasIndexName() bool {
	return o != nil && o.IndexName.IsSet()
}

// SetIndexName gets a reference to the given common.NullableString and assigns it to the IndexName field.
func (o *MonitorDataSink) SetIndexName(v string) {
	o.IndexName.Set(&v)
}
// SetIndexNameNil sets the value for IndexName to be an explicit nil.
func (o *MonitorDataSink) SetIndexNameNil() {
	o.IndexName.Set(nil)
}

// UnsetIndexName ensures that no value is present for IndexName, not even an explicit nil.
func (o *MonitorDataSink) UnsetIndexName() {
	o.IndexName.Unset()
}


// GetMonitorDataSinkType returns the MonitorDataSinkType field value.
func (o *MonitorDataSink) GetMonitorDataSinkType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.MonitorDataSinkType
}

// GetMonitorDataSinkTypeOk returns a tuple with the MonitorDataSinkType field value
// and a boolean to check if the value has been set.
func (o *MonitorDataSink) GetMonitorDataSinkTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MonitorDataSinkType, true
}

// SetMonitorDataSinkType sets field value.
func (o *MonitorDataSink) SetMonitorDataSinkType(v string) {
	o.MonitorDataSinkType = v
}



// MarshalJSON serializes the struct using spec logic.
func (o MonitorDataSink) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["id"] = o.Id
	toSerialize["exporterUrl"] = o.ExporterUrl
	toSerialize["environmentID"] = o.EnvironmentId
	if o.IndexName.IsSet() {
		toSerialize["indexName"] = o.IndexName.Get()
	}
	toSerialize["monitorDataSinkType"] = o.MonitorDataSinkType

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MonitorDataSink) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id *string `json:"id"`
		ExporterUrl *string `json:"exporterUrl"`
		EnvironmentId *string `json:"environmentID"`
		IndexName common.NullableString `json:"indexName,omitempty"`
		MonitorDataSinkType *string `json:"monitorDataSinkType"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.ExporterUrl == nil {
		return fmt.Errorf("required field exporterUrl missing")
	}
	if all.EnvironmentId == nil {
		return fmt.Errorf("required field environmentID missing")
	}
	if all.MonitorDataSinkType == nil {
		return fmt.Errorf("required field monitorDataSinkType missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{ "id", "exporterUrl", "environmentID", "indexName", "monitorDataSinkType",  })
	} else {
		return err
	}
	o.Id = *all.Id
	o.ExporterUrl = *all.ExporterUrl
	o.EnvironmentId = *all.EnvironmentId
	o.IndexName = all.IndexName
	o.MonitorDataSinkType = *all.MonitorDataSinkType

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
