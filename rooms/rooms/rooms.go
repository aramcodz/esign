// Copyright 2019 James Cote
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by gen-esign; DO NOT EDIT.

// Package rooms implements the DocuSign SDK
// category Rooms.
// 
// A room can hold documents, envelopes, a list of tasks comprising a workflow, and other related information. You can invite others to this space and assign them permissions on a per-room basis. The documentation in this section shows you how to perform these and other tasks.
//
// Service Api documentation may be found at:
// https://developers.docusign.com/docs/rooms-api/reference/Rooms
// Usage example:
//
//   import (
//       "github.com/jfcote87/esign"
//       "github.com/jfcote87/esign/rooms"
//   ) 
//   ...
//   roomsService := rooms.New(esignCredential)
package rooms // import "github.com/jfcote87/esignrooms//rooms"

import (
    "context"
    "fmt"
    "io"
    "net/url"
    "strings"
    
    "github.com/jfcote87/esign"
    "github.com/jfcote87/esign/rooms"
)

// Service implements DocuSign Rooms Category API operations
type Service struct {
    credential esign.Credential 
}

// New initializes a rooms service using cred to authorize ops.
func New(cred esign.Credential) *Service {
    return &Service{credential: cred}
}

// AddDocumentToRoom adds a document to a room.
// 
// https://developers.docusign.com/docs/rooms-api/reference/rooms/rooms/adddocumenttoroom
// 
// SDK Method Rooms::AddDocumentToRoom
func (s *Service) AddDocumentToRoom(roomID string, body *rooms.Document) *AddDocumentToRoomOp {
    return &AddDocumentToRoomOp{
        Credential: s.credential,
        Method:  "POST",
        Path: strings.Join([]string{"rooms",roomID,"documents"},"/"),
        Payload: body,
        Accept: "application/json-patch+json, application/json, text/json, application/*+json",
        QueryOpts: make(url.Values),
        Version: esign.RoomsV2,
    }
}

// AddDocumentToRoomOp implements DocuSign API SDK Rooms::AddDocumentToRoom
type AddDocumentToRoomOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *AddDocumentToRoomOp) Do(ctx context.Context)  (*rooms.RoomDocument, error) {
    var res *rooms.RoomDocument
    return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// AddDocumentToRoomViaFileUpload uploads the contents of a file as a document to a room.
// If media is an io.ReadCloser, Do() will close media.
// 
// https://developers.docusign.com/docs/rooms-api/reference/rooms/rooms/adddocumenttoroomviafileupload
// 
// SDK Method Rooms::AddDocumentToRoomViaFileUpload
func (s *Service) AddDocumentToRoomViaFileUpload(roomID string, media io.Reader, mimeType string) *AddDocumentToRoomViaFileUploadOp {
    return &AddDocumentToRoomViaFileUploadOp{
        Credential: s.credential,
        Method:  "POST",
        Path: strings.Join([]string{"rooms",roomID,"documents","contents"},"/"),
        Payload: &esign.UploadFile{ Reader: media, ContentType: mimeType },
        Accept: "multipart/form-data",
        QueryOpts: make(url.Values),
        Version: esign.RoomsV2,
    }
}

// AddDocumentToRoomViaFileUploadOp implements DocuSign API SDK Rooms::AddDocumentToRoomViaFileUpload
type AddDocumentToRoomViaFileUploadOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *AddDocumentToRoomViaFileUploadOp) Do(ctx context.Context)  (*rooms.RoomDocument, error) {
    var res *rooms.RoomDocument
    return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// AddFormToRoom adds a form to a room.
// 
// https://developers.docusign.com/docs/rooms-api/reference/rooms/rooms/addformtoroom
// 
// SDK Method Rooms::AddFormToRoom
func (s *Service) AddFormToRoom(roomID string, body *rooms.FormForAdd) *AddFormToRoomOp {
    return &AddFormToRoomOp{
        Credential: s.credential,
        Method:  "POST",
        Path: strings.Join([]string{"rooms",roomID,"forms"},"/"),
        Payload: body,
        Accept: "application/json",
        QueryOpts: make(url.Values),
        Version: esign.RoomsV2,
    }
}

// AddFormToRoomOp implements DocuSign API SDK Rooms::AddFormToRoom
type AddFormToRoomOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *AddFormToRoomOp) Do(ctx context.Context)  (*rooms.RoomDocument, error) {
    var res *rooms.RoomDocument
    return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// CreateRoom creates a room.
// 
// https://developers.docusign.com/docs/rooms-api/reference/rooms/rooms/createroom
// 
// SDK Method Rooms::CreateRoom
func (s *Service) CreateRoom(body *rooms.RoomForCreate) *CreateRoomOp {
    return &CreateRoomOp{
        Credential: s.credential,
        Method:  "POST",
        Path: "rooms",
        Payload: body,
        Accept: "application/json-patch+json, application/json, text/json, application/*+json",
        QueryOpts: make(url.Values),
        Version: esign.RoomsV2,
    }
}

// CreateRoomOp implements DocuSign API SDK Rooms::CreateRoom
type CreateRoomOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *CreateRoomOp) Do(ctx context.Context)  (*rooms.Room, error) {
    var res *rooms.Room
    return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// DeleteRoom deletes a room.
// 
// https://developers.docusign.com/docs/rooms-api/reference/rooms/rooms/deleteroom
// 
// SDK Method Rooms::DeleteRoom
func (s *Service) DeleteRoom(roomID string) *DeleteRoomOp {
    return &DeleteRoomOp{
        Credential: s.credential,
        Method:  "DELETE",
        Path: strings.Join([]string{"rooms",roomID},"/"),
        QueryOpts: make(url.Values),
        Version: esign.RoomsV2,
    }
}

// DeleteRoomOp implements DocuSign API SDK Rooms::DeleteRoom
type DeleteRoomOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *DeleteRoomOp) Do(ctx context.Context)  error {
    return ((*esign.Op)(op)).Do(ctx, nil)
}

// GetAssignableRoles gets assignable room-level roles in v6.
// 
// https://developers.docusign.com/docs/rooms-api/reference/rooms/rooms/getassignableroles
// 
// SDK Method Rooms::GetAssignableRoles
func (s *Service) GetAssignableRoles(roomID string) *GetAssignableRolesOp {
    return &GetAssignableRolesOp{
        Credential: s.credential,
        Method:  "GET",
        Path: strings.Join([]string{"rooms",roomID,"assignable_roles"},"/"),
        Accept: "application/json",
        QueryOpts: make(url.Values),
        Version: esign.RoomsV2,
    }
}

// GetAssignableRolesOp implements DocuSign API SDK Rooms::GetAssignableRoles
type GetAssignableRolesOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *GetAssignableRolesOp) Do(ctx context.Context)  (*rooms.AssignableRoles, error) {
    var res *rooms.AssignableRoles
    return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// AssigneeEmail (Optional) The email address of a specific member. Using this parameter returns only the roles that the current user can assign to the member with that email address.
func (op *GetAssignableRolesOp) AssigneeEmail(val string) *GetAssignableRolesOp {
    if op != nil {
        op.QueryOpts.Set("assigneeEmail", val)
    }
    return op
}

// Filter (Optional) A search filter that returns assignable roles by the beginning of the role name.
// 
// **Note**: You do not enter a wildcard (*) at the end of the name fragment.
func (op *GetAssignableRolesOp) Filter(val string) *GetAssignableRolesOp {
    if op != nil {
        op.QueryOpts.Set("filter", val)
    }
    return op
}

// StartPosition (Optional) The index position within the total result set from which to start returning values. The default value is `0`.
func (op *GetAssignableRolesOp) StartPosition(val int) *GetAssignableRolesOp {
    if op != nil {
        op.QueryOpts.Set("startPosition", fmt.Sprintf("%d", val ))
    }
    return op
}

// Count (Optional) The number of results to return. This value must be a number between `1` and `100` (default).
func (op *GetAssignableRolesOp) Count(val int) *GetAssignableRolesOp {
    if op != nil {
        op.QueryOpts.Set("count", fmt.Sprintf("%d", val ))
    }
    return op
}

// GetDocuments gets a list of documents in a room.
// 
// https://developers.docusign.com/docs/rooms-api/reference/rooms/rooms/getdocuments
// 
// SDK Method Rooms::GetDocuments
func (s *Service) GetDocuments(roomID string) *GetDocumentsOp {
    return &GetDocumentsOp{
        Credential: s.credential,
        Method:  "GET",
        Path: strings.Join([]string{"rooms",roomID,"documents"},"/"),
        Accept: "application/json",
        QueryOpts: make(url.Values),
        Version: esign.RoomsV2,
    }
}

// GetDocumentsOp implements DocuSign API SDK Rooms::GetDocuments
type GetDocumentsOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *GetDocumentsOp) Do(ctx context.Context)  (*rooms.RoomDocumentList, error) {
    var res *rooms.RoomDocumentList
    return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// Count (Optional) The number of results to return. This value must be a number between `1` and `100` (default).
func (op *GetDocumentsOp) Count(val int) *GetDocumentsOp {
    if op != nil {
        op.QueryOpts.Set("count", fmt.Sprintf("%d", val ))
    }
    return op
}

// StartPosition (Optional) The index position within the total result set from which to start returning values. The default value is `0`.
func (op *GetDocumentsOp) StartPosition(val int) *GetDocumentsOp {
    if op != nil {
        op.QueryOpts.Set("startPosition", fmt.Sprintf("%d", val ))
    }
    return op
}

// GetRoom gets a room.
// 
// https://developers.docusign.com/docs/rooms-api/reference/rooms/rooms/getroom
// 
// SDK Method Rooms::GetRoom
func (s *Service) GetRoom(roomID string) *GetRoomOp {
    return &GetRoomOp{
        Credential: s.credential,
        Method:  "GET",
        Path: strings.Join([]string{"rooms",roomID},"/"),
        Accept: "application/json",
        QueryOpts: make(url.Values),
        Version: esign.RoomsV2,
    }
}

// GetRoomOp implements DocuSign API SDK Rooms::GetRoom
type GetRoomOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *GetRoomOp) Do(ctx context.Context)  (*rooms.Room, error) {
    var res *rooms.Room
    return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// IncludeFieldData (Optional) When set to **true**, the response includes the field data from the room. This is the information that appears on the room's **Details** tab.
func (op *GetRoomOp) IncludeFieldData() *GetRoomOp {
    if op != nil {
        op.QueryOpts.Set("includeFieldData", "true")
    }
    return op
}

// GetRoomFieldData gets a room's field data.
// 
// https://developers.docusign.com/docs/rooms-api/reference/rooms/rooms/getroomfielddata
// 
// SDK Method Rooms::GetRoomFieldData
func (s *Service) GetRoomFieldData(roomID string) *GetRoomFieldDataOp {
    return &GetRoomFieldDataOp{
        Credential: s.credential,
        Method:  "GET",
        Path: strings.Join([]string{"rooms",roomID,"field_data"},"/"),
        Accept: "application/json",
        QueryOpts: make(url.Values),
        Version: esign.RoomsV2,
    }
}

// GetRoomFieldDataOp implements DocuSign API SDK Rooms::GetRoomFieldData
type GetRoomFieldDataOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *GetRoomFieldDataOp) Do(ctx context.Context)  (*rooms.FieldData, error) {
    var res *rooms.FieldData
    return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// GetRoomFieldSet gets the field set for a room.
// 
// https://developers.docusign.com/docs/rooms-api/reference/rooms/rooms/getroomfieldset
// 
// SDK Method Rooms::GetRoomFieldSet
func (s *Service) GetRoomFieldSet(roomID string) *GetRoomFieldSetOp {
    return &GetRoomFieldSetOp{
        Credential: s.credential,
        Method:  "GET",
        Path: strings.Join([]string{"rooms",roomID,"field_set"},"/"),
        Accept: "application/json",
        QueryOpts: make(url.Values),
        Version: esign.RoomsV2,
    }
}

// GetRoomFieldSetOp implements DocuSign API SDK Rooms::GetRoomFieldSet
type GetRoomFieldSetOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *GetRoomFieldSetOp) Do(ctx context.Context)  (*rooms.FieldSet, error) {
    var res *rooms.FieldSet
    return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// GetRoomUsers gets a room's users.
// 
// https://developers.docusign.com/docs/rooms-api/reference/rooms/rooms/getroomusers
// 
// SDK Method Rooms::GetRoomUsers
func (s *Service) GetRoomUsers(roomID string) *GetRoomUsersOp {
    return &GetRoomUsersOp{
        Credential: s.credential,
        Method:  "GET",
        Path: strings.Join([]string{"rooms",roomID,"users"},"/"),
        Accept: "application/json",
        QueryOpts: make(url.Values),
        Version: esign.RoomsV2,
    }
}

// GetRoomUsersOp implements DocuSign API SDK Rooms::GetRoomUsers
type GetRoomUsersOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *GetRoomUsersOp) Do(ctx context.Context)  (*rooms.RoomUsersResult, error) {
    var res *rooms.RoomUsersResult
    return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// Count (Optional) The number of results to return. This value must be a number between `1` and `100` (default).
func (op *GetRoomUsersOp) Count(val int) *GetRoomUsersOp {
    if op != nil {
        op.QueryOpts.Set("count", fmt.Sprintf("%d", val ))
    }
    return op
}

// StartPosition (Optional) The index position within the total result set from which to start returning values. The default value is `0`.
func (op *GetRoomUsersOp) StartPosition(val int) *GetRoomUsersOp {
    if op != nil {
        op.QueryOpts.Set("startPosition", fmt.Sprintf("%d", val ))
    }
    return op
}

// Filter (Optional) A search filter that returns users by the beginning of the user's first name, last name, or email address. You can enter the beginning of the name or email only to return all of the users whose names or email addresses begin with the text that you entered. 
// 
// **Note**: You do not enter a wildcard (*) at the end of the name or email fragment.
func (op *GetRoomUsersOp) Filter(val string) *GetRoomUsersOp {
    if op != nil {
        op.QueryOpts.Set("filter", val)
    }
    return op
}

// Sort (Optional) The order in which to return results. Valid values are:
// 
// - `firstNameAsc`: Sort on first name in ascending order. 
// - `firstNameDesc`:  Sort on first name in descending order. 
// - `lastNameAsc`: Sort on last name in ascending order. 
// - `lastNameDesc`: Sort on last name in descending order. This is the default value.
func (op *GetRoomUsersOp) Sort(val string) *GetRoomUsersOp {
    if op != nil {
        op.QueryOpts.Set("sort", val)
    }
    return op
}

// GetRooms returns a list of rooms.
// 
// https://developers.docusign.com/docs/rooms-api/reference/rooms/rooms/getrooms
// 
// SDK Method Rooms::GetRooms
func (s *Service) GetRooms() *GetRoomsOp {
    return &GetRoomsOp{
        Credential: s.credential,
        Method:  "GET",
        Path: "rooms",
        Accept: "application/json",
        QueryOpts: make(url.Values),
        Version: esign.RoomsV2,
    }
}

// GetRoomsOp implements DocuSign API SDK Rooms::GetRooms
type GetRoomsOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *GetRoomsOp) Do(ctx context.Context)  (*rooms.RoomSummaryList, error) {
    var res *rooms.RoomSummaryList
    return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// Count (Optional) The number of results. When this property is used as a request parameter specifying the number of results to return, the value must be a number between 1 and 100 (default).
func (op *GetRoomsOp) Count(val int) *GetRoomsOp {
    if op != nil {
        op.QueryOpts.Set("count", fmt.Sprintf("%d", val ))
    }
    return op
}

// StartPosition (Optional) The index position within the total result set from which to start returning values. The default value is `0`.
func (op *GetRoomsOp) StartPosition(val int) *GetRoomsOp {
    if op != nil {
        op.QueryOpts.Set("startPosition", fmt.Sprintf("%d", val ))
    }
    return op
}

// RoomStatus (Optional) The status of the room. Valid values are:
// 
// - `active`: This is the default value.
// - `pending`
// - `open`: Includes both `active` and `pending` statuses.
// - `closed`
// 
// To return rooms with multiple statuses, enter a comma-separated list of statuses. 
// 
// Example:
// 
// `closed,open`
func (op *GetRoomsOp) RoomStatus(val string) *GetRoomsOp {
    if op != nil {
        op.QueryOpts.Set("roomStatus", val)
    }
    return op
}

// OfficeID (Optional) The id of the office.
func (op *GetRoomsOp) OfficeID(val int) *GetRoomsOp {
    if op != nil {
        op.QueryOpts.Set("officeId", fmt.Sprintf("%d", val ))
    }
    return op
}

// FieldDataChangedStartDate (Optional) Starting date and time to filter rooms whose field data has changed after this date. Date and time is always given as UTC. If the time (`hh:mm:ss`) is omitted, it defaults to `00:00:00`.
// 
// Valid formats:
// 
// - `yyyy-mm-dd hh:mm:ss`
// - `yyyy/mm/dd hh:mm:ss`
// 
// The default start date is the beginning of time.
func (op *GetRoomsOp) FieldDataChangedStartDate(val string) *GetRoomsOp {
    if op != nil {
        op.QueryOpts.Set("fieldDataChangedStartDate", val)
    }
    return op
}

// FieldDataChangedEndDate (Optional) Ending date and time to filter rooms whose field data has changed before this date. Date and time is always given as UTC. If the time (`hh:mm:ss`) is omitted, it defaults to `00:00:00`.
// 
// 
// Valid formats:
// 
// - `yyyy-mm-dd hh:mm:ss`
// - `yyyy/mm/dd hh:mm:ss`
// 
// If this query parameter is omitted, the default end date is now.
func (op *GetRoomsOp) FieldDataChangedEndDate(val string) *GetRoomsOp {
    if op != nil {
        op.QueryOpts.Set("fieldDataChangedEndDate", val)
    }
    return op
}

// RoomClosedStartDate (Optional) Starting date and time to filter rooms that were closed this date. Date and time is always given as UTC. If the time (`hh:mm:ss`) is omitted, it defaults to `00:00:00`.
// 
// Valid formats:
// 
// - `yyyy-mm-dd hh:mm:ss`
// - `yyyy/mm/dd hh:mm:ss`
// 
// The default start date is the beginning of time.
func (op *GetRoomsOp) RoomClosedStartDate(val string) *GetRoomsOp {
    if op != nil {
        op.QueryOpts.Set("roomClosedStartDate", val)
    }
    return op
}

// RoomClosedEndDate (Optional) Ending date and time to filter rooms that were closed before this date. Date and time is always given as UTC. If the time (`hh:mm:ss`) is omitted, it defaults to `00:00:00`.
// 
// Valid formats:
// 
// - `yyyy-mm-dd hh:mm:ss`
// - `yyyy/mm/dd hh:mm:ss`
// 
// If this query parameter is omitted, the default end date is now.
func (op *GetRoomsOp) RoomClosedEndDate(val string) *GetRoomsOp {
    if op != nil {
        op.QueryOpts.Set("roomClosedEndDate", val)
    }
    return op
}

// InviteUser invites a user to a room.
// 
// https://developers.docusign.com/docs/rooms-api/reference/rooms/rooms/inviteuser
// 
// SDK Method Rooms::InviteUser
func (s *Service) InviteUser(roomID string, body *rooms.RoomInvite) *InviteUserOp {
    return &InviteUserOp{
        Credential: s.credential,
        Method:  "POST",
        Path: strings.Join([]string{"rooms",roomID,"users"},"/"),
        Payload: body,
        Accept: "application/json",
        QueryOpts: make(url.Values),
        Version: esign.RoomsV2,
    }
}

// InviteUserOp implements DocuSign API SDK Rooms::InviteUser
type InviteUserOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *InviteUserOp) Do(ctx context.Context)  (*rooms.RoomInviteResponse, error) {
    var res *rooms.RoomInviteResponse
    return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// PutRoomUser updates a room user.
// 
// https://developers.docusign.com/docs/rooms-api/reference/rooms/rooms/putroomuser
// 
// SDK Method Rooms::PutRoomUser
func (s *Service) PutRoomUser(roomID string, userID string, body *rooms.RoomUserForUpdate) *PutRoomUserOp {
    return &PutRoomUserOp{
        Credential: s.credential,
        Method:  "PUT",
        Path: strings.Join([]string{"rooms",roomID,"users",userID},"/"),
        Payload: body,
        Accept: "application/json",
        QueryOpts: make(url.Values),
        Version: esign.RoomsV2,
    }
}

// PutRoomUserOp implements DocuSign API SDK Rooms::PutRoomUser
type PutRoomUserOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *PutRoomUserOp) Do(ctx context.Context)  (*rooms.RoomUser, error) {
    var res *rooms.RoomUser
    return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// RestoreRoomUserAccess restores the specified user's access to the room.
// If media is an io.ReadCloser, Do() will close media.
// 
// https://developers.docusign.com/docs/rooms-api/reference/rooms/rooms/restoreroomuseraccess
// 
// SDK Method Rooms::RestoreRoomUserAccess
func (s *Service) RestoreRoomUserAccess(roomID string, userID string, media io.Reader, mimeType string) *RestoreRoomUserAccessOp {
    return &RestoreRoomUserAccessOp{
        Credential: s.credential,
        Method:  "POST",
        Path: strings.Join([]string{"rooms",roomID,"users",userID,"restore_access"},"/"),
        Payload: &esign.UploadFile{ Reader: media, ContentType: mimeType },
        QueryOpts: make(url.Values),
        Version: esign.RoomsV2,
    }
}

// RestoreRoomUserAccessOp implements DocuSign API SDK Rooms::RestoreRoomUserAccess
type RestoreRoomUserAccessOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *RestoreRoomUserAccessOp) Do(ctx context.Context)  error {
    return ((*esign.Op)(op)).Do(ctx, nil)
}

// RevokeRoomUserAccess revokes the specified user's access to the room.
// 
// https://developers.docusign.com/docs/rooms-api/reference/rooms/rooms/revokeroomuseraccess
// 
// SDK Method Rooms::RevokeRoomUserAccess
func (s *Service) RevokeRoomUserAccess(roomID string, userID string, body *rooms.RoomUserRemovalDetail) *RevokeRoomUserAccessOp {
    return &RevokeRoomUserAccessOp{
        Credential: s.credential,
        Method:  "POST",
        Path: strings.Join([]string{"rooms",roomID,"users",userID,"revoke_access"},"/"),
        Payload: body,
        Accept: "application/json-patch+json, application/json, text/json, application/*+json",
        QueryOpts: make(url.Values),
        Version: esign.RoomsV2,
    }
}

// RevokeRoomUserAccessOp implements DocuSign API SDK Rooms::RevokeRoomUserAccess
type RevokeRoomUserAccessOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *RevokeRoomUserAccessOp) Do(ctx context.Context)  error {
    return ((*esign.Op)(op)).Do(ctx, nil)
}

// UpdatePicture updates the picture for a room.
// If media is an io.ReadCloser, Do() will close media.
// 
// https://developers.docusign.com/docs/rooms-api/reference/rooms/rooms/updatepicture
// 
// SDK Method Rooms::UpdatePicture
func (s *Service) UpdatePicture(roomID string, media io.Reader, mimeType string) *UpdatePictureOp {
    return &UpdatePictureOp{
        Credential: s.credential,
        Method:  "PUT",
        Path: strings.Join([]string{"rooms",roomID,"picture"},"/"),
        Payload: &esign.UploadFile{ Reader: media, ContentType: mimeType },
        Accept: "application/json",
        QueryOpts: make(url.Values),
        Version: esign.RoomsV2,
    }
}

// UpdatePictureOp implements DocuSign API SDK Rooms::UpdatePicture
type UpdatePictureOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *UpdatePictureOp) Do(ctx context.Context)  (*rooms.RoomPicture, error) {
    var res *rooms.RoomPicture
    return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// UpdateRoomFieldData updates a room's field data.
// 
// https://developers.docusign.com/docs/rooms-api/reference/rooms/rooms/updateroomfielddata
// 
// SDK Method Rooms::UpdateRoomFieldData
func (s *Service) UpdateRoomFieldData(roomID string, body *rooms.FieldDataForUpdate) *UpdateRoomFieldDataOp {
    return &UpdateRoomFieldDataOp{
        Credential: s.credential,
        Method:  "PUT",
        Path: strings.Join([]string{"rooms",roomID,"field_data"},"/"),
        Payload: body,
        Accept: "application/json",
        QueryOpts: make(url.Values),
        Version: esign.RoomsV2,
    }
}

// UpdateRoomFieldDataOp implements DocuSign API SDK Rooms::UpdateRoomFieldData
type UpdateRoomFieldDataOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *UpdateRoomFieldDataOp) Do(ctx context.Context)  (*rooms.FieldData, error) {
    var res *rooms.FieldData
    return res, ((*esign.Op)(op)).Do(ctx, &res)
}

