# arn:aws:ec2:ap-northeast-1:111111111111:spot-fleet-request/spot-fleet-request-id
output "ec2_spot_fleet_request" {
  value = provider::arn::ec2_spot_fleet_request("spot-fleet-request-id")
}
