package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	game_session_types "github.com/PretendoNetwork/nex-protocols-go/v2/ubisoft/game-session/types"
)

// GetParticipantCountRequest holds the request parameters of Game Session::GetParticipantCount
type GetParticipantCountRequest struct {
	GameSessionKey game_session_types.GameSessionKey
}

// WriteTo writes the GetParticipantCountRequest to the given writable
func (m GetParticipantCountRequest) WriteTo(writable types.Writable) {
	m.GameSessionKey.WriteTo(writable)
}

// ExtractFrom extracts the GetParticipantCountRequest from the given readable
func (m *GetParticipantCountRequest) ExtractFrom(readable types.Readable) error {
	if err := m.GameSessionKey.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read GetParticipantCount.gameSessionKey. %s", err.Error())
	}

	return nil
}

// GetParticipantCountResponse holds the response parameters of Game Session::GetParticipantCount
type GetParticipantCountResponse struct {
	Count types.UInt32
}

// WriteTo writes the GetParticipantCountResponse to the given writable
func (m GetParticipantCountResponse) WriteTo(writable types.Writable) {
	m.Count.WriteTo(writable)
}

// ExtractFrom extracts the GetParticipantCountResponse from the given readable
func (m *GetParticipantCountResponse) ExtractFrom(readable types.Readable) error {
	if err := m.Count.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read GetParticipantCount.count. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handleGetParticipantCount(packet nex.PacketInterface) {
	if protocol.GetParticipantCount == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Game Session::GetParticipantCount not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var params GetParticipantCountRequest

	if err := params.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.GetParticipantCount(err, packet, callID, params.GameSessionKey)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetParticipantCount(nil, packet, callID, params.GameSessionKey)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
