package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

// BrowseContentsByTypeRequest holds the request parameters of Title Storage::BrowseContentsByType
type BrowseContentsByTypeRequest struct {
	ContentType           types.String
	Flags                 types.UInt32
	FlagsMaskType         types.UInt8
	ExtendedFlags         types.UInt32
	ExtendedFlagsMaskType types.UInt8
}

// WriteTo writes the BrowseContentsByTypeRequest to the given writable
func (m BrowseContentsByTypeRequest) WriteTo(writable types.Writable) {
	m.ContentType.WriteTo(writable)
	m.Flags.WriteTo(writable)
	m.FlagsMaskType.WriteTo(writable)
	m.ExtendedFlags.WriteTo(writable)
	m.ExtendedFlagsMaskType.WriteTo(writable)
}

// ExtractFrom extracts the BrowseContentsByTypeRequest from the given readable
func (m *BrowseContentsByTypeRequest) ExtractFrom(readable types.Readable) error {
	if err := m.ContentType.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read BrowseContentsByType.contentType. %s", err.Error())
	}

	if err := m.Flags.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read BrowseContentsByType.flags. %s", err.Error())
	}

	if err := m.FlagsMaskType.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read BrowseContentsByType.flagsMaskType. %s", err.Error())
	}

	if err := m.ExtendedFlags.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read BrowseContentsByType.extendedFlags. %s", err.Error())
	}

	if err := m.ExtendedFlagsMaskType.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read BrowseContentsByType.extendedFlagsMaskType. %s", err.Error())
	}

	return nil
}

// BrowseContentsByTypeResponse holds the response parameters of Title Storage::BrowseContentsByType
type BrowseContentsByTypeResponse struct {
	Contents types.List[types.DataHolder]
}

// WriteTo writes the BrowseContentsByTypeResponse to the given writable
func (m BrowseContentsByTypeResponse) WriteTo(writable types.Writable) {
	m.Contents.WriteTo(writable)
}

// ExtractFrom extracts the BrowseContentsByTypeResponse from the given readable
func (m *BrowseContentsByTypeResponse) ExtractFrom(readable types.Readable) error {
	if err := m.Contents.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read BrowseContentsByType.contents. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handleBrowseContentsByType(packet nex.PacketInterface) {
	if protocol.BrowseContentsByType == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Title Storage::BrowseContentsByType not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var params BrowseContentsByTypeRequest

	if err := params.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.BrowseContentsByType(err, packet, callID, params.ContentType, params.Flags, params.FlagsMaskType, params.ExtendedFlags, params.ExtendedFlagsMaskType)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.BrowseContentsByType(nil, packet, callID, params.ContentType, params.Flags, params.FlagsMaskType, params.ExtendedFlags, params.ExtendedFlagsMaskType)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
