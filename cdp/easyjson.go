// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package cdp

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

func easyjsonC5a4559bDecodeGithubComKnqChromedpCdp(in *jlexer.Lexer, out *RGBA) {
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
		case "r":
			out.R = int64(in.Int64())
		case "g":
			out.G = int64(in.Int64())
		case "b":
			out.B = int64(in.Int64())
		case "a":
			out.A = float64(in.Float64())
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
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdp(out *jwriter.Writer, in RGBA) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"r\":")
	out.Int64(int64(in.R))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"g\":")
	out.Int64(int64(in.G))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"b\":")
	out.Int64(int64(in.B))
	if in.A != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"a\":")
		out.Float64(float64(in.A))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v RGBA) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdp(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v RGBA) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdp(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *RGBA) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdp(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RGBA) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdp(l, v)
}
func easyjsonC5a4559bDecodeGithubComKnqChromedpCdp1(in *jlexer.Lexer, out *Node) {
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
		case "nodeId":
			(out.NodeID).UnmarshalEasyJSON(in)
		case "parentId":
			(out.ParentID).UnmarshalEasyJSON(in)
		case "backendNodeId":
			(out.BackendNodeID).UnmarshalEasyJSON(in)
		case "nodeType":
			(out.NodeType).UnmarshalEasyJSON(in)
		case "nodeName":
			out.NodeName = string(in.String())
		case "localName":
			out.LocalName = string(in.String())
		case "nodeValue":
			out.NodeValue = string(in.String())
		case "childNodeCount":
			out.ChildNodeCount = int64(in.Int64())
		case "children":
			if in.IsNull() {
				in.Skip()
				out.Children = nil
			} else {
				in.Delim('[')
				if out.Children == nil {
					if !in.IsDelim(']') {
						out.Children = make([]*Node, 0, 8)
					} else {
						out.Children = []*Node{}
					}
				} else {
					out.Children = (out.Children)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *Node
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(Node)
						}
						(*v1).UnmarshalEasyJSON(in)
					}
					out.Children = append(out.Children, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "attributes":
			if in.IsNull() {
				in.Skip()
				out.Attributes = nil
			} else {
				in.Delim('[')
				if out.Attributes == nil {
					if !in.IsDelim(']') {
						out.Attributes = make([]string, 0, 4)
					} else {
						out.Attributes = []string{}
					}
				} else {
					out.Attributes = (out.Attributes)[:0]
				}
				for !in.IsDelim(']') {
					var v2 string
					v2 = string(in.String())
					out.Attributes = append(out.Attributes, v2)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "documentURL":
			out.DocumentURL = string(in.String())
		case "baseURL":
			out.BaseURL = string(in.String())
		case "publicId":
			out.PublicID = string(in.String())
		case "systemId":
			out.SystemID = string(in.String())
		case "internalSubset":
			out.InternalSubset = string(in.String())
		case "xmlVersion":
			out.XMLVersion = string(in.String())
		case "name":
			out.Name = string(in.String())
		case "value":
			out.Value = string(in.String())
		case "pseudoType":
			(out.PseudoType).UnmarshalEasyJSON(in)
		case "shadowRootType":
			(out.ShadowRootType).UnmarshalEasyJSON(in)
		case "frameId":
			(out.FrameID).UnmarshalEasyJSON(in)
		case "contentDocument":
			if in.IsNull() {
				in.Skip()
				out.ContentDocument = nil
			} else {
				if out.ContentDocument == nil {
					out.ContentDocument = new(Node)
				}
				(*out.ContentDocument).UnmarshalEasyJSON(in)
			}
		case "shadowRoots":
			if in.IsNull() {
				in.Skip()
				out.ShadowRoots = nil
			} else {
				in.Delim('[')
				if out.ShadowRoots == nil {
					if !in.IsDelim(']') {
						out.ShadowRoots = make([]*Node, 0, 8)
					} else {
						out.ShadowRoots = []*Node{}
					}
				} else {
					out.ShadowRoots = (out.ShadowRoots)[:0]
				}
				for !in.IsDelim(']') {
					var v3 *Node
					if in.IsNull() {
						in.Skip()
						v3 = nil
					} else {
						if v3 == nil {
							v3 = new(Node)
						}
						(*v3).UnmarshalEasyJSON(in)
					}
					out.ShadowRoots = append(out.ShadowRoots, v3)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "templateContent":
			if in.IsNull() {
				in.Skip()
				out.TemplateContent = nil
			} else {
				if out.TemplateContent == nil {
					out.TemplateContent = new(Node)
				}
				(*out.TemplateContent).UnmarshalEasyJSON(in)
			}
		case "pseudoElements":
			if in.IsNull() {
				in.Skip()
				out.PseudoElements = nil
			} else {
				in.Delim('[')
				if out.PseudoElements == nil {
					if !in.IsDelim(']') {
						out.PseudoElements = make([]*Node, 0, 8)
					} else {
						out.PseudoElements = []*Node{}
					}
				} else {
					out.PseudoElements = (out.PseudoElements)[:0]
				}
				for !in.IsDelim(']') {
					var v4 *Node
					if in.IsNull() {
						in.Skip()
						v4 = nil
					} else {
						if v4 == nil {
							v4 = new(Node)
						}
						(*v4).UnmarshalEasyJSON(in)
					}
					out.PseudoElements = append(out.PseudoElements, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "importedDocument":
			if in.IsNull() {
				in.Skip()
				out.ImportedDocument = nil
			} else {
				if out.ImportedDocument == nil {
					out.ImportedDocument = new(Node)
				}
				(*out.ImportedDocument).UnmarshalEasyJSON(in)
			}
		case "distributedNodes":
			if in.IsNull() {
				in.Skip()
				out.DistributedNodes = nil
			} else {
				in.Delim('[')
				if out.DistributedNodes == nil {
					if !in.IsDelim(']') {
						out.DistributedNodes = make([]*BackendNode, 0, 8)
					} else {
						out.DistributedNodes = []*BackendNode{}
					}
				} else {
					out.DistributedNodes = (out.DistributedNodes)[:0]
				}
				for !in.IsDelim(']') {
					var v5 *BackendNode
					if in.IsNull() {
						in.Skip()
						v5 = nil
					} else {
						if v5 == nil {
							v5 = new(BackendNode)
						}
						(*v5).UnmarshalEasyJSON(in)
					}
					out.DistributedNodes = append(out.DistributedNodes, v5)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "isSVG":
			out.IsSVG = bool(in.Bool())
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
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdp1(out *jwriter.Writer, in Node) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"nodeId\":")
	out.Int64(int64(in.NodeID))
	if in.ParentID != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"parentId\":")
		out.Int64(int64(in.ParentID))
	}
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"backendNodeId\":")
	out.Int64(int64(in.BackendNodeID))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"nodeType\":")
	(in.NodeType).MarshalEasyJSON(out)
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"nodeName\":")
	out.String(string(in.NodeName))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"localName\":")
	out.String(string(in.LocalName))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"nodeValue\":")
	out.String(string(in.NodeValue))
	if in.ChildNodeCount != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"childNodeCount\":")
		out.Int64(int64(in.ChildNodeCount))
	}
	if len(in.Children) != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"children\":")
		if in.Children == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v6, v7 := range in.Children {
				if v6 > 0 {
					out.RawByte(',')
				}
				if v7 == nil {
					out.RawString("null")
				} else {
					(*v7).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	if len(in.Attributes) != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"attributes\":")
		if in.Attributes == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v8, v9 := range in.Attributes {
				if v8 > 0 {
					out.RawByte(',')
				}
				out.String(string(v9))
			}
			out.RawByte(']')
		}
	}
	if in.DocumentURL != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"documentURL\":")
		out.String(string(in.DocumentURL))
	}
	if in.BaseURL != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"baseURL\":")
		out.String(string(in.BaseURL))
	}
	if in.PublicID != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"publicId\":")
		out.String(string(in.PublicID))
	}
	if in.SystemID != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"systemId\":")
		out.String(string(in.SystemID))
	}
	if in.InternalSubset != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"internalSubset\":")
		out.String(string(in.InternalSubset))
	}
	if in.XMLVersion != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"xmlVersion\":")
		out.String(string(in.XMLVersion))
	}
	if in.Name != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"name\":")
		out.String(string(in.Name))
	}
	if in.Value != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"value\":")
		out.String(string(in.Value))
	}
	if in.PseudoType != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"pseudoType\":")
		(in.PseudoType).MarshalEasyJSON(out)
	}
	if in.ShadowRootType != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"shadowRootType\":")
		(in.ShadowRootType).MarshalEasyJSON(out)
	}
	if in.FrameID != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"frameId\":")
		out.String(string(in.FrameID))
	}
	if in.ContentDocument != nil {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"contentDocument\":")
		if in.ContentDocument == nil {
			out.RawString("null")
		} else {
			(*in.ContentDocument).MarshalEasyJSON(out)
		}
	}
	if len(in.ShadowRoots) != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"shadowRoots\":")
		if in.ShadowRoots == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v10, v11 := range in.ShadowRoots {
				if v10 > 0 {
					out.RawByte(',')
				}
				if v11 == nil {
					out.RawString("null")
				} else {
					(*v11).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	if in.TemplateContent != nil {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"templateContent\":")
		if in.TemplateContent == nil {
			out.RawString("null")
		} else {
			(*in.TemplateContent).MarshalEasyJSON(out)
		}
	}
	if len(in.PseudoElements) != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"pseudoElements\":")
		if in.PseudoElements == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v12, v13 := range in.PseudoElements {
				if v12 > 0 {
					out.RawByte(',')
				}
				if v13 == nil {
					out.RawString("null")
				} else {
					(*v13).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	if in.ImportedDocument != nil {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"importedDocument\":")
		if in.ImportedDocument == nil {
			out.RawString("null")
		} else {
			(*in.ImportedDocument).MarshalEasyJSON(out)
		}
	}
	if len(in.DistributedNodes) != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"distributedNodes\":")
		if in.DistributedNodes == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v14, v15 := range in.DistributedNodes {
				if v14 > 0 {
					out.RawByte(',')
				}
				if v15 == nil {
					out.RawString("null")
				} else {
					(*v15).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	if in.IsSVG {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"isSVG\":")
		out.Bool(bool(in.IsSVG))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Node) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdp1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Node) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdp1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Node) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdp1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Node) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdp1(l, v)
}
func easyjsonC5a4559bDecodeGithubComKnqChromedpCdp2(in *jlexer.Lexer, out *MessageError) {
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
		case "code":
			out.Code = int64(in.Int64())
		case "message":
			out.Message = string(in.String())
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
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdp2(out *jwriter.Writer, in MessageError) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"code\":")
	out.Int64(int64(in.Code))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"message\":")
	out.String(string(in.Message))
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v MessageError) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdp2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v MessageError) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdp2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *MessageError) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdp2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *MessageError) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdp2(l, v)
}
func easyjsonC5a4559bDecodeGithubComKnqChromedpCdp3(in *jlexer.Lexer, out *Message) {
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
		case "id":
			out.ID = int64(in.Int64())
		case "method":
			(out.Method).UnmarshalEasyJSON(in)
		case "params":
			(out.Params).UnmarshalEasyJSON(in)
		case "result":
			(out.Result).UnmarshalEasyJSON(in)
		case "error":
			if in.IsNull() {
				in.Skip()
				out.Error = nil
			} else {
				if out.Error == nil {
					out.Error = new(MessageError)
				}
				(*out.Error).UnmarshalEasyJSON(in)
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
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdp3(out *jwriter.Writer, in Message) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ID != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"id\":")
		out.Int64(int64(in.ID))
	}
	if in.Method != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"method\":")
		(in.Method).MarshalEasyJSON(out)
	}
	if (in.Params).IsDefined() {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"params\":")
		(in.Params).MarshalEasyJSON(out)
	}
	if (in.Result).IsDefined() {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"result\":")
		(in.Result).MarshalEasyJSON(out)
	}
	if in.Error != nil {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"error\":")
		if in.Error == nil {
			out.RawString("null")
		} else {
			(*in.Error).MarshalEasyJSON(out)
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Message) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdp3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Message) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdp3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Message) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdp3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Message) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdp3(l, v)
}
func easyjsonC5a4559bDecodeGithubComKnqChromedpCdp4(in *jlexer.Lexer, out *Frame) {
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
		case "id":
			(out.ID).UnmarshalEasyJSON(in)
		case "parentId":
			(out.ParentID).UnmarshalEasyJSON(in)
		case "loaderId":
			out.LoaderID = LoaderID(in.String())
		case "name":
			out.Name = string(in.String())
		case "url":
			out.URL = string(in.String())
		case "securityOrigin":
			out.SecurityOrigin = string(in.String())
		case "mimeType":
			out.MimeType = string(in.String())
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
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdp4(out *jwriter.Writer, in Frame) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"id\":")
	out.String(string(in.ID))
	if in.ParentID != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"parentId\":")
		out.String(string(in.ParentID))
	}
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"loaderId\":")
	out.String(string(in.LoaderID))
	if in.Name != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"name\":")
		out.String(string(in.Name))
	}
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"url\":")
	out.String(string(in.URL))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"securityOrigin\":")
	out.String(string(in.SecurityOrigin))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"mimeType\":")
	out.String(string(in.MimeType))
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Frame) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdp4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Frame) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdp4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Frame) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdp4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Frame) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdp4(l, v)
}
func easyjsonC5a4559bDecodeGithubComKnqChromedpCdp5(in *jlexer.Lexer, out *BackendNode) {
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
		case "nodeType":
			(out.NodeType).UnmarshalEasyJSON(in)
		case "nodeName":
			out.NodeName = string(in.String())
		case "backendNodeId":
			(out.BackendNodeID).UnmarshalEasyJSON(in)
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
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdp5(out *jwriter.Writer, in BackendNode) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"nodeType\":")
	(in.NodeType).MarshalEasyJSON(out)
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"nodeName\":")
	out.String(string(in.NodeName))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"backendNodeId\":")
	out.Int64(int64(in.BackendNodeID))
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v BackendNode) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdp5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v BackendNode) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdp5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *BackendNode) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdp5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *BackendNode) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdp5(l, v)
}
