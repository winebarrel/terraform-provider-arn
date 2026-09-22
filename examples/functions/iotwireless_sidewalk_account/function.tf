# arn:aws:iotwireless:ap-northeast-1:111111111111:SidewalkAccount/sidewalk-account-id
output "iotwireless_sidewalk_account" {
  value = provider::arn::iotwireless_sidewalk_account("sidewalk-account-id")
}
