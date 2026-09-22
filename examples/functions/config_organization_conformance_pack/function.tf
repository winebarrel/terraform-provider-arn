# arn:aws:config:ap-northeast-1:111111111111:organization-conformance-pack/organization-conformance-pack-id
output "config_organization_conformance_pack" {
  value = provider::arn::config_organization_conformance_pack("organization-conformance-pack-id")
}
