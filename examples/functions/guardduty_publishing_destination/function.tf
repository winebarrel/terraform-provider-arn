# arn:aws:guardduty:ap-northeast-1:111111111111:detector/detector-id/publishingdestination/publishing-destination-id
output "guardduty_publishing_destination" {
  value = provider::arn::guardduty_publishing_destination("detector-id", "publishing-destination-id")
}
