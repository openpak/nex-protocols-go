package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	game_session_types "github.com/PretendoNetwork/nex-protocols-go/v2/ubisoft/game-session/types"
)

// SearchSocialSessionsRequest holds the request parameters of Game Session::SearchSocialSessions
type SearchSocialSessionsRequest struct {
	GameSessionSocialQuery game_session_types.GameSessionSocialQuery
}

// WriteTo writes the SearchSocialSessionsRequest to the given writable
func (m SearchSocialSessionsRequest) WriteTo(writable types.Writable) {
	m.GameSessionSocialQuery.WriteTo(writable)
}

// ExtractFrom extracts the SearchSocialSessionsRequest from the given readable
func (m *SearchSocialSessionsRequest) ExtractFrom(readable types.Readable) error {
	if err := m.GameSessionSocialQuery.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read SearchSocialSessions.gameSessionSocialQuery. %s", err.Error())
	}

	return nil
}

// SearchSocialSessionsResponse holds the response parameters of Game Session::SearchSocialSessions
type SearchSocialSessionsResponse struct {
	SearchResults types.List[game_session_types.GameSessionSearchWithParticipantsResult]
}

// WriteTo writes the SearchSocialSessionsResponse to the given writable
func (m SearchSocialSessionsResponse) WriteTo(writable types.Writable) {
	m.SearchResults.WriteTo(writable)
}

// ExtractFrom extracts the SearchSocialSessionsResponse from the given readable
func (m *SearchSocialSessionsResponse) ExtractFrom(readable types.Readable) error {
	if err := m.SearchResults.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read SearchSocialSessions.searchResults. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handleSearchSocialSessions(packet nex.PacketInterface) {
	if protocol.SearchSocialSessions == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Game Session::SearchSocialSessions not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var params SearchSocialSessionsRequest

	if err := params.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.SearchSocialSessions(err, packet, callID, params.GameSessionSocialQuery)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.SearchSocialSessions(nil, packet, callID, params.GameSessionSocialQuery)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
