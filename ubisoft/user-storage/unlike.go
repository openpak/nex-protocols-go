package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	user_storage_types "github.com/PretendoNetwork/nex-protocols-go/v2/ubisoft/user-storage/types"
)

// UnlikeRequest holds the request parameters of User Storage::Unlike
type UnlikeRequest struct {
	ContentKey user_storage_types.UserContentKey
}

// WriteTo writes the UnlikeRequest to the given writable
func (m UnlikeRequest) WriteTo(writable types.Writable) {
	m.ContentKey.WriteTo(writable)
}

// ExtractFrom extracts the UnlikeRequest from the given readable
func (m *UnlikeRequest) ExtractFrom(readable types.Readable) error {
	if err := m.ContentKey.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read Unlike.contentKey. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handleUnlike(packet nex.PacketInterface) {
	if protocol.Unlike == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "User Storage::Unlike not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var params UnlikeRequest

	if err := params.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.Unlike(err, packet, callID, params.ContentKey)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.Unlike(nil, packet, callID, params.ContentKey)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
