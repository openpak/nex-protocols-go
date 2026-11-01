package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	user_storage_types "github.com/PretendoNetwork/nex-protocols-go/v2/ubisoft/user-storage/types"
)

// GetContentsToModerateWithThresholdRequest holds the request parameters of User Storage Admin::GetContentsToModerateWithThreshold
type GetContentsToModerateWithThresholdRequest struct {
	TypeID    types.UInt32
	Threshold types.UInt32
	Offset    types.UInt32
	Size      types.UInt32
}

// WriteTo writes the GetContentsToModerateWithThresholdRequest to the given writable
func (m GetContentsToModerateWithThresholdRequest) WriteTo(writable types.Writable) {
	m.TypeID.WriteTo(writable)
	m.Threshold.WriteTo(writable)
	m.Offset.WriteTo(writable)
	m.Size.WriteTo(writable)
}

// ExtractFrom extracts the GetContentsToModerateWithThresholdRequest from the given readable
func (m *GetContentsToModerateWithThresholdRequest) ExtractFrom(readable types.Readable) error {
	if err := m.TypeID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read GetContentsToModerateWithThreshold.typeID. %s", err.Error())
	}

	if err := m.Threshold.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read GetContentsToModerateWithThreshold.threshold. %s", err.Error())
	}

	if err := m.Offset.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read GetContentsToModerateWithThreshold.offset. %s", err.Error())
	}

	if err := m.Size.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read GetContentsToModerateWithThreshold.size. %s", err.Error())
	}

	return nil
}

// GetContentsToModerateWithThresholdResponse holds the response parameters of User Storage Admin::GetContentsToModerateWithThreshold
type GetContentsToModerateWithThresholdResponse struct {
	Contents     types.List[user_storage_types.UserContent]
	TotalResults types.UInt32
}

// WriteTo writes the GetContentsToModerateWithThresholdResponse to the given writable
func (m GetContentsToModerateWithThresholdResponse) WriteTo(writable types.Writable) {
	m.Contents.WriteTo(writable)
	m.TotalResults.WriteTo(writable)
}

// ExtractFrom extracts the GetContentsToModerateWithThresholdResponse from the given readable
func (m *GetContentsToModerateWithThresholdResponse) ExtractFrom(readable types.Readable) error {
	if err := m.Contents.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read GetContentsToModerateWithThreshold.contents. %s", err.Error())
	}

	if err := m.TotalResults.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read GetContentsToModerateWithThreshold.totalResults. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handleGetContentsToModerateWithThreshold(packet nex.PacketInterface) {
	if protocol.GetContentsToModerateWithThreshold == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "User Storage Admin::GetContentsToModerateWithThreshold not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var params GetContentsToModerateWithThresholdRequest

	if err := params.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.GetContentsToModerateWithThreshold(err, packet, callID, params.TypeID, params.Threshold, params.Offset, params.Size)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetContentsToModerateWithThreshold(nil, packet, callID, params.TypeID, params.Threshold, params.Offset, params.Size)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
