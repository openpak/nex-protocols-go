package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	game_session_types "github.com/PretendoNetwork/nex-protocols-go/v2/ubisoft/game-session/types"
)

// SearchSessionsWithParticipantsRequest holds the request parameters of Game Session::SearchSessionsWithParticipants
type SearchSessionsWithParticipantsRequest struct {
	GameSessionTypeID types.UInt32
	ParticipantIDs    types.List[types.UInt32]
}

// WriteTo writes the SearchSessionsWithParticipantsRequest to the given writable
func (m SearchSessionsWithParticipantsRequest) WriteTo(writable types.Writable) {
	m.GameSessionTypeID.WriteTo(writable)
	m.ParticipantIDs.WriteTo(writable)
}

// ExtractFrom extracts the SearchSessionsWithParticipantsRequest from the given readable
func (m *SearchSessionsWithParticipantsRequest) ExtractFrom(readable types.Readable) error {
	if err := m.GameSessionTypeID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read SearchSessionsWithParticipants.gameSessionTypeID. %s", err.Error())
	}

	if err := m.ParticipantIDs.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read SearchSessionsWithParticipants.participantIDs. %s", err.Error())
	}

	return nil
}

// SearchSessionsWithParticipantsResponse holds the response parameters of Game Session::SearchSessionsWithParticipants
type SearchSessionsWithParticipantsResponse struct {
	SearchResults types.List[game_session_types.GameSessionSearchWithParticipantsResult]
}

// WriteTo writes the SearchSessionsWithParticipantsResponse to the given writable
func (m SearchSessionsWithParticipantsResponse) WriteTo(writable types.Writable) {
	m.SearchResults.WriteTo(writable)
}

// ExtractFrom extracts the SearchSessionsWithParticipantsResponse from the given readable
func (m *SearchSessionsWithParticipantsResponse) ExtractFrom(readable types.Readable) error {
	if err := m.SearchResults.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read SearchSessionsWithParticipants.searchResults. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handleSearchSessionsWithParticipants(packet nex.PacketInterface) {
	if protocol.SearchSessionsWithParticipants == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Game Session::SearchSessionsWithParticipants not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var params SearchSessionsWithParticipantsRequest

	if err := params.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.SearchSessionsWithParticipants(err, packet, callID, params.GameSessionTypeID, params.ParticipantIDs)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.SearchSessionsWithParticipants(nil, packet, callID, params.GameSessionTypeID, params.ParticipantIDs)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
