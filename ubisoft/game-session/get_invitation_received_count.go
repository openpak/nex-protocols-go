package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

// GetInvitationReceivedCountRequest holds the request parameters of Game Session::GetInvitationReceivedCount
type GetInvitationReceivedCountRequest struct {
	GameSessionTypeID types.UInt32
}

// WriteTo writes the GetInvitationReceivedCountRequest to the given writable
func (m GetInvitationReceivedCountRequest) WriteTo(writable types.Writable) {
	m.GameSessionTypeID.WriteTo(writable)
}

// ExtractFrom extracts the GetInvitationReceivedCountRequest from the given readable
func (m *GetInvitationReceivedCountRequest) ExtractFrom(readable types.Readable) error {
	if err := m.GameSessionTypeID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read GetInvitationReceivedCount.gameSessionTypeID. %s", err.Error())
	}

	return nil
}

// GetInvitationReceivedCountResponse holds the response parameters of Game Session::GetInvitationReceivedCount
type GetInvitationReceivedCountResponse struct {
	Count types.UInt32
}

// WriteTo writes the GetInvitationReceivedCountResponse to the given writable
func (m GetInvitationReceivedCountResponse) WriteTo(writable types.Writable) {
	m.Count.WriteTo(writable)
}

// ExtractFrom extracts the GetInvitationReceivedCountResponse from the given readable
func (m *GetInvitationReceivedCountResponse) ExtractFrom(readable types.Readable) error {
	if err := m.Count.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read GetInvitationReceivedCount.count. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handleGetInvitationReceivedCount(packet nex.PacketInterface) {
	if protocol.GetInvitationReceivedCount == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Game Session::GetInvitationReceivedCount not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var params GetInvitationReceivedCountRequest

	if err := params.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.GetInvitationReceivedCount(err, packet, callID, params.GameSessionTypeID)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetInvitationReceivedCount(nil, packet, callID, params.GameSessionTypeID)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
