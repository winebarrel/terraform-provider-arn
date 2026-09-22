# arn:aws:controltower:ap-northeast-1:111111111111:landingzone/landing-zone-id
output "controltower_landing_zone" {
  value = provider::arn::controltower_landing_zone("landing-zone-id")
}
