// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package client

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

func easyjsonC5a4559bDecodeGithubComKnqChromedpClient(in *jlexer.Lexer, out *Chrome) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "description":
			out.Description = string(in.String())
		case "devtoolsFrontendUrl":
			out.DevtoolsURL = string(in.String())
		case "id":
			out.ID = string(in.String())
		case "title":
			out.Title = string(in.String())
		case "type":
			(out.Type).UnmarshalEasyJSON(in)
		case "url":
			out.URL = string(in.String())
		case "webSocketDebuggerUrl":
			out.WebsocketURL = string(in.String())
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
func easyjsonC5a4559bEncodeGithubComKnqChromedpClient(out *jwriter.Writer, in Chrome) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Description != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"description\":")
		out.String(string(in.Description))
	}
	if in.DevtoolsURL != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"devtoolsFrontendUrl\":")
		out.String(string(in.DevtoolsURL))
	}
	if in.ID != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"id\":")
		out.String(string(in.ID))
	}
	if in.Title != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"title\":")
		out.String(string(in.Title))
	}
	if in.Type != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"type\":")
		(in.Type).MarshalEasyJSON(out)
	}
	if in.URL != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"url\":")
		out.String(string(in.URL))
	}
	if in.WebsocketURL != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"webSocketDebuggerUrl\":")
		out.String(string(in.WebsocketURL))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Chrome) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpClient(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Chrome) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpClient(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Chrome) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpClient(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Chrome) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpClient(l, v)
}
