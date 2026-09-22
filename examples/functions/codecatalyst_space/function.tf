# arn:aws:codecatalyst:::space/space-id
output "codecatalyst_space" {
  value = provider::arn::codecatalyst_space("space-id")
}
