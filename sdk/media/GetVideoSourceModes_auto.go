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

// Call_GetVideoSourceModes forwards the call to dev.CallMethod() then parses the payload of the reply as a GetVideoSourceModesResponse.
func Call_GetVideoSourceModes(ctx context.Context, dev *onvif.Device, request media.GetVideoSourceModes) (media.GetVideoSourceModesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetVideoSourceModesResponse media.GetVideoSourceModesResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetVideoSourceModesResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetVideoSourceModes")
		return reply.Body.GetVideoSourceModesResponse, errors.Annotate(err, "reply")
	}
}
