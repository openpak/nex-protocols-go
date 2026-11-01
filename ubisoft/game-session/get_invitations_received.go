package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	game_session_types "github.com/PretendoNetwork/nex-protocols-go/v2/ubisoft/game-session/types"
)

// GetInvitationsReceivedRequest holds the request parameters of Game Session::GetInvitationsReceived
type GetInvitationsReceivedRequest struct {
	GameSessionTypeID types.UInt32
	ResultRange       types.ResultRange
}

// WriteTo writes the GetInvitationsReceivedRequest to the given writable
func (m GetInvitationsReceivedRequest) WriteTo(writable types.Writable) {
	m.GameSessionTypeID.WriteTo(writable)
	m.ResultRange.WriteTo(writable)
}

// ExtractFrom extracts the GetInvitationsReceivedRequest from the given readable
func (m *GetInvitationsReceivedRequest) ExtractFrom(readable types.Readable) error {
	if err := m.GameSessionTypeID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read GetInvitationsReceived.gameSessionTypeID. %s", err.Error())
	}

	if err := m.ResultRange.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read GetInvitationsReceived.resultRange. %s", err.Error())
	}

	return nil
}

// GetInvitationsReceivedResponse holds the response parameters of Game Session::GetInvitationsReceived
type GetInvitationsReceivedResponse struct {
	Invitations types.List[game_session_types.GameSessionInvitationReceived]
}

// WriteTo writes the GetInvitationsReceivedResponse to the given writable
func (m GetInvitationsReceivedResponse) WriteTo(writable types.Writable) {
	m.Invitations.WriteTo(writable)
}

// ExtractFrom extracts the GetInvitationsReceivedResponse from the given readable
func (m *GetInvitationsReceivedResponse) ExtractFrom(readable types.Readable) error {
	if err := m.Invitations.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read GetInvitationsReceived.invitations. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handleGetInvitationsReceived(packet nex.PacketInterface) {
	if protocol.GetInvitationsReceived == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Game Session::GetInvitationsReceived not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var params GetInvitationsReceivedRequest

	if err := params.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.GetInvitationsReceived(err, packet, callID, params.GameSessionTypeID, params.ResultRange)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetInvitationsReceived(nil, packet, callID, params.GameSessionTypeID, params.ResultRange)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
