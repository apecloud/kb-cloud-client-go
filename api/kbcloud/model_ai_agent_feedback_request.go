// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiAgentFeedbackRequest struct {
	Rating  AiAgentFeedbackRating   `json:"rating"`
	Reasons []AiAgentFeedbackReason `json:"reasons,omitempty"`
	Comment *string                 `json:"comment,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiAgentFeedbackRequest instantiates a new AiAgentFeedbackRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiAgentFeedbackRequest(rating AiAgentFeedbackRating) *AiAgentFeedbackRequest {
	this := AiAgentFeedbackRequest{}
	this.Rating = rating
	return &this
}

// NewAiAgentFeedbackRequestWithDefaults instantiates a new AiAgentFeedbackRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiAgentFeedbackRequestWithDefaults() *AiAgentFeedbackRequest {
	this := AiAgentFeedbackRequest{}
	return &this
}

// GetRating returns the Rating field value.
func (o *AiAgentFeedbackRequest) GetRating() AiAgentFeedbackRating {
	if o == nil {
		var ret AiAgentFeedbackRating
		return ret
	}
	return o.Rating
}

// GetRatingOk returns a tuple with the Rating field value
// and a boolean to check if the value has been set.
func (o *AiAgentFeedbackRequest) GetRatingOk() (*AiAgentFeedbackRating, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rating, true
}

// SetRating sets field value.
func (o *AiAgentFeedbackRequest) SetRating(v AiAgentFeedbackRating) {
	o.Rating = v
}

// GetReasons returns the Reasons field value if set, zero value otherwise.
func (o *AiAgentFeedbackRequest) GetReasons() []AiAgentFeedbackReason {
	if o == nil || o.Reasons == nil {
		var ret []AiAgentFeedbackReason
		return ret
	}
	return o.Reasons
}

// GetReasonsOk returns a tuple with the Reasons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentFeedbackRequest) GetReasonsOk() (*[]AiAgentFeedbackReason, bool) {
	if o == nil || o.Reasons == nil {
		return nil, false
	}
	return &o.Reasons, true
}

// HasReasons returns a boolean if a field has been set.
func (o *AiAgentFeedbackRequest) HasReasons() bool {
	return o != nil && o.Reasons != nil
}

// SetReasons gets a reference to the given []AiAgentFeedbackReason and assigns it to the Reasons field.
func (o *AiAgentFeedbackRequest) SetReasons(v []AiAgentFeedbackReason) {
	o.Reasons = v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *AiAgentFeedbackRequest) GetComment() string {
	if o == nil || o.Comment == nil {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentFeedbackRequest) GetCommentOk() (*string, bool) {
	if o == nil || o.Comment == nil {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *AiAgentFeedbackRequest) HasComment() bool {
	return o != nil && o.Comment != nil
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *AiAgentFeedbackRequest) SetComment(v string) {
	o.Comment = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiAgentFeedbackRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["rating"] = o.Rating
	if o.Reasons != nil {
		toSerialize["reasons"] = o.Reasons
	}
	if o.Comment != nil {
		toSerialize["comment"] = o.Comment
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiAgentFeedbackRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Rating  *AiAgentFeedbackRating  `json:"rating"`
		Reasons []AiAgentFeedbackReason `json:"reasons,omitempty"`
		Comment *string                 `json:"comment,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Rating == nil {
		return fmt.Errorf("required field rating missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"rating", "reasons", "comment"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.Rating.IsValid() {
		hasInvalidField = true
	} else {
		o.Rating = *all.Rating
	}
	o.Reasons = all.Reasons
	o.Comment = all.Comment

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
