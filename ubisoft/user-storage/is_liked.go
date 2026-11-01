package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	user_storage_types "github.com/PretendoNetwork/nex-protocols-go/v2/ubisoft/user-storage/types"
)

// IsLikedRequest holds the request parameters of User Storage::IsLiked
type IsLikedRequest struct {
	ContentKey user_storage_types.UserContentKey
}

// WriteTo writes the IsLikedRequest to the given writable
func (m IsLikedRequest) WriteTo(writable types.Writable) {
	m.ContentKey.WriteTo(writable)
}

// ExtractFrom extracts the IsLikedRequest from the given readable
func (m *IsLikedRequest) ExtractFrom(readable types.Readable) error {
	if err := m.ContentKey.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read IsLiked.contentKey. %s", err.Error())
	}

	return nil
}

// IsLikedResponse holds the response parameters of User Storage::IsLiked
type IsLikedResponse struct {
	Liked types.Bool
}

// WriteTo writes the IsLikedResponse to the given writable
func (m IsLikedResponse) WriteTo(writable types.Writable) {
	m.Liked.WriteTo(writable)
}

// ExtractFrom extracts the IsLikedResponse from the given readable
func (m *IsLikedResponse) ExtractFrom(readable types.Readable) error {
	if err := m.Liked.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read IsLiked.liked. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handleIsLiked(packet nex.PacketInterface) {
	if protocol.IsLiked == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "User Storage::IsLiked not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var params IsLikedRequest

	if err := params.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.IsLiked(err, packet, callID, params.ContentKey)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.IsLiked(nil, packet, callID, params.ContentKey)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
