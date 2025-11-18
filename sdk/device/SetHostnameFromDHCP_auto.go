// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package device

import (
	"context"

	"github.com/gowvp/onvif"
	"github.com/gowvp/onvif/device"
	"github.com/gowvp/onvif/sdk"
	"github.com/juju/errors"
)

// Call_SetHostnameFromDHCP forwards the call to dev.CallMethod() then parses the payload of the reply as a SetHostnameFromDHCPResponse.
func Call_SetHostnameFromDHCP(ctx context.Context, dev *onvif.Device, request device.SetHostnameFromDHCP) (device.SetHostnameFromDHCPResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetHostnameFromDHCPResponse device.SetHostnameFromDHCPResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.SetHostnameFromDHCPResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "SetHostnameFromDHCP")
		return reply.Body.SetHostnameFromDHCPResponse, errors.Annotate(err, "reply")
	}
}
