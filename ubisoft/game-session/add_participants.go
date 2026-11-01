package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	game_session_types "github.com/PretendoNetwork/nex-protocols-go/v2/ubisoft/game-session/types"
)

// AddParticipantsRequest holds the request parameters of Game Session::AddParticipants
type AddParticipantsRequest struct {
	GameSessionKey        game_session_types.GameSessionKey
	PublicParticipantIDs  types.List[types.UInt32]
	PrivateParticipantIDs types.List[types.UInt32]
}

// WriteTo writes the AddParticipantsRequest to the given writable
func (m AddParticipantsRequest) WriteTo(writable types.Writable) {
	m.GameSessionKey.WriteTo(writable)
	m.PublicParticipantIDs.WriteTo(writable)
	m.PrivateParticipantIDs.WriteTo(writable)
}

// ExtractFrom extracts the AddParticipantsRequest from the given readable
func (m *AddParticipantsRequest) ExtractFrom(readable types.Readable) error {
	if err := m.GameSessionKey.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read AddParticipants.gameSessionKey. %s", err.Error())
	}

	if err := m.PublicParticipantIDs.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read AddParticipants.publicParticipantIDs. %s", err.Error())
	}

	if err := m.PrivateParticipantIDs.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read AddParticipants.privateParticipantIDs. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handleAddParticipants(packet nex.PacketInterface) {
	if protocol.AddParticipants == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Game Session::AddParticipants not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var params AddParticipantsRequest

	if err := params.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.AddParticipants(err, packet, callID, params.GameSessionKey, params.PublicParticipantIDs, params.PrivateParticipantIDs)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.AddParticipants(nil, packet, callID, params.GameSessionKey, params.PublicParticipantIDs, params.PrivateParticipantIDs)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
