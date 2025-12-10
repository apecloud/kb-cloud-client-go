// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type CustomEndpoint struct {
	// display connection string information, including connection details such as url, database, etc
	ConnectionString *string         `json:"connectionString,omitempty"`
	Connection       *string         `json:"connection,omitempty"`
	AuthType         *AuthType       `json:"authType,omitempty"`
	BasicAuth        *BasicAuthModel `json:"basicAuth,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCustomEndpoint instantiates a new CustomEndpoint object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCustomEndpoint() *CustomEndpoint {
	this := CustomEndpoint{}
	return &this
}

// NewCustomEndpointWithDefaults instantiates a new CustomEndpoint object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCustomEndpointWithDefaults() *CustomEndpoint {
	this := CustomEndpoint{}
	return &this
}

// GetConnectionString returns the ConnectionString field value if set, zero value otherwise.
func (o *CustomEndpoint) GetConnectionString() string {
	if o == nil || o.ConnectionString == nil {
		var ret string
		return ret
	}
	return *o.ConnectionString
}

// GetConnectionStringOk returns a tuple with the ConnectionString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomEndpoint) GetConnectionStringOk() (*string, bool) {
	if o == nil || o.ConnectionString == nil {
		return nil, false
	}
	return o.ConnectionString, true
}

// HasConnectionString returns a boolean if a field has been set.
func (o *CustomEndpoint) HasConnectionString() bool {
	return o != nil && o.ConnectionString != nil
}

// SetConnectionString gets a reference to the given string and assigns it to the ConnectionString field.
func (o *CustomEndpoint) SetConnectionString(v string) {
	o.ConnectionString = &v
}

// GetConnection returns the Connection field value if set, zero value otherwise.
func (o *CustomEndpoint) GetConnection() string {
	if o == nil || o.Connection == nil {
		var ret string
		return ret
	}
	return *o.Connection
}

// GetConnectionOk returns a tuple with the Connection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomEndpoint) GetConnectionOk() (*string, bool) {
	if o == nil || o.Connection == nil {
		return nil, false
	}
	return o.Connection, true
}

// HasConnection returns a boolean if a field has been set.
func (o *CustomEndpoint) HasConnection() bool {
	return o != nil && o.Connection != nil
}

// SetConnection gets a reference to the given string and assigns it to the Connection field.
func (o *CustomEndpoint) SetConnection(v string) {
	o.Connection = &v
}

// GetAuthType returns the AuthType field value if set, zero value otherwise.
func (o *CustomEndpoint) GetAuthType() AuthType {
	if o == nil || o.AuthType == nil {
		var ret AuthType
		return ret
	}
	return *o.AuthType
}

// GetAuthTypeOk returns a tuple with the AuthType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomEndpoint) GetAuthTypeOk() (*AuthType, bool) {
	if o == nil || o.AuthType == nil {
		return nil, false
	}
	return o.AuthType, true
}

// HasAuthType returns a boolean if a field has been set.
func (o *CustomEndpoint) HasAuthType() bool {
	return o != nil && o.AuthType != nil
}

// SetAuthType gets a reference to the given AuthType and assigns it to the AuthType field.
func (o *CustomEndpoint) SetAuthType(v AuthType) {
	o.AuthType = &v
}

// GetBasicAuth returns the BasicAuth field value if set, zero value otherwise.
func (o *CustomEndpoint) GetBasicAuth() BasicAuthModel {
	if o == nil || o.BasicAuth == nil {
		var ret BasicAuthModel
		return ret
	}
	return *o.BasicAuth
}

// GetBasicAuthOk returns a tuple with the BasicAuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomEndpoint) GetBasicAuthOk() (*BasicAuthModel, bool) {
	if o == nil || o.BasicAuth == nil {
		return nil, false
	}
	return o.BasicAuth, true
}

// HasBasicAuth returns a boolean if a field has been set.
func (o *CustomEndpoint) HasBasicAuth() bool {
	return o != nil && o.BasicAuth != nil
}

// SetBasicAuth gets a reference to the given BasicAuthModel and assigns it to the BasicAuth field.
func (o *CustomEndpoint) SetBasicAuth(v BasicAuthModel) {
	o.BasicAuth = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CustomEndpoint) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.ConnectionString != nil {
		toSerialize["connectionString"] = o.ConnectionString
	}
	if o.Connection != nil {
		toSerialize["connection"] = o.Connection
	}
	if o.AuthType != nil {
		toSerialize["authType"] = o.AuthType
	}
	if o.BasicAuth != nil {
		toSerialize["basicAuth"] = o.BasicAuth
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CustomEndpoint) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ConnectionString *string         `json:"connectionString,omitempty"`
		Connection       *string         `json:"connection,omitempty"`
		AuthType         *AuthType       `json:"authType,omitempty"`
		BasicAuth        *BasicAuthModel `json:"basicAuth,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"connectionString", "connection", "authType", "basicAuth"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ConnectionString = all.ConnectionString
	o.Connection = all.Connection
	if all.AuthType != nil && !all.AuthType.IsValid() {
		hasInvalidField = true
	} else {
		o.AuthType = all.AuthType
	}
	if all.BasicAuth != nil && all.BasicAuth.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.BasicAuth = all.BasicAuth

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
