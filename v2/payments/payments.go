// Copyright 2019 James Cote
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by gen-esign; DO NOT EDIT.

// Package payments implements the DocuSign SDK
// category Payments.
// 
// This category includes resources for managing payment gateways. Payment information is added to envelopes via methods in the Envelopes category.
//
// Service Api documentation may be found at:
// https://developers.docusign.com/docs/esign-api/v2/reference/Payments
// Usage example:
//
//   import (
//       "github.com/jfcote87/esign"
//       "github.com/jfcote87/esign/v2/model"
//   ) 
//   ...
//   paymentsService := payments.New(esignCredential)
package payments // import "github.com/jfcote87/esignv2/payments"

import (
    "context"
    "net/url"
    
    "github.com/jfcote87/esign"
    "github.com/jfcote87/esign/v2/model"
)

// Service implements DocuSign Payments Category API operations
type Service struct {
    credential esign.Credential 
}

// New initializes a payments service using cred to authorize ops.
func New(cred esign.Credential) *Service {
    return &Service{credential: cred}
}

// GatewayAccountsList list payment gateway account information
// 
// https://developers.docusign.com/docs/esign-api/v2/reference/payments/paymentgatewayaccounts/list
// 
// SDK Method Payments::getAllPaymentGatewayAccounts
func (s *Service) GatewayAccountsList() *GatewayAccountsListOp {
    return &GatewayAccountsListOp{
        Credential: s.credential,
        Method:  "GET",
        Path: "payment_gateway_accounts",
        Accept: "application/json",
        QueryOpts: make(url.Values),
        Version: esign.APIv2,
    }
}

// GatewayAccountsListOp implements DocuSign API SDK Payments::getAllPaymentGatewayAccounts
type GatewayAccountsListOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *GatewayAccountsListOp) Do(ctx context.Context)  (*model.PaymentGatewayAccountsInfo, error) {
    var res *model.PaymentGatewayAccountsInfo
    return res, ((*esign.Op)(op)).Do(ctx, &res)
}

