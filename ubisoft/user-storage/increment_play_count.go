package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	user_storage_types "github.com/PretendoNetwork/nex-protocols-go/v2/ubisoft/user-storage/types"
)

// IncrementPlayCountRequest holds the request parameters of User Storage::IncrementPlayCount
type IncrementPlayCountRequest struct {
	ContentKey user_storage_types.UserContentKey
}

// WriteTo writes the IncrementPlayCountRequest to the given writable
func (m IncrementPlayCountRequest) WriteTo(writable types.Writable) {
	m.ContentKey.WriteTo(writable)
}

// ExtractFrom extracts the IncrementPlayCountRequest from the given readable
func (m *IncrementPlayCountRequest) ExtractFrom(readable types.Readable) error {
	if err := m.ContentKey.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read IncrementPlayCount.contentKey. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handleIncrementPlayCount(packet nex.PacketInterface) {
	if protocol.IncrementPlayCount == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "User Storage::IncrementPlayCount not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var params IncrementPlayCountRequest

	if err := params.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.IncrementPlayCount(err, packet, callID, params.ContentKey)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.IncrementPlayCount(nil, packet, callID, params.ContentKey)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
