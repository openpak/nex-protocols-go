package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

// BanUserFromContentTypeRequest holds the request parameters of User Storage Admin::BanUserFromContentType
type BanUserFromContentTypeRequest struct {
	TypeID      types.UInt32
	Pid         types.UInt32
	Reason      types.String
	BanContents types.Bool
	ExpireDate  types.DateTime
}

// WriteTo writes the BanUserFromContentTypeRequest to the given writable
func (m BanUserFromContentTypeRequest) WriteTo(writable types.Writable) {
	m.TypeID.WriteTo(writable)
	m.Pid.WriteTo(writable)
	m.Reason.WriteTo(writable)
	m.BanContents.WriteTo(writable)
	m.ExpireDate.WriteTo(writable)
}

// ExtractFrom extracts the BanUserFromContentTypeRequest from the given readable
func (m *BanUserFromContentTypeRequest) ExtractFrom(readable types.Readable) error {
	if err := m.TypeID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read BanUserFromContentType.typeID. %s", err.Error())
	}

	if err := m.Pid.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read BanUserFromContentType.pid. %s", err.Error())
	}

	if err := m.Reason.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read BanUserFromContentType.reason. %s", err.Error())
	}

	if err := m.BanContents.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read BanUserFromContentType.banContents. %s", err.Error())
	}

	if err := m.ExpireDate.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read BanUserFromContentType.expireDate. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handleBanUserFromContentType(packet nex.PacketInterface) {
	if protocol.BanUserFromContentType == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "User Storage Admin::BanUserFromContentType not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var params BanUserFromContentTypeRequest

	if err := params.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.BanUserFromContentType(err, packet, callID, params.TypeID, params.Pid, params.Reason, params.BanContents, params.ExpireDate)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.BanUserFromContentType(nil, packet, callID, params.TypeID, params.Pid, params.Reason, params.BanContents, params.ExpireDate)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
