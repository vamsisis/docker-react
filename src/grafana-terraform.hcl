{{ define "CustomBodyfor5XXAlerting" -}}
{{ if gt (len .Alerts.Firing) 0 -}}
**🚨 Firing 5XX Alerts 🚨**

{{ range .Alerts.Firing -}}
{{ if eq .Labels.alert_name "High 5XX Errors" -}}
**Alert Details:**
- **Load Balancer**: {{ .Labels.LoadBalancer }}
- **Current 5XX Error Rate**: 🚨 {{ index .Values "C" }}% (Above Threshold)
{{ end -}}
{{ end -}}

{{ else if gt (len .Alerts.Resolved) 0 -}}
**✅ Resolved 5XX Alerts ✅**

{{ range .Alerts.Resolved -}}
{{ if eq .Labels.alert_name "High 5XX Errors" -}}
**Alert Details:**
- **Load Balancer**: {{ .Labels.LoadBalancer }}
- **Resolved 5XX Error Rate**: {{ index .Values "C" }}%
- **Status**: Resolved ✅
{{ end -}}
{{ end -}}

{{ else -}}
No active 5XX alerts at the moment.
{{ end -}}
{{ end -}}
