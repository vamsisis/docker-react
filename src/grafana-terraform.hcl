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


======================================================================

resource "grafana_rule_group" "rule_group_0000" {
  org_id           = 1
  name             = "Target_response_time"
  folder_uid       = "be5v8t186bzswe"
  interval_seconds = 10

  rule {
    name      = "Target Response time for app/rps-rem-services-svclb/6f638396b8b2d714"
    condition = "C"

    data {
      ref_id = "A"

      relative_time_range {
        from = 300
        to   = 0
      }

      datasource_uid = "JXIxxq6Iz"
      model          = "{\"dimensions\":{\"LoadBalancer\":\"app/rps-rem-services-svclb6f638396b8b2d714\"},\"expression\":\"\",\"id\":\"\",\"intervalMs\":1000,\"label\":\"\",\"logGroups\":[],\"matchExact\":true,\"maxDataPoints\":43200,\"metricEditorMode\":0,\"metricName\":\"TargetResponseTime\",\"metricQueryType\":0,\"namespace\":\"AWS/ApplicationELB\",\"period\":\"\",\"queryMode\":\"Metrics\",\"refId\":\"A\",\"region\":\"default\",\"sqlExpression\":\"\",\"statistic\":\"Average\"}"
    }

    data {
      ref_id = "B"

      relative_time_range {
        from = 300
        to   = 0
      }

      datasource_uid = "__expr__"
      model          = "{\"conditions\":[{\"evaluator\":{\"params\":[],\"type\":\"gt\"},\"operator\":{\"type\":\"and\"},\"query\":{\"params\":[\"B\"]},\"reducer\":{\"params\":[],\"type\":\"last\"},\"type\":\"query\"}],\"datasource\":{\"type\":\"__expr__\",\"uid\":\"__expr__\"},\"expression\":\"A\",\"intervalMs\":1000,\"maxDataPoints\":43200,\"reducer\":\"last\",\"refId\":\"B\",\"type\":\"reduce\"}"
    }

    data {
      ref_id = "C"

      relative_time_range {
        from = 300
        to   = 0
      }

      datasource_uid = "__expr__"
      model          = "{\"conditions\":[{\"evaluator\":{\"params\":[0.4],\"type\":\"gt\"},\"operator\":{\"type\":\"and\"},\"query\":{\"params\":[\"C\"]},\"reducer\":{\"params\":[],\"type\":\"last\"},\"type\":\"query\"}],\"datasource\":{\"type\":\"__expr__\",\"uid\":\"__expr__\"},\"expression\":\"B\",\"intervalMs\":1000,\"maxDataPoints\":43200,\"refId\":\"C\",\"type\":\"threshold\"}"
    }

    no_data_state  = "NoData"
    exec_err_state = "Error"
    for            = "1m"

    annotations = {
      description = "High target response time detected"
      runbook_url = "https://runbook.example.com/high-response-time"
      summary     = "ALB Target Response Time is too high"
    }

    labels = {
      alert_name = "High Target Response Time"
    }

    is_paused = false

    notification_settings {
      receiver            = "your-notification-receiver"
      group_by            = null
      mute_time_intervals = []
    }
  }
}

