// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: lookoutmetrics
// Source: https://servicereference.us-east-1.amazonaws.com/v1/lookoutmetrics/lookoutmetrics.json
// Functions: 3
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "lookoutmetrics_alert", Service: "lookoutmetrics", Resource: "Alert", Template: "arn:${Partition}:lookoutmetrics:${Region}:${Account}:Alert:${AlertName}"},
		{Name: "lookoutmetrics_anomaly_detector", Service: "lookoutmetrics", Resource: "AnomalyDetector", Template: "arn:${Partition}:lookoutmetrics:${Region}:${Account}:AnomalyDetector:${AnomalyDetectorName}"},
		{Name: "lookoutmetrics_metric_set", Service: "lookoutmetrics", Resource: "MetricSet", Template: "arn:${Partition}:lookoutmetrics:${Region}:${Account}:MetricSet/${AnomalyDetectorName}/${MetricSetName}"},
	})
}
