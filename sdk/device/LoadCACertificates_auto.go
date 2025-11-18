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

// Call_LoadCACertificates forwards the call to dev.CallMethod() then parses the payload of the reply as a LoadCACertificatesResponse.
func Call_LoadCACertificates(ctx context.Context, dev *onvif.Device, request device.LoadCACertificates) (device.LoadCACertificatesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			LoadCACertificatesResponse device.LoadCACertificatesResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.LoadCACertificatesResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "LoadCACertificates")
		return reply.Body.LoadCACertificatesResponse, errors.Annotate(err, "reply")
	}
}
