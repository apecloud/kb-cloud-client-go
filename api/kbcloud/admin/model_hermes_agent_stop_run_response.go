// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type HermesAgentStopRunResponse struct {
	RunId  string               `json:"runId"`
	Status HermesAgentRunStatus `json:"status"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewHermesAgentStopRunResponse instantiates a new HermesAgentStopRunResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewHermesAgentStopRunResponse(runId string, status HermesAgentRunStatus) *HermesAgentStopRunResponse {
	this := HermesAgentStopRunResponse{}
	this.RunId = runId
	this.Status = status
	return &this
}

// NewHermesAgentStopRunResponseWithDefaults instantiates a new HermesAgentStopRunResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewHermesAgentStopRunResponseWithDefaults() *HermesAgentStopRunResponse {
	this := HermesAgentStopRunResponse{}
	return &this
}

// GetRunId returns the RunId field value.
func (o *HermesAgentStopRunResponse) GetRunId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.RunId
}

// GetRunIdOk returns a tuple with the RunId field value
// and a boolean to check if the value has been set.
func (o *HermesAgentStopRunResponse) GetRunIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RunId, true
}

// SetRunId sets field value.
func (o *HermesAgentStopRunResponse) SetRunId(v string) {
	o.RunId = v
}

// GetStatus returns the Status field value.
func (o *HermesAgentStopRunResponse) GetStatus() HermesAgentRunStatus {
	if o == nil {
		var ret HermesAgentRunStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *HermesAgentStopRunResponse) GetStatusOk() (*HermesAgentRunStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *HermesAgentStopRunResponse) SetStatus(v HermesAgentRunStatus) {
	o.Status = v
}

// MarshalJSON serializes the struct using spec logic.
func (o HermesAgentStopRunResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["runId"] = o.RunId
	toSerialize["status"] = o.Status

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *HermesAgentStopRunResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		RunId  *string               `json:"runId"`
		Status *HermesAgentRunStatus `json:"status"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.RunId == nil {
		return fmt.Errorf("required field runId missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"runId", "status"})
	} else {
		return err
	}

	hasInvalidField := false
	o.RunId = *all.RunId
	if !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = *all.Status
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
