package input

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/mailru/easyjson"
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

// Code generated by chromedp-gen. DO NOT EDIT.

// TouchPoint [no description].
type TouchPoint struct {
	State         TouchState `json:"state"`                   // State of the touch point.
	X             int64      `json:"x"`                       // X coordinate of the event relative to the main frame's viewport.
	Y             int64      `json:"y"`                       // Y coordinate of the event relative to the main frame's viewport. 0 refers to the top of the viewport and Y increases as it proceeds towards the bottom of the viewport.
	RadiusX       int64      `json:"radiusX,omitempty"`       // X radius of the touch area (default: 1).
	RadiusY       int64      `json:"radiusY,omitempty"`       // Y radius of the touch area (default: 1).
	RotationAngle float64    `json:"rotationAngle,omitempty"` // Rotation angle (default: 0.0).
	Force         float64    `json:"force,omitempty"`         // Force (default: 1.0).
	ID            float64    `json:"id,omitempty"`            // Identifier used to track touch sources between events, must be unique within an event.
}

// GestureType [no description].
type GestureType string

// String returns the GestureType as string value.
func (t GestureType) String() string {
	return string(t)
}

// GestureType values.
const (
	GestureDefault GestureType = "default"
	GestureTouch   GestureType = "touch"
	GestureMouse   GestureType = "mouse"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t GestureType) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t GestureType) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *GestureType) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch GestureType(in.String()) {
	case GestureDefault:
		*t = GestureDefault
	case GestureTouch:
		*t = GestureTouch
	case GestureMouse:
		*t = GestureMouse

	default:
		in.AddError(errors.New("unknown GestureType value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *GestureType) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// TimeSinceEpoch uTC time in seconds, counted from January 1, 1970.
type TimeSinceEpoch time.Time

// Time returns the TimeSinceEpoch as time.Time value.
func (t TimeSinceEpoch) Time() time.Time {
	return time.Time(t)
}

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t TimeSinceEpoch) MarshalEasyJSON(out *jwriter.Writer) {
	v := float64(time.Time(t).UnixNano() / int64(time.Second))

	out.Buffer.EnsureSpace(20)
	out.Buffer.Buf = strconv.AppendFloat(out.Buffer.Buf, v, 'f', -1, 64)
}

// MarshalJSON satisfies json.Marshaler.
func (t TimeSinceEpoch) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *TimeSinceEpoch) UnmarshalEasyJSON(in *jlexer.Lexer) {
	*t = TimeSinceEpoch(time.Unix(0, int64(in.Float64()*float64(time.Second))))
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *TimeSinceEpoch) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// Modifier input key modifier type.
type Modifier int64

// Int64 returns the Modifier as int64 value.
func (t Modifier) Int64() int64 {
	return int64(t)
}

// Modifier values.
const (
	ModifierNone  Modifier = 0
	ModifierAlt   Modifier = 1
	ModifierCtrl  Modifier = 2
	ModifierMeta  Modifier = 4
	ModifierShift Modifier = 8
)

// String returns the Modifier as string value.
func (t Modifier) String() string {
	switch t {
	case ModifierNone:
		return "None"
	case ModifierAlt:
		return "Alt"
	case ModifierCtrl:
		return "Ctrl"
	case ModifierMeta:
		return "Meta"
	case ModifierShift:
		return "Shift"
	}

	return fmt.Sprintf("Modifier(%d)", t)
}

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t Modifier) MarshalEasyJSON(out *jwriter.Writer) {
	out.Int64(int64(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t Modifier) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *Modifier) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch Modifier(in.Int64()) {
	case ModifierNone:
		*t = ModifierNone
	case ModifierAlt:
		*t = ModifierAlt
	case ModifierCtrl:
		*t = ModifierCtrl
	case ModifierMeta:
		*t = ModifierMeta
	case ModifierShift:
		*t = ModifierShift

	default:
		in.AddError(errors.New("unknown Modifier value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *Modifier) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// ModifierCommand is an alias for ModifierMeta.
const ModifierCommand Modifier = ModifierMeta

// TouchState state of the touch point.
type TouchState string

// String returns the TouchState as string value.
func (t TouchState) String() string {
	return string(t)
}

// TouchState values.
const (
	TouchPressed    TouchState = "touchPressed"
	TouchReleased   TouchState = "touchReleased"
	TouchMoved      TouchState = "touchMoved"
	TouchStationary TouchState = "touchStationary"
	TouchCanceled   TouchState = "touchCancelled"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t TouchState) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t TouchState) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *TouchState) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch TouchState(in.String()) {
	case TouchPressed:
		*t = TouchPressed
	case TouchReleased:
		*t = TouchReleased
	case TouchMoved:
		*t = TouchMoved
	case TouchStationary:
		*t = TouchStationary
	case TouchCanceled:
		*t = TouchCanceled

	default:
		in.AddError(errors.New("unknown TouchState value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *TouchState) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// KeyType type of the key event.
type KeyType string

// String returns the KeyType as string value.
func (t KeyType) String() string {
	return string(t)
}

// KeyType values.
const (
	KeyDown    KeyType = "keyDown"
	KeyUp      KeyType = "keyUp"
	KeyRawDown KeyType = "rawKeyDown"
	KeyChar    KeyType = "char"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t KeyType) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t KeyType) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *KeyType) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch KeyType(in.String()) {
	case KeyDown:
		*t = KeyDown
	case KeyUp:
		*t = KeyUp
	case KeyRawDown:
		*t = KeyRawDown
	case KeyChar:
		*t = KeyChar

	default:
		in.AddError(errors.New("unknown KeyType value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *KeyType) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// MouseType type of the mouse event.
type MouseType string

// String returns the MouseType as string value.
func (t MouseType) String() string {
	return string(t)
}

// MouseType values.
const (
	MousePressed  MouseType = "mousePressed"
	MouseReleased MouseType = "mouseReleased"
	MouseMoved    MouseType = "mouseMoved"
	MouseWheel    MouseType = "mouseWheel"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t MouseType) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t MouseType) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *MouseType) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch MouseType(in.String()) {
	case MousePressed:
		*t = MousePressed
	case MouseReleased:
		*t = MouseReleased
	case MouseMoved:
		*t = MouseMoved
	case MouseWheel:
		*t = MouseWheel

	default:
		in.AddError(errors.New("unknown MouseType value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *MouseType) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// ButtonType mouse button (default: "none").
type ButtonType string

// String returns the ButtonType as string value.
func (t ButtonType) String() string {
	return string(t)
}

// ButtonType values.
const (
	ButtonNone   ButtonType = "none"
	ButtonLeft   ButtonType = "left"
	ButtonMiddle ButtonType = "middle"
	ButtonRight  ButtonType = "right"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t ButtonType) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t ButtonType) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *ButtonType) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch ButtonType(in.String()) {
	case ButtonNone:
		*t = ButtonNone
	case ButtonLeft:
		*t = ButtonLeft
	case ButtonMiddle:
		*t = ButtonMiddle
	case ButtonRight:
		*t = ButtonRight

	default:
		in.AddError(errors.New("unknown ButtonType value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *ButtonType) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// TouchType type of the touch event.
type TouchType string

// String returns the TouchType as string value.
func (t TouchType) String() string {
	return string(t)
}

// TouchType values.
const (
	TouchStart TouchType = "touchStart"
	TouchEnd   TouchType = "touchEnd"
	TouchMove  TouchType = "touchMove"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t TouchType) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t TouchType) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *TouchType) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch TouchType(in.String()) {
	case TouchStart:
		*t = TouchStart
	case TouchEnd:
		*t = TouchEnd
	case TouchMove:
		*t = TouchMove

	default:
		in.AddError(errors.New("unknown TouchType value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *TouchType) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}
