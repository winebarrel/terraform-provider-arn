# arn:aws:ec2:ap-northeast-1:111111111111:capacity-reservation-cancellation-quote/capacity-reservation-cancellation-quote-id
output "ec2_capacity_reservation_cancellation_quote" {
  value = provider::arn::ec2_capacity_reservation_cancellation_quote("capacity-reservation-cancellation-quote-id")
}
