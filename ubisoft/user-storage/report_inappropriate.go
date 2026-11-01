package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	user_storage_types "github.com/PretendoNetwork/nex-protocols-go/v2/ubisoft/user-storage/types"
)

// ReportInappropriateRequest holds the request parameters of User Storage::ReportInappropriate
type ReportInappropriateRequest struct {
	ContentKey user_storage_types.UserContentKey
	Reason     types.String
}

// WriteTo writes the ReportInappropriateRequest to the given writable
func (m ReportInappropriateRequest) WriteTo(writable types.Writable) {
	m.ContentKey.WriteTo(writable)
	m.Reason.WriteTo(writable)
}

// ExtractFrom extracts the ReportInappropriateRequest from the given readable
func (m *ReportInappropriateRequest) ExtractFrom(readable types.Readable) error {
	if err := m.ContentKey.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read ReportInappropriate.contentKey. %s", err.Error())
	}

	if err := m.Reason.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read ReportInappropriate.reason. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handleReportInappropriate(packet nex.PacketInterface) {
	if protocol.ReportInappropriate == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "User Storage::ReportInappropriate not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var params ReportInappropriateRequest

	if err := params.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.ReportInappropriate(err, packet, callID, params.ContentKey, params.Reason)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.ReportInappropriate(nil, packet, callID, params.ContentKey, params.Reason)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
