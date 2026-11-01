package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	game_session_types "github.com/PretendoNetwork/nex-protocols-go/v2/ubisoft/game-session/types"
)

// AcceptInvitationRequest holds the request parameters of Game Session::AcceptInvitation
type AcceptInvitationRequest struct {
	GameSessionInvitation game_session_types.GameSessionInvitationReceived
}

// WriteTo writes the AcceptInvitationRequest to the given writable
func (m AcceptInvitationRequest) WriteTo(writable types.Writable) {
	m.GameSessionInvitation.WriteTo(writable)
}

// ExtractFrom extracts the AcceptInvitationRequest from the given readable
func (m *AcceptInvitationRequest) ExtractFrom(readable types.Readable) error {
	if err := m.GameSessionInvitation.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read AcceptInvitation.gameSessionInvitation. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handleAcceptInvitation(packet nex.PacketInterface) {
	if protocol.AcceptInvitation == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Game Session::AcceptInvitation not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var params AcceptInvitationRequest

	if err := params.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.AcceptInvitation(err, packet, callID, params.GameSessionInvitation)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.AcceptInvitation(nil, packet, callID, params.GameSessionInvitation)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
