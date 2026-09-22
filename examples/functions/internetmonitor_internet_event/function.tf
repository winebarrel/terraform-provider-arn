# arn:aws:internetmonitor::111111111111:internet-event/internet-event-id
output "internetmonitor_internet_event" {
  value = provider::arn::internetmonitor_internet_event("internet-event-id")
}
