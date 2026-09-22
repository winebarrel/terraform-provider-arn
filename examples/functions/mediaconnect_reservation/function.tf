# arn:aws:mediaconnect:ap-northeast-1:111111111111:reservation:reservation-id:reservation-name
output "mediaconnect_reservation" {
  value = provider::arn::mediaconnect_reservation("reservation-id", "reservation-name")
}
