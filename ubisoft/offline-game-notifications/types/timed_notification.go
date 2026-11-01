// Package types implements the structures used by the Offline Game Notifications protocol

package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
	notifications_types "github.com/PretendoNetwork/nex-protocols-go/v2/notifications/types"
)

// TimedNotification is a type within the Offline Game Notifications protocol
type TimedNotification struct {
	types.Structure
	Timestamp    types.DateTime
	Notification notifications_types.NotificationEvent
}

// WriteTo writes the TimedNotification to the given writable
func (tn TimedNotification) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	tn.Timestamp.WriteTo(contentWritable)
	tn.Notification.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	tn.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the TimedNotification from the given readable
func (tn *TimedNotification) ExtractFrom(readable types.Readable) error {
	if err := tn.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("failed to extract TimedNotification header. %s", err.Error())
	}

	if err := tn.Timestamp.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract TimedNotification.Timestamp. %s", err.Error())
	}

	if err := tn.Notification.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract TimedNotification.Notification. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of TimedNotification
func (tn TimedNotification) Copy() types.RVType {
	copied := NewTimedNotification()

	copied.StructureVersion = tn.StructureVersion
	copied.Timestamp = tn.Timestamp.Copy().(types.DateTime)
	copied.Notification = tn.Notification.Copy().(notifications_types.NotificationEvent)

	return copied
}

// Equals checks if the given TimedNotification contains the same data as the current TimedNotification
func (tn TimedNotification) Equals(o types.RVType) bool {
	other, ok := o.(TimedNotification)
	if !ok {
		return false
	}

	if tn.StructureVersion != other.StructureVersion {
		return false
	}

	if !(tn.Timestamp.Equals(other.Timestamp)) {
		return false
	}

	if !(tn.Notification.Equals(other.Notification)) {
		return false
	}

	return true
}

// CopyRef copies the current value of the TimedNotification
// and returns a pointer to the new copy
func (tn TimedNotification) CopyRef() types.RVTypePtr {
	copied := tn.Copy().(TimedNotification)
	return &copied
}

// Deref takes a pointer to the TimedNotification
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (tn *TimedNotification) Deref() types.RVType {
	return *tn
}

// String returns the string representation of the TimedNotification
func (tn TimedNotification) String() string {
	return tn.FormatToString(0)
}

// FormatToString pretty-prints the TimedNotification using the provided indentation level
func (tn TimedNotification) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("TimedNotification{\n")
	fmt.Fprintf(&b, "%sTimestamp: %s,\n", indentationValues, tn.Timestamp.FormatToString(indentationLevel+1))
	fmt.Fprintf(&b, "%sNotification: %s,\n", indentationValues, tn.Notification.FormatToString(indentationLevel+1))
	fmt.Fprintf(&b, "%s}", indentationEnd)

	return b.String()
}

// NewTimedNotification returns a new TimedNotification
func NewTimedNotification() TimedNotification {
	return TimedNotification{
		Timestamp:    types.NewDateTime(0),
		Notification: notifications_types.NewNotificationEvent(),
	}
}
