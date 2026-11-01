package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	offline_game_notifications_types "github.com/PretendoNetwork/nex-protocols-go/v2/ubisoft/offline-game-notifications/types"
)

// PollAnyOfflineNotificationsRequest holds the request parameters of Offline Game Notifications::PollAnyOfflineNotifications
type PollAnyOfflineNotificationsRequest struct{}

// WriteTo writes the PollAnyOfflineNotificationsRequest to the given writable
func (m PollAnyOfflineNotificationsRequest) WriteTo(writable types.Writable) {
	_ = writable
}

// ExtractFrom extracts the PollAnyOfflineNotificationsRequest from the given readable
func (m *PollAnyOfflineNotificationsRequest) ExtractFrom(readable types.Readable) error {
	_ = readable
	return nil
}

// PollAnyOfflineNotificationsResponse holds the response parameters of Offline Game Notifications::PollAnyOfflineNotifications
type PollAnyOfflineNotificationsResponse struct {
	ListTimedNotification types.List[offline_game_notifications_types.TimedNotification]
	NbRemainingNotifs     types.UInt32
}

// WriteTo writes the PollAnyOfflineNotificationsResponse to the given writable
func (m PollAnyOfflineNotificationsResponse) WriteTo(writable types.Writable) {
	m.ListTimedNotification.WriteTo(writable)
	m.NbRemainingNotifs.WriteTo(writable)
}

// ExtractFrom extracts the PollAnyOfflineNotificationsResponse from the given readable
func (m *PollAnyOfflineNotificationsResponse) ExtractFrom(readable types.Readable) error {
	if err := m.ListTimedNotification.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read PollAnyOfflineNotifications.listTimedNotification. %s", err.Error())
	}

	if err := m.NbRemainingNotifs.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read PollAnyOfflineNotifications.nbRemainingNotifs. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handlePollAnyOfflineNotifications(packet nex.PacketInterface) {
	if protocol.PollAnyOfflineNotifications == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Offline Game Notifications::PollAnyOfflineNotifications not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID

	rmcMessage, rmcError := protocol.PollAnyOfflineNotifications(nil, packet, callID)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
