// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

type MongoShellSession struct {
	// shell session id
	SessionId *string `json:"sessionId,omitempty"`
	// current mongosh prompt
	Prompt *string `json:"prompt,omitempty"`
	// session state
	State *string `json:"state,omitempty"`
	// creation time
	CreatedAt *string `json:"createdAt,omitempty"`
	// expiration time
	ExpiresAt *string `json:"expiresAt,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMongoShellSession instantiates a new MongoShellSession object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMongoShellSession() *MongoShellSession {
	this := MongoShellSession{}
	return &this
}

// NewMongoShellSessionWithDefaults instantiates a new MongoShellSession object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMongoShellSessionWithDefaults() *MongoShellSession {
	this := MongoShellSession{}
	return &this
}

// GetSessionId returns the SessionId field value if set, zero value otherwise.
func (o *MongoShellSession) GetSessionId() string {
	if o == nil || o.SessionId == nil {
		var ret string
		return ret
	}
	return *o.SessionId
}

// GetSessionIdOk returns a tuple with the SessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoShellSession) GetSessionIdOk() (*string, bool) {
	if o == nil || o.SessionId == nil {
		return nil, false
	}
	return o.SessionId, true
}

// HasSessionId returns a boolean if a field has been set.
func (o *MongoShellSession) HasSessionId() bool {
	return o != nil && o.SessionId != nil
}

// SetSessionId gets a reference to the given string and assigns it to the SessionId field.
func (o *MongoShellSession) SetSessionId(v string) {
	o.SessionId = &v
}

// GetPrompt returns the Prompt field value if set, zero value otherwise.
func (o *MongoShellSession) GetPrompt() string {
	if o == nil || o.Prompt == nil {
		var ret string
		return ret
	}
	return *o.Prompt
}

// GetPromptOk returns a tuple with the Prompt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoShellSession) GetPromptOk() (*string, bool) {
	if o == nil || o.Prompt == nil {
		return nil, false
	}
	return o.Prompt, true
}

// HasPrompt returns a boolean if a field has been set.
func (o *MongoShellSession) HasPrompt() bool {
	return o != nil && o.Prompt != nil
}

// SetPrompt gets a reference to the given string and assigns it to the Prompt field.
func (o *MongoShellSession) SetPrompt(v string) {
	o.Prompt = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *MongoShellSession) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoShellSession) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *MongoShellSession) HasState() bool {
	return o != nil && o.State != nil
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *MongoShellSession) SetState(v string) {
	o.State = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *MongoShellSession) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoShellSession) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *MongoShellSession) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *MongoShellSession) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *MongoShellSession) GetExpiresAt() string {
	if o == nil || o.ExpiresAt == nil {
		var ret string
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoShellSession) GetExpiresAtOk() (*string, bool) {
	if o == nil || o.ExpiresAt == nil {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *MongoShellSession) HasExpiresAt() bool {
	return o != nil && o.ExpiresAt != nil
}

// SetExpiresAt gets a reference to the given string and assigns it to the ExpiresAt field.
func (o *MongoShellSession) SetExpiresAt(v string) {
	o.ExpiresAt = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o MongoShellSession) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.SessionId != nil {
		toSerialize["sessionId"] = o.SessionId
	}
	if o.Prompt != nil {
		toSerialize["prompt"] = o.Prompt
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.CreatedAt != nil {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if o.ExpiresAt != nil {
		toSerialize["expiresAt"] = o.ExpiresAt
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MongoShellSession) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		SessionId *string `json:"sessionId,omitempty"`
		Prompt    *string `json:"prompt,omitempty"`
		State     *string `json:"state,omitempty"`
		CreatedAt *string `json:"createdAt,omitempty"`
		ExpiresAt *string `json:"expiresAt,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"sessionId", "prompt", "state", "createdAt", "expiresAt"})
	} else {
		return err
	}
	o.SessionId = all.SessionId
	o.Prompt = all.Prompt
	o.State = all.State
	o.CreatedAt = all.CreatedAt
	o.ExpiresAt = all.ExpiresAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
