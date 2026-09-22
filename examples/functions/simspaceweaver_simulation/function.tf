# arn:aws:simspaceweaver:ap-northeast-1:111111111111:simulation/simulation-name
output "simspaceweaver_simulation" {
  value = provider::arn::simspaceweaver_simulation("simulation-name")
}
