{{ define "CustomBodyForAlerting" }}

{{ if gt (len .Alerts.Firing) 0 }}
**🚨 Firing Alerts 🚨**

{{ range .Alerts.Firing }}
**Alert Name:** {{ .Labels.alert_name }}

{{ if eq .Labels.alert_name "unhealthyhost" }}

- **Load Balancer:** {{ .Labels.LoadBalancer }}

  {{ range $k, $v := .Values }}
  - **Target Group / Metric Key:** {{ $k }}
    - **Unhealthy Host Count:** {{ $v }}
  {{ end }}

{{ end }}
{{ end }}
{{ end }}

{{ if gt (len .Alerts.Resolved) 0 }}
**✅ Resolved Alerts**

{{ range .Alerts.Resolved }}
**Alert Name:** {{ .Labels.alert_name }}

{{ if eq .Labels.alert_name "unhealthyhost" }}

- **Load Balancer:** {{ .Labels.LoadBalancer }}

  {{ range $k, $v := .Values }}
  - **Target Group / Metric Key:** {{ $k }}
    - **Resolved Unhealthy Host Count:** {{ $v }}
  {{ end }}

{{ end }}
{{ end }}
{{ end }}

{{ end }}


=============================================================
{{{ define "CustomBodyfor5XXAlerting" -}}
{{ if gt (len .Alerts.Firing) 0 }}
**Alert: 5XX Error Rate Above Threshold**

**Summary**: 5XX error rate for Load Balancer `{{ .Labels.LoadBalancer }}` exceeded the threshold of 10%.

**Description**:
- **Alert Name**: {{ .Labels.alert_name }}
- **Load Balancer**: {{ .Labels.LoadBalancer }}
- **Total 5XX Errors**: {{ with index .EvalMatches 0 }}{{ .Value }}{{ end }}
- **Total Requests**: {{ with index .EvalMatches 1 }}{{ .Value }}{{ end }}
- **Current 5XX Error Rate**: {{ with index .EvalMatches 2 }}{{ printf "%.2f%%" .Value }}{{ end }}
- **Threshold**: >10%

{{ else if gt (len .Alerts.Resolved) 0 }}
**Resolved Alert: 5XX Error Rate**

**Summary**: 5XX error rate for Load Balancer `{{ .Labels.LoadBalancer }}` has dropped below the threshold.

**Description**:
- **Alert Name**: {{ .Labels.alert_name }}
- **Load Balancer**: {{ .Labels.LoadBalancer }}
{{ end }}
{{ end }}


