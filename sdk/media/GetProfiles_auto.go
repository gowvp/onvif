// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package media

import (
	"context"

	"github.com/gowvp/onvif"
	"github.com/gowvp/onvif/media"
	"github.com/gowvp/onvif/sdk"
	"github.com/juju/errors"
)

// Call_GetProfiles forwards the call to dev.CallMethod() then parses the payload of the reply as a GetProfilesResponse.
func Call_GetProfiles(ctx context.Context, dev *onvif.Device, request media.GetProfiles) (media.GetProfilesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetProfilesResponse media.GetProfilesResponse
			Fault               media.GetDeviceInformationFault
		}
	}
	var reply Envelope
	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.GetProfilesResponse, errors.Annotate(err, "call")
	}
	if err := sdk.ReadAndParse(ctx, httpReply, &reply, "GetProfiles"); err != nil {
		return reply.Body.GetProfilesResponse, errors.Annotate(err, "reply")
	}
	if httpReply.StatusCode == 200 {
		return reply.Body.GetProfilesResponse, nil
	}
	return reply.Body.GetProfilesResponse, errors.New(reply.Body.Fault.Reason.Text)
}

type Fault struct {
	Code struct {
		Value   string `xml:"Value"`
		Subcode struct {
			Value string `xml:"Value"`
		} `xml:"Subcode"`
	} `xml:"Code"`
	Reason struct {
		Text string `xml:"Text"`
	} `xml:"Reason"`
}
