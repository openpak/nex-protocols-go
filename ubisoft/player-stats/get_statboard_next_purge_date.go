package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

// GetStatboardNextPurgeDateRequest holds the request parameters of Player Stats::GetStatboardNextPurgeDate
type GetStatboardNextPurgeDateRequest struct {
	BoardID        types.UInt32
	ResetFrequency types.UInt32
}

// WriteTo writes the GetStatboardNextPurgeDateRequest to the given writable
func (m GetStatboardNextPurgeDateRequest) WriteTo(writable types.Writable) {
	m.BoardID.WriteTo(writable)
	m.ResetFrequency.WriteTo(writable)
}

// ExtractFrom extracts the GetStatboardNextPurgeDateRequest from the given readable
func (m *GetStatboardNextPurgeDateRequest) ExtractFrom(readable types.Readable) error {
	if err := m.BoardID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read GetStatboardNextPurgeDate.boardID. %s", err.Error())
	}

	if err := m.ResetFrequency.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read GetStatboardNextPurgeDate.resetFrequency. %s", err.Error())
	}

	return nil
}

// GetStatboardNextPurgeDateResponse holds the response parameters of Player Stats::GetStatboardNextPurgeDate
type GetStatboardNextPurgeDateResponse struct {
	PurgeDate types.DateTime
}

// WriteTo writes the GetStatboardNextPurgeDateResponse to the given writable
func (m GetStatboardNextPurgeDateResponse) WriteTo(writable types.Writable) {
	m.PurgeDate.WriteTo(writable)
}

// ExtractFrom extracts the GetStatboardNextPurgeDateResponse from the given readable
func (m *GetStatboardNextPurgeDateResponse) ExtractFrom(readable types.Readable) error {
	if err := m.PurgeDate.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read GetStatboardNextPurgeDate.purgeDate. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handleGetStatboardNextPurgeDate(packet nex.PacketInterface) {
	if protocol.GetStatboardNextPurgeDate == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Player Stats::GetStatboardNextPurgeDate not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var params GetStatboardNextPurgeDateRequest

	if err := params.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.GetStatboardNextPurgeDate(err, packet, callID, params.BoardID, params.ResetFrequency)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetStatboardNextPurgeDate(nil, packet, callID, params.BoardID, params.ResetFrequency)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
