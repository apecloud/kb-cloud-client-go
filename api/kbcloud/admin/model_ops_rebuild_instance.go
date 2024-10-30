// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd


package admin

import (
	"github.com/google/uuid"
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api"

)



// OpsRebuildInstance rebuild the instances of the cluster. 
type OpsRebuildInstance struct {
	// will ignore role check during rebuilding instance.
	IgnoreRoleCheck *bool `json:"ignoreRoleCheck,omitempty"`
	Requests []OpsRebuildInstanceRequestsItem `json:"requests,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}


// NewOpsRebuildInstance instantiates a new OpsRebuildInstance object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOpsRebuildInstance() *OpsRebuildInstance {
	this := OpsRebuildInstance{}
	return &this
}

// NewOpsRebuildInstanceWithDefaults instantiates a new OpsRebuildInstance object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOpsRebuildInstanceWithDefaults() *OpsRebuildInstance {
	this := OpsRebuildInstance{}
	return &this
}
// GetIgnoreRoleCheck returns the IgnoreRoleCheck field value if set, zero value otherwise.
func (o *OpsRebuildInstance) GetIgnoreRoleCheck() bool {
	if o == nil || o.IgnoreRoleCheck == nil {
		var ret bool
		return ret
	}
	return *o.IgnoreRoleCheck
}

// GetIgnoreRoleCheckOk returns a tuple with the IgnoreRoleCheck field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsRebuildInstance) GetIgnoreRoleCheckOk() (*bool, bool) {
	if o == nil || o.IgnoreRoleCheck == nil {
		return nil, false
	}
	return o.IgnoreRoleCheck, true
}

// HasIgnoreRoleCheck returns a boolean if a field has been set.
func (o *OpsRebuildInstance) HasIgnoreRoleCheck() bool {
	return o != nil && o.IgnoreRoleCheck != nil
}

// SetIgnoreRoleCheck gets a reference to the given bool and assigns it to the IgnoreRoleCheck field.
func (o *OpsRebuildInstance) SetIgnoreRoleCheck(v bool) {
	o.IgnoreRoleCheck = &v
}


// GetRequests returns the Requests field value if set, zero value otherwise.
func (o *OpsRebuildInstance) GetRequests() []OpsRebuildInstanceRequestsItem {
	if o == nil || o.Requests == nil {
		var ret []OpsRebuildInstanceRequestsItem
		return ret
	}
	return o.Requests
}

// GetRequestsOk returns a tuple with the Requests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsRebuildInstance) GetRequestsOk() (*[]OpsRebuildInstanceRequestsItem, bool) {
	if o == nil || o.Requests == nil {
		return nil, false
	}
	return &o.Requests, true
}

// HasRequests returns a boolean if a field has been set.
func (o *OpsRebuildInstance) HasRequests() bool {
	return o != nil && o.Requests != nil
}

// SetRequests gets a reference to the given []OpsRebuildInstanceRequestsItem and assigns it to the Requests field.
func (o *OpsRebuildInstance) SetRequests(v []OpsRebuildInstanceRequestsItem) {
	o.Requests = v
}



// MarshalJSON serializes the struct using spec logic.
func (o OpsRebuildInstance) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.IgnoreRoleCheck != nil {
		toSerialize["ignoreRoleCheck"] = o.IgnoreRoleCheck
	}
	if o.Requests != nil {
		toSerialize["requests"] = o.Requests
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OpsRebuildInstance) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		IgnoreRoleCheck *bool `json:"ignoreRoleCheck,omitempty"`
		Requests []OpsRebuildInstanceRequestsItem `json:"requests,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{ "ignoreRoleCheck", "requests",  })
	} else {
		return err
	}
	o.IgnoreRoleCheck = all.IgnoreRoleCheck
	o.Requests = all.Requests

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
