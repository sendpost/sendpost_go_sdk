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

## Documentation for API Endpoints

All URIs are relative to *https://api.sendpost.io/api/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*EmailApi* | [**SendEmail**](docs/EmailApi.md#sendemail) | **Post** /subaccount/email/ | 
*EmailApi* | [**SendEmailWithTemplate**](docs/EmailApi.md#sendemailwithtemplate) | **Post** /subaccount/email/template | 
*SuppressionApi* | [**Count**](docs/SuppressionApi.md#count) | **Get** /subaccount/suppression/count | 
*SuppressionApi* | [**CreateSuppressions**](docs/SuppressionApi.md#createsuppressions) | **Post** /subaccount/suppression/ | 
*SuppressionApi* | [**DeleteSuppression**](docs/SuppressionApi.md#deletesuppression) | **Delete** /subaccount/suppression/ | 
*SuppressionApi* | [**GetSuppressions**](docs/SuppressionApi.md#getsuppressions) | **Get** /subaccount/suppression/ | 


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

dev@sendpost.io
