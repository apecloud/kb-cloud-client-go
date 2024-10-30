// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd


package admin

import (
	"github.com/google/uuid"
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api"

)



// MonitorDataSinkCreate External Endpoint create option 
type MonitorDataSinkCreate struct {
	// url for exporter(eg. url for ElasticSearch)
	ExporterUrl string `json:"exporterUrl"`
	// Username
	Username common.NullableString `json:"username,omitempty"`
	// Name for relative Environment
	EnvironmentName string `json:"environmentName"`
	// type of monitor data sink(logs or metrics)
	MonitorDataSinkType string `json:"monitorDataSinkType"`
	// Password
	Password common.NullableString `json:"password,omitempty"`
	// apiKey
	ApiKey common.NullableString `json:"apiKey,omitempty"`
	// indexName for ElasticSearch
	IndexName common.NullableString `json:"indexName,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}


// NewMonitorDataSinkCreate instantiates a new MonitorDataSinkCreate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMonitorDataSinkCreate(exporterUrl string, environmentName string, monitorDataSinkType string) *MonitorDataSinkCreate {
	this := MonitorDataSinkCreate{}
	this.ExporterUrl = exporterUrl
	this.EnvironmentName = environmentName
	this.MonitorDataSinkType = monitorDataSinkType
	return &this
}

// NewMonitorDataSinkCreateWithDefaults instantiates a new MonitorDataSinkCreate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMonitorDataSinkCreateWithDefaults() *MonitorDataSinkCreate {
	this := MonitorDataSinkCreate{}
	return &this
}
// GetExporterUrl returns the ExporterUrl field value.
func (o *MonitorDataSinkCreate) GetExporterUrl() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ExporterUrl
}

// GetExporterUrlOk returns a tuple with the ExporterUrl field value
// and a boolean to check if the value has been set.
func (o *MonitorDataSinkCreate) GetExporterUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExporterUrl, true
}

// SetExporterUrl sets field value.
func (o *MonitorDataSinkCreate) SetExporterUrl(v string) {
	o.ExporterUrl = v
}


// GetUsername returns the Username field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonitorDataSinkCreate) GetUsername() string {
	if o == nil || o.Username.Get() == nil {
		var ret string
		return ret
	}
	return *o.Username.Get()
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MonitorDataSinkCreate) GetUsernameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Username.Get(), o.Username.IsSet()
}

// HasUsername returns a boolean if a field has been set.
func (o *MonitorDataSinkCreate) HasUsername() bool {
	return o != nil && o.Username.IsSet()
}

// SetUsername gets a reference to the given common.NullableString and assigns it to the Username field.
func (o *MonitorDataSinkCreate) SetUsername(v string) {
	o.Username.Set(&v)
}
// SetUsernameNil sets the value for Username to be an explicit nil.
func (o *MonitorDataSinkCreate) SetUsernameNil() {
	o.Username.Set(nil)
}

// UnsetUsername ensures that no value is present for Username, not even an explicit nil.
func (o *MonitorDataSinkCreate) UnsetUsername() {
	o.Username.Unset()
}


// GetEnvironmentName returns the EnvironmentName field value.
func (o *MonitorDataSinkCreate) GetEnvironmentName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.EnvironmentName
}

// GetEnvironmentNameOk returns a tuple with the EnvironmentName field value
// and a boolean to check if the value has been set.
func (o *MonitorDataSinkCreate) GetEnvironmentNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnvironmentName, true
}

// SetEnvironmentName sets field value.
func (o *MonitorDataSinkCreate) SetEnvironmentName(v string) {
	o.EnvironmentName = v
}


// GetMonitorDataSinkType returns the MonitorDataSinkType field value.
func (o *MonitorDataSinkCreate) GetMonitorDataSinkType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.MonitorDataSinkType
}

// GetMonitorDataSinkTypeOk returns a tuple with the MonitorDataSinkType field value
// and a boolean to check if the value has been set.
func (o *MonitorDataSinkCreate) GetMonitorDataSinkTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MonitorDataSinkType, true
}

// SetMonitorDataSinkType sets field value.
func (o *MonitorDataSinkCreate) SetMonitorDataSinkType(v string) {
	o.MonitorDataSinkType = v
}


// GetPassword returns the Password field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonitorDataSinkCreate) GetPassword() string {
	if o == nil || o.Password.Get() == nil {
		var ret string
		return ret
	}
	return *o.Password.Get()
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MonitorDataSinkCreate) GetPasswordOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Password.Get(), o.Password.IsSet()
}

// HasPassword returns a boolean if a field has been set.
func (o *MonitorDataSinkCreate) HasPassword() bool {
	return o != nil && o.Password.IsSet()
}

// SetPassword gets a reference to the given common.NullableString and assigns it to the Password field.
func (o *MonitorDataSinkCreate) SetPassword(v string) {
	o.Password.Set(&v)
}
// SetPasswordNil sets the value for Password to be an explicit nil.
func (o *MonitorDataSinkCreate) SetPasswordNil() {
	o.Password.Set(nil)
}

// UnsetPassword ensures that no value is present for Password, not even an explicit nil.
func (o *MonitorDataSinkCreate) UnsetPassword() {
	o.Password.Unset()
}


// GetApiKey returns the ApiKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonitorDataSinkCreate) GetApiKey() string {
	if o == nil || o.ApiKey.Get() == nil {
		var ret string
		return ret
	}
	return *o.ApiKey.Get()
}

// GetApiKeyOk returns a tuple with the ApiKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MonitorDataSinkCreate) GetApiKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ApiKey.Get(), o.ApiKey.IsSet()
}

// HasApiKey returns a boolean if a field has been set.
func (o *MonitorDataSinkCreate) HasApiKey() bool {
	return o != nil && o.ApiKey.IsSet()
}

// SetApiKey gets a reference to the given common.NullableString and assigns it to the ApiKey field.
func (o *MonitorDataSinkCreate) SetApiKey(v string) {
	o.ApiKey.Set(&v)
}
// SetApiKeyNil sets the value for ApiKey to be an explicit nil.
func (o *MonitorDataSinkCreate) SetApiKeyNil() {
	o.ApiKey.Set(nil)
}

// UnsetApiKey ensures that no value is present for ApiKey, not even an explicit nil.
func (o *MonitorDataSinkCreate) UnsetApiKey() {
	o.ApiKey.Unset()
}


// GetIndexName returns the IndexName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonitorDataSinkCreate) GetIndexName() string {
	if o == nil || o.IndexName.Get() == nil {
		var ret string
		return ret
	}
	return *o.IndexName.Get()
}

// GetIndexNameOk returns a tuple with the IndexName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MonitorDataSinkCreate) GetIndexNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IndexName.Get(), o.IndexName.IsSet()
}

// HasIndexName returns a boolean if a field has been set.
func (o *MonitorDataSinkCreate) HasIndexName() bool {
	return o != nil && o.IndexName.IsSet()
}

// SetIndexName gets a reference to the given common.NullableString and assigns it to the IndexName field.
func (o *MonitorDataSinkCreate) SetIndexName(v string) {
	o.IndexName.Set(&v)
}
// SetIndexNameNil sets the value for IndexName to be an explicit nil.
func (o *MonitorDataSinkCreate) SetIndexNameNil() {
	o.IndexName.Set(nil)
}

// UnsetIndexName ensures that no value is present for IndexName, not even an explicit nil.
func (o *MonitorDataSinkCreate) UnsetIndexName() {
	o.IndexName.Unset()
}



// MarshalJSON serializes the struct using spec logic.
func (o MonitorDataSinkCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["exporterUrl"] = o.ExporterUrl
	if o.Username.IsSet() {
		toSerialize["username"] = o.Username.Get()
	}
	toSerialize["environmentName"] = o.EnvironmentName
	toSerialize["monitorDataSinkType"] = o.MonitorDataSinkType
	if o.Password.IsSet() {
		toSerialize["password"] = o.Password.Get()
	}
	if o.ApiKey.IsSet() {
		toSerialize["apiKey"] = o.ApiKey.Get()
	}
	if o.IndexName.IsSet() {
		toSerialize["indexName"] = o.IndexName.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MonitorDataSinkCreate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ExporterUrl *string `json:"exporterUrl"`
		Username common.NullableString `json:"username,omitempty"`
		EnvironmentName *string `json:"environmentName"`
		MonitorDataSinkType *string `json:"monitorDataSinkType"`
		Password common.NullableString `json:"password,omitempty"`
		ApiKey common.NullableString `json:"apiKey,omitempty"`
		IndexName common.NullableString `json:"indexName,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ExporterUrl == nil {
		return fmt.Errorf("required field exporterUrl missing")
	}
	if all.EnvironmentName == nil {
		return fmt.Errorf("required field environmentName missing")
	}
	if all.MonitorDataSinkType == nil {
		return fmt.Errorf("required field monitorDataSinkType missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{ "exporterUrl", "username", "environmentName", "monitorDataSinkType", "password", "apiKey", "indexName",  })
	} else {
		return err
	}
	o.ExporterUrl = *all.ExporterUrl
	o.Username = all.Username
	o.EnvironmentName = *all.EnvironmentName
	o.MonitorDataSinkType = *all.MonitorDataSinkType
	o.Password = all.Password
	o.ApiKey = all.ApiKey
	o.IndexName = all.IndexName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
