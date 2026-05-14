// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiAgentFeedbackResponse struct {
	FeedbackId string     `json:"feedbackId"`
	CreatedAt  *time.Time `json:"createdAt,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiAgentFeedbackResponse instantiates a new AiAgentFeedbackResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiAgentFeedbackResponse(feedbackId string) *AiAgentFeedbackResponse {
	this := AiAgentFeedbackResponse{}
	this.FeedbackId = feedbackId
	return &this
}

// NewAiAgentFeedbackResponseWithDefaults instantiates a new AiAgentFeedbackResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiAgentFeedbackResponseWithDefaults() *AiAgentFeedbackResponse {
	this := AiAgentFeedbackResponse{}
	return &this
}

// GetFeedbackId returns the FeedbackId field value.
func (o *AiAgentFeedbackResponse) GetFeedbackId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.FeedbackId
}

// GetFeedbackIdOk returns a tuple with the FeedbackId field value
// and a boolean to check if the value has been set.
func (o *AiAgentFeedbackResponse) GetFeedbackIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FeedbackId, true
}

// SetFeedbackId sets field value.
func (o *AiAgentFeedbackResponse) SetFeedbackId(v string) {
	o.FeedbackId = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *AiAgentFeedbackResponse) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentFeedbackResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AiAgentFeedbackResponse) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *AiAgentFeedbackResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiAgentFeedbackResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["feedbackId"] = o.FeedbackId
	if o.CreatedAt != nil {
		if o.CreatedAt.Nanosecond() == 0 {
			toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiAgentFeedbackResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		FeedbackId *string    `json:"feedbackId"`
		CreatedAt  *time.Time `json:"createdAt,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.FeedbackId == nil {
		return fmt.Errorf("required field feedbackId missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"feedbackId", "createdAt"})
	} else {
		return err
	}
	o.FeedbackId = *all.FeedbackId
	o.CreatedAt = all.CreatedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
