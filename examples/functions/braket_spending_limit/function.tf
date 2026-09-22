# arn:aws:braket:ap-northeast-1:111111111111:spending-limit/random-id
output "braket_spending_limit" {
  value = provider::arn::braket_spending_limit("random-id")
}
