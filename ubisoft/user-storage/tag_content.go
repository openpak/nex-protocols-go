package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	user_storage_types "github.com/PretendoNetwork/nex-protocols-go/v2/ubisoft/user-storage/types"
)

// TagContentRequest holds the request parameters of User Storage::TagContent
type TagContentRequest struct {
	ContentKey user_storage_types.UserContentKey
	NewTagIds  types.List[types.UInt32]
}

// WriteTo writes the TagContentRequest to the given writable
func (m TagContentRequest) WriteTo(writable types.Writable) {
	m.ContentKey.WriteTo(writable)
	m.NewTagIds.WriteTo(writable)
}

// ExtractFrom extracts the TagContentRequest from the given readable
func (m *TagContentRequest) ExtractFrom(readable types.Readable) error {
	if err := m.ContentKey.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read TagContent.contentKey. %s", err.Error())
	}

	if err := m.NewTagIds.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read TagContent.newTagIds. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handleTagContent(packet nex.PacketInterface) {
	if protocol.TagContent == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "User Storage::TagContent not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var params TagContentRequest

	if err := params.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.TagContent(err, packet, callID, params.ContentKey, params.NewTagIds)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.TagContent(nil, packet, callID, params.ContentKey, params.NewTagIds)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
