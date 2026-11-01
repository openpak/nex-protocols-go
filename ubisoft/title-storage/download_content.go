package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

// DownloadContentRequest holds the request parameters of Title Storage::DownloadContent
type DownloadContentRequest struct {
	ContentID types.UInt32
}

// WriteTo writes the DownloadContentRequest to the given writable
func (m DownloadContentRequest) WriteTo(writable types.Writable) {
	m.ContentID.WriteTo(writable)
}

// ExtractFrom extracts the DownloadContentRequest from the given readable
func (m *DownloadContentRequest) ExtractFrom(readable types.Readable) error {
	if err := m.ContentID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read DownloadContent.contentID. %s", err.Error())
	}

	return nil
}

// DownloadContentResponse holds the response parameters of Title Storage::DownloadContent
type DownloadContentResponse struct {
	Proto types.String
	Host  types.String
	Path  types.String
}

// WriteTo writes the DownloadContentResponse to the given writable
func (m DownloadContentResponse) WriteTo(writable types.Writable) {
	m.Proto.WriteTo(writable)
	m.Host.WriteTo(writable)
	m.Path.WriteTo(writable)
}

// ExtractFrom extracts the DownloadContentResponse from the given readable
func (m *DownloadContentResponse) ExtractFrom(readable types.Readable) error {
	if err := m.Proto.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read DownloadContent.proto. %s", err.Error())
	}

	if err := m.Host.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read DownloadContent.host. %s", err.Error())
	}

	if err := m.Path.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read DownloadContent.path. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handleDownloadContent(packet nex.PacketInterface) {
	if protocol.DownloadContent == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Title Storage::DownloadContent not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var params DownloadContentRequest

	if err := params.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.DownloadContent(err, packet, callID, params.ContentID)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.DownloadContent(nil, packet, callID, params.ContentID)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
