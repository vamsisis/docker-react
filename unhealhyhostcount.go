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
{{ define "CustomBodyfor5XXAlerting" -}}

{{ if gt (len .Alerts.Firing) 0 }}
**Alert: 5XX Error Rate Above Threshold**

{{ range .Alerts.Firing }}
**Summary**: 5XX error rate for Load Balancer `{{ .Labels.LoadBalancer }}` exceeded the threshold of 10%.

**Description**:
- **Alert Name**: {{ .Labels.alert_name }}
- **Load Balancer**: {{ .Labels.LoadBalancer }}
- **Total 5XX Errors**: {{ with index .Values.B 0 }}{{ . }}{{ end }}
- **Total Requests**: {{ with index .Values.C 0 }}{{ . }}{{ end }}
- **Current 5XX Error Rate**: {{ with index .Values.E 0 }}{{ printf "%.2f%%" . }}{{ end }}
- **Threshold**: >10%
{{ end }}

{{ else if gt (len .Alerts.Resolved) 0 }}
**Resolved Alert: 5XX Error Rate**

{{ range .Alerts.Resolved }}
**Summary**: 5XX error rate for Load Balancer `{{ .Labels.LoadBalancer }}` is back to normal.

**Description**:
- **Alert Name**: {{ .Labels.alert_name }}
- **Load Balancer**: {{ .Labels.LoadBalancer }}
{{ end }}

{{ end }}
{{ end }}


