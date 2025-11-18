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

// Call_GetCertificateInformation forwards the call to dev.CallMethod() then parses the payload of the reply as a GetCertificateInformationResponse.
func Call_GetCertificateInformation(ctx context.Context, dev *onvif.Device, request device.GetCertificateInformation) (device.GetCertificateInformationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetCertificateInformationResponse device.GetCertificateInformationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetCertificateInformationResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetCertificateInformation")
		return reply.Body.GetCertificateInformationResponse, errors.Annotate(err, "reply")
	}
}
