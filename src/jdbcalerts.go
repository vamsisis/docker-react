{{ define "custom_title" }}
  {{ if gt (len .Alerts.Firing) 0 }}🚨 Firing alert(s): JDBC Connection Usage Alert: {{ .CommonLabels.alert_name }} - {{ .CommonLabels.job }} - {{ .CommonLabels.value }}%{{ end }}
  {{ if gt (len .Alerts.Resolved) 0 }}✅ Resolved alert(s): JDBC Connection Usage Alert: {{ .CommonLabels.alert_name }} - {{ .CommonLabels.job }} - {{ .CommonLabels.value }}%{{ end }}
{{ end }}
