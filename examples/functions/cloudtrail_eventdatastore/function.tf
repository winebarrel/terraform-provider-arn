# arn:aws:cloudtrail:ap-northeast-1:111111111111:eventdatastore/event-data-store-id
output "cloudtrail_eventdatastore" {
  value = provider::arn::cloudtrail_eventdatastore("event-data-store-id")
}
