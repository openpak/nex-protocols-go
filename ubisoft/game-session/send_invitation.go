package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	game_session_types "github.com/PretendoNetwork/nex-protocols-go/v2/ubisoft/game-session/types"
)

// SendInvitationRequest holds the request parameters of Game Session::SendInvitation
type SendInvitationRequest struct {
	Invitation game_session_types.GameSessionInvitation
}

// WriteTo writes the SendInvitationRequest to the given writable
func (m SendInvitationRequest) WriteTo(writable types.Writable) {
	m.Invitation.WriteTo(writable)
}

// ExtractFrom extracts the SendInvitationRequest from the given readable
func (m *SendInvitationRequest) ExtractFrom(readable types.Readable) error {
	if err := m.Invitation.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read SendInvitation.invitation. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handleSendInvitation(packet nex.PacketInterface) {
	if protocol.SendInvitation == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Game Session::SendInvitation not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var params SendInvitationRequest

	if err := params.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.SendInvitation(err, packet, callID, params.Invitation)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.SendInvitation(nil, packet, callID, params.Invitation)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
