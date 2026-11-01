package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	game_session_types "github.com/PretendoNetwork/nex-protocols-go/v2/ubisoft/game-session/types"
)

// CancelInvitationRequest holds the request parameters of Game Session::CancelInvitation
type CancelInvitationRequest struct {
	GameSessionInvitation game_session_types.GameSessionInvitationSent
}

// WriteTo writes the CancelInvitationRequest to the given writable
func (m CancelInvitationRequest) WriteTo(writable types.Writable) {
	m.GameSessionInvitation.WriteTo(writable)
}

// ExtractFrom extracts the CancelInvitationRequest from the given readable
func (m *CancelInvitationRequest) ExtractFrom(readable types.Readable) error {
	if err := m.GameSessionInvitation.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read CancelInvitation.gameSessionInvitation. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handleCancelInvitation(packet nex.PacketInterface) {
	if protocol.CancelInvitation == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Game Session::CancelInvitation not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var params CancelInvitationRequest

	if err := params.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.CancelInvitation(err, packet, callID, params.GameSessionInvitation)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.CancelInvitation(nil, packet, callID, params.GameSessionInvitation)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
