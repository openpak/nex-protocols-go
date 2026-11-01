package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	user_storage_types "github.com/PretendoNetwork/nex-protocols-go/v2/ubisoft/user-storage/types"
)

// GetSlotCountRequest holds the request parameters of User Storage::GetSlotCount
type GetSlotCountRequest struct {
	TypeID types.UInt32
}

// WriteTo writes the GetSlotCountRequest to the given writable
func (m GetSlotCountRequest) WriteTo(writable types.Writable) {
	m.TypeID.WriteTo(writable)
}

// ExtractFrom extracts the GetSlotCountRequest from the given readable
func (m *GetSlotCountRequest) ExtractFrom(readable types.Readable) error {
	if err := m.TypeID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read GetSlotCount.typeID. %s", err.Error())
	}

	return nil
}

// GetSlotCountResponse holds the response parameters of User Storage::GetSlotCount
type GetSlotCountResponse struct {
	SlotCount user_storage_types.UserSlotCount
}

// WriteTo writes the GetSlotCountResponse to the given writable
func (m GetSlotCountResponse) WriteTo(writable types.Writable) {
	m.SlotCount.WriteTo(writable)
}

// ExtractFrom extracts the GetSlotCountResponse from the given readable
func (m *GetSlotCountResponse) ExtractFrom(readable types.Readable) error {
	if err := m.SlotCount.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read GetSlotCount.slotCount. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handleGetSlotCount(packet nex.PacketInterface) {
	if protocol.GetSlotCount == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "User Storage::GetSlotCount not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var params GetSlotCountRequest

	if err := params.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.GetSlotCount(err, packet, callID, params.TypeID)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetSlotCount(nil, packet, callID, params.TypeID)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
