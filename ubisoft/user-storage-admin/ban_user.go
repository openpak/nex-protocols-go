package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

// BanUserRequest holds the request parameters of User Storage Admin::BanUser
type BanUserRequest struct {
	Pid         types.UInt32
	Reason      types.String
	BanContents types.Bool
	ExpireDate  types.DateTime
}

// WriteTo writes the BanUserRequest to the given writable
func (m BanUserRequest) WriteTo(writable types.Writable) {
	m.Pid.WriteTo(writable)
	m.Reason.WriteTo(writable)
	m.BanContents.WriteTo(writable)
	m.ExpireDate.WriteTo(writable)
}

// ExtractFrom extracts the BanUserRequest from the given readable
func (m *BanUserRequest) ExtractFrom(readable types.Readable) error {
	if err := m.Pid.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read BanUser.pid. %s", err.Error())
	}

	if err := m.Reason.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read BanUser.reason. %s", err.Error())
	}

	if err := m.BanContents.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read BanUser.banContents. %s", err.Error())
	}

	if err := m.ExpireDate.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read BanUser.expireDate. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handleBanUser(packet nex.PacketInterface) {
	if protocol.BanUser == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "User Storage Admin::BanUser not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var params BanUserRequest

	if err := params.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.BanUser(err, packet, callID, params.Pid, params.Reason, params.BanContents, params.ExpireDate)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.BanUser(nil, packet, callID, params.Pid, params.Reason, params.BanContents, params.ExpireDate)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
