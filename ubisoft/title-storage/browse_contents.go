package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	title_storage_types "github.com/PretendoNetwork/nex-protocols-go/v2/ubisoft/title-storage/types"
)

// BrowseContentsRequest holds the request parameters of Title Storage::BrowseContents
type BrowseContentsRequest struct {
	Flags                 types.UInt32
	FlagsMaskType         types.UInt8
	ExtendedFlags         types.UInt32
	ExtendedFlagsMaskType types.UInt8
}

// WriteTo writes the BrowseContentsRequest to the given writable
func (m BrowseContentsRequest) WriteTo(writable types.Writable) {
	m.Flags.WriteTo(writable)
	m.FlagsMaskType.WriteTo(writable)
	m.ExtendedFlags.WriteTo(writable)
	m.ExtendedFlagsMaskType.WriteTo(writable)
}

// ExtractFrom extracts the BrowseContentsRequest from the given readable
func (m *BrowseContentsRequest) ExtractFrom(readable types.Readable) error {
	if err := m.Flags.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read BrowseContents.flags. %s", err.Error())
	}

	if err := m.FlagsMaskType.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read BrowseContents.flagsMaskType. %s", err.Error())
	}

	if err := m.ExtendedFlags.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read BrowseContents.extendedFlags. %s", err.Error())
	}

	if err := m.ExtendedFlagsMaskType.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read BrowseContents.extendedFlagsMaskType. %s", err.Error())
	}

	return nil
}

// BrowseContentsResponse holds the response parameters of Title Storage::BrowseContents
type BrowseContentsResponse struct {
	Contents types.List[title_storage_types.TitleContent]
}

// WriteTo writes the BrowseContentsResponse to the given writable
func (m BrowseContentsResponse) WriteTo(writable types.Writable) {
	m.Contents.WriteTo(writable)
}

// ExtractFrom extracts the BrowseContentsResponse from the given readable
func (m *BrowseContentsResponse) ExtractFrom(readable types.Readable) error {
	if err := m.Contents.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read BrowseContents.contents. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handleBrowseContents(packet nex.PacketInterface) {
	if protocol.BrowseContents == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Title Storage::BrowseContents not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var params BrowseContentsRequest

	if err := params.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.BrowseContents(err, packet, callID, params.Flags, params.FlagsMaskType, params.ExtendedFlags, params.ExtendedFlagsMaskType)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.BrowseContents(nil, packet, callID, params.Flags, params.FlagsMaskType, params.ExtendedFlags, params.ExtendedFlagsMaskType)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
