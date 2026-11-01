package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

// UpdateContentMetaDataRequest holds the request parameters of Title Storage Admin::UpdateContentMetaData
type UpdateContentMetaDataRequest struct {
	ContentId     types.UInt32
	Name          types.String
	Description   types.String
	DateAvailable types.DateTime
	DateExpired   types.DateTime
	Flags         types.UInt32
	ExtendedFlags types.UInt32
	Extension     types.DataHolder
}

// WriteTo writes the UpdateContentMetaDataRequest to the given writable
func (m UpdateContentMetaDataRequest) WriteTo(writable types.Writable) {
	m.ContentId.WriteTo(writable)
	m.Name.WriteTo(writable)
	m.Description.WriteTo(writable)
	m.DateAvailable.WriteTo(writable)
	m.DateExpired.WriteTo(writable)
	m.Flags.WriteTo(writable)
	m.ExtendedFlags.WriteTo(writable)
	m.Extension.WriteTo(writable)
}

// ExtractFrom extracts the UpdateContentMetaDataRequest from the given readable
func (m *UpdateContentMetaDataRequest) ExtractFrom(readable types.Readable) error {
	if err := m.ContentId.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read UpdateContentMetaData.contentId. %s", err.Error())
	}

	if err := m.Name.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read UpdateContentMetaData.name. %s", err.Error())
	}

	if err := m.Description.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read UpdateContentMetaData.description. %s", err.Error())
	}

	if err := m.DateAvailable.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read UpdateContentMetaData.dateAvailable. %s", err.Error())
	}

	if err := m.DateExpired.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read UpdateContentMetaData.dateExpired. %s", err.Error())
	}

	if err := m.Flags.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read UpdateContentMetaData.flags. %s", err.Error())
	}

	if err := m.ExtendedFlags.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read UpdateContentMetaData.extendedFlags. %s", err.Error())
	}

	if err := m.Extension.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read UpdateContentMetaData.extension. %s", err.Error())
	}

	return nil
}

// UpdateContentMetaDataResponse holds the response parameters of Title Storage Admin::UpdateContentMetaData
type UpdateContentMetaDataResponse struct {
	Retval types.Bool
}

// WriteTo writes the UpdateContentMetaDataResponse to the given writable
func (m UpdateContentMetaDataResponse) WriteTo(writable types.Writable) {
	m.Retval.WriteTo(writable)
}

// ExtractFrom extracts the UpdateContentMetaDataResponse from the given readable
func (m *UpdateContentMetaDataResponse) ExtractFrom(readable types.Readable) error {
	if err := m.Retval.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read UpdateContentMetaData.retval. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handleUpdateContentMetaData(packet nex.PacketInterface) {
	if protocol.UpdateContentMetaData == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Title Storage Admin::UpdateContentMetaData not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var params UpdateContentMetaDataRequest

	if err := params.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.UpdateContentMetaData(err, packet, callID, params.ContentId, params.Name, params.Description, params.DateAvailable, params.DateExpired, params.Flags, params.ExtendedFlags, params.Extension)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.UpdateContentMetaData(nil, packet, callID, params.ContentId, params.Name, params.Description, params.DateAvailable, params.DateExpired, params.Flags, params.ExtendedFlags, params.Extension)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
