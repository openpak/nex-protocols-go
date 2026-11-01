package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// GameSessionSearchResult is a type within the Game Session protocol
type GameSessionSearchResult struct {
	types.Structure
	SessionKey GameSessionKey
	HostPID    types.UInt32
	HostURLs   types.List[types.StationURL]
	Attributes types.List[Property]
}

// WriteTo writes the GameSessionSearchResult to the given writable
func (gssr GameSessionSearchResult) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	gssr.SessionKey.WriteTo(contentWritable)
	gssr.HostPID.WriteTo(contentWritable)
	gssr.HostURLs.WriteTo(contentWritable)
	gssr.Attributes.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	gssr.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the GameSessionSearchResult from the given readable
func (gssr *GameSessionSearchResult) ExtractFrom(readable types.Readable) error {
	if err := gssr.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("failed to extract GameSessionSearchResult header. %s", err.Error())
	}

	if err := gssr.SessionKey.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract GameSessionSearchResult.SessionKey. %s", err.Error())
	}

	if err := gssr.HostPID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract GameSessionSearchResult.HostPID. %s", err.Error())
	}

	if err := gssr.HostURLs.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract GameSessionSearchResult.HostURLs. %s", err.Error())
	}

	if err := gssr.Attributes.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract GameSessionSearchResult.Attributes. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of GameSessionSearchResult
func (gssr GameSessionSearchResult) Copy() types.RVType {
	copied := NewGameSessionSearchResult()

	copied.StructureVersion = gssr.StructureVersion
	copied.SessionKey = gssr.SessionKey.Copy().(GameSessionKey)
	copied.HostPID = gssr.HostPID.Copy().(types.UInt32)
	copied.HostURLs = gssr.HostURLs.Copy().(types.List[types.StationURL])
	copied.Attributes = gssr.Attributes.Copy().(types.List[Property])

	return copied
}

// Equals checks if the given GameSessionSearchResult contains the same data as the current GameSessionSearchResult
func (gssr GameSessionSearchResult) Equals(o types.RVType) bool {
	other, ok := o.(GameSessionSearchResult)
	if !ok {
		return false
	}

	if gssr.StructureVersion != other.StructureVersion {
		return false
	}

	if !(gssr.SessionKey.Equals(other.SessionKey)) {
		return false
	}

	if !(gssr.HostPID.Equals(other.HostPID)) {
		return false
	}

	if !(gssr.HostURLs.Equals(other.HostURLs)) {
		return false
	}

	if !(gssr.Attributes.Equals(other.Attributes)) {
		return false
	}

	return true
}

// CopyRef copies the current value of the GameSessionSearchResult
// and returns a pointer to the new copy
func (gssr GameSessionSearchResult) CopyRef() types.RVTypePtr {
	copied := gssr.Copy().(GameSessionSearchResult)
	return &copied
}

// Deref takes a pointer to the GameSessionSearchResult
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (gssr *GameSessionSearchResult) Deref() types.RVType {
	return *gssr
}

// String returns the string representation of the GameSessionSearchResult
func (gssr GameSessionSearchResult) String() string {
	return gssr.FormatToString(0)
}

// FormatToString pretty-prints the GameSessionSearchResult using the provided indentation level
func (gssr GameSessionSearchResult) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("GameSessionSearchResult{\n")
	fmt.Fprintf(&b, "%sSessionKey: %s,\n", indentationValues, gssr.SessionKey.FormatToString(indentationLevel+1))
	fmt.Fprintf(&b, "%sHostPID: %s,\n", indentationValues, gssr.HostPID)
	fmt.Fprintf(&b, "%sHostURLs: %s,\n", indentationValues, gssr.HostURLs)
	fmt.Fprintf(&b, "%sAttributes: %s,\n", indentationValues, gssr.Attributes)
	fmt.Fprintf(&b, "%s}", indentationEnd)

	return b.String()
}

// NewGameSessionSearchResult returns a new GameSessionSearchResult
func NewGameSessionSearchResult() GameSessionSearchResult {
	return GameSessionSearchResult{
		SessionKey: NewGameSessionKey(),
		HostPID:    types.NewUInt32(0),
		HostURLs:   types.NewList[types.StationURL](),
		Attributes: types.NewList[Property](),
	}
}
