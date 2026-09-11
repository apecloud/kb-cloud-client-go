// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type MssqlWait struct {
	SessionId int64 `json:"sessionId"`
	// Request ID, or -1 when a waiting task cannot be associated with a live request.
	RequestId int64 `json:"requestId"`
	// SQL Server blocking owner. Negative values are special owners, not session IDs.
	BlockingSessionId int64  `json:"blockingSessionId"`
	WaitType          string `json:"waitType"`
	WaitMs            int64  `json:"waitMs"`
	Resource          string `json:"resource"`
	Source            string `json:"source"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMssqlWait instantiates a new MssqlWait object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMssqlWait(sessionId int64, requestId int64, blockingSessionId int64, waitType string, waitMs int64, resource string, source string) *MssqlWait {
	this := MssqlWait{}
	this.SessionId = sessionId
	this.RequestId = requestId
	this.BlockingSessionId = blockingSessionId
	this.WaitType = waitType
	this.WaitMs = waitMs
	this.Resource = resource
	this.Source = source
	return &this
}

// NewMssqlWaitWithDefaults instantiates a new MssqlWait object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMssqlWaitWithDefaults() *MssqlWait {
	this := MssqlWait{}
	return &this
}

// GetSessionId returns the SessionId field value.
func (o *MssqlWait) GetSessionId() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.SessionId
}

// GetSessionIdOk returns a tuple with the SessionId field value
// and a boolean to check if the value has been set.
func (o *MssqlWait) GetSessionIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionId, true
}

// SetSessionId sets field value.
func (o *MssqlWait) SetSessionId(v int64) {
	o.SessionId = v
}

// GetRequestId returns the RequestId field value.
func (o *MssqlWait) GetRequestId() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *MssqlWait) GetRequestIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value.
func (o *MssqlWait) SetRequestId(v int64) {
	o.RequestId = v
}

// GetBlockingSessionId returns the BlockingSessionId field value.
func (o *MssqlWait) GetBlockingSessionId() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.BlockingSessionId
}

// GetBlockingSessionIdOk returns a tuple with the BlockingSessionId field value
// and a boolean to check if the value has been set.
func (o *MssqlWait) GetBlockingSessionIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockingSessionId, true
}

// SetBlockingSessionId sets field value.
func (o *MssqlWait) SetBlockingSessionId(v int64) {
	o.BlockingSessionId = v
}

// GetWaitType returns the WaitType field value.
func (o *MssqlWait) GetWaitType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.WaitType
}

// GetWaitTypeOk returns a tuple with the WaitType field value
// and a boolean to check if the value has been set.
func (o *MssqlWait) GetWaitTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WaitType, true
}

// SetWaitType sets field value.
func (o *MssqlWait) SetWaitType(v string) {
	o.WaitType = v
}

// GetWaitMs returns the WaitMs field value.
func (o *MssqlWait) GetWaitMs() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.WaitMs
}

// GetWaitMsOk returns a tuple with the WaitMs field value
// and a boolean to check if the value has been set.
func (o *MssqlWait) GetWaitMsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WaitMs, true
}

// SetWaitMs sets field value.
func (o *MssqlWait) SetWaitMs(v int64) {
	o.WaitMs = v
}

// GetResource returns the Resource field value.
func (o *MssqlWait) GetResource() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Resource
}

// GetResourceOk returns a tuple with the Resource field value
// and a boolean to check if the value has been set.
func (o *MssqlWait) GetResourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Resource, true
}

// SetResource sets field value.
func (o *MssqlWait) SetResource(v string) {
	o.Resource = v
}

// GetSource returns the Source field value.
func (o *MssqlWait) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *MssqlWait) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value.
func (o *MssqlWait) SetSource(v string) {
	o.Source = v
}

// MarshalJSON serializes the struct using spec logic.
func (o MssqlWait) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["sessionId"] = o.SessionId
	toSerialize["requestId"] = o.RequestId
	toSerialize["blockingSessionId"] = o.BlockingSessionId
	toSerialize["waitType"] = o.WaitType
	toSerialize["waitMs"] = o.WaitMs
	toSerialize["resource"] = o.Resource
	toSerialize["source"] = o.Source

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MssqlWait) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		SessionId         *int64  `json:"sessionId"`
		RequestId         *int64  `json:"requestId"`
		BlockingSessionId *int64  `json:"blockingSessionId"`
		WaitType          *string `json:"waitType"`
		WaitMs            *int64  `json:"waitMs"`
		Resource          *string `json:"resource"`
		Source            *string `json:"source"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.SessionId == nil {
		return fmt.Errorf("required field sessionId missing")
	}
	if all.RequestId == nil {
		return fmt.Errorf("required field requestId missing")
	}
	if all.BlockingSessionId == nil {
		return fmt.Errorf("required field blockingSessionId missing")
	}
	if all.WaitType == nil {
		return fmt.Errorf("required field waitType missing")
	}
	if all.WaitMs == nil {
		return fmt.Errorf("required field waitMs missing")
	}
	if all.Resource == nil {
		return fmt.Errorf("required field resource missing")
	}
	if all.Source == nil {
		return fmt.Errorf("required field source missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"sessionId", "requestId", "blockingSessionId", "waitType", "waitMs", "resource", "source"})
	} else {
		return err
	}
	o.SessionId = *all.SessionId
	o.RequestId = *all.RequestId
	o.BlockingSessionId = *all.BlockingSessionId
	o.WaitType = *all.WaitType
	o.WaitMs = *all.WaitMs
	o.Resource = *all.Resource
	o.Source = *all.Source

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
