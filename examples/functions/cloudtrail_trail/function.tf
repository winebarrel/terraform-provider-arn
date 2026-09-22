# arn:aws:cloudtrail:ap-northeast-1:111111111111:trail/trail-name
output "cloudtrail_trail" {
  value = provider::arn::cloudtrail_trail("trail-name")
}
