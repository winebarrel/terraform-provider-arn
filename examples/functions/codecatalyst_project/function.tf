# arn:aws:codecatalyst:::space/space-id/project/project-id
output "codecatalyst_project" {
  value = provider::arn::codecatalyst_project("space-id", "project-id")
}
