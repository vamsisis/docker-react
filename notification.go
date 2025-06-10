{{ define "CustomBodyfor5XXAlerting" -}}
{{ if gt (len .Alerts.Firing) 0 -}}
{{ range .Alerts.Firing -}}
{{ if or (eq .Labels.alert_name "unhealthyhost") (eq .Labels.alert_name "4XX Error Rate") (eq .Labels.alert_name "High Target Response Time") (eq .Labels.alert_name "5XX Error Rate") (eq .Labels.alert_name "ApiGateway_5XX") (eq .Labels.alert_name "rds_cpu_alerts") -}}

{{ if eq .Labels.alert_name "unhealthyhost" -}}
Alert Name: PROD::UNHEALTHY TARGET GROUP ALERT
Description:
- Load Balancer: {{ .Labels.LoadBalancer }}
- Current Unhealthy Host Count: {{ .Labels.unhealthy_host_count }}

{{ else if eq .Labels.alert_name "5XX Error Rate" }}
Alert Name: {{ .Labels.alert_name }} for {{ .Labels.LoadBalancer }}
Description:
- Load Balancer: {{ .Labels.LoadBalancer }}
- Total 5XX Errors: {{ .Labels.errors }}
- Total Requests: {{ .Labels.total_requests }}
- Total 5XX Errors: {{ .Labels.error_rate }}% 
- Threshold: >10%

{{ else if eq .Labels.alert_name "4XX Error Rate" }}
Alert Name: {{ .Labels.alert_name }} for {{ .Labels.LoadBalancer }}
Description:
- Load Balancer: {{ .Labels.LoadBalancer }}
- Total 4XX Errors: {{ .Labels.errors }}
- Total Requests: {{ .Labels.total_requests }}
- Total 4XX Errors: {{ .Labels.error_rate }}% 
- Threshold: >10%

{{ else if eq .Labels.alert_name "rds_cpu_alerts" }}
Alert Name: High CPU Utiliziation  for {{ .Labels.DBClusterIdentifier }}
Description:
- DB Cluster: {{ .Labels.DBClusterIdentifier }}
- Role: {{ .Labels.Role }}
- CPU Utiliziation: {{ .Labels.cpu_utilization }}
- Threshold: >70%

{{ else if eq .Labels.alert_name "ApiGateway_5XX" }}
Alert Name: {{ .Labels.alert_name }} for {{ .Labels.ApiName }}
Description:
- API: {{ .Labels.ApiName }}
- HTTP Method: {{ .Labels.Method }}
- Resource Path: {{ .Labels.Resource }}

Metrics
- Total 5XX Errors: {{ .Labels.errors }}
- Total Requests: {{ .Labels.total_requests }}
- Total 5XX Errors: {{ .Labels.error_rate }}% 
- Threshold: >10%

{{ end -}}
{{ end -}}
{{ end -}}

{{ else if gt (len .Alerts.Resolved) 0 -}}
**Resolved Alerts**
{{ range .Alerts.Resolved -}}
{{ if or (eq .Labels.alert_name "unhealthyhost") (eq .Labels.alert_name "High Request Count") (eq .Labels.alert_name "High Target Response Time") (eq .Labels.alert_name "5XX Error Rate") -}}
**Alert Details:**
- **Alert Name**: {{ .Labels.alert_name }}

{{ if eq .Labels.alert_name "unhealthyhost" -}}
- **Load Balancer**: {{ .Labels.LoadBalancer }}
- **Target Group**: {{ .Labels.TargetGroup }}
- **Resolved Unhealthy Host Count**: {{ range $refID, $value := .Values -}}{{ $value }}{{ end }}

{{ else if eq .Labels.alert_name "5XX Error Rate" }}
- **Load Balancer**: {{ .Labels.LoadBalancer }}
- **Total 5XX Errors**: {{ .Labels.error_rate }}% 

{{ end -}}
{{ end -}}
{{ end -}}

{{ else -}}
No active alerts.
{{ end -}}
{{ end -}}
