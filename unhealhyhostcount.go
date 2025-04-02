{{ define "CustomBodyForAlerting" }}

{{ if gt (len .Alerts.Firing) 0 }}
**🚨 Firing Alerts 🚨**

{{ range .Alerts.Firing }}
**Alert Name:** {{ .Labels.alert_name }}

{{ if eq .Labels.alert_name "unhealthyhost" }}

**Load Balancer:** {{ .Labels.LoadBalancer }}

{{ range $tg, $value := .Values }}
  {{ if hasPrefix $tg "TargetGroup" }}
  - **Target Group:** {{ $tg }}
    - **Unhealthy Host Count:** {{ $value }}
  {{ end }}
{{ end }}

{{ end }}
{{ end }}
{{ end }}

{{ if gt (len .Alerts.Resolved) 0 }}
**✅ Resolved Alerts**

{{ range .Alerts.Resolved }}
**Alert Name:** {{ .Labels.alert_name }}
**Load Balancer:** {{ .Labels.LoadBalancer }}

{{ range $tg, $value := .Values }}
  {{ if hasPrefix $tg "TargetGroup" }}
  - **Target Group:** {{ $tg }}
    - **Unhealthy Host Count (Resolved):** {{ $value }}
  {{ end }}
{{ end }}

{{ end }}
{{ end }}

{{ end }}
