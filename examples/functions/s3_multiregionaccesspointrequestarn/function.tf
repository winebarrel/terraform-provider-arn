# arn:aws:s3:us-west-2:111111111111:async-request/mrap/operation/token
output "s3_multiregionaccesspointrequestarn" {
  value = provider::arn::s3_multiregionaccesspointrequestarn("operation", "token")
}
