# arn:aws:memorydb:ap-northeast-1:111111111111:reservednode/reservation-id
output "memorydb_reservednode" {
  value = provider::arn::memorydb_reservednode("reservation-id")
}
