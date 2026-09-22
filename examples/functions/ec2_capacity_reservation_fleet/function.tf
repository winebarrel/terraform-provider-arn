# arn:aws:ec2:ap-northeast-1:111111111111:capacity-reservation-fleet/capacity-reservation-fleet-id
output "ec2_capacity_reservation_fleet" {
  value = provider::arn::ec2_capacity_reservation_fleet("capacity-reservation-fleet-id")
}
