package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	user_storage_types "github.com/PretendoNetwork/nex-protocols-go/v2/ubisoft/user-storage/types"
)

// GetContentsToModerateRequest holds the request parameters of User Storage Admin::GetContentsToModerate
type GetContentsToModerateRequest struct {
	TypeID types.UInt32
	Offset types.UInt32
	Size   types.UInt32
}

// WriteTo writes the GetContentsToModerateRequest to the given writable
func (m GetContentsToModerateRequest) WriteTo(writable types.Writable) {
	m.TypeID.WriteTo(writable)
	m.Offset.WriteTo(writable)
	m.Size.WriteTo(writable)
}

// ExtractFrom extracts the GetContentsToModerateRequest from the given readable
func (m *GetContentsToModerateRequest) ExtractFrom(readable types.Readable) error {
	if err := m.TypeID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read GetContentsToModerate.typeID. %s", err.Error())
	}

	if err := m.Offset.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read GetContentsToModerate.offset. %s", err.Error())
	}

	if err := m.Size.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read GetContentsToModerate.size. %s", err.Error())
	}

	return nil
}

// GetContentsToModerateResponse holds the response parameters of User Storage Admin::GetContentsToModerate
type GetContentsToModerateResponse struct {
	Contents     types.List[user_storage_types.UserContent]
	TotalResults types.UInt32
}

// WriteTo writes the GetContentsToModerateResponse to the given writable
func (m GetContentsToModerateResponse) WriteTo(writable types.Writable) {
	m.Contents.WriteTo(writable)
	m.TotalResults.WriteTo(writable)
}

// ExtractFrom extracts the GetContentsToModerateResponse from the given readable
func (m *GetContentsToModerateResponse) ExtractFrom(readable types.Readable) error {
	if err := m.Contents.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read GetContentsToModerate.contents. %s", err.Error())
	}

	if err := m.TotalResults.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read GetContentsToModerate.totalResults. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handleGetContentsToModerate(packet nex.PacketInterface) {
	if protocol.GetContentsToModerate == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "User Storage Admin::GetContentsToModerate not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var params GetContentsToModerateRequest

	if err := params.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.GetContentsToModerate(err, packet, callID, params.TypeID, params.Offset, params.Size)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetContentsToModerate(nil, packet, callID, params.TypeID, params.Offset, params.Size)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
