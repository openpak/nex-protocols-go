package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

// UnbanUserRequest holds the request parameters of User Storage Admin::UnbanUser
type UnbanUserRequest struct {
	Pid types.UInt32
}

// WriteTo writes the UnbanUserRequest to the given writable
func (m UnbanUserRequest) WriteTo(writable types.Writable) {
	m.Pid.WriteTo(writable)
}

// ExtractFrom extracts the UnbanUserRequest from the given readable
func (m *UnbanUserRequest) ExtractFrom(readable types.Readable) error {
	if err := m.Pid.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read UnbanUser.pid. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handleUnbanUser(packet nex.PacketInterface) {
	if protocol.UnbanUser == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "User Storage Admin::UnbanUser not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var params UnbanUserRequest

	if err := params.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.UnbanUser(err, packet, callID, params.Pid)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.UnbanUser(nil, packet, callID, params.Pid)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
