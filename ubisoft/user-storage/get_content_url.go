package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	user_storage_types "github.com/PretendoNetwork/nex-protocols-go/v2/ubisoft/user-storage/types"
)

// GetContentURLRequest holds the request parameters of User Storage::GetContentURL
type GetContentURLRequest struct {
	ContentKey user_storage_types.UserContentKey
}

// WriteTo writes the GetContentURLRequest to the given writable
func (m GetContentURLRequest) WriteTo(writable types.Writable) {
	m.ContentKey.WriteTo(writable)
}

// ExtractFrom extracts the GetContentURLRequest from the given readable
func (m *GetContentURLRequest) ExtractFrom(readable types.Readable) error {
	if err := m.ContentKey.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read GetContentURL.contentKey. %s", err.Error())
	}

	return nil
}

// GetContentURLResponse holds the response parameters of User Storage::GetContentURL
type GetContentURLResponse struct {
	DownloadInfo user_storage_types.UserContentURL
}

// WriteTo writes the GetContentURLResponse to the given writable
func (m GetContentURLResponse) WriteTo(writable types.Writable) {
	m.DownloadInfo.WriteTo(writable)
}

// ExtractFrom extracts the GetContentURLResponse from the given readable
func (m *GetContentURLResponse) ExtractFrom(readable types.Readable) error {
	if err := m.DownloadInfo.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read GetContentURL.downloadInfo. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handleGetContentURL(packet nex.PacketInterface) {
	if protocol.GetContentURL == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "User Storage::GetContentURL not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var params GetContentURLRequest

	if err := params.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.GetContentURL(err, packet, callID, params.ContentKey)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetContentURL(nil, packet, callID, params.ContentKey)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
