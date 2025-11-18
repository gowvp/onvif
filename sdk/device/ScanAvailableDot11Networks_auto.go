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

// Call_ScanAvailableDot11Networks forwards the call to dev.CallMethod() then parses the payload of the reply as a ScanAvailableDot11NetworksResponse.
func Call_ScanAvailableDot11Networks(ctx context.Context, dev *onvif.Device, request device.ScanAvailableDot11Networks) (device.ScanAvailableDot11NetworksResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			ScanAvailableDot11NetworksResponse device.ScanAvailableDot11NetworksResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.ScanAvailableDot11NetworksResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "ScanAvailableDot11Networks")
		return reply.Body.ScanAvailableDot11NetworksResponse, errors.Annotate(err, "reply")
	}
}
