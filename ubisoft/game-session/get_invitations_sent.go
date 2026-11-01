package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	game_session_types "github.com/PretendoNetwork/nex-protocols-go/v2/ubisoft/game-session/types"
)

// GetInvitationsSentRequest holds the request parameters of Game Session::GetInvitationsSent
type GetInvitationsSentRequest struct {
	GameSessionTypeID types.UInt32
	ResultRange       types.ResultRange
}

// WriteTo writes the GetInvitationsSentRequest to the given writable
func (m GetInvitationsSentRequest) WriteTo(writable types.Writable) {
	m.GameSessionTypeID.WriteTo(writable)
	m.ResultRange.WriteTo(writable)
}

// ExtractFrom extracts the GetInvitationsSentRequest from the given readable
func (m *GetInvitationsSentRequest) ExtractFrom(readable types.Readable) error {
	if err := m.GameSessionTypeID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read GetInvitationsSent.gameSessionTypeID. %s", err.Error())
	}

	if err := m.ResultRange.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read GetInvitationsSent.resultRange. %s", err.Error())
	}

	return nil
}

// GetInvitationsSentResponse holds the response parameters of Game Session::GetInvitationsSent
type GetInvitationsSentResponse struct {
	Invitations types.List[game_session_types.GameSessionInvitationSent]
}

// WriteTo writes the GetInvitationsSentResponse to the given writable
func (m GetInvitationsSentResponse) WriteTo(writable types.Writable) {
	m.Invitations.WriteTo(writable)
}

// ExtractFrom extracts the GetInvitationsSentResponse from the given readable
func (m *GetInvitationsSentResponse) ExtractFrom(readable types.Readable) error {
	if err := m.Invitations.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read GetInvitationsSent.invitations. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handleGetInvitationsSent(packet nex.PacketInterface) {
	if protocol.GetInvitationsSent == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Game Session::GetInvitationsSent not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var params GetInvitationsSentRequest

	if err := params.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.GetInvitationsSent(err, packet, callID, params.GameSessionTypeID, params.ResultRange)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetInvitationsSent(nil, packet, callID, params.GameSessionTypeID, params.ResultRange)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
