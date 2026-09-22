# arn:aws:pipes:ap-northeast-1:111111111111:pipe/name
output "pipes_pipe" {
  value = provider::arn::pipes_pipe("name")
}
