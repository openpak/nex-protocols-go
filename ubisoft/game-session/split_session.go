package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	game_session_types "github.com/PretendoNetwork/nex-protocols-go/v2/ubisoft/game-session/types"
)

// SplitSessionRequest holds the request parameters of Game Session::SplitSession
type SplitSessionRequest struct {
	GameSessionKey game_session_types.GameSessionKey
}

// WriteTo writes the SplitSessionRequest to the given writable
func (m SplitSessionRequest) WriteTo(writable types.Writable) {
	m.GameSessionKey.WriteTo(writable)
}

// ExtractFrom extracts the SplitSessionRequest from the given readable
func (m *SplitSessionRequest) ExtractFrom(readable types.Readable) error {
	if err := m.GameSessionKey.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read SplitSession.gameSessionKey. %s", err.Error())
	}

	return nil
}

// SplitSessionResponse holds the response parameters of Game Session::SplitSession
type SplitSessionResponse struct {
	GameSessionKeyMigrated game_session_types.GameSessionKey
}

// WriteTo writes the SplitSessionResponse to the given writable
func (m SplitSessionResponse) WriteTo(writable types.Writable) {
	m.GameSessionKeyMigrated.WriteTo(writable)
}

// ExtractFrom extracts the SplitSessionResponse from the given readable
func (m *SplitSessionResponse) ExtractFrom(readable types.Readable) error {
	if err := m.GameSessionKeyMigrated.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read SplitSession.gameSessionKeyMigrated. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handleSplitSession(packet nex.PacketInterface) {
	if protocol.SplitSession == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Game Session::SplitSession not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var params SplitSessionRequest

	if err := params.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.SplitSession(err, packet, callID, params.GameSessionKey)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.SplitSession(nil, packet, callID, params.GameSessionKey)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
