# arn:aws:interconnect:ap-northeast-1:111111111111:connection/id
output "interconnect_connection" {
  value = provider::arn::interconnect_connection("id")
}
