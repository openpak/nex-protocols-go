package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	offline_game_notifications_types "github.com/PretendoNetwork/nex-protocols-go/v2/ubisoft/offline-game-notifications/types"
)

// PollSpecificOfflineNotificationsRequest holds the request parameters of Offline Game Notifications::PollSpecificOfflineNotifications
type PollSpecificOfflineNotificationsRequest struct {
	Majortype types.List[types.UInt32]
}

// WriteTo writes the PollSpecificOfflineNotificationsRequest to the given writable
func (m PollSpecificOfflineNotificationsRequest) WriteTo(writable types.Writable) {
	m.Majortype.WriteTo(writable)
}

// ExtractFrom extracts the PollSpecificOfflineNotificationsRequest from the given readable
func (m *PollSpecificOfflineNotificationsRequest) ExtractFrom(readable types.Readable) error {
	if err := m.Majortype.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read PollSpecificOfflineNotifications.majortype. %s", err.Error())
	}

	return nil
}

// PollSpecificOfflineNotificationsResponse holds the response parameters of Offline Game Notifications::PollSpecificOfflineNotifications
type PollSpecificOfflineNotificationsResponse struct {
	ListTimedNotification types.List[offline_game_notifications_types.TimedNotification]
	Ret                   types.UInt32
}

// WriteTo writes the PollSpecificOfflineNotificationsResponse to the given writable
func (m PollSpecificOfflineNotificationsResponse) WriteTo(writable types.Writable) {
	m.ListTimedNotification.WriteTo(writable)
	m.Ret.WriteTo(writable)
}

// ExtractFrom extracts the PollSpecificOfflineNotificationsResponse from the given readable
func (m *PollSpecificOfflineNotificationsResponse) ExtractFrom(readable types.Readable) error {
	if err := m.ListTimedNotification.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read PollSpecificOfflineNotifications.listTimedNotification. %s", err.Error())
	}

	if err := m.Ret.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read PollSpecificOfflineNotifications.ret. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handlePollSpecificOfflineNotifications(packet nex.PacketInterface) {
	if protocol.PollSpecificOfflineNotifications == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Offline Game Notifications::PollSpecificOfflineNotifications not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var params PollSpecificOfflineNotificationsRequest

	if err := params.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.PollSpecificOfflineNotifications(err, packet, callID, params.Majortype)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.PollSpecificOfflineNotifications(nil, packet, callID, params.Majortype)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
