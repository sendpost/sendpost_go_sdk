# Go API client for sendpost

Email API and SMTP relay to not just send and measure email sending, but also alert and optimise. We provide you with tools, expertise and support needed to reliably deliver emails to your customers inboxes on time, every time.

## Installation

Install the following dependencies:

```shell
go get github.com/sendpost/sendpost_go_sdk
```

Put the package under your project folder and add the following in import:

```golang
import (
	sendpost "github.com/sendpost/sendpost_go_sdk"
)
```

## Getting Started

Please follow the [installation](#installation) instruction and execute the following Go code:

```go
cfg := sendpost.NewConfiguration()
client := sendpost.NewAPIClient(cfg)

emailMessage := sendpost.EmailMessage{}
emailMessage.SetSubject("Hello World")
emailMessage.SetHtmlBody("<strong>it works!</strong>")
emailMessage.SetIppool("PiedPiper")
emailMessage.From = &sendpost.From{}
emailMessage.From.SetEmail("richard@piedpiper.com")
tos := make([]sendpost.To, 0)
to := &sendpost.To{}
to.SetEmail("gavin@hooli.com")
tos = append(tos, *to)
emailMessage.To = tos

emailRequest := sendpost.ApiSendEmailRequest{}
emailRequest = emailRequest.XSubAccountApiKey("your_api_key")
emailRequest = emailRequest.EmailMessage(emailMessage)
res, _, err := client.EmailApi.SendEmailExecute(emailRequest)
```

Example with cc, bcc and template:

```go
cfg := sendpost.NewConfiguration()
client := sendpost.NewAPIClient(cfg)

emailMessage := sendpost.EmailMessage{}
emailMessage.SetSubject("Hello World")
emailMessage.SetHtmlBody("<strong>it works!</strong>")
emailMessage.SetIppool("PiedPiper")
emailMessage.From = &sendpost.From{}
emailMessage.From.SetEmail("richard@piedpiper.com")

emailMessage.SetTemplate("Welcome Mail")

tos := make([]sendpost.To, 0)
to := &sendpost.To{}
to.SetEmail("gavin@hooli.com")

ccs := make([]sendpost.CopyTo, 0)
cc := &sendpost.CopyTo{}
cc.SetEmail("dinesh@bachmanity.com")
ccs = append(ccs, *cc)
to.SetCc(ccs)
bccs := make([]sendpost.CopyTo, 0)
bcc := &sendpost.CopyTo{}
bcc.SetEmail("jian@bachmanity.com")
bccs = append(bccs, *bcc)
to.SetBcc(bccs)

tos = append(tos, *to)
emailMessage.To = tos

emailRequest := sendpost.ApiSendEmailWithTemplateRequest{}
emailRequest = emailRequest.XSubAccountApiKey("your_api_key")
emailRequest = emailRequest.EmailMessage(emailMessage)
res, _, err := client.EmailApi.SendEmailWithTemplateExecute(emailRequest)
```

## Suppressions

Create Suppressions

```go
cfg := sendpost.NewConfiguration()
client := sendpost.NewAPIClient(cfg)

rSuppression := sendpost.RSuppression{}
suppressionEmails := make([]sendpost.SuppressionEmail, 0)
email := "richard@piedpiper_fake.com"
suppressionEmails = append(suppressionEmails, sendpost.SuppressionEmail{
	Email: &email,
})
rSuppression.SetHardBounce(suppressionEmails)

// fields are optional, but you have to send at least one of them.

//rSuppression.SetManual(suppressionEmails)
//rSuppression.SetUnsubscribe(suppressionEmails)
//rSuppression.SetSpamComplaint(suppressionEmails)

createSuppressionReq := sendpost.ApiCreateSuppressionsRequest{}
createSuppressionReq = createSuppressionReq.XSubAccountApiKey("your_api_key")
createSuppressionReq = createSuppressionReq.RSuppression(rSuppression)
res, _, err := client.SuppressionApi.CreateSuppressionsExecute(createSuppressionReq)
```

Get Suppressions

```go
cfg := sendpost.NewConfiguration()
client := sendpost.NewAPIClient(cfg)

getSuppressionReq := sendpost.ApiGetSuppressionsRequest{}
getSuppressionReq = getSuppressionReq.XSubAccountApiKey("your_api_key")
getSuppressionReq = getSuppressionReq.From("2023-06-07")
getSuppressionReq = getSuppressionReq.To("2023-08-08")
getSuppressionReq = getSuppressionReq.Offset(0)
getSuppressionReq = getSuppressionReq.Limit(10)
getSuppressionReq = getSuppressionReq.Search("")

res, _, err := client.SuppressionApi.GetSuppressionsExecute(getSuppressionReq)
```

Delete Suppressions

```go
cfg := sendpost.NewConfiguration()
client := sendpost.NewAPIClient(cfg)

rDSuppression := sendpost.RDSuppression{}
suppressionEmails := make([]sendpost.SuppressionEmail, 0)
email := "richard@piedpiper_fake.com"
suppressionEmails = append(suppressionEmails, sendpost.SuppressionEmail{
	Email: &email,
})
rDSuppression.SetSuppressions(suppressionEmails)

deleteSuppressionReq := sendpost.ApiDeleteSuppressionRequest{}
deleteSuppressionReq = deleteSuppressionReq.XSubAccountApiKey("your_api_key")
deleteSuppressionReq = deleteSuppressionReq.RDSuppression(rDSuppression)
res, _, err := client.SuppressionApi.DeleteSuppressionExecute(deleteSuppressionReq)
```

Count Suppressions

```go
cfg := sendpost.NewConfiguration()
client := sendpost.NewAPIClient(cfg)

countSuppressionReq := sendpost.ApiCountRequest{}
countSuppressionReq = countSuppressionReq.XSubAccountApiKey("your_api_key")
countSuppressionReq = countSuppressionReq.From("2023-06-07")
countSuppressionReq = countSuppressionReq.To("2023-08-08")

res, _, err := client.SuppressionApi.CountExecute(countSuppressionReq)
```

## Documentation for API Endpoints

All URIs are relative to *https://api.sendpost.io/api/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*EmailAPI* | [**SendEmail**](docs/EmailAPI.md#sendemail) | **Post** /subaccount/email/ | 
*EmailAPI* | [**SendEmailWithTemplate**](docs/EmailAPI.md#sendemailwithtemplate) | **Post** /subaccount/email/template | 
*SuppressionAPI* | [**Count**](docs/SuppressionAPI.md#count) | **Get** /subaccount/suppression/count | 
*SuppressionAPI* | [**CreateSuppressions**](docs/SuppressionAPI.md#createsuppressions) | **Post** /subaccount/suppression/ | 
*SuppressionAPI* | [**DeleteSuppression**](docs/SuppressionAPI.md#deletesuppression) | **Delete** /subaccount/suppression/ | 
*SuppressionAPI* | [**GetSuppressions**](docs/SuppressionAPI.md#getsuppressions) | **Get** /subaccount/suppression/ | 


## Documentation For Models

 - [Attachment](docs/Attachment.md)
 - [City](docs/City.md)
 - [CopyTo](docs/CopyTo.md)
 - [CountStat](docs/CountStat.md)
 - [DeleteResponse](docs/DeleteResponse.md)
 - [Device](docs/Device.md)
 - [EmailMessage](docs/EmailMessage.md)
 - [EmailResponse](docs/EmailResponse.md)
 - [EventMetadata](docs/EventMetadata.md)
 - [From](docs/From.md)
 - [Os](docs/Os.md)
 - [QEmailMessage](docs/QEmailMessage.md)
 - [QEvent](docs/QEvent.md)
 - [RDSuppression](docs/RDSuppression.md)
 - [RSuppression](docs/RSuppression.md)
 - [ReplyTo](docs/ReplyTo.md)
 - [Suppression](docs/Suppression.md)
 - [SuppressionEmail](docs/SuppressionEmail.md)
 - [To](docs/To.md)
 - [UserAgent](docs/UserAgent.md)
 - [WebhookEvent](docs/WebhookEvent.md)


## Documentation For Authorization

Endpoints do not require authorization.


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

hello@sendpost.io

