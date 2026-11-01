package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	user_storage_types "github.com/PretendoNetwork/nex-protocols-go/v2/ubisoft/user-storage/types"
)

// GetOwnContentsRequest holds the request parameters of User Storage::GetOwnContents
type GetOwnContentsRequest struct {
	TypeID types.UInt32
}

// WriteTo writes the GetOwnContentsRequest to the given writable
func (m GetOwnContentsRequest) WriteTo(writable types.Writable) {
	m.TypeID.WriteTo(writable)
}

// ExtractFrom extracts the GetOwnContentsRequest from the given readable
func (m *GetOwnContentsRequest) ExtractFrom(readable types.Readable) error {
	if err := m.TypeID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read GetOwnContents.typeID. %s", err.Error())
	}

	return nil
}

// GetOwnContentsResponse holds the response parameters of User Storage::GetOwnContents
type GetOwnContentsResponse struct {
	Results types.List[user_storage_types.UserContent]
}

// WriteTo writes the GetOwnContentsResponse to the given writable
func (m GetOwnContentsResponse) WriteTo(writable types.Writable) {
	m.Results.WriteTo(writable)
}

// ExtractFrom extracts the GetOwnContentsResponse from the given readable
func (m *GetOwnContentsResponse) ExtractFrom(readable types.Readable) error {
	if err := m.Results.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read GetOwnContents.results. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handleGetOwnContents(packet nex.PacketInterface) {
	if protocol.GetOwnContents == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "User Storage::GetOwnContents not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var params GetOwnContentsRequest

	if err := params.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.GetOwnContents(err, packet, callID, params.TypeID)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetOwnContents(nil, packet, callID, params.TypeID)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
