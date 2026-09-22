# arn:aws:athena:ap-northeast-1:111111111111:capacity-reservation/capacity-reservation-name
output "athena_capacity_reservation" {
  value = provider::arn::athena_capacity_reservation("capacity-reservation-name")
}
