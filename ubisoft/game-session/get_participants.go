package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	game_session_types "github.com/PretendoNetwork/nex-protocols-go/v2/ubisoft/game-session/types"
)

// GetParticipantsRequest holds the request parameters of Game Session::GetParticipants
type GetParticipantsRequest struct {
	GameSessionKey game_session_types.GameSessionKey
	ResultRange    types.ResultRange
}

// WriteTo writes the GetParticipantsRequest to the given writable
func (m GetParticipantsRequest) WriteTo(writable types.Writable) {
	m.GameSessionKey.WriteTo(writable)
	m.ResultRange.WriteTo(writable)
}

// ExtractFrom extracts the GetParticipantsRequest from the given readable
func (m *GetParticipantsRequest) ExtractFrom(readable types.Readable) error {
	if err := m.GameSessionKey.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read GetParticipants.gameSessionKey. %s", err.Error())
	}

	if err := m.ResultRange.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read GetParticipants.resultRange. %s", err.Error())
	}

	return nil
}

// GetParticipantsResponse holds the response parameters of Game Session::GetParticipants
type GetParticipantsResponse struct {
	Participants types.List[game_session_types.GameSessionParticipant]
}

// WriteTo writes the GetParticipantsResponse to the given writable
func (m GetParticipantsResponse) WriteTo(writable types.Writable) {
	m.Participants.WriteTo(writable)
}

// ExtractFrom extracts the GetParticipantsResponse from the given readable
func (m *GetParticipantsResponse) ExtractFrom(readable types.Readable) error {
	if err := m.Participants.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read GetParticipants.participants. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handleGetParticipants(packet nex.PacketInterface) {
	if protocol.GetParticipants == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Game Session::GetParticipants not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var params GetParticipantsRequest

	if err := params.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.GetParticipants(err, packet, callID, params.GameSessionKey, params.ResultRange)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetParticipants(nil, packet, callID, params.GameSessionKey, params.ResultRange)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
