package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	game_session_types "github.com/PretendoNetwork/nex-protocols-go/v2/ubisoft/game-session/types"
)

// DeclineInvitationRequest holds the request parameters of Game Session::DeclineInvitation
type DeclineInvitationRequest struct {
	GameSessionInvitation game_session_types.GameSessionInvitationReceived
}

// WriteTo writes the DeclineInvitationRequest to the given writable
func (m DeclineInvitationRequest) WriteTo(writable types.Writable) {
	m.GameSessionInvitation.WriteTo(writable)
}

// ExtractFrom extracts the DeclineInvitationRequest from the given readable
func (m *DeclineInvitationRequest) ExtractFrom(readable types.Readable) error {
	if err := m.GameSessionInvitation.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read DeclineInvitation.gameSessionInvitation. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handleDeclineInvitation(packet nex.PacketInterface) {
	if protocol.DeclineInvitation == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Game Session::DeclineInvitation not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var params DeclineInvitationRequest

	if err := params.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.DeclineInvitation(err, packet, callID, params.GameSessionInvitation)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.DeclineInvitation(nil, packet, callID, params.GameSessionInvitation)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
