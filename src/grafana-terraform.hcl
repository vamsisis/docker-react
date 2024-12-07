# Define the Grafana provider
provider "grafana" {
  url  = "https://<YOUR-AMG-ENDPOINT>" # Replace with your AMG endpoint
  auth = "<YOUR-API-TOKEN>"           # Replace with an API token generated in AMG
}

# Create a Grafana alert rule
resource "grafana_alert_rule" "elb_target_response_alert" {
  name        = "ELB_Target_Response_Alert"
  folder_uid  = "default" # Default folder in Grafana
  interval    = "1m"      # Alert evaluation interval

  data {
    ref_id = "A"
    datasource_uid = "<DATASOURCE-UID>" # CloudWatch datasource UID in AMG
    model = jsonencode({
      "datasource" : { "type" : "cloudwatch", "uid" : "<DATASOURCE-UID>" },
      "namespace"  : "AWS/ApplicationELB",
      "metricName" : "TargetResponseTime",
      "stat"       : "Average",
      "period"     : 60,
      "dimensions" : {
        "LoadBalancer" : "<YOUR-LOAD-BALANCER-NAME>"
      }
    })
  }

  condition {
    evaluator {
      type  = "gt"
      value = [1] # Threshold: Target Response Time > 1 second
    }
    operator = "and"
    reducer  = "avg"
    query    = "A"
  }

  notifications {
    uid = "<NOTIFICATION-CHANNEL-UID>" # Replace with SNS or email channel UID
  }
}
