package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	game_session_types "github.com/PretendoNetwork/nex-protocols-go/v2/ubisoft/game-session/types"
)

// UpdateSessionRequest holds the request parameters of Game Session::UpdateSession
type UpdateSessionRequest struct {
	GameSessionUpdate game_session_types.GameSessionUpdate
}

// WriteTo writes the UpdateSessionRequest to the given writable
func (m UpdateSessionRequest) WriteTo(writable types.Writable) {
	m.GameSessionUpdate.WriteTo(writable)
}

// ExtractFrom extracts the UpdateSessionRequest from the given readable
func (m *UpdateSessionRequest) ExtractFrom(readable types.Readable) error {
	if err := m.GameSessionUpdate.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read UpdateSession.gameSessionUpdate. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handleUpdateSession(packet nex.PacketInterface) {
	if protocol.UpdateSession == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Game Session::UpdateSession not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var params UpdateSessionRequest

	if err := params.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.UpdateSession(err, packet, callID, params.GameSessionUpdate)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.UpdateSession(nil, packet, callID, params.GameSessionUpdate)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
