// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type AiDataGatewayTestDataSourceRequest struct {
	DatasourceId *string                      `json:"datasourceId,omitempty"`
	Spec         *AiDataGatewayDataSourceSpec `json:"spec,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiDataGatewayTestDataSourceRequest instantiates a new AiDataGatewayTestDataSourceRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiDataGatewayTestDataSourceRequest() *AiDataGatewayTestDataSourceRequest {
	this := AiDataGatewayTestDataSourceRequest{}
	return &this
}

// NewAiDataGatewayTestDataSourceRequestWithDefaults instantiates a new AiDataGatewayTestDataSourceRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiDataGatewayTestDataSourceRequestWithDefaults() *AiDataGatewayTestDataSourceRequest {
	this := AiDataGatewayTestDataSourceRequest{}
	return &this
}

// GetDatasourceId returns the DatasourceId field value if set, zero value otherwise.
func (o *AiDataGatewayTestDataSourceRequest) GetDatasourceId() string {
	if o == nil || o.DatasourceId == nil {
		var ret string
		return ret
	}
	return *o.DatasourceId
}

// GetDatasourceIdOk returns a tuple with the DatasourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayTestDataSourceRequest) GetDatasourceIdOk() (*string, bool) {
	if o == nil || o.DatasourceId == nil {
		return nil, false
	}
	return o.DatasourceId, true
}

// HasDatasourceId returns a boolean if a field has been set.
func (o *AiDataGatewayTestDataSourceRequest) HasDatasourceId() bool {
	return o != nil && o.DatasourceId != nil
}

// SetDatasourceId gets a reference to the given string and assigns it to the DatasourceId field.
func (o *AiDataGatewayTestDataSourceRequest) SetDatasourceId(v string) {
	o.DatasourceId = &v
}

// GetSpec returns the Spec field value if set, zero value otherwise.
func (o *AiDataGatewayTestDataSourceRequest) GetSpec() AiDataGatewayDataSourceSpec {
	if o == nil || o.Spec == nil {
		var ret AiDataGatewayDataSourceSpec
		return ret
	}
	return *o.Spec
}

// GetSpecOk returns a tuple with the Spec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayTestDataSourceRequest) GetSpecOk() (*AiDataGatewayDataSourceSpec, bool) {
	if o == nil || o.Spec == nil {
		return nil, false
	}
	return o.Spec, true
}

// HasSpec returns a boolean if a field has been set.
func (o *AiDataGatewayTestDataSourceRequest) HasSpec() bool {
	return o != nil && o.Spec != nil
}

// SetSpec gets a reference to the given AiDataGatewayDataSourceSpec and assigns it to the Spec field.
func (o *AiDataGatewayTestDataSourceRequest) SetSpec(v AiDataGatewayDataSourceSpec) {
	o.Spec = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiDataGatewayTestDataSourceRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.DatasourceId != nil {
		toSerialize["datasourceId"] = o.DatasourceId
	}
	if o.Spec != nil {
		toSerialize["spec"] = o.Spec
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiDataGatewayTestDataSourceRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DatasourceId *string                      `json:"datasourceId,omitempty"`
		Spec         *AiDataGatewayDataSourceSpec `json:"spec,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"datasourceId", "spec"})
	} else {
		return err
	}

	hasInvalidField := false
	o.DatasourceId = all.DatasourceId
	if all.Spec != nil && all.Spec.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Spec = all.Spec

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
