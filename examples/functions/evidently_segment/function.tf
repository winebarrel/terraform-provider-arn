# arn:aws:evidently:ap-northeast-1:111111111111:segment/segment-name
output "evidently_segment" {
  value = provider::arn::evidently_segment("segment-name")
}
