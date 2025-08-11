// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type CustomEndpointCreate struct {
	Connection *string           `json:"connection,omitempty"`
	AuthType   *AuthType         `json:"authType,omitempty"`
	BasicAuth  *BasicAuthModel   `json:"basicAuth,omitempty"`
	ExtraCfgs  map[string]string `json:"extraCfgs,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCustomEndpointCreate instantiates a new CustomEndpointCreate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCustomEndpointCreate() *CustomEndpointCreate {
	this := CustomEndpointCreate{}
	return &this
}

// NewCustomEndpointCreateWithDefaults instantiates a new CustomEndpointCreate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCustomEndpointCreateWithDefaults() *CustomEndpointCreate {
	this := CustomEndpointCreate{}
	return &this
}

// GetConnection returns the Connection field value if set, zero value otherwise.
func (o *CustomEndpointCreate) GetConnection() string {
	if o == nil || o.Connection == nil {
		var ret string
		return ret
	}
	return *o.Connection
}

// GetConnectionOk returns a tuple with the Connection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomEndpointCreate) GetConnectionOk() (*string, bool) {
	if o == nil || o.Connection == nil {
		return nil, false
	}
	return o.Connection, true
}

// HasConnection returns a boolean if a field has been set.
func (o *CustomEndpointCreate) HasConnection() bool {
	return o != nil && o.Connection != nil
}

// SetConnection gets a reference to the given string and assigns it to the Connection field.
func (o *CustomEndpointCreate) SetConnection(v string) {
	o.Connection = &v
}

// GetAuthType returns the AuthType field value if set, zero value otherwise.
func (o *CustomEndpointCreate) GetAuthType() AuthType {
	if o == nil || o.AuthType == nil {
		var ret AuthType
		return ret
	}
	return *o.AuthType
}

// GetAuthTypeOk returns a tuple with the AuthType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomEndpointCreate) GetAuthTypeOk() (*AuthType, bool) {
	if o == nil || o.AuthType == nil {
		return nil, false
	}
	return o.AuthType, true
}

// HasAuthType returns a boolean if a field has been set.
func (o *CustomEndpointCreate) HasAuthType() bool {
	return o != nil && o.AuthType != nil
}

// SetAuthType gets a reference to the given AuthType and assigns it to the AuthType field.
func (o *CustomEndpointCreate) SetAuthType(v AuthType) {
	o.AuthType = &v
}

// GetBasicAuth returns the BasicAuth field value if set, zero value otherwise.
func (o *CustomEndpointCreate) GetBasicAuth() BasicAuthModel {
	if o == nil || o.BasicAuth == nil {
		var ret BasicAuthModel
		return ret
	}
	return *o.BasicAuth
}

// GetBasicAuthOk returns a tuple with the BasicAuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomEndpointCreate) GetBasicAuthOk() (*BasicAuthModel, bool) {
	if o == nil || o.BasicAuth == nil {
		return nil, false
	}
	return o.BasicAuth, true
}

// HasBasicAuth returns a boolean if a field has been set.
func (o *CustomEndpointCreate) HasBasicAuth() bool {
	return o != nil && o.BasicAuth != nil
}

// SetBasicAuth gets a reference to the given BasicAuthModel and assigns it to the BasicAuth field.
func (o *CustomEndpointCreate) SetBasicAuth(v BasicAuthModel) {
	o.BasicAuth = &v
}

// GetExtraCfgs returns the ExtraCfgs field value if set, zero value otherwise.
func (o *CustomEndpointCreate) GetExtraCfgs() map[string]string {
	if o == nil || o.ExtraCfgs == nil {
		var ret map[string]string
		return ret
	}
	return o.ExtraCfgs
}

// GetExtraCfgsOk returns a tuple with the ExtraCfgs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomEndpointCreate) GetExtraCfgsOk() (*map[string]string, bool) {
	if o == nil || o.ExtraCfgs == nil {
		return nil, false
	}
	return &o.ExtraCfgs, true
}

// HasExtraCfgs returns a boolean if a field has been set.
func (o *CustomEndpointCreate) HasExtraCfgs() bool {
	return o != nil && o.ExtraCfgs != nil
}

// SetExtraCfgs gets a reference to the given map[string]string and assigns it to the ExtraCfgs field.
func (o *CustomEndpointCreate) SetExtraCfgs(v map[string]string) {
	o.ExtraCfgs = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CustomEndpointCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
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
	if o.ExtraCfgs != nil {
		toSerialize["extraCfgs"] = o.ExtraCfgs
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CustomEndpointCreate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Connection *string           `json:"connection,omitempty"`
		AuthType   *AuthType         `json:"authType,omitempty"`
		BasicAuth  *BasicAuthModel   `json:"basicAuth,omitempty"`
		ExtraCfgs  map[string]string `json:"extraCfgs,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"connection", "authType", "basicAuth", "extraCfgs"})
	} else {
		return err
	}

	hasInvalidField := false
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
	o.ExtraCfgs = all.ExtraCfgs

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
