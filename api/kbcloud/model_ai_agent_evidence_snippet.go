// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiAgentEvidenceSnippet struct {
	SnippetId string     `json:"snippetId"`
	Time      *time.Time `json:"time,omitempty"`
	Text      string     `json:"text"`
	Source    *string    `json:"source,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiAgentEvidenceSnippet instantiates a new AiAgentEvidenceSnippet object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiAgentEvidenceSnippet(snippetId string, text string) *AiAgentEvidenceSnippet {
	this := AiAgentEvidenceSnippet{}
	this.SnippetId = snippetId
	this.Text = text
	return &this
}

// NewAiAgentEvidenceSnippetWithDefaults instantiates a new AiAgentEvidenceSnippet object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiAgentEvidenceSnippetWithDefaults() *AiAgentEvidenceSnippet {
	this := AiAgentEvidenceSnippet{}
	return &this
}

// GetSnippetId returns the SnippetId field value.
func (o *AiAgentEvidenceSnippet) GetSnippetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SnippetId
}

// GetSnippetIdOk returns a tuple with the SnippetId field value
// and a boolean to check if the value has been set.
func (o *AiAgentEvidenceSnippet) GetSnippetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SnippetId, true
}

// SetSnippetId sets field value.
func (o *AiAgentEvidenceSnippet) SetSnippetId(v string) {
	o.SnippetId = v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *AiAgentEvidenceSnippet) GetTime() time.Time {
	if o == nil || o.Time == nil {
		var ret time.Time
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentEvidenceSnippet) GetTimeOk() (*time.Time, bool) {
	if o == nil || o.Time == nil {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *AiAgentEvidenceSnippet) HasTime() bool {
	return o != nil && o.Time != nil
}

// SetTime gets a reference to the given time.Time and assigns it to the Time field.
func (o *AiAgentEvidenceSnippet) SetTime(v time.Time) {
	o.Time = &v
}

// GetText returns the Text field value.
func (o *AiAgentEvidenceSnippet) GetText() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Text
}

// GetTextOk returns a tuple with the Text field value
// and a boolean to check if the value has been set.
func (o *AiAgentEvidenceSnippet) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Text, true
}

// SetText sets field value.
func (o *AiAgentEvidenceSnippet) SetText(v string) {
	o.Text = v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *AiAgentEvidenceSnippet) GetSource() string {
	if o == nil || o.Source == nil {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentEvidenceSnippet) GetSourceOk() (*string, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *AiAgentEvidenceSnippet) HasSource() bool {
	return o != nil && o.Source != nil
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *AiAgentEvidenceSnippet) SetSource(v string) {
	o.Source = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiAgentEvidenceSnippet) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["snippetId"] = o.SnippetId
	if o.Time != nil {
		if o.Time.Nanosecond() == 0 {
			toSerialize["time"] = o.Time.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["time"] = o.Time.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	toSerialize["text"] = o.Text
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiAgentEvidenceSnippet) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		SnippetId *string    `json:"snippetId"`
		Time      *time.Time `json:"time,omitempty"`
		Text      *string    `json:"text"`
		Source    *string    `json:"source,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.SnippetId == nil {
		return fmt.Errorf("required field snippetId missing")
	}
	if all.Text == nil {
		return fmt.Errorf("required field text missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"snippetId", "time", "text", "source"})
	} else {
		return err
	}
	o.SnippetId = *all.SnippetId
	o.Time = all.Time
	o.Text = *all.Text
	o.Source = all.Source

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
