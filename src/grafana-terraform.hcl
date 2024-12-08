resource "grafana_folder" "target_response_time_folder" {
  title = "Target Response Time"
}

resource "grafana_rule_group" "target_response_time_group" {
  name             = "Target Response Time"
  folder_uid       = grafana_folder.target_response_time_folder.uid
  interval_seconds = 240
  org_id           = 1

  rule {
    name           = "Target Response Time for Load Balancer"
    for            = "2m"
    condition      = "B"
    no_data_state  = "NoData"
    exec_err_state = "Alerting"

    annotations = {
      "severity" = "critical"
      "service"  = "LoadBalancer"
    }

    labels = {
      "team"   = "SRE"
      "region" = "us-east-1"
    }

    is_paused = false

    # Data source: AWS CloudWatch for ALB metrics
    data {
      ref_id     = "A"
      query_type = ""
      relative_time_range {
        from = 600
        to   = 0
      }
      datasource_uid = "PD8C576611E62080A" # Replace with your AWS CloudWatch data source UID
      model = jsonencode({
        datasource = {
          "type" = "cloudwatch"
          "uid"  = "PD8C576611E62080A"
        }
        refId         = "A"
        hide          = false
        intervalMs    = 1000
        maxDataPoints = 43200
        namespace     = "AWS/ApplicationELB"
        metricName    = "TargetResponseTime"
        dimensions    = {
          "LoadBalancer" = "your-load-balancer-arn" # Replace with the specific Load Balancer ARN or Name
        }
        statistic = "Average"
        period    = 60
      })
    }

    data {
      ref_id     = "B"
      query_type = ""
      relative_time_range {
        from = 0
        to   = 0
      }
      datasource_uid = "-100"
      model = <<EOT
{
    "conditions": [
        {
        "evaluator": {
            "params": [
            0.5
            ],
            "type": "gt"
        },
        "operator": {
            "type": "and"
        },
        "query": {
            "params": [
            "A"
            ]
        },
        "reducer": {
            "params": [],
            "type": "last"
        },
        "type": "query"
        }
    ],
    "datasource": {
        "type": "__expr__",
        "uid": "-100"
    },
    "hide": false,
    "intervalMs": 1000,
    "maxDataPoints": 43200,
    "refId": "B",
    "type": "classic_conditions"
}
EOT
    }
  }
}




