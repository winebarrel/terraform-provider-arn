// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: iq
// Source: https://servicereference.us-east-1.amazonaws.com/v1/iq/iq.json
// Functions: 14
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "iq_attachment", Service: "iq", Resource: "attachment", Template: "arn:${Partition}:iq:${Region}::attachment/${AttachmentId}"},
		{Name: "iq_buyer", Service: "iq", Resource: "buyer", Template: "arn:${Partition}:iq:${Region}::buyer/${BuyerId}"},
		{Name: "iq_call", Service: "iq", Resource: "call", Template: "arn:${Partition}:iq:${Region}::call/${CallId}"},
		{Name: "iq_company", Service: "iq", Resource: "company", Template: "arn:${Partition}:iq:${Region}::company/${CompanyId}"},
		{Name: "iq_conversation", Service: "iq", Resource: "conversation", Template: "arn:${Partition}:iq:${Region}::conversation/${ConversationId}"},
		{Name: "iq_expert", Service: "iq", Resource: "expert", Template: "arn:${Partition}:iq:${Region}::expert/${ExpertId}"},
		{Name: "iq_listing", Service: "iq", Resource: "listing", Template: "arn:${Partition}:iq:${Region}::listing/${ListingId}"},
		{Name: "iq_payment_request", Service: "iq", Resource: "paymentRequest", Template: "arn:${Partition}:iq:${Region}::paymentRequest/${ConversationId}/${ProposalId}/${PaymentRequestId}"},
		{Name: "iq_payment_schedule", Service: "iq", Resource: "paymentSchedule", Template: "arn:${Partition}:iq:${Region}::paymentSchedule/${ConversationId}/${ProposalId}/${VersionId}"},
		{Name: "iq_permission", Service: "iq", Resource: "permission", Template: "arn:${Partition}:iq-permission:${Region}::permission/${PermissionRequestId}"},
		{Name: "iq_proposal", Service: "iq", Resource: "proposal", Template: "arn:${Partition}:iq:${Region}::proposal/${ConversationId}/${ProposalId}"},
		{Name: "iq_request", Service: "iq", Resource: "request", Template: "arn:${Partition}:iq:${Region}::request/${RequestId}"},
		{Name: "iq_seller", Service: "iq", Resource: "seller", Template: "arn:${Partition}:iq:${Region}::seller/${SellerAwsAccountId}"},
		{Name: "iq_token", Service: "iq", Resource: "token", Template: "arn:${Partition}:iq:${Region}::token/${TokenId}"},
	})
}
