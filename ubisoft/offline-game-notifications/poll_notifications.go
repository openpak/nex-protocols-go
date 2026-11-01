package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	notifications_types "github.com/PretendoNetwork/nex-protocols-go/v2/notifications/types"
)

// PollNotificationsRequest holds the request parameters of Offline Game Notifications::PollNotifications
type PollNotificationsRequest struct{}

// WriteTo writes the PollNotificationsRequest to the given writable
func (m PollNotificationsRequest) WriteTo(writable types.Writable) {
	_ = writable
}

// ExtractFrom extracts the PollNotificationsRequest from the given readable
func (m *PollNotificationsRequest) ExtractFrom(readable types.Readable) error {
	_ = readable
	return nil
}

// PollNotificationsResponse holds the response parameters of Offline Game Notifications::PollNotifications
type PollNotificationsResponse struct {
	ListNotifications types.List[notifications_types.NotificationEvent]
	NbRemainingNotifs types.UInt32
}

// WriteTo writes the PollNotificationsResponse to the given writable
func (m PollNotificationsResponse) WriteTo(writable types.Writable) {
	m.ListNotifications.WriteTo(writable)
	m.NbRemainingNotifs.WriteTo(writable)
}

// ExtractFrom extracts the PollNotificationsResponse from the given readable
func (m *PollNotificationsResponse) ExtractFrom(readable types.Readable) error {
	if err := m.ListNotifications.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read PollNotifications.listNotifications. %s", err.Error())
	}

	if err := m.NbRemainingNotifs.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read PollNotifications.nbRemainingNotifs. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handlePollNotifications(packet nex.PacketInterface) {
	if protocol.PollNotifications == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Offline Game Notifications::PollNotifications not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID

	rmcMessage, rmcError := protocol.PollNotifications(nil, packet, callID)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
