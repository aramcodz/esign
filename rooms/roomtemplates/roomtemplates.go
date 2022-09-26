// Copyright 2019 James Cote
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by gen-esign; DO NOT EDIT.

// Package roomtemplates implements the DocuSign SDK
// category RoomTemplates.
// 
// This section shows you how to retrieve room templates. You can perform additional room template tasks in the console.
//
// Service Api documentation may be found at:
// https://developers.docusign.com/docs/rooms-api/reference/RoomTemplates
// Usage example:
//
//   import (
//       "github.com/jfcote87/esign"
//       "github.com/jfcote87/esign/rooms"
//   ) 
//   ...
//   roomtemplatesService := roomtemplates.New(esignCredential)
package roomtemplates // import "github.com/jfcote87/esignrooms//roomtemplates"

import (
    "context"
    "fmt"
    "net/url"
    
    "github.com/jfcote87/esign"
    "github.com/jfcote87/esign/rooms"
)

// Service implements DocuSign RoomTemplates Category API operations
type Service struct {
    credential esign.Credential 
}

// New initializes a roomtemplates service using cred to authorize ops.
func New(cred esign.Credential) *Service {
    return &Service{credential: cred}
}

// GetRoomTemplates gets room templates.
// 
// https://developers.docusign.com/docs/rooms-api/reference/roomtemplates/roomtemplates/getroomtemplates
// 
// SDK Method RoomTemplates::GetRoomTemplates
func (s *Service) GetRoomTemplates() *GetRoomTemplatesOp {
    return &GetRoomTemplatesOp{
        Credential: s.credential,
        Method:  "GET",
        Path: "room_templates",
        Accept: "application/json",
        QueryOpts: make(url.Values),
        Version: esign.RoomsV2,
    }
}

// GetRoomTemplatesOp implements DocuSign API SDK RoomTemplates::GetRoomTemplates
type GetRoomTemplatesOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *GetRoomTemplatesOp) Do(ctx context.Context)  (*rooms.RoomTemplatesSummaryList, error) {
    var res *rooms.RoomTemplatesSummaryList
    return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// OfficeID (Optional) The ID of the office for which the user wants to create a room. When you pass in a value for this parameter, only room templates that are valid for that office appear in the results. For users who are not Admins, the default is the id of the user's default office.
// However, you can specify a value if the user belongs to multiple offices.
// 
// If the user is an Admin, set the `forAdmin` search parameter to **true** instead and omit the `officeId` parameter.
func (op *GetRoomTemplatesOp) OfficeID(val int) *GetRoomTemplatesOp {
    if op != nil {
        op.QueryOpts.Set("officeId", fmt.Sprintf("%d", val ))
    }
    return op
}

// OnlyAssignable (Optional) When set to **true**, returns only the roles that the current user can assign to someone else. The default value is **false**.
func (op *GetRoomTemplatesOp) OnlyAssignable() *GetRoomTemplatesOp {
    if op != nil {
        op.QueryOpts.Set("onlyAssignable", "true")
    }
    return op
}

// OnlyEnabled when set to true, only returns room templates that are not disabled.
func (op *GetRoomTemplatesOp) OnlyEnabled() *GetRoomTemplatesOp {
    if op != nil {
        op.QueryOpts.Set("onlyEnabled", "true")
    }
    return op
}

// Count (Optional) The number of results to return. This value must be a number between `1` and `100` (default).
func (op *GetRoomTemplatesOp) Count(val int) *GetRoomTemplatesOp {
    if op != nil {
        op.QueryOpts.Set("count", fmt.Sprintf("%d", val ))
    }
    return op
}

// StartPosition (Optional) The index position within the total result set from which to start returning values. The default value is `0`.
func (op *GetRoomTemplatesOp) StartPosition(val int) *GetRoomTemplatesOp {
    if op != nil {
        op.QueryOpts.Set("startPosition", fmt.Sprintf("%d", val ))
    }
    return op
}

