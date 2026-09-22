# arn:aws:braket:ap-northeast-1:111111111111:job/random-id
output "braket_job" {
  value = provider::arn::braket_job("random-id")
}
