// Copyright 2019 James Cote
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by gen-esign; DO NOT EDIT. 
package click


type ClickwrapAgreementsResponse struct { 
    // User agreements from this datetime.
    BeginCreatedOn interface{} `json:"beginCreatedOn,omitempty"`
    // Number of pages remaining in the response.
    MinimumPagesRemaining int32 `json:"minimumPagesRemaining,omitempty"`
    // The number of the current page.
    Page int32 `json:"page,omitempty"`
    // The number of items per page.
    PageSize int32 `json:"pageSize,omitempty"`
    // An array of user agreements.
    UserAgreements []UserAgreementResponse `json:"userAgreements,omitempty"`
}

type ClickwrapDeleteResponse struct { 
    // The ID of the clickwrap.
    ClickwrapID string `json:"clickwrapId,omitempty"`
    // The name of the clickwrap.
    ClickwrapName string `json:"clickwrapName,omitempty"`
    // A message describing the result of deletion request. One of:
    // 
    // - `alreadyDeleted`: Clickwrap is already deleted.
    // - `deletionSuccess`: Successfully deleted the clickwrap.
    // - `deletionFailure`: Failed to delete the clickwrap.
    // - `cannotDelete`: Active clickwrap version cannot be deleted.
    DeletionMessage string `json:"deletionMessage,omitempty"`
    // **True** if the clickwrap was deleted successfully. **False** otherwise.
    DeletionSuccess bool `json:"deletionSuccess,omitempty"`
    // Clickwrap status. Possible values:
    // 
    // - `active`
    // - `inactive`
    // - `deleted`
    Status string `json:"status,omitempty"`
}

// ClickwrapRequest request body for working with clickwrap.
type ClickwrapRequest struct { 
    // The name of the clickwrap.
    ClickwrapName string `json:"clickwrapName,omitempty"`
    // Display settings for this clickwrap.
    DisplaySettings *DisplaySettings `json:"displaySettings,omitempty"`
    // An array of documents.
    Documents []Document `json:"documents,omitempty"`
    // Specifies whether `scheduledReacceptance` and `scheduledDate` should be cleared. May be one of:
    // 
    // - `"scheduledReacceptance"`
    // - `"scheduledDate"`
    // - `"scheduledReacceptance,scheduledDate"`
    // 
    FieldsToNull string `json:"fieldsToNull,omitempty"`
    // When **true**, the next version created is a major version. When **false** the next version created is minor.
    IsMajorVersion bool `json:"isMajorVersion,omitempty"`
    // 
    IsShared bool `json:"isShared,omitempty"`
    // Name of the clickwrap.
    Name string `json:"name,omitempty"`
    // When **true**, requires signers who have previously agreed to this
    // clickwrap to sign again. The version number is incremented.
    RequireReacceptance bool `json:"requireReacceptance,omitempty"`
    // The time and date when this clickwrap is activated.
    ScheduledDate interface{} `json:"scheduledDate,omitempty"`
    // Specifies the interval between reacceptances in days, weeks, months, or years.
    ScheduledReacceptance *ClickwrapScheduledReacceptance `json:"scheduledReacceptance,omitempty"`
    // Clickwrap status. Possible values:
    // 
    // - `active`
    // - `inactive`
    // - `deleted`
    Status interface{} `json:"status,omitempty"`
    // The user ID of current owner of the clickwrap.
    TransferFromUserID string `json:"transferFromUserId,omitempty"`
    // The user ID of the new owner of the clickwrap.
    TransferToUserID string `json:"transferToUserId,omitempty"`
}

type ClickwrapScheduledReacceptance struct { 
    // The time between recurrences specified in `recurrenceIntervalType` units.
    // 
    // The minimum and maximum values depend on `recurrenceIntervalType`:
    // 
    // - `days`: 1 - 365
    // - `weeks`: 1 - 52
    // - `months`: 1 - 12
    // - `years`: 1
    RecurrenceInterval int32 `json:"recurrenceInterval,omitempty"`
    // The units of the `recurrenceInterval`. Must be one of:
    // 
    // - `days`
    // - `weeks`
    // - `month`
    // - `years`
    // 
    RecurrenceIntervalType string `json:"recurrenceIntervalType,omitempty"`
    // The date when the recurrence interval starts.
    StartDateTime interface{} `json:"startDateTime,omitempty"`
}

type ClickwrapTransferRequest struct { 
    // ID of the user to transfer from.
    TransferFromUserID string `json:"transferFromUserId,omitempty"`
    // ID of the user to transfer to.
    TransferToUserID string `json:"transferToUserId,omitempty"`
}

type ClickwrapVersion struct { 
    // The unique version ID, a GUID, of this clickwrap version.
    ClickwrapVersionID string `json:"clickwrapVersionId,omitempty"`
    // The time that the clickwrap was created.
    CreatedTime interface{} `json:"createdTime,omitempty"`
    // The time that the clickwrap was last modified.
    LastModified interface{} `json:"lastModified,omitempty"`
    // The user ID of the last user who modified this clickwrap.
    LastModifiedBy string `json:"lastModifiedBy,omitempty"`
    // The user ID of the owner of this clickwrap.
    OwnerUserID string `json:"ownerUserId,omitempty"`
    // When **true**, requires signers who have previously agreed to this
    // clickwrap to sign again. The version number is incremented.
    RequireReacceptance bool `json:"requireReacceptance,omitempty"`
    // The time and date when this clickwrap is activated.
    ScheduledDate interface{} `json:"scheduledDate,omitempty"`
    // Specifies the interval between reacceptances in days, weeks, months, or years.
    ScheduledReacceptance *ClickwrapScheduledReacceptance `json:"scheduledReacceptance,omitempty"`
    // Clickwrap status. Possible values:
    // 
    // - `active`
    // - `inactive`
    // - `deleted`
    Status string `json:"status,omitempty"`
    // The ID of the version.
    VersionID string `json:"versionId,omitempty"`
    // Version of the clickwrap.
    VersionNumber string `json:"versionNumber,omitempty"`
}

type ClickwrapVersionDeleteResponse struct { 
    // The unique version ID, a GUID, of this clickwrap version.
    ClickwrapVersionID string `json:"clickwrapVersionId,omitempty"`
    // The time that the clickwrap was created.
    CreatedTime interface{} `json:"createdTime,omitempty"`
    // A message describing the result of deletion request. One of:
    // 
    // - `alreadyDeleted`: Clickwrap is already deleted.
    // - `deletionSuccess`: Successfully deleted the clickwrap.
    // - `deletionFailure`: Failed to delete the clickwrap.
    // - `cannotDelete`: Active clickwrap version cannot be deleted.
    DeletionMessage string `json:"deletionMessage,omitempty"`
    // **True** if the clickwrap was deleted successfully. **False** otherwise.
    DeletionSuccess bool `json:"deletionSuccess,omitempty"`
    // The time that the clickwrap was last modified.
    LastModified interface{} `json:"lastModified,omitempty"`
    // The user ID of the last user who modified this clickwrap.
    LastModifiedBy string `json:"lastModifiedBy,omitempty"`
    // The user ID of the owner of this clickwrap.
    OwnerUserID string `json:"ownerUserId,omitempty"`
    // When **true**, requires signers who have previously agreed to this
    // clickwrap to sign again. The version number is incremented.
    RequireReacceptance bool `json:"requireReacceptance,omitempty"`
    // The time and date when this clickwrap is activated.
    ScheduledDate interface{} `json:"scheduledDate,omitempty"`
    // Specifies the interval between reacceptances in days, weeks, months, or years.
    ScheduledReacceptance *ClickwrapScheduledReacceptance `json:"scheduledReacceptance,omitempty"`
    // Clickwrap status. Possible values:
    // 
    // - `active`
    // - `inactive`
    // - `deleted`
    Status string `json:"status,omitempty"`
    // The ID of the version.
    VersionID string `json:"versionId,omitempty"`
    // Version of the clickwrap.
    VersionNumber string `json:"versionNumber,omitempty"`
}

type ClickwrapVersionResponse struct { 
    // A GUID that identifies your account.
    // This value is automatically generated by
    // DocuSign for any account you create. Copy the
    // value from the **API Account ID** field in
    // the **API and Keys** page in
    // eSignature Settings.
    // 
    AccountID string `json:"accountId,omitempty"`
    // The ID of the clickwrap.
    ClickwrapID string `json:"clickwrapId,omitempty"`
    // The name of the clickwrap.
    ClickwrapName string `json:"clickwrapName,omitempty"`
    // The unique version ID, a GUID, of this clickwrap version.
    ClickwrapVersionID string `json:"clickwrapVersionId,omitempty"`
    // The time that the clickwrap was created.
    CreatedTime interface{} `json:"createdTime,omitempty"`
    // Display settings for a clickwrap.
    DisplaySettings *DisplaySettings `json:"displaySettings,omitempty"`
    // An array of documents.
    Documents []Document `json:"documents,omitempty"`
    // The time that the clickwrap was last modified.
    LastModified interface{} `json:"lastModified,omitempty"`
    // The user ID of the last user who modified this clickwrap.
    LastModifiedBy string `json:"lastModifiedBy,omitempty"`
    // The user ID of the owner of this clickwrap.
    OwnerUserID string `json:"ownerUserId,omitempty"`
    // When **true**, requires signers who have previously agreed to this
    // clickwrap to sign again. The version number is incremented.
    RequireReacceptance bool `json:"requireReacceptance,omitempty"`
    // The time and date when this clickwrap is activated.
    ScheduledDate interface{} `json:"scheduledDate,omitempty"`
    // Specifies the interval between reacceptances in days, weeks, months, or years.
    ScheduledReacceptance *ClickwrapScheduledReacceptance `json:"scheduledReacceptance,omitempty"`
    // Clickwrap status. Possible values:
    // 
    // - `active`
    // - `inactive`
    // - `deleted`
    Status string `json:"status,omitempty"`
    // The ID of the version.
    VersionID string `json:"versionId,omitempty"`
    // Version of the clickwrap.
    VersionNumber string `json:"versionNumber,omitempty"`
}

type ClickwrapVersionSummaryResponse struct { 
    // A GUID that identifies your account.
    // This value is automatically generated by
    // DocuSign for any account you create. Copy the
    // value from the **API Account ID** field in
    // the **API and Keys** page in
    // eSignature Settings.
    // 
    AccountID string `json:"accountId,omitempty"`
    // The ID of the clickwrap.
    ClickwrapID string `json:"clickwrapId,omitempty"`
    // The name of the clickwrap.
    ClickwrapName string `json:"clickwrapName,omitempty"`
    // The unique version ID, a GUID, of this clickwrap version.
    ClickwrapVersionID string `json:"clickwrapVersionId,omitempty"`
    // The time that the clickwrap was created.
    CreatedTime interface{} `json:"createdTime,omitempty"`
    // The time that the clickwrap was last modified.
    LastModified interface{} `json:"lastModified,omitempty"`
    // The user ID of the last user who modified this clickwrap.
    LastModifiedBy string `json:"lastModifiedBy,omitempty"`
    // The user ID of the owner of this clickwrap.
    OwnerUserID string `json:"ownerUserId,omitempty"`
    // When **true**, requires signers who have previously agreed to this
    // clickwrap to sign again. The version number is incremented.
    RequireReacceptance bool `json:"requireReacceptance,omitempty"`
    // The time and date when this clickwrap is activated.
    ScheduledDate interface{} `json:"scheduledDate,omitempty"`
    // Specifies the interval between reacceptances in days, weeks, months, or years.
    ScheduledReacceptance *ClickwrapScheduledReacceptance `json:"scheduledReacceptance,omitempty"`
    // Clickwrap status. Possible values:
    // 
    // - `active`
    // - `inactive`
    // - `deleted`
    Status string `json:"status,omitempty"`
    // The ID of the version.
    VersionID string `json:"versionId,omitempty"`
    // Version of the clickwrap.
    VersionNumber string `json:"versionNumber,omitempty"`
}

type ClickwrapVersionsDeleteResponse struct { 
    // The ID of the clickwrap.
    ClickwrapID string `json:"clickwrapId,omitempty"`
    // The name of the clickwrap.
    ClickwrapName string `json:"clickwrapName,omitempty"`
    // An array delete responses.
    Versions []ClickwrapVersionDeleteResponse `json:"versions,omitempty"`
}

type ClickwrapVersionsPagedResponse struct { 
    // A GUID that identifies your account.
    // This value is automatically generated by
    // DocuSign for any account you create. Copy the
    // value from the **API Account ID** field in
    // the **API and Keys** page in
    // eSignature Settings.
    // 
    AccountID string `json:"accountId,omitempty"`
    // The ID of the clickwrap.
    ClickwrapID string `json:"clickwrapId,omitempty"`
    // The name of the clickwrap.
    ClickwrapName string `json:"clickwrapName,omitempty"`
    // Number of pages remaining in the response.
    MinimumPagesRemaining int32 `json:"minimumPagesRemaining,omitempty"`
    // The number of the current page.
    Page int32 `json:"page,omitempty"`
    // The number of items per page.
    PageSize int32 `json:"pageSize,omitempty"`
    // An array of clickwrap versions.
    Versions []ClickwrapVersion `json:"versions,omitempty"`
}

type ClickwrapVersionsResponse struct { 
    // An array of `clickwrapVersionSummaryResponse` objects.
    Clickwraps []ClickwrapVersionSummaryResponse `json:"clickwraps,omitempty"`
    // Number of pages remaining in the response.
    MinimumPagesRemaining int32 `json:"minimumPagesRemaining,omitempty"`
    // The number of the current page.
    Page int32 `json:"page,omitempty"`
    // The number of items per page.
    PageSize int32 `json:"pageSize,omitempty"`
}

type ClickwrapsDeleteResponse struct { 
    // 
    Clickwraps []ClickwrapDeleteResponse `json:"clickwraps,omitempty"`
}

// DisplaySettings information about how an agreement is displayed.
type DisplaySettings struct { 
    // Position of the Accept button in the agreement. One of 
    // 
    // - `right`
    // - `left`
    // 
    ActionButtonAlignment string `json:"actionButtonAlignment,omitempty"`
    // When **true**, this agreement can be be used in client-only integrations
    AllowClientOnly bool `json:"allowClientOnly,omitempty"`
    // Hosts that can host the clickwrap.
    // 
    // It is an error if the clickwrap didn't come from one of these hosts.
    // 
    AllowedHosts []string `json:"allowedHosts,omitempty"`
    // The signing brand ID.
    BrandID string `json:"brandId,omitempty"`
    // Text on the agree button.
    ConsentButtonText string `json:"consentButtonText,omitempty"`
    // The text on agree button.
    ConsentText string `json:"consentText,omitempty"`
    // The text on the decline button.
    DeclineButtonText string `json:"declineButtonText,omitempty"`
    // The display name of the user agreement.
    DisplayName string `json:"displayName,omitempty"`
    // Display type: link or document
    DocumentDisplay string `json:"documentDisplay,omitempty"`
    // **True** if the agreement is downloadable.
    Downloadable bool `json:"downloadable,omitempty"`
    // Display format: inline or modal.
    Format string `json:"format,omitempty"`
    // **True** if the agreement has a decline checkbox.
    HasDeclineButton bool `json:"hasDeclineButton,omitempty"`
    // The host origin.
    // 
    HostOrigin string `json:"hostOrigin,omitempty"`
    // **True** if the user needs to scroll to the end of the document.
    // 
    MustRead bool `json:"mustRead,omitempty"`
    // **True** if the user must view the document.
    // 
    MustView bool `json:"mustView,omitempty"`
    // When **true**, this agreement records decline actions.
    RecordDeclineResponses bool `json:"recordDeclineResponses,omitempty"`
    // **True** if accept is required.
    // 
    RequireAccept bool `json:"requireAccept,omitempty"`
    // **True** if send to email is applicable.
    // 
    SendToEmail bool `json:"sendToEmail,omitempty"`
}

// Document information about a document.
type Document struct { 
    // The base64-encoded contents of the document.
    DocumentBase64 string `json:"documentBase64,omitempty"`
    // The HTML representation of the document.
    DocumentHTML string `json:"documentHtml,omitempty"`
    // The name of the document.
    DocumentName string `json:"documentName,omitempty"`
    // The file extension of the document.
    FileExtension string `json:"fileExtension,omitempty"`
    // The order of document layout.
    Order int32 `json:"order,omitempty"`
}

// ErrorDetails error details.
type ErrorDetails struct { 
    // The error code.
    ErrorCode string `json:"errorCode,omitempty"`
    // The error message.
    Message string `json:"message,omitempty"`
}

type ServiceInformation struct { 
    // 
    BuildBranch string `json:"buildBranch,omitempty"`
    // 
    BuildBranchDeployedDateTime string `json:"buildBranchDeployedDateTime,omitempty"`
    // 
    BuildSHA string `json:"buildSHA,omitempty"`
    // The internal build version information.
    BuildVersion string `json:"buildVersion,omitempty"`
    // An array of URLs (strings) of related sites.
    LinkedSites []string `json:"linkedSites,omitempty"`
    // An array of `serviceVersion` objects.
    ServiceVersions []ServiceVersion `json:"serviceVersions,omitempty"`
}

type ServiceVersion struct { 
    // The human-readable semver version string.
    Version string `json:"version,omitempty"`
    // The URL where this version of the API can be found.
    VersionURL string `json:"versionUrl,omitempty"`
}

type UserAgreementRequest struct { 
    // The user ID of the client.
    ClientUserID string `json:"clientUserId,omitempty"`
    // The host origin.
    // 
    HostOrigin string `json:"hostOrigin,omitempty"`
    // A customer-defined string you can use in requests. This string will appear in the corresponding response.
    Metadata string `json:"metadata,omitempty"`
}

type UserAgreementResponse struct { 
    // A GUID that identifies your account.
    // This value is automatically generated by
    // DocuSign for any account you create. Copy the
    // value from the **API Account ID** field in
    // the **API and Keys** page in
    // eSignature Settings.
    // 
    AccountID string `json:"accountId,omitempty"`
    // Date that the client last completed the agreement.
    // 
    // This property is null if `agreementUrl` is not null and `status` is not  `agreed`.
    AgreedOn interface{} `json:"agreedOn,omitempty"`
    // The agreement ID.
    AgreementID string `json:"agreementId,omitempty"`
    // When not null, an agreement is required for user specified by  `clientUserId`.
    // 
    // When missing the user specified by `clientUserId`
    // has already agreed and does not require a new acceptance.
    // 
    // Use this URL to render the agreement in a web page.
    // 
    // <!--
    // or redirected to when providing redirect_url as a query paramter.
    // -->
    // 
    AgreementURL string `json:"agreementUrl,omitempty"`
    // The ID of the clickwrap.
    ClickwrapID string `json:"clickwrapId,omitempty"`
    // The user ID of the client.
    ClientUserID string `json:"clientUserId,omitempty"`
    // The customer-branded HTML with the Electronic Record and Signature Disclosure information
    ConsumerDisclosureHTML string `json:"consumerDisclosureHtml,omitempty"`
    // The date when the clickwrap was created. May be null.
    CreatedOn interface{} `json:"createdOn,omitempty"`
    // The date when the user declined the most recent required agreement.
    // 
    // This property is valid only when `status` is `declined`. Otherwise it is null.
    DeclinedOn interface{} `json:"declinedOn,omitempty"`
    // An array of documents.
    Documents []Document `json:"documents,omitempty"`
    // A customer-defined string you can use in requests. This string will appear in the corresponding response.
    Metadata string `json:"metadata,omitempty"`
    // The display settings for this agreement.
    Settings *DisplaySettings `json:"settings,omitempty"`
    // User agreement status. One of:
    // 
    // - `created`
    // - `agreed`
    // - `declined`
    Status string `json:"status,omitempty"`
    // The human-readable semver version string.
    Version string `json:"version,omitempty"`
    // The ID of the version.
    VersionID string `json:"versionId,omitempty"`
    // Version of the clickwrap.
    VersionNumber int32 `json:"versionNumber,omitempty"`
}

