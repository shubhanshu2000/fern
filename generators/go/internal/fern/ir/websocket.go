// This file was auto-generated by Fern from our API Definition.

package ir

import (
	json "encoding/json"
	fmt "fmt"

	internal "github.com/fern-api/fern-go/internal/fern/ir/internal"
)

type ExampleWebSocketMessage struct {
	Type WebSocketMessageId           `json:"type" url:"type"`
	Body *ExampleWebSocketMessageBody `json:"body,omitempty" url:"body,omitempty"`

	extraProperties map[string]interface{}
}

func (e *ExampleWebSocketMessage) GetType() WebSocketMessageId {
	if e == nil {
		return ""
	}
	return e.Type
}

func (e *ExampleWebSocketMessage) GetBody() *ExampleWebSocketMessageBody {
	if e == nil {
		return nil
	}
	return e.Body
}

func (e *ExampleWebSocketMessage) GetExtraProperties() map[string]interface{} {
	return e.extraProperties
}

func (e *ExampleWebSocketMessage) UnmarshalJSON(data []byte) error {
	type unmarshaler ExampleWebSocketMessage
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*e = ExampleWebSocketMessage(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *e)
	if err != nil {
		return err
	}
	e.extraProperties = extraProperties
	return nil
}

func (e *ExampleWebSocketMessage) String() string {
	if value, err := internal.StringifyJSON(e); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", e)
}

type ExampleWebSocketMessageBody struct {
	Type        string
	InlinedBody *ExampleInlinedRequestBody
	Reference   *ExampleTypeReference
}

func NewExampleWebSocketMessageBodyFromInlinedBody(value *ExampleInlinedRequestBody) *ExampleWebSocketMessageBody {
	return &ExampleWebSocketMessageBody{Type: "inlinedBody", InlinedBody: value}
}

func NewExampleWebSocketMessageBodyFromReference(value *ExampleTypeReference) *ExampleWebSocketMessageBody {
	return &ExampleWebSocketMessageBody{Type: "reference", Reference: value}
}

func (e *ExampleWebSocketMessageBody) GetType() string {
	if e == nil {
		return ""
	}
	return e.Type
}

func (e *ExampleWebSocketMessageBody) GetInlinedBody() *ExampleInlinedRequestBody {
	if e == nil {
		return nil
	}
	return e.InlinedBody
}

func (e *ExampleWebSocketMessageBody) GetReference() *ExampleTypeReference {
	if e == nil {
		return nil
	}
	return e.Reference
}

func (e *ExampleWebSocketMessageBody) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	e.Type = unmarshaler.Type
	if unmarshaler.Type == "" {
		return fmt.Errorf("%T did not include discriminant type", e)
	}
	switch unmarshaler.Type {
	case "inlinedBody":
		value := new(ExampleInlinedRequestBody)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		e.InlinedBody = value
	case "reference":
		value := new(ExampleTypeReference)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		e.Reference = value
	}
	return nil
}

func (e ExampleWebSocketMessageBody) MarshalJSON() ([]byte, error) {
	if err := e.validate(); err != nil {
		return nil, err
	}
	switch e.Type {
	default:
		return nil, fmt.Errorf("invalid type %s in %T", e.Type, e)
	case "inlinedBody":
		return internal.MarshalJSONWithExtraProperty(e.InlinedBody, "type", "inlinedBody")
	case "reference":
		return internal.MarshalJSONWithExtraProperty(e.Reference, "type", "reference")
	}
}

type ExampleWebSocketMessageBodyVisitor interface {
	VisitInlinedBody(*ExampleInlinedRequestBody) error
	VisitReference(*ExampleTypeReference) error
}

func (e *ExampleWebSocketMessageBody) Accept(visitor ExampleWebSocketMessageBodyVisitor) error {
	switch e.Type {
	default:
		return fmt.Errorf("invalid type %s in %T", e.Type, e)
	case "inlinedBody":
		return visitor.VisitInlinedBody(e.InlinedBody)
	case "reference":
		return visitor.VisitReference(e.Reference)
	}
}

func (e *ExampleWebSocketMessageBody) validate() error {
	if e == nil {
		return fmt.Errorf("type %T is nil", e)
	}
	var fields []string
	if e.InlinedBody != nil {
		fields = append(fields, "inlinedBody")
	}
	if e.Reference != nil {
		fields = append(fields, "reference")
	}
	if len(fields) == 0 {
		if e.Type != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", e, e.Type)
		}
		return fmt.Errorf("type %T is empty", e)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", e, fields)
	}
	if e.Type != "" {
		field := fields[0]
		if e.Type != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				e,
				e.Type,
				e,
			)
		}
	}
	return nil
}

type ExampleWebSocketSession struct {
	Docs            *string                    `json:"docs,omitempty" url:"docs,omitempty"`
	Name            *Name                      `json:"name,omitempty" url:"name,omitempty"`
	Url             string                     `json:"url" url:"url"`
	PathParameters  []*ExamplePathParameter    `json:"pathParameters,omitempty" url:"pathParameters,omitempty"`
	Headers         []*ExampleHeader           `json:"headers,omitempty" url:"headers,omitempty"`
	QueryParameters []*ExampleQueryParameter   `json:"queryParameters,omitempty" url:"queryParameters,omitempty"`
	Messages        []*ExampleWebSocketMessage `json:"messages,omitempty" url:"messages,omitempty"`

	extraProperties map[string]interface{}
}

func (e *ExampleWebSocketSession) GetDocs() *string {
	if e == nil {
		return nil
	}
	return e.Docs
}

func (e *ExampleWebSocketSession) GetName() *Name {
	if e == nil {
		return nil
	}
	return e.Name
}

func (e *ExampleWebSocketSession) GetUrl() string {
	if e == nil {
		return ""
	}
	return e.Url
}

func (e *ExampleWebSocketSession) GetPathParameters() []*ExamplePathParameter {
	if e == nil {
		return nil
	}
	return e.PathParameters
}

func (e *ExampleWebSocketSession) GetHeaders() []*ExampleHeader {
	if e == nil {
		return nil
	}
	return e.Headers
}

func (e *ExampleWebSocketSession) GetQueryParameters() []*ExampleQueryParameter {
	if e == nil {
		return nil
	}
	return e.QueryParameters
}

func (e *ExampleWebSocketSession) GetMessages() []*ExampleWebSocketMessage {
	if e == nil {
		return nil
	}
	return e.Messages
}

func (e *ExampleWebSocketSession) GetExtraProperties() map[string]interface{} {
	return e.extraProperties
}

func (e *ExampleWebSocketSession) UnmarshalJSON(data []byte) error {
	type unmarshaler ExampleWebSocketSession
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*e = ExampleWebSocketSession(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *e)
	if err != nil {
		return err
	}
	e.extraProperties = extraProperties
	return nil
}

func (e *ExampleWebSocketSession) String() string {
	if value, err := internal.StringifyJSON(e); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", e)
}

type InlinedWebSocketMessageBody struct {
	Name       *Name                                  `json:"name,omitempty" url:"name,omitempty"`
	Extends    []*DeclaredTypeName                    `json:"extends,omitempty" url:"extends,omitempty"`
	Properties []*InlinedWebSocketMessageBodyProperty `json:"properties,omitempty" url:"properties,omitempty"`

	extraProperties map[string]interface{}
}

func (i *InlinedWebSocketMessageBody) GetName() *Name {
	if i == nil {
		return nil
	}
	return i.Name
}

func (i *InlinedWebSocketMessageBody) GetExtends() []*DeclaredTypeName {
	if i == nil {
		return nil
	}
	return i.Extends
}

func (i *InlinedWebSocketMessageBody) GetProperties() []*InlinedWebSocketMessageBodyProperty {
	if i == nil {
		return nil
	}
	return i.Properties
}

func (i *InlinedWebSocketMessageBody) GetExtraProperties() map[string]interface{} {
	return i.extraProperties
}

func (i *InlinedWebSocketMessageBody) UnmarshalJSON(data []byte) error {
	type unmarshaler InlinedWebSocketMessageBody
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*i = InlinedWebSocketMessageBody(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *i)
	if err != nil {
		return err
	}
	i.extraProperties = extraProperties
	return nil
}

func (i *InlinedWebSocketMessageBody) String() string {
	if value, err := internal.StringifyJSON(i); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", i)
}

type InlinedWebSocketMessageBodyProperty struct {
	Docs         *string           `json:"docs,omitempty" url:"docs,omitempty"`
	Availability *Availability     `json:"availability,omitempty" url:"availability,omitempty"`
	Name         *NameAndWireValue `json:"name,omitempty" url:"name,omitempty"`
	ValueType    *TypeReference    `json:"valueType,omitempty" url:"valueType,omitempty"`

	extraProperties map[string]interface{}
}

func (i *InlinedWebSocketMessageBodyProperty) GetDocs() *string {
	if i == nil {
		return nil
	}
	return i.Docs
}

func (i *InlinedWebSocketMessageBodyProperty) GetAvailability() *Availability {
	if i == nil {
		return nil
	}
	return i.Availability
}

func (i *InlinedWebSocketMessageBodyProperty) GetName() *NameAndWireValue {
	if i == nil {
		return nil
	}
	return i.Name
}

func (i *InlinedWebSocketMessageBodyProperty) GetValueType() *TypeReference {
	if i == nil {
		return nil
	}
	return i.ValueType
}

func (i *InlinedWebSocketMessageBodyProperty) GetExtraProperties() map[string]interface{} {
	return i.extraProperties
}

func (i *InlinedWebSocketMessageBodyProperty) UnmarshalJSON(data []byte) error {
	type unmarshaler InlinedWebSocketMessageBodyProperty
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*i = InlinedWebSocketMessageBodyProperty(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *i)
	if err != nil {
		return err
	}
	i.extraProperties = extraProperties
	return nil
}

func (i *InlinedWebSocketMessageBodyProperty) String() string {
	if value, err := internal.StringifyJSON(i); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", i)
}

type WebSocketChannel struct {
	Docs            *string               `json:"docs,omitempty" url:"docs,omitempty"`
	Availability    *Availability         `json:"availability,omitempty" url:"availability,omitempty"`
	Name            WebSocketName         `json:"name,omitempty" url:"name,omitempty"`
	DisplayName     *string               `json:"displayName,omitempty" url:"displayName,omitempty"`
	BaseUrl         *EnvironmentBaseUrlId `json:"baseUrl,omitempty" url:"baseUrl,omitempty"`
	Path            *HttpPath             `json:"path,omitempty" url:"path,omitempty"`
	Auth            bool                  `json:"auth" url:"auth"`
	Headers         []*HttpHeader         `json:"headers,omitempty" url:"headers,omitempty"`
	QueryParameters []*QueryParameter     `json:"queryParameters,omitempty" url:"queryParameters,omitempty"`
	PathParameters  []*PathParameter      `json:"pathParameters,omitempty" url:"pathParameters,omitempty"`
	// The messages that can be sent and received on this channel
	Messages []*WebSocketMessage        `json:"messages,omitempty" url:"messages,omitempty"`
	Examples []*ExampleWebSocketSession `json:"examples,omitempty" url:"examples,omitempty"`

	extraProperties map[string]interface{}
}

func (w *WebSocketChannel) GetDocs() *string {
	if w == nil {
		return nil
	}
	return w.Docs
}

func (w *WebSocketChannel) GetAvailability() *Availability {
	if w == nil {
		return nil
	}
	return w.Availability
}

func (w *WebSocketChannel) GetName() WebSocketName {
	if w == nil {
		return nil
	}
	return w.Name
}

func (w *WebSocketChannel) GetDisplayName() *string {
	if w == nil {
		return nil
	}
	return w.DisplayName
}

func (w *WebSocketChannel) GetBaseUrl() *EnvironmentBaseUrlId {
	if w == nil {
		return nil
	}
	return w.BaseUrl
}

func (w *WebSocketChannel) GetPath() *HttpPath {
	if w == nil {
		return nil
	}
	return w.Path
}

func (w *WebSocketChannel) GetAuth() bool {
	if w == nil {
		return false
	}
	return w.Auth
}

func (w *WebSocketChannel) GetHeaders() []*HttpHeader {
	if w == nil {
		return nil
	}
	return w.Headers
}

func (w *WebSocketChannel) GetQueryParameters() []*QueryParameter {
	if w == nil {
		return nil
	}
	return w.QueryParameters
}

func (w *WebSocketChannel) GetPathParameters() []*PathParameter {
	if w == nil {
		return nil
	}
	return w.PathParameters
}

func (w *WebSocketChannel) GetMessages() []*WebSocketMessage {
	if w == nil {
		return nil
	}
	return w.Messages
}

func (w *WebSocketChannel) GetExamples() []*ExampleWebSocketSession {
	if w == nil {
		return nil
	}
	return w.Examples
}

func (w *WebSocketChannel) GetExtraProperties() map[string]interface{} {
	return w.extraProperties
}

func (w *WebSocketChannel) UnmarshalJSON(data []byte) error {
	type unmarshaler WebSocketChannel
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*w = WebSocketChannel(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *w)
	if err != nil {
		return err
	}
	w.extraProperties = extraProperties
	return nil
}

func (w *WebSocketChannel) String() string {
	if value, err := internal.StringifyJSON(w); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", w)
}

type WebSocketMessage struct {
	Docs         *string                `json:"docs,omitempty" url:"docs,omitempty"`
	Availability *Availability          `json:"availability,omitempty" url:"availability,omitempty"`
	Type         WebSocketMessageId     `json:"type" url:"type"`
	DisplayName  *string                `json:"displayName,omitempty" url:"displayName,omitempty"`
	Origin       WebSocketMessageOrigin `json:"origin" url:"origin"`
	Body         *WebSocketMessageBody  `json:"body,omitempty" url:"body,omitempty"`

	extraProperties map[string]interface{}
}

func (w *WebSocketMessage) GetDocs() *string {
	if w == nil {
		return nil
	}
	return w.Docs
}

func (w *WebSocketMessage) GetAvailability() *Availability {
	if w == nil {
		return nil
	}
	return w.Availability
}

func (w *WebSocketMessage) GetType() WebSocketMessageId {
	if w == nil {
		return ""
	}
	return w.Type
}

func (w *WebSocketMessage) GetDisplayName() *string {
	if w == nil {
		return nil
	}
	return w.DisplayName
}

func (w *WebSocketMessage) GetOrigin() WebSocketMessageOrigin {
	if w == nil {
		return ""
	}
	return w.Origin
}

func (w *WebSocketMessage) GetBody() *WebSocketMessageBody {
	if w == nil {
		return nil
	}
	return w.Body
}

func (w *WebSocketMessage) GetExtraProperties() map[string]interface{} {
	return w.extraProperties
}

func (w *WebSocketMessage) UnmarshalJSON(data []byte) error {
	type unmarshaler WebSocketMessage
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*w = WebSocketMessage(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *w)
	if err != nil {
		return err
	}
	w.extraProperties = extraProperties
	return nil
}

func (w *WebSocketMessage) String() string {
	if value, err := internal.StringifyJSON(w); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", w)
}

type WebSocketMessageBody struct {
	Type        string
	InlinedBody *InlinedWebSocketMessageBody
	Reference   *WebSocketMessageBodyReference
}

func NewWebSocketMessageBodyFromInlinedBody(value *InlinedWebSocketMessageBody) *WebSocketMessageBody {
	return &WebSocketMessageBody{Type: "inlinedBody", InlinedBody: value}
}

func NewWebSocketMessageBodyFromReference(value *WebSocketMessageBodyReference) *WebSocketMessageBody {
	return &WebSocketMessageBody{Type: "reference", Reference: value}
}

func (w *WebSocketMessageBody) GetType() string {
	if w == nil {
		return ""
	}
	return w.Type
}

func (w *WebSocketMessageBody) GetInlinedBody() *InlinedWebSocketMessageBody {
	if w == nil {
		return nil
	}
	return w.InlinedBody
}

func (w *WebSocketMessageBody) GetReference() *WebSocketMessageBodyReference {
	if w == nil {
		return nil
	}
	return w.Reference
}

func (w *WebSocketMessageBody) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	w.Type = unmarshaler.Type
	if unmarshaler.Type == "" {
		return fmt.Errorf("%T did not include discriminant type", w)
	}
	switch unmarshaler.Type {
	case "inlinedBody":
		value := new(InlinedWebSocketMessageBody)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		w.InlinedBody = value
	case "reference":
		value := new(WebSocketMessageBodyReference)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		w.Reference = value
	}
	return nil
}

func (w WebSocketMessageBody) MarshalJSON() ([]byte, error) {
	if err := w.validate(); err != nil {
		return nil, err
	}
	switch w.Type {
	default:
		return nil, fmt.Errorf("invalid type %s in %T", w.Type, w)
	case "inlinedBody":
		return internal.MarshalJSONWithExtraProperty(w.InlinedBody, "type", "inlinedBody")
	case "reference":
		return internal.MarshalJSONWithExtraProperty(w.Reference, "type", "reference")
	}
}

type WebSocketMessageBodyVisitor interface {
	VisitInlinedBody(*InlinedWebSocketMessageBody) error
	VisitReference(*WebSocketMessageBodyReference) error
}

func (w *WebSocketMessageBody) Accept(visitor WebSocketMessageBodyVisitor) error {
	switch w.Type {
	default:
		return fmt.Errorf("invalid type %s in %T", w.Type, w)
	case "inlinedBody":
		return visitor.VisitInlinedBody(w.InlinedBody)
	case "reference":
		return visitor.VisitReference(w.Reference)
	}
}

func (w *WebSocketMessageBody) validate() error {
	if w == nil {
		return fmt.Errorf("type %T is nil", w)
	}
	var fields []string
	if w.InlinedBody != nil {
		fields = append(fields, "inlinedBody")
	}
	if w.Reference != nil {
		fields = append(fields, "reference")
	}
	if len(fields) == 0 {
		if w.Type != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", w, w.Type)
		}
		return fmt.Errorf("type %T is empty", w)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", w, fields)
	}
	if w.Type != "" {
		field := fields[0]
		if w.Type != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				w,
				w.Type,
				w,
			)
		}
	}
	return nil
}

type WebSocketMessageBodyReference struct {
	Docs     *string        `json:"docs,omitempty" url:"docs,omitempty"`
	BodyType *TypeReference `json:"bodyType,omitempty" url:"bodyType,omitempty"`

	extraProperties map[string]interface{}
}

func (w *WebSocketMessageBodyReference) GetDocs() *string {
	if w == nil {
		return nil
	}
	return w.Docs
}

func (w *WebSocketMessageBodyReference) GetBodyType() *TypeReference {
	if w == nil {
		return nil
	}
	return w.BodyType
}

func (w *WebSocketMessageBodyReference) GetExtraProperties() map[string]interface{} {
	return w.extraProperties
}

func (w *WebSocketMessageBodyReference) UnmarshalJSON(data []byte) error {
	type unmarshaler WebSocketMessageBodyReference
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*w = WebSocketMessageBodyReference(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *w)
	if err != nil {
		return err
	}
	w.extraProperties = extraProperties
	return nil
}

func (w *WebSocketMessageBodyReference) String() string {
	if value, err := internal.StringifyJSON(w); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", w)
}

type WebSocketMessageId = string

type WebSocketMessageOrigin string

const (
	WebSocketMessageOriginClient WebSocketMessageOrigin = "client"
	WebSocketMessageOriginServer WebSocketMessageOrigin = "server"
)

func NewWebSocketMessageOriginFromString(s string) (WebSocketMessageOrigin, error) {
	switch s {
	case "client":
		return WebSocketMessageOriginClient, nil
	case "server":
		return WebSocketMessageOriginServer, nil
	}
	var t WebSocketMessageOrigin
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (w WebSocketMessageOrigin) Ptr() *WebSocketMessageOrigin {
	return &w
}

type WebSocketName = *Name
