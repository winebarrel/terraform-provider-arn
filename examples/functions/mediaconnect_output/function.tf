# arn:aws:mediaconnect:ap-northeast-1:111111111111:output:output-id:output-name
output "mediaconnect_output" {
  value = provider::arn::mediaconnect_output("output-id", "output-name")
}
