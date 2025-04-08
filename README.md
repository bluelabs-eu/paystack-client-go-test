# Go API client for openapi

The OpenAPI specification of the Paystack API that merchants and developers can harness to build financial solutions in Africa.

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 1.0.0
- Generator version: 7.12.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://paystack.com/docs](https://paystack.com/docs)

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import openapi "github.com/GIT_USER_ID/GIT_REPO_ID"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `openapi.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), openapi.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `openapi.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), openapi.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `openapi.ContextOperationServerIndices` and `openapi.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), openapi.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), openapi.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.paystack.co*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*BalanceAPI* | [**BalanceFetch**](docs/BalanceAPI.md#balancefetch) | **Get** /balance | Fetch Balance
*BalanceAPI* | [**BalanceLedger**](docs/BalanceAPI.md#balanceledger) | **Get** /balance/ledger | Balance Ledger
*BulkChargeAPI* | [**BulkChargeCharges**](docs/BulkChargeAPI.md#bulkchargecharges) | **Get** /bulkcharge/{code}/charges | Fetch Charges in a Batch
*BulkChargeAPI* | [**BulkChargeFetch**](docs/BulkChargeAPI.md#bulkchargefetch) | **Get** /bulkcharge/{code} | Fetch Bulk Charge Batch
*BulkChargeAPI* | [**BulkChargeInitiate**](docs/BulkChargeAPI.md#bulkchargeinitiate) | **Post** /bulkcharge | Initiate Bulk Charge
*BulkChargeAPI* | [**BulkChargeList**](docs/BulkChargeAPI.md#bulkchargelist) | **Get** /bulkcharge | List Bulk Charge Batches
*BulkChargeAPI* | [**BulkChargePause**](docs/BulkChargeAPI.md#bulkchargepause) | **Get** /bulkcharge/pause/{code} | Pause Bulk Charge Batch
*BulkChargeAPI* | [**BulkChargeResume**](docs/BulkChargeAPI.md#bulkchargeresume) | **Get** /bulkcharge/resume/{code} | Resume Bulk Charge Batch
*ChargeAPI* | [**ChargeCheck**](docs/ChargeAPI.md#chargecheck) | **Get** /charge/{reference} | Check pending charge
*ChargeAPI* | [**ChargeCreate**](docs/ChargeAPI.md#chargecreate) | **Post** /charge | Create Charge
*ChargeAPI* | [**ChargeSubmitAddress**](docs/ChargeAPI.md#chargesubmitaddress) | **Post** /charge/submit_address | Submit Address
*ChargeAPI* | [**ChargeSubmitBirthday**](docs/ChargeAPI.md#chargesubmitbirthday) | **Post** /charge/submit_birthday | Submit Birthday
*ChargeAPI* | [**ChargeSubmitOtp**](docs/ChargeAPI.md#chargesubmitotp) | **Post** /charge/submit_otp | Submit OTP
*ChargeAPI* | [**ChargeSubmitPhone**](docs/ChargeAPI.md#chargesubmitphone) | **Post** /charge/submit_phone | Submit Phone
*ChargeAPI* | [**ChargeSubmitPin**](docs/ChargeAPI.md#chargesubmitpin) | **Post** /charge/submit_pin | Submit PIN
*CustomerAPI* | [**CustomerCreate**](docs/CustomerAPI.md#customercreate) | **Post** /customer | Create Customer
*CustomerAPI* | [**CustomerDeactivateAuthorization**](docs/CustomerAPI.md#customerdeactivateauthorization) | **Post** /customer/deactivate_authorization | Deactivate Authorization
*CustomerAPI* | [**CustomerFetch**](docs/CustomerAPI.md#customerfetch) | **Get** /customer/{code} | Fetch Customer
*CustomerAPI* | [**CustomerList**](docs/CustomerAPI.md#customerlist) | **Get** /customer | List Customers
*CustomerAPI* | [**CustomerRiskAction**](docs/CustomerAPI.md#customerriskaction) | **Post** /customer/set_risk_action | White/blacklist Customer
*CustomerAPI* | [**CustomerUpdate**](docs/CustomerAPI.md#customerupdate) | **Put** /customer/{code} | Update Customer
*CustomerAPI* | [**CustomerValidate**](docs/CustomerAPI.md#customervalidate) | **Post** /customer/{code}/identification | Validate Customer
*DedicatedVirtualAccountAPI* | [**DedicatedAccountAddSplit**](docs/DedicatedVirtualAccountAPI.md#dedicatedaccountaddsplit) | **Post** /dedicated_account/split | Split Dedicated Account Transaction
*DedicatedVirtualAccountAPI* | [**DedicatedAccountAvailableProviders**](docs/DedicatedVirtualAccountAPI.md#dedicatedaccountavailableproviders) | **Get** /dedicated_account/available_providers | Fetch Bank Providers
*DedicatedVirtualAccountAPI* | [**DedicatedAccountCreate**](docs/DedicatedVirtualAccountAPI.md#dedicatedaccountcreate) | **Post** /dedicated_account | Create Dedicated Account
*DedicatedVirtualAccountAPI* | [**DedicatedAccountDeactivate**](docs/DedicatedVirtualAccountAPI.md#dedicatedaccountdeactivate) | **Delete** /dedicated_account/{account_id} | Deactivate Dedicated Account
*DedicatedVirtualAccountAPI* | [**DedicatedAccountFetch**](docs/DedicatedVirtualAccountAPI.md#dedicatedaccountfetch) | **Get** /dedicated_account/{account_id} | Fetch Dedicated Account
*DedicatedVirtualAccountAPI* | [**DedicatedAccountList**](docs/DedicatedVirtualAccountAPI.md#dedicatedaccountlist) | **Get** /dedicated_account | List Dedicated Accounts
*DedicatedVirtualAccountAPI* | [**DedicatedAccountRemoveSplit**](docs/DedicatedVirtualAccountAPI.md#dedicatedaccountremovesplit) | **Delete** /dedicated_account/split | Remove Split from Dedicated Account
*DisputeAPI* | [**DisputeDownload**](docs/DisputeAPI.md#disputedownload) | **Get** /dispute/export | Export Disputes
*DisputeAPI* | [**DisputeEvidence**](docs/DisputeAPI.md#disputeevidence) | **Post** /dispute/{id}/evidence | Add Evidence
*DisputeAPI* | [**DisputeFetch**](docs/DisputeAPI.md#disputefetch) | **Get** /dispute/{id} | Fetch Dispute
*DisputeAPI* | [**DisputeList**](docs/DisputeAPI.md#disputelist) | **Get** /dispute | List Disputes
*DisputeAPI* | [**DisputeResolve**](docs/DisputeAPI.md#disputeresolve) | **Put** /dispute/{id}/resolve | Resolve a Dispute
*DisputeAPI* | [**DisputeTransaction**](docs/DisputeAPI.md#disputetransaction) | **Get** /dispute/transaction/{id} | List Transaction Disputes
*DisputeAPI* | [**DisputeUpdate**](docs/DisputeAPI.md#disputeupdate) | **Put** /dispute/{id} | Update Dispute
*DisputeAPI* | [**DisputeUploadUrl**](docs/DisputeAPI.md#disputeuploadurl) | **Get** /dispute/{id}/upload_url | Get Upload URL
*IntegrationAPI* | [**IntegrationFetchPaymentSessionTimeout**](docs/IntegrationAPI.md#integrationfetchpaymentsessiontimeout) | **Get** /integration/payment_session_timeout | Fetch Payment Session Timeout
*IntegrationAPI* | [**IntegrationUpdatePaymentSessionTimeout**](docs/IntegrationAPI.md#integrationupdatepaymentsessiontimeout) | **Put** /integration/payment_session_timeout | Update Payment Session Timeout
*PageAPI* | [**PageAddProducts**](docs/PageAPI.md#pageaddproducts) | **Post** /page/{id}/product | Add Products
*PageAPI* | [**PageCheckSlugAvailability**](docs/PageAPI.md#pagecheckslugavailability) | **Get** /page/check_slug_availability/{slug} | Check Slug Availability
*PageAPI* | [**PageCreate**](docs/PageAPI.md#pagecreate) | **Post** /page | Create Page
*PageAPI* | [**PageFetch**](docs/PageAPI.md#pagefetch) | **Get** /page/{id} | Fetch Page
*PageAPI* | [**PageList**](docs/PageAPI.md#pagelist) | **Get** /page | List Pages
*PageAPI* | [**PageUpdate**](docs/PageAPI.md#pageupdate) | **Put** /page/{id} | Update Page
*PaymentRequestAPI* | [**PaymentRequestArchive**](docs/PaymentRequestAPI.md#paymentrequestarchive) | **Post** /paymentrequest/archive/{id} | Archive Payment Request
*PaymentRequestAPI* | [**PaymentRequestCreate**](docs/PaymentRequestAPI.md#paymentrequestcreate) | **Post** /paymentrequest | Create Payment Request
*PaymentRequestAPI* | [**PaymentRequestFetch**](docs/PaymentRequestAPI.md#paymentrequestfetch) | **Get** /paymentrequest/{id} | Fetch Payment Request
*PaymentRequestAPI* | [**PaymentRequestFinalize**](docs/PaymentRequestAPI.md#paymentrequestfinalize) | **Post** /paymentrequest/finalize/{id} | Finalize Payment Request
*PaymentRequestAPI* | [**PaymentRequestList**](docs/PaymentRequestAPI.md#paymentrequestlist) | **Get** /paymentrequest | List Payment Request
*PaymentRequestAPI* | [**PaymentRequestNotify**](docs/PaymentRequestAPI.md#paymentrequestnotify) | **Post** /paymentrequest/notify/{id} | Send Notification
*PaymentRequestAPI* | [**PaymentRequestTotals**](docs/PaymentRequestAPI.md#paymentrequesttotals) | **Get** /paymentrequest/totals | Payment Request Total
*PaymentRequestAPI* | [**PaymentRequestUpdate**](docs/PaymentRequestAPI.md#paymentrequestupdate) | **Put** /paymentrequest/{id} | Update Payment Request
*PaymentRequestAPI* | [**PaymentRequestVerify**](docs/PaymentRequestAPI.md#paymentrequestverify) | **Get** /paymentrequest/verify/{id} | Verify Payment Request
*PlanAPI* | [**PlanCreate**](docs/PlanAPI.md#plancreate) | **Post** /plan | Create Plan
*PlanAPI* | [**PlanFetch**](docs/PlanAPI.md#planfetch) | **Get** /plan/{code} | Fetch Plan
*PlanAPI* | [**PlanList**](docs/PlanAPI.md#planlist) | **Get** /plan | List Plans
*PlanAPI* | [**PlanUpdate**](docs/PlanAPI.md#planupdate) | **Put** /plan/{code} | Update Plan
*ProductAPI* | [**ProductCreate**](docs/ProductAPI.md#productcreate) | **Post** /product | Create Product
*ProductAPI* | [**ProductDelete**](docs/ProductAPI.md#productdelete) | **Delete** /product/{id} | Delete Product
*ProductAPI* | [**ProductFetch**](docs/ProductAPI.md#productfetch) | **Get** /product/{id} | Fetch Product
*ProductAPI* | [**ProductList**](docs/ProductAPI.md#productlist) | **Get** /product | List Products
*ProductAPI* | [**ProductUpdate**](docs/ProductAPI.md#productupdate) | **Put** /product/{id} | Update product
*RefundAPI* | [**RefundCreate**](docs/RefundAPI.md#refundcreate) | **Post** /refund | Create Refund
*RefundAPI* | [**RefundFetch**](docs/RefundAPI.md#refundfetch) | **Get** /refund/{id} | Fetch Refund
*RefundAPI* | [**RefundList**](docs/RefundAPI.md#refundlist) | **Get** /refund | List Refunds
*SettlementAPI* | [**SettlementsFetch**](docs/SettlementAPI.md#settlementsfetch) | **Get** /settlement | Fetch Settlements
*SettlementAPI* | [**SettlementsTransaction**](docs/SettlementAPI.md#settlementstransaction) | **Get** /settlement/{id}/transaction | Settlement Transactions
*SplitAPI* | [**SplitAddSubaccount**](docs/SplitAPI.md#splitaddsubaccount) | **Post** /split/{id}/subaccount/add | Add Subaccount to Split
*SplitAPI* | [**SplitCreate**](docs/SplitAPI.md#splitcreate) | **Post** /split | Create Split
*SplitAPI* | [**SplitFetch**](docs/SplitAPI.md#splitfetch) | **Get** /split/{id} | Fetch Split
*SplitAPI* | [**SplitList**](docs/SplitAPI.md#splitlist) | **Get** /split | List/Search Splits
*SplitAPI* | [**SplitRemoveSubaccount**](docs/SplitAPI.md#splitremovesubaccount) | **Post** /split/{id}/subaccount/remove | Remove Subaccount from split
*SplitAPI* | [**SplitUpdate**](docs/SplitAPI.md#splitupdate) | **Put** /split/{id} | Update Split
*SubaccountAPI* | [**SubaccountCreate**](docs/SubaccountAPI.md#subaccountcreate) | **Post** /subaccount | Create Subaccount
*SubaccountAPI* | [**SubaccountFetch**](docs/SubaccountAPI.md#subaccountfetch) | **Get** /subaccount/{code} | Fetch Subaccount
*SubaccountAPI* | [**SubaccountList**](docs/SubaccountAPI.md#subaccountlist) | **Get** /subaccount | List Subaccounts
*SubaccountAPI* | [**SubaccountUpdate**](docs/SubaccountAPI.md#subaccountupdate) | **Put** /subaccount/{code} | Update Subaccount
*SubscriptionAPI* | [**SubscriptionCreate**](docs/SubscriptionAPI.md#subscriptioncreate) | **Post** /subscription | Create Subscription
*SubscriptionAPI* | [**SubscriptionDisable**](docs/SubscriptionAPI.md#subscriptiondisable) | **Post** /subscription/disable | Disable Subscription
*SubscriptionAPI* | [**SubscriptionEnable**](docs/SubscriptionAPI.md#subscriptionenable) | **Post** /subscription/enable | Enable Subscription
*SubscriptionAPI* | [**SubscriptionFetch**](docs/SubscriptionAPI.md#subscriptionfetch) | **Get** /subscription/{code} | Fetch Subscription
*SubscriptionAPI* | [**SubscriptionList**](docs/SubscriptionAPI.md#subscriptionlist) | **Get** /subscription | List Subscriptions
*SubscriptionAPI* | [**SubscriptionManageEmail**](docs/SubscriptionAPI.md#subscriptionmanageemail) | **Post** /subscription/{code}/manage/email | Send Update Subscription Link
*SubscriptionAPI* | [**SubscriptionManageLink**](docs/SubscriptionAPI.md#subscriptionmanagelink) | **Get** /subscription/{code}/manage/link | Generate Update Subscription Link
*TransactionAPI* | [**TransactionChargeAuthorization**](docs/TransactionAPI.md#transactionchargeauthorization) | **Post** /transaction/charge_authorization | Charge Authorization
*TransactionAPI* | [**TransactionCheckAuthorization**](docs/TransactionAPI.md#transactioncheckauthorization) | **Post** /transaction/check_authorization | Check Authorization
*TransactionAPI* | [**TransactionDownload**](docs/TransactionAPI.md#transactiondownload) | **Get** /transaction/export | Export Transactions
*TransactionAPI* | [**TransactionEvent**](docs/TransactionAPI.md#transactionevent) | **Get** /transaction/{id}/event | Get Transaction Event
*TransactionAPI* | [**TransactionFetch**](docs/TransactionAPI.md#transactionfetch) | **Get** /transaction/{id} | Fetch Transaction
*TransactionAPI* | [**TransactionInitialize**](docs/TransactionAPI.md#transactioninitialize) | **Post** /transaction/initialize | Initialize Transaction
*TransactionAPI* | [**TransactionList**](docs/TransactionAPI.md#transactionlist) | **Get** /transaction | List Transactions
*TransactionAPI* | [**TransactionPartialDebit**](docs/TransactionAPI.md#transactionpartialdebit) | **Post** /transaction/partial_debit | Partial Debit
*TransactionAPI* | [**TransactionSession**](docs/TransactionAPI.md#transactionsession) | **Get** /transaction/{id}/session | Get Transaction Session
*TransactionAPI* | [**TransactionTimeline**](docs/TransactionAPI.md#transactiontimeline) | **Get** /transaction/timeline/{id_or_reference} | Fetch Transaction Timeline
*TransactionAPI* | [**TransactionTotals**](docs/TransactionAPI.md#transactiontotals) | **Get** /transaction/totals | Transaction Totals
*TransactionAPI* | [**TransactionVerify**](docs/TransactionAPI.md#transactionverify) | **Get** /transaction/verify/{reference} | Verify Transaction
*TransferAPI* | [**TransferBulk**](docs/TransferAPI.md#transferbulk) | **Post** /transfer/bulk | Initiate Bulk Transfer
*TransferAPI* | [**TransferDisableOtp**](docs/TransferAPI.md#transferdisableotp) | **Post** /transfer/disable_otp | Disable OTP requirement for Transfers
*TransferAPI* | [**TransferDisableOtpFinalize**](docs/TransferAPI.md#transferdisableotpfinalize) | **Post** /transfer/disable_otp_finalize | Finalize Disabling of OTP requirement for Transfers
*TransferAPI* | [**TransferDownload**](docs/TransferAPI.md#transferdownload) | **Get** /transfer/export | Export Transfers
*TransferAPI* | [**TransferEnableOtp**](docs/TransferAPI.md#transferenableotp) | **Post** /transfer/enable_otp | Enable OTP requirement for Transfers
*TransferAPI* | [**TransferFetch**](docs/TransferAPI.md#transferfetch) | **Get** /transfer/{code} | Fetch Transfer
*TransferAPI* | [**TransferFinalize**](docs/TransferAPI.md#transferfinalize) | **Post** /transfer/finalize_transfer | Finalize Transfer
*TransferAPI* | [**TransferInitiate**](docs/TransferAPI.md#transferinitiate) | **Post** /transfer | Initiate Transfer
*TransferAPI* | [**TransferList**](docs/TransferAPI.md#transferlist) | **Get** /transfer | List Transfers
*TransferAPI* | [**TransferResendOtp**](docs/TransferAPI.md#transferresendotp) | **Post** /transfer/resend_otp | Resend OTP for Transfer
*TransferAPI* | [**TransferVerify**](docs/TransferAPI.md#transferverify) | **Get** /transfer/verify/{reference} | Verify Transfer
*TransferRecipientAPI* | [**TransferrecipientBulk**](docs/TransferRecipientAPI.md#transferrecipientbulk) | **Post** /transferrecipient/bulk | Bulk Create Transfer Recipient
*TransferRecipientAPI* | [**TransferrecipientCodeDelete**](docs/TransferRecipientAPI.md#transferrecipientcodedelete) | **Delete** /transferrecipient/{code} | Delete Transfer Recipient
*TransferRecipientAPI* | [**TransferrecipientCodePut**](docs/TransferRecipientAPI.md#transferrecipientcodeput) | **Put** /transferrecipient/{code} | Update Transfer recipient
*TransferRecipientAPI* | [**TransferrecipientCreate**](docs/TransferRecipientAPI.md#transferrecipientcreate) | **Post** /transferrecipient | Create Transfer Recipient
*TransferRecipientAPI* | [**TransferrecipientFetch**](docs/TransferRecipientAPI.md#transferrecipientfetch) | **Get** /transferrecipient/{code} | Fetch Transfer recipient
*TransferRecipientAPI* | [**TransferrecipientList**](docs/TransferRecipientAPI.md#transferrecipientlist) | **Get** /transferrecipient | List Transfer Recipients
*VerificationAPI* | [**VerificationAvs**](docs/VerificationAPI.md#verificationavs) | **Get** /address_verification/states | List States (AVS)
*VerificationAPI* | [**VerificationFetchBanks**](docs/VerificationAPI.md#verificationfetchbanks) | **Get** /bank | Fetch Banks
*VerificationAPI* | [**VerificationListCountries**](docs/VerificationAPI.md#verificationlistcountries) | **Get** /country | List Countries
*VerificationAPI* | [**VerificationResolveAccountNumber**](docs/VerificationAPI.md#verificationresolveaccountnumber) | **Get** /bank/resolve | Resolve Account Number
*VerificationAPI* | [**VerificationResolveCardBin**](docs/VerificationAPI.md#verificationresolvecardbin) | **Get** /decision/bin/{bin} | Resolve Card BIN


## Documentation For Models

 - [Accepted](docs/Accepted.md)
 - [Bank](docs/Bank.md)
 - [BulkChargeInitiate](docs/BulkChargeInitiate.md)
 - [BulkChargeInitiateRequestInner1](docs/BulkChargeInitiateRequestInner1.md)
 - [ChargeCreateRequest](docs/ChargeCreateRequest.md)
 - [ChargeSubmitAddress](docs/ChargeSubmitAddress.md)
 - [ChargeSubmitBirthday](docs/ChargeSubmitBirthday.md)
 - [ChargeSubmitOTP](docs/ChargeSubmitOTP.md)
 - [ChargeSubmitPhone](docs/ChargeSubmitPhone.md)
 - [ChargeSubmitPin](docs/ChargeSubmitPin.md)
 - [CustomerCreate](docs/CustomerCreate.md)
 - [CustomerDeactivateAuthorization](docs/CustomerDeactivateAuthorization.md)
 - [CustomerRiskAction](docs/CustomerRiskAction.md)
 - [CustomerUpdate](docs/CustomerUpdate.md)
 - [CustomerValidate](docs/CustomerValidate.md)
 - [CustomerValidation](docs/CustomerValidation.md)
 - [DedicatedVirtualAccountCreate](docs/DedicatedVirtualAccountCreate.md)
 - [DedicatedVirtualAccountSplit](docs/DedicatedVirtualAccountSplit.md)
 - [DisputeEvidence](docs/DisputeEvidence.md)
 - [DisputeResolve](docs/DisputeResolve.md)
 - [DisputeUpdate](docs/DisputeUpdate.md)
 - [EFT](docs/EFT.md)
 - [Error](docs/Error.md)
 - [MobileMoney](docs/MobileMoney.md)
 - [PageCreate](docs/PageCreate.md)
 - [PageProduct](docs/PageProduct.md)
 - [PageUpdate](docs/PageUpdate.md)
 - [PaymentRequestCreate](docs/PaymentRequestCreate.md)
 - [PaymentRequestUpdate](docs/PaymentRequestUpdate.md)
 - [PlanCreate](docs/PlanCreate.md)
 - [PlanUpdate](docs/PlanUpdate.md)
 - [ProductCreate](docs/ProductCreate.md)
 - [ProductUpdate](docs/ProductUpdate.md)
 - [RefundCreate](docs/RefundCreate.md)
 - [Response](docs/Response.md)
 - [SplitCreate](docs/SplitCreate.md)
 - [SplitSubaccounts](docs/SplitSubaccounts.md)
 - [SplitUpdate](docs/SplitUpdate.md)
 - [SubaccountCreate](docs/SubaccountCreate.md)
 - [SubaccountUpdate](docs/SubaccountUpdate.md)
 - [SubscriptionCreate](docs/SubscriptionCreate.md)
 - [SubscriptionToggle](docs/SubscriptionToggle.md)
 - [TransactionChargeAuthorization](docs/TransactionChargeAuthorization.md)
 - [TransactionCheckAuthorization](docs/TransactionCheckAuthorization.md)
 - [TransactionInitialize](docs/TransactionInitialize.md)
 - [TransactionPartialDebit](docs/TransactionPartialDebit.md)
 - [TransferBulk](docs/TransferBulk.md)
 - [TransferFinalize](docs/TransferFinalize.md)
 - [TransferFinalizeDisableOTP](docs/TransferFinalizeDisableOTP.md)
 - [TransferInitiate](docs/TransferInitiate.md)
 - [TransferRecipientBulk](docs/TransferRecipientBulk.md)
 - [TransferRecipientCreate](docs/TransferRecipientCreate.md)
 - [TransferRecipientUpdate](docs/TransferRecipientUpdate.md)
 - [TransferResendOTP](docs/TransferResendOTP.md)
 - [USSD](docs/USSD.md)
 - [VerificationBVNMatch](docs/VerificationBVNMatch.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### bearerAuth

- **Type**: HTTP Bearer token authentication

Example

```go
auth := context.WithValue(context.Background(), openapi.ContextAccessToken, "BEARER_TOKEN_STRING")
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

techsupport@paystack.com

