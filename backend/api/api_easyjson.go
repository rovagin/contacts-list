// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package api

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonC1cedd36DecodeContactsListApi(in *jlexer.Lexer, out *Response) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "rid":
			out.RID = string(in.String())
		case "code":
			out.Code = int(in.Int())
		case "payload":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Payload).UnmarshalJSON(data))
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC1cedd36EncodeContactsListApi(out *jwriter.Writer, in Response) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"rid\":"
		out.RawString(prefix[1:])
		out.String(string(in.RID))
	}
	{
		const prefix string = ",\"code\":"
		out.RawString(prefix)
		out.Int(int(in.Code))
	}
	{
		const prefix string = ",\"payload\":"
		out.RawString(prefix)
		out.Raw((in.Payload).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Response) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC1cedd36EncodeContactsListApi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Response) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC1cedd36EncodeContactsListApi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Response) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC1cedd36DecodeContactsListApi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Response) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC1cedd36DecodeContactsListApi(l, v)
}
func easyjsonC1cedd36DecodeContactsListApi1(in *jlexer.Lexer, out *Request) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "rid":
			out.RID = string(in.String())
		case "payload":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Payload).UnmarshalJSON(data))
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC1cedd36EncodeContactsListApi1(out *jwriter.Writer, in Request) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"rid\":"
		out.RawString(prefix[1:])
		out.String(string(in.RID))
	}
	{
		const prefix string = ",\"payload\":"
		out.RawString(prefix)
		out.Raw((in.Payload).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Request) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC1cedd36EncodeContactsListApi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Request) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC1cedd36EncodeContactsListApi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Request) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC1cedd36DecodeContactsListApi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Request) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC1cedd36DecodeContactsListApi1(l, v)
}
func easyjsonC1cedd36DecodeContactsListApi2(in *jlexer.Lexer, out *CreateContactResponse) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.ID = int(in.Int())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC1cedd36EncodeContactsListApi2(out *jwriter.Writer, in CreateContactResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Int(int(in.ID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CreateContactResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC1cedd36EncodeContactsListApi2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CreateContactResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC1cedd36EncodeContactsListApi2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CreateContactResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC1cedd36DecodeContactsListApi2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CreateContactResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC1cedd36DecodeContactsListApi2(l, v)
}
func easyjsonC1cedd36DecodeContactsListApi3(in *jlexer.Lexer, out *CreateContactRequest) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "first_name":
			out.FirstName = string(in.String())
		case "last_name":
			out.LastName = string(in.String())
		case "phone":
			out.Phone = string(in.String())
		case "email":
			out.Email = string(in.String())
		case "note":
			out.Note = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC1cedd36EncodeContactsListApi3(out *jwriter.Writer, in CreateContactRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"first_name\":"
		out.RawString(prefix[1:])
		out.String(string(in.FirstName))
	}
	{
		const prefix string = ",\"last_name\":"
		out.RawString(prefix)
		out.String(string(in.LastName))
	}
	{
		const prefix string = ",\"phone\":"
		out.RawString(prefix)
		out.String(string(in.Phone))
	}
	{
		const prefix string = ",\"email\":"
		out.RawString(prefix)
		out.String(string(in.Email))
	}
	{
		const prefix string = ",\"note\":"
		out.RawString(prefix)
		out.String(string(in.Note))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CreateContactRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC1cedd36EncodeContactsListApi3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CreateContactRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC1cedd36EncodeContactsListApi3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CreateContactRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC1cedd36DecodeContactsListApi3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CreateContactRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC1cedd36DecodeContactsListApi3(l, v)
}
func easyjsonC1cedd36DecodeContactsListApi4(in *jlexer.Lexer, out *Contact) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.ID = int(in.Int())
		case "first_name":
			out.FirstName = string(in.String())
		case "last_name":
			out.LastName = string(in.String())
		case "phone":
			out.Phone = string(in.String())
		case "email":
			out.Email = string(in.String())
		case "note":
			out.Note = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC1cedd36EncodeContactsListApi4(out *jwriter.Writer, in Contact) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Int(int(in.ID))
	}
	{
		const prefix string = ",\"first_name\":"
		out.RawString(prefix)
		out.String(string(in.FirstName))
	}
	{
		const prefix string = ",\"last_name\":"
		out.RawString(prefix)
		out.String(string(in.LastName))
	}
	{
		const prefix string = ",\"phone\":"
		out.RawString(prefix)
		out.String(string(in.Phone))
	}
	{
		const prefix string = ",\"email\":"
		out.RawString(prefix)
		out.String(string(in.Email))
	}
	{
		const prefix string = ",\"note\":"
		out.RawString(prefix)
		out.String(string(in.Note))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Contact) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC1cedd36EncodeContactsListApi4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Contact) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC1cedd36EncodeContactsListApi4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Contact) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC1cedd36DecodeContactsListApi4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Contact) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC1cedd36DecodeContactsListApi4(l, v)
}
