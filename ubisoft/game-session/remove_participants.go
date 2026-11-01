package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	game_session_types "github.com/PretendoNetwork/nex-protocols-go/v2/ubisoft/game-session/types"
)

// RemoveParticipantsRequest holds the request parameters of Game Session::RemoveParticipants
type RemoveParticipantsRequest struct {
	GameSessionKey game_session_types.GameSessionKey
	ParticipantIDs types.List[types.UInt32]
}

// WriteTo writes the RemoveParticipantsRequest to the given writable
func (m RemoveParticipantsRequest) WriteTo(writable types.Writable) {
	m.GameSessionKey.WriteTo(writable)
	m.ParticipantIDs.WriteTo(writable)
}

// ExtractFrom extracts the RemoveParticipantsRequest from the given readable
func (m *RemoveParticipantsRequest) ExtractFrom(readable types.Readable) error {
	if err := m.GameSessionKey.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read RemoveParticipants.gameSessionKey. %s", err.Error())
	}

	if err := m.ParticipantIDs.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read RemoveParticipants.participantIDs. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handleRemoveParticipants(packet nex.PacketInterface) {
	if protocol.RemoveParticipants == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Game Session::RemoveParticipants not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var params RemoveParticipantsRequest

	if err := params.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.RemoveParticipants(err, packet, callID, params.GameSessionKey, params.ParticipantIDs)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.RemoveParticipants(nil, packet, callID, params.GameSessionKey, params.ParticipantIDs)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
