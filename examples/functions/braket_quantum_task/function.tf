# arn:aws:braket:ap-northeast-1:111111111111:quantum-task/random-id
output "braket_quantum_task" {
  value = provider::arn::braket_quantum_task("random-id")
}
