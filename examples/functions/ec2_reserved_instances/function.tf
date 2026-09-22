# arn:aws:ec2:ap-northeast-1:111111111111:reserved-instances/reservation-id
output "ec2_reserved_instances" {
  value = provider::arn::ec2_reserved_instances("reservation-id")
}
