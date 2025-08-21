// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type DataChannelListEndpoint struct {
	EngineName   *string                       `json:"engineName,omitempty"`
	EndpointType *DataReplication_endpointType `json:"endpointType,omitempty"`
	Custom       *CustomEndpoint               `json:"custom,omitempty"`
	Kubeblocks   *KubeblocksEndpoint           `json:"kubeblocks,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDataChannelListEndpoint instantiates a new DataChannelListEndpoint object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDataChannelListEndpoint() *DataChannelListEndpoint {
	this := DataChannelListEndpoint{}
	return &this
}

// NewDataChannelListEndpointWithDefaults instantiates a new DataChannelListEndpoint object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDataChannelListEndpointWithDefaults() *DataChannelListEndpoint {
	this := DataChannelListEndpoint{}
	return &this
}

// GetEngineName returns the EngineName field value if set, zero value otherwise.
func (o *DataChannelListEndpoint) GetEngineName() string {
	if o == nil || o.EngineName == nil {
		var ret string
		return ret
	}
	return *o.EngineName
}

// GetEngineNameOk returns a tuple with the EngineName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataChannelListEndpoint) GetEngineNameOk() (*string, bool) {
	if o == nil || o.EngineName == nil {
		return nil, false
	}
	return o.EngineName, true
}

// HasEngineName returns a boolean if a field has been set.
func (o *DataChannelListEndpoint) HasEngineName() bool {
	return o != nil && o.EngineName != nil
}

// SetEngineName gets a reference to the given string and assigns it to the EngineName field.
func (o *DataChannelListEndpoint) SetEngineName(v string) {
	o.EngineName = &v
}

// GetEndpointType returns the EndpointType field value if set, zero value otherwise.
func (o *DataChannelListEndpoint) GetEndpointType() DataReplication_endpointType {
	if o == nil || o.EndpointType == nil {
		var ret DataReplication_endpointType
		return ret
	}
	return *o.EndpointType
}

// GetEndpointTypeOk returns a tuple with the EndpointType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataChannelListEndpoint) GetEndpointTypeOk() (*DataReplication_endpointType, bool) {
	if o == nil || o.EndpointType == nil {
		return nil, false
	}
	return o.EndpointType, true
}

// HasEndpointType returns a boolean if a field has been set.
func (o *DataChannelListEndpoint) HasEndpointType() bool {
	return o != nil && o.EndpointType != nil
}

// SetEndpointType gets a reference to the given DataReplication_endpointType and assigns it to the EndpointType field.
func (o *DataChannelListEndpoint) SetEndpointType(v DataReplication_endpointType) {
	o.EndpointType = &v
}

// GetCustom returns the Custom field value if set, zero value otherwise.
func (o *DataChannelListEndpoint) GetCustom() CustomEndpoint {
	if o == nil || o.Custom == nil {
		var ret CustomEndpoint
		return ret
	}
	return *o.Custom
}

// GetCustomOk returns a tuple with the Custom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataChannelListEndpoint) GetCustomOk() (*CustomEndpoint, bool) {
	if o == nil || o.Custom == nil {
		return nil, false
	}
	return o.Custom, true
}

// HasCustom returns a boolean if a field has been set.
func (o *DataChannelListEndpoint) HasCustom() bool {
	return o != nil && o.Custom != nil
}

// SetCustom gets a reference to the given CustomEndpoint and assigns it to the Custom field.
func (o *DataChannelListEndpoint) SetCustom(v CustomEndpoint) {
	o.Custom = &v
}

// GetKubeblocks returns the Kubeblocks field value if set, zero value otherwise.
func (o *DataChannelListEndpoint) GetKubeblocks() KubeblocksEndpoint {
	if o == nil || o.Kubeblocks == nil {
		var ret KubeblocksEndpoint
		return ret
	}
	return *o.Kubeblocks
}

// GetKubeblocksOk returns a tuple with the Kubeblocks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataChannelListEndpoint) GetKubeblocksOk() (*KubeblocksEndpoint, bool) {
	if o == nil || o.Kubeblocks == nil {
		return nil, false
	}
	return o.Kubeblocks, true
}

// HasKubeblocks returns a boolean if a field has been set.
func (o *DataChannelListEndpoint) HasKubeblocks() bool {
	return o != nil && o.Kubeblocks != nil
}

// SetKubeblocks gets a reference to the given KubeblocksEndpoint and assigns it to the Kubeblocks field.
func (o *DataChannelListEndpoint) SetKubeblocks(v KubeblocksEndpoint) {
	o.Kubeblocks = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DataChannelListEndpoint) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.EngineName != nil {
		toSerialize["engineName"] = o.EngineName
	}
	if o.EndpointType != nil {
		toSerialize["endpointType"] = o.EndpointType
	}
	if o.Custom != nil {
		toSerialize["custom"] = o.Custom
	}
	if o.Kubeblocks != nil {
		toSerialize["kubeblocks"] = o.Kubeblocks
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DataChannelListEndpoint) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		EngineName   *string                       `json:"engineName,omitempty"`
		EndpointType *DataReplication_endpointType `json:"endpointType,omitempty"`
		Custom       *CustomEndpoint               `json:"custom,omitempty"`
		Kubeblocks   *KubeblocksEndpoint           `json:"kubeblocks,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"engineName", "endpointType", "custom", "kubeblocks"})
	} else {
		return err
	}

	hasInvalidField := false
	o.EngineName = all.EngineName
	if all.EndpointType != nil && !all.EndpointType.IsValid() {
		hasInvalidField = true
	} else {
		o.EndpointType = all.EndpointType
	}
	if all.Custom != nil && all.Custom.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Custom = all.Custom
	if all.Kubeblocks != nil && all.Kubeblocks.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Kubeblocks = all.Kubeblocks

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
