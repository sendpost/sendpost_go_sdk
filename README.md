# Go API client for swagger

Email API and SMTP relay to not just send and measure email sending, but also alert and optimise. We provide you with tools, expertise and support needed to reliably deliver emails to your customers inboxes on time, every time.

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 1.0.0
- Build package: io.swagger.codegen.languages.GoClientCodegen

## Installation
Put the package under your project folder and add the following in import:
```golang
import "./swagger"
```

## Documentation for API Endpoints

All URIs are relative to *https://api.sendpost.io/api/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AccountadminApi* | [**AccountAdminRouterAssumeAccountGetToken**](docs/AccountadminApi.md#accountadminrouterassumeaccountgettoken) | **Get** /account/admin/assume | 
*AccountalertApi* | [**AlertRouterCount**](docs/AccountalertApi.md#alertroutercount) | **Get** /account/alert/count | 
*AccountalertApi* | [**AlertRouterCreateAlert**](docs/AccountalertApi.md#alertroutercreatealert) | **Post** /account/alert/ | 
*AccountalertApi* | [**AlertRouterDelete**](docs/AccountalertApi.md#alertrouterdelete) | **Delete** /account/alert/{alertId} | 
*AccountalertApi* | [**AlertRouterGetAll**](docs/AccountalertApi.md#alertroutergetall) | **Get** /account/alert/ | 
*AccountalertApi* | [**AlertRouterUpdate**](docs/AccountalertApi.md#alertrouterupdate) | **Put** /account/alert/{alertId} | 
*AccountdomainApi* | [**AccountDomainRouterCount**](docs/AccountdomainApi.md#accountdomainroutercount) | **Get** /account/domain/count | 
*AccountdomainApi* | [**AccountDomainRouterGetAll**](docs/AccountdomainApi.md#accountdomainroutergetall) | **Get** /account/domain/ | 
*AccountdomainReportApi* | [**DomainReportRouterReputation**](docs/AccountdomainReportApi.md#domainreportrouterreputation) | **Get** /account/domainReport/reputation | 
*AccountdomainStatApi* | [**DomainStatRouterGetAllAggregateDomainStatsByGroup**](docs/AccountdomainStatApi.md#domainstatroutergetallaggregatedomainstatsbygroup) | **Get** /account/domainStat/{domainId}/aggregate/provider | 
*AccountdomainStatApi* | [**DomainStatRouterGetStatsForASingleDomainStats**](docs/AccountdomainStatApi.md#domainstatroutergetstatsforasingledomainstats) | **Get** /account/domainStat/{domainId}/aggregate | 
*AccounteventApi* | [**EventRouterCountAllEventsFromAAccountForAGivenTimeRange**](docs/AccounteventApi.md#eventroutercountalleventsfromaaccountforagiventimerange) | **Get** /account/event/count | 
*AccounteventApi* | [**EventRouterCountAllEventsFromANodeOfASubAccountForAGivenTimeRange**](docs/AccounteventApi.md#eventroutercountalleventsfromanodeofasubaccountforagiventimerange) | **Get** /account/event/node/count | 
*AccounteventApi* | [**EventRouterGet**](docs/AccounteventApi.md#eventrouterget) | **Get** /account/event/{eventId} | 
*AccounteventApi* | [**EventRouterGetAllEventTimestampKeysOfASubAccountFromASpecificNodeForAGivenTimeRange**](docs/AccounteventApi.md#eventroutergetalleventtimestampkeysofasubaccountfromaspecificnodeforagiventimerange) | **Get** /account/event/node/timestampkeys | 
*AccounteventApi* | [**EventRouterGetAllEventsFromAAccountForAGivenTimeRange**](docs/AccounteventApi.md#eventroutergetalleventsfromaaccountforagiventimerange) | **Get** /account/event/ | 
*AccounteventApi* | [**EventRouterGetAllEventsFromAnAccountWhichHasOnlyProccessed**](docs/AccounteventApi.md#eventroutergetalleventsfromanaccountwhichhasonlyproccessed) | **Get** /account/event/node/eventType | 
*AccounteventApi* | [**EventRouterGetAllEventsOfAAccountFromASpecificNode**](docs/AccounteventApi.md#eventroutergetalleventsofaaccountfromaspecificnode) | **Post** /account/event/node | 
*AccounteventApi* | [**EventRouterGetEventInNode**](docs/AccounteventApi.md#eventroutergeteventinnode) | **Get** /account/event/node/{eventId} | 
*AccountincidentApi* | [**IncidentRouterAdd**](docs/AccountincidentApi.md#incidentrouteradd) | **Post** /account/incident/{incidentId}/comment | 
*AccountincidentApi* | [**IncidentRouterCount**](docs/AccountincidentApi.md#incidentroutercount) | **Get** /account/incident/count | 
*AccountincidentApi* | [**IncidentRouterCreate**](docs/AccountincidentApi.md#incidentroutercreate) | **Post** /account/incident/ | 
*AccountincidentApi* | [**IncidentRouterGetAll**](docs/AccountincidentApi.md#incidentroutergetall) | **Get** /account/incident/ | 
*AccountincidentApi* | [**IncidentRouterGetAllComments**](docs/AccountincidentApi.md#incidentroutergetallcomments) | **Get** /account/incident/{incidentId}/comment | 
*AccountincidentApi* | [**IncidentRouterGetIncident**](docs/AccountincidentApi.md#incidentroutergetincident) | **Get** /account/incident/{incidentId} | 
*AccountincidentApi* | [**IncidentRouterUpdate**](docs/AccountincidentApi.md#incidentrouterupdate) | **Put** /account/incident/{incidentId} | 
*AccountintegrationApi* | [**AccountIntegrationRouterCount**](docs/AccountintegrationApi.md#accountintegrationroutercount) | **Get** /account/integration/count | 
*AccountintegrationApi* | [**AccountIntegrationRouterCreate**](docs/AccountintegrationApi.md#accountintegrationroutercreate) | **Post** /account/integration/{itype} | 
*AccountintegrationApi* | [**AccountIntegrationRouterDelete**](docs/AccountintegrationApi.md#accountintegrationrouterdelete) | **Delete** /account/integration/{itype} | 
*AccountintegrationApi* | [**AccountIntegrationRouterDisableGlockappsIPMonitoring**](docs/AccountintegrationApi.md#accountintegrationrouterdisableglockappsipmonitoring) | **Delete** /account/integration/glockapps/monitor/{ipid} | 
*AccountintegrationApi* | [**AccountIntegrationRouterEnableGlockappsIPMonitoring**](docs/AccountintegrationApi.md#accountintegrationrouterenableglockappsipmonitoring) | **Post** /account/integration/glockapps/monitor/{ipid} | 
*AccountintegrationApi* | [**AccountIntegrationRouterGetAll**](docs/AccountintegrationApi.md#accountintegrationroutergetall) | **Get** /account/integration/ | 
*AccountintegrationApi* | [**AccountIntegrationRouterGetMonitoredIPStats**](docs/AccountintegrationApi.md#accountintegrationroutergetmonitoredipstats) | **Get** /account/integration/glockapps/monitor/stat/{ipid} | 
*AccountintegrationApi* | [**AccountIntegrationRouterUpdate**](docs/AccountintegrationApi.md#accountintegrationrouterupdate) | **Put** /account/integration/{itype} | 
*AccountinvitationApi* | [**InvitationRouterCreate**](docs/AccountinvitationApi.md#invitationroutercreate) | **Post** /account/invitation/ | 
*AccountinvitationApi* | [**InvitationRouterDelete**](docs/AccountinvitationApi.md#invitationrouterdelete) | **Delete** /account/invitation/{invitationId} | 
*AccountinvitationApi* | [**InvitationRouterGetAll**](docs/AccountinvitationApi.md#invitationroutergetall) | **Get** /account/invitation/ | 
*AccountipApi* | [**IPRouterAllocateIP**](docs/AccountipApi.md#iprouterallocateip) | **Post** /account/ip/allocate | 
*AccountipApi* | [**IPRouterCount**](docs/AccountipApi.md#iproutercount) | **Get** /account/ip/count | 
*AccountipApi* | [**IPRouterDelete**](docs/AccountipApi.md#iprouterdelete) | **Delete** /account/ip/{ipid} | 
*AccountipApi* | [**IPRouterGet**](docs/AccountipApi.md#iprouterget) | **Get** /account/ip/{ipid} | 
*AccountipApi* | [**IPRouterGetAll**](docs/AccountipApi.md#iproutergetall) | **Get** /account/ip/ | 
*AccountipApi* | [**IPRouterGetAllIPIncidents**](docs/AccountipApi.md#iproutergetallipincidents) | **Get** /account/ip/{ipid}/incident | 
*AccountipApi* | [**IPRouterGetIpHealth**](docs/AccountipApi.md#iproutergetiphealth) | **Get** /account/ip/{ipid}/health | 
*AccountipApi* | [**IPRouterUpdate**](docs/AccountipApi.md#iprouterupdate) | **Put** /account/ip/{ipid} | 
*AccountippoolApi* | [**AccountIPPoolRouterCount**](docs/AccountippoolApi.md#accountippoolroutercount) | **Get** /account/ippool/count | 
*AccountippoolApi* | [**AccountIPPoolRouterCreate**](docs/AccountippoolApi.md#accountippoolroutercreate) | **Post** /account/ippool/ | 
*AccountippoolApi* | [**AccountIPPoolRouterDelete**](docs/AccountippoolApi.md#accountippoolrouterdelete) | **Delete** /account/ippool/{ippoolid} | 
*AccountippoolApi* | [**AccountIPPoolRouterGet**](docs/AccountippoolApi.md#accountippoolrouterget) | **Get** /account/ippool/{ippoolid} | 
*AccountippoolApi* | [**AccountIPPoolRouterGetAll**](docs/AccountippoolApi.md#accountippoolroutergetall) | **Get** /account/ippool/ | 
*AccountippoolApi* | [**AccountIPPoolRouterUpdate**](docs/AccountippoolApi.md#accountippoolrouterupdate) | **Put** /account/ippool/{ippoolid} | 
*AccountipstatApi* | [**IPStatRouterGetAllAggregateIPStats**](docs/AccountipstatApi.md#ipstatroutergetallaggregateipstats) | **Get** /account/ip/stat/{ipid}/aggregate | 
*AccountipstatApi* | [**IPStatRouterGetAllAggregateIPStatsByProvider**](docs/AccountipstatApi.md#ipstatroutergetallaggregateipstatsbyprovider) | **Get** /account/ip/stat/{ipid}/aggregate/provider | 
*AccountipstatApi* | [**IPStatRouterGetAllAggregatedProviderStatsForAIP**](docs/AccountipstatApi.md#ipstatroutergetallaggregatedproviderstatsforaip) | **Get** /account/ip/stat/{ipid}/aggregate/providers | 
*AccountipstatApi* | [**IPStatRouterGetAllAggregatedProviderStatsForASpecificSubAccountOfAIP**](docs/AccountipstatApi.md#ipstatroutergetallaggregatedproviderstatsforaspecificsubaccountofaip) | **Get** /account/ip/stat/{ipid}/aggregate/sid/{sid}/providers | 
*AccountipstatApi* | [**IPStatRouterGetAllAggregatedSubAccountStatsForAnIP**](docs/AccountipstatApi.md#ipstatroutergetallaggregatedsubaccountstatsforanip) | **Get** /account/ip/stat/{ipid}/aggregate/subaccounts | 
*AccountipstatApi* | [**IPStatRouterGetAllIPStats**](docs/AccountipstatApi.md#ipstatroutergetallipstats) | **Get** /account/ip/stat/{ipid} | 
*AccountlabelApi* | [**LabelRouterCount**](docs/AccountlabelApi.md#labelroutercount) | **Get** /account/label/count | 
*AccountlabelApi* | [**LabelRouterCreate**](docs/AccountlabelApi.md#labelroutercreate) | **Post** /account/label/ | 
*AccountlabelApi* | [**LabelRouterDelete**](docs/AccountlabelApi.md#labelrouterdelete) | **Delete** /account/label/{labelId} | 
*AccountlabelApi* | [**LabelRouterGet**](docs/AccountlabelApi.md#labelrouterget) | **Get** /account/label/{labelId} | 
*AccountlabelApi* | [**LabelRouterGetAll**](docs/AccountlabelApi.md#labelroutergetall) | **Get** /account/label/ | 
*AccountlabelApi* | [**LabelRouterUpdate**](docs/AccountlabelApi.md#labelrouterupdate) | **Put** /account/label/{labelId} | 
*AccountmailReportApi* | [**MailReportRouterMailReport**](docs/AccountmailReportApi.md#mailreportroutermailreport) | **Post** /account/mailReport/ | 
*AccountmailReportApi* | [**MailReportRouterMailReportGetSingleReport**](docs/AccountmailReportApi.md#mailreportroutermailreportgetsinglereport) | **Get** /account/mailReport/{reportId} | 
*AccountmailReportApi* | [**MailReportRouterMailReportProviders**](docs/AccountmailReportApi.md#mailreportroutermailreportproviders) | **Get** /account/mailReport/provider | 
*AccountmemberApi* | [**MemberRouterDelete**](docs/AccountmemberApi.md#memberrouterdelete) | **Delete** /account/member/{memberId} | 
*AccountmemberApi* | [**MemberRouterGet**](docs/AccountmemberApi.md#memberrouterget) | **Get** /account/member/{memberId} | 
*AccountmemberApi* | [**MemberRouterGetAll**](docs/AccountmemberApi.md#memberroutergetall) | **Get** /account/member/ | 
*AccountmemberApi* | [**MemberRouterUpdate**](docs/AccountmemberApi.md#memberrouterupdate) | **Put** /account/member/{memberId} | 
*AccountmemberApi* | [**MemberRouterVerifyByEmailRequest**](docs/AccountmemberApi.md#memberrouterverifybyemailrequest) | **Post** /account/member/{memberId}/verify/email | 
*AccountmessageApi* | [**MessageRouterGet**](docs/AccountmessageApi.md#messagerouterget) | **Get** /account/message/{messageId} | 
*AccountmessageApi* | [**MessageRouterGetAllEventsForAMessageId**](docs/AccountmessageApi.md#messageroutergetalleventsforamessageid) | **Get** /account/message/{messageId}/events | 
*AccountmessageApi* | [**MessageRouterGetAllEventsForAMessageIdFromANode**](docs/AccountmessageApi.md#messageroutergetalleventsforamessageidfromanode) | **Get** /account/message/node/{messageId}/events | 
*AccountmessageApi* | [**MessageRouterGetMessageFromNode**](docs/AccountmessageApi.md#messageroutergetmessagefromnode) | **Get** /account/message/node/{messageId} | 
*AccountonboardingApi* | [**OnboardingRouterGetOnboardingChecklist**](docs/AccountonboardingApi.md#onboardingroutergetonboardingchecklist) | **Get** /account/onboarding/checklist | 
*AccountpaymentApi* | [**PaymentRouterApplyCouponToStripeCustomer**](docs/AccountpaymentApi.md#paymentrouterapplycoupontostripecustomer) | **Post** /account/payment/customer/coupon | 
*AccountpaymentApi* | [**PaymentRouterCreateCustomerPortal**](docs/AccountpaymentApi.md#paymentroutercreatecustomerportal) | **Post** /account/payment/portal | 
*AccountpaymentApi* | [**PaymentRouterCreatePaymentSubscription**](docs/AccountpaymentApi.md#paymentroutercreatepaymentsubscription) | **Post** /account/payment/subscription | 
*AccountpaymentApi* | [**PaymentRouterHandlePaymentWebhook**](docs/AccountpaymentApi.md#paymentrouterhandlepaymentwebhook) | **Post** /account/payment/webhook | 
*AccountrecipientApi* | [**RecipientRouterGetAllMessagesForARecipient**](docs/AccountrecipientApi.md#recipientroutergetallmessagesforarecipient) | **Get** /account/recipient/{recipient}/messages | 
*AccountrecipientApi* | [**RecipientRouterGetAllMessagesForARecipientFromANode**](docs/AccountrecipientApi.md#recipientroutergetallmessagesforarecipientfromanode) | **Get** /account/recipient/node/{recipient}/messages | 
*AccountsettingApi* | [**AccountSettingRouterUpdate**](docs/AccountsettingApi.md#accountsettingrouterupdate) | **Put** /account/setting/ | 
*AccountsmtpstatApi* | [**SMTPStatRouterGetAllAggregateIPProviderSMTPStats**](docs/AccountsmtpstatApi.md#smtpstatroutergetallaggregateipprovidersmtpstats) | **Get** /account/smtp/stat/ip/{ipid}/provider/{provider}/aggregate | 
*AccountsmtpstatApi* | [**SMTPStatRouterGetAllAggregateIPSMTPStats**](docs/AccountsmtpstatApi.md#smtpstatroutergetallaggregateipsmtpstats) | **Get** /account/smtp/stat/ip/{ipid}/aggregate | 
*AccountsmtpstatApi* | [**SMTPStatRouterGetAllAggregateIPSMTPStatsForSubAccount**](docs/AccountsmtpstatApi.md#smtpstatroutergetallaggregateipsmtpstatsforsubaccount) | **Get** /account/smtp/stat/ip/{ipid}/subaccount/{sid}/aggregate | 
*AccountsmtpstatApi* | [**SMTPStatRouterGetAllAggregateSubAccountProviderSMTPStats**](docs/AccountsmtpstatApi.md#smtpstatroutergetallaggregatesubaccountprovidersmtpstats) | **Get** /account/smtp/stat/subaccount/{sid}/provider/{provider}/aggregate | 
*AccountsmtpstatApi* | [**SMTPStatRouterGetAllAggregateSubAccountSMTPStats**](docs/AccountsmtpstatApi.md#smtpstatroutergetallaggregatesubaccountsmtpstats) | **Get** /account/smtp/stat/subaccount/{sid}/aggregate | 
*AccountsmtpstatApi* | [**SMTPStatRouterGetAllAggregateSubAccountSMTPStatsForIP**](docs/AccountsmtpstatApi.md#smtpstatroutergetallaggregatesubaccountsmtpstatsforip) | **Get** /account/smtp/stat/subaccount/{sid}/ip/{ipid}/aggregate | 
*AccountstatApi* | [**AccountStatRouterGetAllAccountStats**](docs/AccountstatApi.md#accountstatroutergetallaccountstats) | **Get** /account/stat/ | 
*AccountstatApi* | [**AccountStatRouterGetAllAccountStatsByGroup**](docs/AccountstatApi.md#accountstatroutergetallaccountstatsbygroup) | **Get** /account/stat/group | 
*AccountstatApi* | [**AccountStatRouterGetAllAggregateAccountStats**](docs/AccountstatApi.md#accountstatroutergetallaggregateaccountstats) | **Get** /account/stat/aggregate | 
*AccountstatApi* | [**AccountStatRouterGetAllAggregateAccountStatsByGroup**](docs/AccountstatApi.md#accountstatroutergetallaggregateaccountstatsbygroup) | **Get** /account/stat/aggregate/group | 
*AccountsubaccountApi* | [**SubAccountRouterCount**](docs/AccountsubaccountApi.md#subaccountroutercount) | **Get** /account/subaccount/count | 
*AccountsubaccountApi* | [**SubAccountRouterCreate**](docs/AccountsubaccountApi.md#subaccountroutercreate) | **Post** /account/subaccount/ | 
*AccountsubaccountApi* | [**SubAccountRouterDelete**](docs/AccountsubaccountApi.md#subaccountrouterdelete) | **Delete** /account/subaccount/{subAccountId} | 
*AccountsubaccountApi* | [**SubAccountRouterGet**](docs/AccountsubaccountApi.md#subaccountrouterget) | **Get** /account/subaccount/{subAccountId} | 
*AccountsubaccountApi* | [**SubAccountRouterGetAll**](docs/AccountsubaccountApi.md#subaccountroutergetall) | **Get** /account/subaccount/ | 
*AccountsubaccountApi* | [**SubAccountRouterGetAllSubAccountIncidents**](docs/AccountsubaccountApi.md#subaccountroutergetallsubaccountincidents) | **Get** /account/subaccount/{subAccountId}/incident | 
*AccountsubaccountApi* | [**SubAccountRouterUpdate**](docs/AccountsubaccountApi.md#subaccountrouterupdate) | **Put** /account/subaccount/{subAccountId} | 
*AccounttagApi* | [**TagRouterCreate**](docs/AccounttagApi.md#tagroutercreate) | **Post** /account/tag/ | 
*AccounttagApi* | [**TagRouterDelete**](docs/AccounttagApi.md#tagrouterdelete) | **Delete** /account/tag/{tagid} | 
*AccounttagApi* | [**TagRouterGetAll**](docs/AccounttagApi.md#tagroutergetall) | **Get** /account/tag/ | 
*AccounttemplateApi* | [**AccountTemplateRouterCopy**](docs/AccounttemplateApi.md#accounttemplateroutercopy) | **Post** /account/template/{templateid}/copy | 
*AccounttemplateApi* | [**AccountTemplateRouterCount**](docs/AccounttemplateApi.md#accounttemplateroutercount) | **Get** /account/template/count | 
*AccounttemplateApi* | [**AccountTemplateRouterCreate**](docs/AccounttemplateApi.md#accounttemplateroutercreate) | **Post** /account/template/ | 
*AccounttemplateApi* | [**AccountTemplateRouterDelete**](docs/AccounttemplateApi.md#accounttemplaterouterdelete) | **Delete** /account/template/{templateid} | 
*AccounttemplateApi* | [**AccountTemplateRouterGet**](docs/AccounttemplateApi.md#accounttemplaterouterget) | **Get** /account/template/{templateid} | 
*AccounttemplateApi* | [**AccountTemplateRouterGetAll**](docs/AccounttemplateApi.md#accounttemplateroutergetall) | **Get** /account/template/ | 
*AccounttemplateApi* | [**AccountTemplateRouterUpdate**](docs/AccounttemplateApi.md#accounttemplaterouterupdate) | **Put** /account/template/{templateid} | 
*AccountvalidationApi* | [**ValidateRouterValidateEmailBulk**](docs/AccountvalidationApi.md#validateroutervalidateemailbulk) | **Post** /account/validation/bulk | 
*AccountvalidationApi* | [**ValidationRouterCount**](docs/AccountvalidationApi.md#validationroutercount) | **Get** /account/validation/count | 
*AccountvalidationApi* | [**ValidationRouterDeleteValidation**](docs/AccountvalidationApi.md#validationrouterdeletevalidation) | **Delete** /account/validation/ | 
*AccountvalidationApi* | [**ValidationRouterGetAll**](docs/AccountvalidationApi.md#validationroutergetall) | **Get** /account/validation/ | 
*AccountvalidationApi* | [**ValidationRouterValidateEmailList**](docs/AccountvalidationApi.md#validationroutervalidateemaillist) | **Post** /account/validation/ | 
*AccountwebhookApi* | [**AccountWebhookRouterCount**](docs/AccountwebhookApi.md#accountwebhookroutercount) | **Get** /account/webhook/count | 
*AccountwebhookApi* | [**AccountWebhookRouterCreate**](docs/AccountwebhookApi.md#accountwebhookroutercreate) | **Post** /account/webhook/ | 
*AccountwebhookApi* | [**AccountWebhookRouterDelete**](docs/AccountwebhookApi.md#accountwebhookrouterdelete) | **Delete** /account/webhook/{webhookId} | 
*AccountwebhookApi* | [**AccountWebhookRouterGet**](docs/AccountwebhookApi.md#accountwebhookrouterget) | **Get** /account/webhook/{webhookId} | 
*AccountwebhookApi* | [**AccountWebhookRouterGetAll**](docs/AccountwebhookApi.md#accountwebhookroutergetall) | **Get** /account/webhook/ | 
*AccountwebhookApi* | [**AccountWebhookRouterUpdate**](docs/AccountwebhookApi.md#accountwebhookrouterupdate) | **Put** /account/webhook/{webhookId} | 
*AuthApi* | [**AuthRouterCreate**](docs/AuthApi.md#authroutercreate) | **Post** /auth/create | 
*AuthApi* | [**AuthRouterGetAuthInfo**](docs/AuthApi.md#authroutergetauthinfo) | **Post** /auth/info | 
*AuthApi* | [**AuthRouterUpdateAuthInfo**](docs/AuthApi.md#authrouterupdateauthinfo) | **Put** /auth/info | 
*ClusterApi* | [**ClusterRouterAddItemsToSuppressionFilterOfEveryNodeInCluster**](docs/ClusterApi.md#clusterrouteradditemstosuppressionfilterofeverynodeincluster) | **Post** /cluster/suppression/filter | 
*ClusterApi* | [**ClusterRouterDeleteItemFromCacheOfEveryNodeInCluster**](docs/ClusterApi.md#clusterrouterdeleteitemfromcacheofeverynodeincluster) | **Delete** /cluster/cache | 
*ClusterApi* | [**ClusterRouterDeleteItemsFromSuppressionFilterOfEveryNodeInCluster**](docs/ClusterApi.md#clusterrouterdeleteitemsfromsuppressionfilterofeverynodeincluster) | **Delete** /cluster/suppression/filter | 
*ClusterApi* | [**ClusterRouterGetItemFromCacheOfEveryNodeInCluster**](docs/ClusterApi.md#clusterroutergetitemfromcacheofeverynodeincluster) | **Get** /cluster/cache | 
*ClusterApi* | [**ClusterRouterGetItemFromCacheOfSpecificNodeInCluster**](docs/ClusterApi.md#clusterroutergetitemfromcacheofspecificnodeincluster) | **Get** /cluster/cache/node | 
*ClusterApi* | [**ClusterRouterGetItemFromCacheOfSpecificNodeInCluster_0**](docs/ClusterApi.md#clusterroutergetitemfromcacheofspecificnodeincluster_0) | **Delete** /cluster/cache/node | 
*EditorApi* | [**EditorApiRouterGetToken**](docs/EditorApi.md#editorapiroutergettoken) | **Get** /editor/ | 
*SmtpApi* | [**SMTPRouterReceiveWebhooksRaisedFromSMTPServers**](docs/SmtpApi.md#smtprouterreceivewebhooksraisedfromsmtpservers) | **Post** /smtp/webhook | 
*SubaccountdomainApi* | [**DomainRouterCount**](docs/SubaccountdomainApi.md#domainroutercount) | **Get** /subaccount/domain/count | 
*SubaccountdomainApi* | [**DomainRouterCreate**](docs/SubaccountdomainApi.md#domainroutercreate) | **Post** /subaccount/domain/ | 
*SubaccountdomainApi* | [**DomainRouterDelete**](docs/SubaccountdomainApi.md#domainrouterdelete) | **Delete** /subaccount/domain/{domainId} | 
*SubaccountdomainApi* | [**DomainRouterGet**](docs/SubaccountdomainApi.md#domainrouterget) | **Get** /subaccount/domain/{domainId} | 
*SubaccountdomainApi* | [**DomainRouterGetAll**](docs/SubaccountdomainApi.md#domainroutergetall) | **Get** /subaccount/domain/ | 
*SubaccountdomainApi* | [**DomainRouterUpdate**](docs/SubaccountdomainApi.md#domainrouterupdate) | **Put** /subaccount/domain/{domainId} | 
*SubaccountdomainApi* | [**DomainRouterVerify**](docs/SubaccountdomainApi.md#domainrouterverify) | **Post** /subaccount/domain/{domainId}/verify | 
*SubaccountdomainApi* | [**DomainRouterVerifyByToken**](docs/SubaccountdomainApi.md#domainrouterverifybytoken) | **Post** /subaccount/domain/{domainId}/verify/email/{token} | 
*SubaccountdomainApi* | [**DomainRouterVerifyRequest**](docs/SubaccountdomainApi.md#domainrouterverifyrequest) | **Post** /subaccount/domain/{domainId}/verify/email | 
*SubaccountemailApi* | [**EmailRouterSendEmail**](docs/SubaccountemailApi.md#emailroutersendemail) | **Post** /subaccount/email/ | 
*SubaccountemailApi* | [**EmailRouterSendEmailWithTemplate**](docs/SubaccountemailApi.md#emailroutersendemailwithtemplate) | **Post** /subaccount/email/template | 
*SubaccountippoolApi* | [**IPPoolRouterCount**](docs/SubaccountippoolApi.md#ippoolroutercount) | **Get** /subaccount/ippool/count | 
*SubaccountippoolApi* | [**IPPoolRouterCreate**](docs/SubaccountippoolApi.md#ippoolroutercreate) | **Post** /subaccount/ippool/ | 
*SubaccountippoolApi* | [**IPPoolRouterDelete**](docs/SubaccountippoolApi.md#ippoolrouterdelete) | **Delete** /subaccount/ippool/{ippoolid} | 
*SubaccountippoolApi* | [**IPPoolRouterGet**](docs/SubaccountippoolApi.md#ippoolrouterget) | **Get** /subaccount/ippool/{ippoolid} | 
*SubaccountippoolApi* | [**IPPoolRouterGetAll**](docs/SubaccountippoolApi.md#ippoolroutergetall) | **Get** /subaccount/ippool/ | 
*SubaccountippoolApi* | [**IPPoolRouterUpdate**](docs/SubaccountippoolApi.md#ippoolrouterupdate) | **Put** /subaccount/ippool/{ippoolid} | 
*SubaccountsenderApi* | [**SenderRouterCount**](docs/SubaccountsenderApi.md#senderroutercount) | **Get** /subaccount/sender/count | 
*SubaccountsenderApi* | [**SenderRouterCreate**](docs/SubaccountsenderApi.md#senderroutercreate) | **Post** /subaccount/sender/ | 
*SubaccountsenderApi* | [**SenderRouterDelete**](docs/SubaccountsenderApi.md#senderrouterdelete) | **Delete** /subaccount/sender/{senderId} | 
*SubaccountsenderApi* | [**SenderRouterGet**](docs/SubaccountsenderApi.md#senderrouterget) | **Get** /subaccount/sender/{senderId} | 
*SubaccountsenderApi* | [**SenderRouterGetAll**](docs/SubaccountsenderApi.md#senderroutergetall) | **Get** /subaccount/sender/ | 
*SubaccountsenderApi* | [**SenderRouterUpdate**](docs/SubaccountsenderApi.md#senderrouterupdate) | **Put** /subaccount/sender/{senderId} | 
*SubaccountstatApi* | [**SubAccountStatRouterGetAllAggregateSubAccountStats**](docs/SubaccountstatApi.md#subaccountstatroutergetallaggregatesubaccountstats) | **Get** /subaccount/stat/aggregate | 
*SubaccountstatApi* | [**SubAccountStatRouterGetAllAggregateSubAccountStatsByGroup**](docs/SubaccountstatApi.md#subaccountstatroutergetallaggregatesubaccountstatsbygroup) | **Get** /subaccount/stat/aggregate/group | 
*SubaccountstatApi* | [**SubAccountStatRouterGetAllAggregatedGroupStatsForASubAccount**](docs/SubaccountstatApi.md#subaccountstatroutergetallaggregatedgroupstatsforasubaccount) | **Get** /subaccount/stat/aggregate/groups | 
*SubaccountstatApi* | [**SubAccountStatRouterGetAllAggregatedIPStatsForASubAccount**](docs/SubaccountstatApi.md#subaccountstatroutergetallaggregatedipstatsforasubaccount) | **Get** /subaccount/stat/aggregate/ips | 
*SubaccountstatApi* | [**SubAccountStatRouterGetAllAggregatedProviderStatsForASpecificIPOfASubAccount**](docs/SubaccountstatApi.md#subaccountstatroutergetallaggregatedproviderstatsforaspecificipofasubaccount) | **Get** /subaccount/stat/aggregate/ip/{ipid}/providers | 
*SubaccountstatApi* | [**SubAccountStatRouterGetAllAggregatedProviderStatsForASubAccount**](docs/SubaccountstatApi.md#subaccountstatroutergetallaggregatedproviderstatsforasubaccount) | **Get** /subaccount/stat/aggregate/providers | 
*SubaccountstatApi* | [**SubAccountStatRouterGetAllSubAccountStats**](docs/SubaccountstatApi.md#subaccountstatroutergetallsubaccountstats) | **Get** /subaccount/stat/ | 
*SubaccountstatApi* | [**SubAccountStatRouterGetAllSubAccountStatsByGroup**](docs/SubaccountstatApi.md#subaccountstatroutergetallsubaccountstatsbygroup) | **Get** /subaccount/stat/group | 
*SubaccountsuppressionApi* | [**SuppressionRouterCount**](docs/SubaccountsuppressionApi.md#suppressionroutercount) | **Get** /subaccount/suppression/count | 
*SubaccountsuppressionApi* | [**SuppressionRouterCreateSuppressions**](docs/SubaccountsuppressionApi.md#suppressionroutercreatesuppressions) | **Post** /subaccount/suppression/ | 
*SubaccountsuppressionApi* | [**SuppressionRouterCreateSuppressionsInSuppressionFilter**](docs/SubaccountsuppressionApi.md#suppressionroutercreatesuppressionsinsuppressionfilter) | **Post** /subaccount/suppression/filter | 
*SubaccountsuppressionApi* | [**SuppressionRouterDeleteSuppression**](docs/SubaccountsuppressionApi.md#suppressionrouterdeletesuppression) | **Delete** /subaccount/suppression/ | 
*SubaccountsuppressionApi* | [**SuppressionRouterDeleteSuppressionsInSuppressionFilter**](docs/SubaccountsuppressionApi.md#suppressionrouterdeletesuppressionsinsuppressionfilter) | **Delete** /subaccount/suppression/filter | 
*SubaccountsuppressionApi* | [**SuppressionRouterGetAllSuppressions**](docs/SubaccountsuppressionApi.md#suppressionroutergetallsuppressions) | **Get** /subaccount/suppression/ | 
*SystemtemplateApi* | [**SystemTemplateRouterGetAllSystemTemplates**](docs/SystemtemplateApi.md#systemtemplateroutergetallsystemtemplates) | **Get** /system/template/ | 
*TrackApi* | [**TrackRouterTrackEmailOpen**](docs/TrackApi.md#trackroutertrackemailopen) | **Get** /track/open/{accountId}/{subAccountId}/{ipId}/{emailType}/{messageId}/1.png | 
*TrackApi* | [**TrackRouterTrackLinkClick**](docs/TrackApi.md#trackroutertracklinkclick) | **Get** /track/click/{accountId}/{subAccountId}/{ipId}/{emailType}/{messageId} | 
*TrackApi* | [**TrackRouterTrackUnsubscribe**](docs/TrackApi.md#trackroutertrackunsubscribe) | **Get** /track/unsubscribe/{accountId}/{subAccountId}/{ipId}/{emailType}/{messageId} | 


## Documentation For Models

 - [Alert](docs/Alert.md)
 - [ApiGlockappsMailReport](docs/ApiGlockappsMailReport.md)
 - [False](docs/False.md)
 - [ModelsAccount](docs/ModelsAccount.md)
 - [ModelsAccountDomain](docs/ModelsAccountDomain.md)
 - [ModelsAccountIpPool](docs/ModelsAccountIpPool.md)
 - [ModelsAccountTemplate](docs/ModelsAccountTemplate.md)
 - [ModelsAccountWebhook](docs/ModelsAccountWebhook.md)
 - [ModelsAgStat](docs/ModelsAgStat.md)
 - [ModelsAipStat](docs/ModelsAipStat.md)
 - [ModelsAlertLabel](docs/ModelsAlertLabel.md)
 - [ModelsAlertRequest](docs/ModelsAlertRequest.md)
 - [ModelsAlertResponse](docs/ModelsAlertResponse.md)
 - [ModelsAllClusterCache](docs/ModelsAllClusterCache.md)
 - [ModelsAttachment](docs/ModelsAttachment.md)
 - [ModelsAuthInfo](docs/ModelsAuthInfo.md)
 - [ModelsAutoWarmupContent](docs/ModelsAutoWarmupContent.md)
 - [ModelsAutoWarmupContentRequest](docs/ModelsAutoWarmupContentRequest.md)
 - [ModelsAutoWarmupDailySchedule](docs/ModelsAutoWarmupDailySchedule.md)
 - [ModelsAutoWarmupDomainSchedule](docs/ModelsAutoWarmupDomainSchedule.md)
 - [ModelsAutoWarmupPlan](docs/ModelsAutoWarmupPlan.md)
 - [ModelsAutoWarmupPlanRequest](docs/ModelsAutoWarmupPlanRequest.md)
 - [ModelsAutoWarmupRecipient](docs/ModelsAutoWarmupRecipient.md)
 - [ModelsAutoWarmupRecipientResponse](docs/ModelsAutoWarmupRecipientResponse.md)
 - [ModelsBackOffConfiguration](docs/ModelsBackOffConfiguration.md)
 - [ModelsBackOffDecreaseType](docs/ModelsBackOffDecreaseType.md)
 - [ModelsBackOffTrigger](docs/ModelsBackOffTrigger.md)
 - [ModelsBillingPortalSession](docs/ModelsBillingPortalSession.md)
 - [ModelsBlackListEngine](docs/ModelsBlackListEngine.md)
 - [ModelsBlackListResult](docs/ModelsBlackListResult.md)
 - [ModelsBlacklistStatus](docs/ModelsBlacklistStatus.md)
 - [ModelsBulkResponse](docs/ModelsBulkResponse.md)
 - [ModelsCertificateDetails](docs/ModelsCertificateDetails.md)
 - [ModelsCertificateExtensions](docs/ModelsCertificateExtensions.md)
 - [ModelsCertificateIssuer](docs/ModelsCertificateIssuer.md)
 - [ModelsCertificateSignature](docs/ModelsCertificateSignature.md)
 - [ModelsCertificateSubject](docs/ModelsCertificateSubject.md)
 - [ModelsCertificateValidity](docs/ModelsCertificateValidity.md)
 - [ModelsCity](docs/ModelsCity.md)
 - [ModelsCleanedList](docs/ModelsCleanedList.md)
 - [ModelsClusterCache](docs/ModelsClusterCache.md)
 - [ModelsComment](docs/ModelsComment.md)
 - [ModelsConsumerStats](docs/ModelsConsumerStats.md)
 - [ModelsContent](docs/ModelsContent.md)
 - [ModelsCopyTo](docs/ModelsCopyTo.md)
 - [ModelsCountStat](docs/ModelsCountStat.md)
 - [ModelsCouponOptions](docs/ModelsCouponOptions.md)
 - [ModelsCreateMailReport](docs/ModelsCreateMailReport.md)
 - [ModelsCustomerQuality](docs/ModelsCustomerQuality.md)
 - [ModelsDeleteResponse](docs/ModelsDeleteResponse.md)
 - [ModelsDetailedAlert](docs/ModelsDetailedAlert.md)
 - [ModelsDnsRecord](docs/ModelsDnsRecord.md)
 - [ModelsDomain](docs/ModelsDomain.md)
 - [ModelsDomainAge](docs/ModelsDomainAge.md)
 - [ModelsDomainCheckResult](docs/ModelsDomainCheckResult.md)
 - [ModelsEAccount](docs/ModelsEAccount.md)
 - [ModelsEAccountMember](docs/ModelsEAccountMember.md)
 - [ModelsEAccountSetting](docs/ModelsEAccountSetting.md)
 - [ModelsEAlert](docs/ModelsEAlert.md)
 - [ModelsEComment](docs/ModelsEComment.md)
 - [ModelsEDomain](docs/ModelsEDomain.md)
 - [ModelsEIncident](docs/ModelsEIncident.md)
 - [ModelsEIntegration](docs/ModelsEIntegration.md)
 - [ModelsEInvitation](docs/ModelsEInvitation.md)
 - [ModelsEMember](docs/ModelsEMember.md)
 - [ModelsESender](docs/ModelsESender.md)
 - [ModelsESubAccount](docs/ModelsESubAccount.md)
 - [ModelsESystemDomain](docs/ModelsESystemDomain.md)
 - [ModelsEValidation](docs/ModelsEValidation.md)
 - [ModelsEWebhook](docs/ModelsEWebhook.md)
 - [ModelsEditorTokenResponse](docs/ModelsEditorTokenResponse.md)
 - [ModelsEip](docs/ModelsEip.md)
 - [ModelsEipPool](docs/ModelsEipPool.md)
 - [ModelsEmailErrorCode](docs/ModelsEmailErrorCode.md)
 - [ModelsEmailList](docs/ModelsEmailList.md)
 - [ModelsEmailMessage](docs/ModelsEmailMessage.md)
 - [ModelsEmailResponse](docs/ModelsEmailResponse.md)
 - [ModelsEventMetadata](docs/ModelsEventMetadata.md)
 - [ModelsEventType](docs/ModelsEventType.md)
 - [ModelsFrequencyType](docs/ModelsFrequencyType.md)
 - [ModelsFrom](docs/ModelsFrom.md)
 - [ModelsGlockappsBlacklist](docs/ModelsGlockappsBlacklist.md)
 - [ModelsGlockappsMonitorStat](docs/ModelsGlockappsMonitorStat.md)
 - [ModelsIeMember](docs/ModelsIeMember.md)
 - [ModelsIeSubAccount](docs/ModelsIeSubAccount.md)
 - [ModelsIeTag](docs/ModelsIeTag.md)
 - [ModelsIeip](docs/ModelsIeip.md)
 - [ModelsIip](docs/ModelsIip.md)
 - [ModelsIipUpdateType](docs/ModelsIipUpdateType.md)
 - [ModelsIncident](docs/ModelsIncident.md)
 - [ModelsIncidentStatus](docs/ModelsIncidentStatus.md)
 - [ModelsInstance](docs/ModelsInstance.md)
 - [ModelsIntegration](docs/ModelsIntegration.md)
 - [ModelsIntegrationSettings](docs/ModelsIntegrationSettings.md)
 - [ModelsIntegrationType](docs/ModelsIntegrationType.md)
 - [ModelsInvitation](docs/ModelsInvitation.md)
 - [ModelsInvitationStatus](docs/ModelsInvitationStatus.md)
 - [ModelsIp](docs/ModelsIp.md)
 - [ModelsIpDomainWarmupStatus](docs/ModelsIpDomainWarmupStatus.md)
 - [ModelsIpHealthResponse](docs/ModelsIpHealthResponse.md)
 - [ModelsIpPool](docs/ModelsIpPool.md)
 - [ModelsIpPoolType](docs/ModelsIpPoolType.md)
 - [ModelsIpProviderSettings](docs/ModelsIpProviderSettings.md)
 - [ModelsIpStat](docs/ModelsIpStat.md)
 - [ModelsIpType](docs/ModelsIpType.md)
 - [ModelsLabel](docs/ModelsLabel.md)
 - [ModelsLabelType](docs/ModelsLabelType.md)
 - [ModelsMailReportResult](docs/ModelsMailReportResult.md)
 - [ModelsMember](docs/ModelsMember.md)
 - [ModelsMemberRole](docs/ModelsMemberRole.md)
 - [ModelsNotificationType](docs/ModelsNotificationType.md)
 - [ModelsOnboardingChecklist](docs/ModelsOnboardingChecklist.md)
 - [ModelsPaymentOptions](docs/ModelsPaymentOptions.md)
 - [ModelsPaymentStatus](docs/ModelsPaymentStatus.md)
 - [ModelsPipStat](docs/ModelsPipStat.md)
 - [ModelsProvider](docs/ModelsProvider.md)
 - [ModelsProviderDetails](docs/ModelsProviderDetails.md)
 - [ModelsProviderResult](docs/ModelsProviderResult.md)
 - [ModelsProviderSettings](docs/ModelsProviderSettings.md)
 - [ModelsProviderType](docs/ModelsProviderType.md)
 - [ModelsQEmailMessage](docs/ModelsQEmailMessage.md)
 - [ModelsQEvent](docs/ModelsQEvent.md)
 - [ModelsRAssumeAccount](docs/ModelsRAssumeAccount.md)
 - [ModelsRGlockappsMonitorStat](docs/ModelsRGlockappsMonitorStat.md)
 - [ModelsRStat](docs/ModelsRStat.md)
 - [ModelsRSuppression](docs/ModelsRSuppression.md)
 - [ModelsRdSuppression](docs/ModelsRdSuppression.md)
 - [ModelsReplyTo](docs/ModelsReplyTo.md)
 - [ModelsReportProvider](docs/ModelsReportProvider.md)
 - [ModelsResponse](docs/ModelsResponse.md)
 - [ModelsRipStat](docs/ModelsRipStat.md)
 - [ModelsSender](docs/ModelsSender.md)
 - [ModelsSingleCleanedMail](docs/ModelsSingleCleanedMail.md)
 - [ModelsSipStat](docs/ModelsSipStat.md)
 - [ModelsSmtpAuth](docs/ModelsSmtpAuth.md)
 - [ModelsSmtpStat](docs/ModelsSmtpStat.md)
 - [ModelsSslCertificate](docs/ModelsSslCertificate.md)
 - [ModelsSslInfo](docs/ModelsSslInfo.md)
 - [ModelsStat](docs/ModelsStat.md)
 - [ModelsSubAccount](docs/ModelsSubAccount.md)
 - [ModelsSubAccountType](docs/ModelsSubAccountType.md)
 - [ModelsSuppression](docs/ModelsSuppression.md)
 - [ModelsSuppressionEmail](docs/ModelsSuppressionEmail.md)
 - [ModelsSuppressionReason](docs/ModelsSuppressionReason.md)
 - [ModelsSystemDnsRecord](docs/ModelsSystemDnsRecord.md)
 - [ModelsSystemDomain](docs/ModelsSystemDomain.md)
 - [ModelsSystemIpPool](docs/ModelsSystemIpPool.md)
 - [ModelsSystemTemplate](docs/ModelsSystemTemplate.md)
 - [ModelsTag](docs/ModelsTag.md)
 - [ModelsTo](docs/ModelsTo.md)
 - [ModelsValidation](docs/ModelsValidation.md)
 - [ModelsValidationReason](docs/ModelsValidationReason.md)
 - [ModelsVerifyByMemberTokenRequest](docs/ModelsVerifyByMemberTokenRequest.md)
 - [ModelsVerifyByTokenRequest](docs/ModelsVerifyByTokenRequest.md)
 - [ModelsWMessage](docs/ModelsWMessage.md)
 - [Provider](docs/Provider.md)
 - [UaparserDevice](docs/UaparserDevice.md)
 - [UaparserOs](docs/UaparserOs.md)
 - [UaparserUserAgent](docs/UaparserUserAgent.md)


## Documentation For Authorization
 Endpoints do not require authorization.


## Author

hello@sendpost.io

