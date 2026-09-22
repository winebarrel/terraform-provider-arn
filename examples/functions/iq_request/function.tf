# arn:aws:iq:ap-northeast-1::request/request-id
output "iq_request" {
  value = provider::arn::iq_request("request-id")
}
