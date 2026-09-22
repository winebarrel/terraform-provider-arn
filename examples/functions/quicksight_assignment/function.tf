# arn:aws:quicksight::111111111111:assignment/resource-id
output "quicksight_assignment" {
  value = provider::arn::quicksight_assignment("resource-id")
}
