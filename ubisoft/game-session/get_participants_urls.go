package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	game_session_types "github.com/PretendoNetwork/nex-protocols-go/v2/ubisoft/game-session/types"
)

// GetParticipantsURLsRequest holds the request parameters of Game Session::GetParticipantsURLs
type GetParticipantsURLsRequest struct {
	GameSessionKey game_session_types.GameSessionKey
	ParticipantIDs types.List[types.UInt32]
}

// WriteTo writes the GetParticipantsURLsRequest to the given writable
func (m GetParticipantsURLsRequest) WriteTo(writable types.Writable) {
	m.GameSessionKey.WriteTo(writable)
	m.ParticipantIDs.WriteTo(writable)
}

// ExtractFrom extracts the GetParticipantsURLsRequest from the given readable
func (m *GetParticipantsURLsRequest) ExtractFrom(readable types.Readable) error {
	if err := m.GameSessionKey.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read GetParticipantsURLs.gameSessionKey. %s", err.Error())
	}

	if err := m.ParticipantIDs.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read GetParticipantsURLs.participantIDs. %s", err.Error())
	}

	return nil
}

// GetParticipantsURLsResponse holds the response parameters of Game Session::GetParticipantsURLs
type GetParticipantsURLsResponse struct {
	Participants types.List[game_session_types.GameSessionParticipant]
}

// WriteTo writes the GetParticipantsURLsResponse to the given writable
func (m GetParticipantsURLsResponse) WriteTo(writable types.Writable) {
	m.Participants.WriteTo(writable)
}

// ExtractFrom extracts the GetParticipantsURLsResponse from the given readable
func (m *GetParticipantsURLsResponse) ExtractFrom(readable types.Readable) error {
	if err := m.Participants.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read GetParticipantsURLs.participants. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handleGetParticipantsURLs(packet nex.PacketInterface) {
	if protocol.GetParticipantsURLs == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Game Session::GetParticipantsURLs not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var params GetParticipantsURLsRequest

	if err := params.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.GetParticipantsURLs(err, packet, callID, params.GameSessionKey, params.ParticipantIDs)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetParticipantsURLs(nil, packet, callID, params.GameSessionKey, params.ParticipantIDs)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
