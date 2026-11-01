package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

// UnbanUserFromContentTypeRequest holds the request parameters of User Storage Admin::UnbanUserFromContentType
type UnbanUserFromContentTypeRequest struct {
	TypeID types.UInt32
	Pid    types.UInt32
}

// WriteTo writes the UnbanUserFromContentTypeRequest to the given writable
func (m UnbanUserFromContentTypeRequest) WriteTo(writable types.Writable) {
	m.TypeID.WriteTo(writable)
	m.Pid.WriteTo(writable)
}

// ExtractFrom extracts the UnbanUserFromContentTypeRequest from the given readable
func (m *UnbanUserFromContentTypeRequest) ExtractFrom(readable types.Readable) error {
	if err := m.TypeID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read UnbanUserFromContentType.typeID. %s", err.Error())
	}

	if err := m.Pid.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read UnbanUserFromContentType.pid. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handleUnbanUserFromContentType(packet nex.PacketInterface) {
	if protocol.UnbanUserFromContentType == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "User Storage Admin::UnbanUserFromContentType not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var params UnbanUserFromContentTypeRequest

	if err := params.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.UnbanUserFromContentType(err, packet, callID, params.TypeID, params.Pid)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.UnbanUserFromContentType(nil, packet, callID, params.TypeID, params.Pid)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
