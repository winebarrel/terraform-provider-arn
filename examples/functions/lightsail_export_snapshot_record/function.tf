# arn:aws:lightsail:ap-northeast-1:111111111111:ExportSnapshotRecord/id
output "lightsail_export_snapshot_record" {
  value = provider::arn::lightsail_export_snapshot_record("id")
}
