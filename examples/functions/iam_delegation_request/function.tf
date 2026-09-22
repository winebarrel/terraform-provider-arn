# arn:aws:iam::111111111111:delegation-request/delegation-request-id
output "iam_delegation_request" {
  value = provider::arn::iam_delegation_request("delegation-request-id")
}
