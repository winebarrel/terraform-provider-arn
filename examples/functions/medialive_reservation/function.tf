# arn:aws:medialive:ap-northeast-1:111111111111:reservation:reservation-id
output "medialive_reservation" {
  value = provider::arn::medialive_reservation("reservation-id")
}
