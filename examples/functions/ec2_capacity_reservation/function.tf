# arn:aws:ec2:ap-northeast-1:111111111111:capacity-reservation/capacity-reservation-id
output "ec2_capacity_reservation" {
  value = provider::arn::ec2_capacity_reservation("capacity-reservation-id")
}
