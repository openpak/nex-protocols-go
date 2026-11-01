package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// GameSessionParticipant is a type within the Game Session protocol
type GameSessionParticipant struct {
	types.Structure
	PID         types.UInt32
	Name        types.String
	StationURLs types.List[types.StationURL]
}

// WriteTo writes the GameSessionParticipant to the given writable
func (gsp GameSessionParticipant) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	gsp.PID.WriteTo(contentWritable)
	gsp.Name.WriteTo(contentWritable)
	gsp.StationURLs.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	gsp.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the GameSessionParticipant from the given readable
func (gsp *GameSessionParticipant) ExtractFrom(readable types.Readable) error {
	if err := gsp.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("failed to extract GameSessionParticipant header. %s", err.Error())
	}

	if err := gsp.PID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract GameSessionParticipant.PID. %s", err.Error())
	}

	if err := gsp.Name.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract GameSessionParticipant.Name. %s", err.Error())
	}

	if err := gsp.StationURLs.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract GameSessionParticipant.StationURLs. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of GameSessionParticipant
func (gsp GameSessionParticipant) Copy() types.RVType {
	copied := NewGameSessionParticipant()

	copied.StructureVersion = gsp.StructureVersion
	copied.PID = gsp.PID.Copy().(types.UInt32)
	copied.Name = gsp.Name.Copy().(types.String)
	copied.StationURLs = gsp.StationURLs.Copy().(types.List[types.StationURL])

	return copied
}

// Equals checks if the given GameSessionParticipant contains the same data as the current GameSessionParticipant
func (gsp GameSessionParticipant) Equals(o types.RVType) bool {
	other, ok := o.(GameSessionParticipant)
	if !ok {
		return false
	}

	if gsp.StructureVersion != other.StructureVersion {
		return false
	}

	if !(gsp.PID.Equals(other.PID)) {
		return false
	}

	if !(gsp.Name.Equals(other.Name)) {
		return false
	}

	if !(gsp.StationURLs.Equals(other.StationURLs)) {
		return false
	}

	return true
}

// CopyRef copies the current value of the GameSessionParticipant
// and returns a pointer to the new copy
func (gsp GameSessionParticipant) CopyRef() types.RVTypePtr {
	copied := gsp.Copy().(GameSessionParticipant)
	return &copied
}

// Deref takes a pointer to the GameSessionParticipant
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (gsp *GameSessionParticipant) Deref() types.RVType {
	return *gsp
}

// String returns the string representation of the GameSessionParticipant
func (gsp GameSessionParticipant) String() string {
	return gsp.FormatToString(0)
}

// FormatToString pretty-prints the GameSessionParticipant using the provided indentation level
func (gsp GameSessionParticipant) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("GameSessionParticipant{\n")
	fmt.Fprintf(&b, "%sPID: %s,\n", indentationValues, gsp.PID)
	fmt.Fprintf(&b, "%sName: %s,\n", indentationValues, gsp.Name)
	fmt.Fprintf(&b, "%sStationURLs: %s,\n", indentationValues, gsp.StationURLs)
	fmt.Fprintf(&b, "%s}", indentationEnd)

	return b.String()
}

// NewGameSessionParticipant returns a new GameSessionParticipant
func NewGameSessionParticipant() GameSessionParticipant {
	return GameSessionParticipant{
		PID:         types.NewUInt32(0),
		Name:        types.NewString(""),
		StationURLs: types.NewList[types.StationURL](),
	}
}
