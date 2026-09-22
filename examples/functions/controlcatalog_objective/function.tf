# arn:aws:controlcatalog:::objective/objective-id
output "controlcatalog_objective" {
  value = provider::arn::controlcatalog_objective("objective-id")
}
