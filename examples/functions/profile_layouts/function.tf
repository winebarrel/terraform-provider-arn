# arn:aws:profile:ap-northeast-1:111111111111:domains/domain-name/layouts/layout-definition-name
output "profile_layouts" {
  value = provider::arn::profile_layouts("domain-name", "layout-definition-name")
}
