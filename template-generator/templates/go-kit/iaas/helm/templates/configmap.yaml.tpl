apiVersion: v1
kind: ConfigMap
metadata:
    name: {{.Values.name}}-cfg
    namespace: {{ .Values.namespace }}
data:
    PORT: "{{ .Values.config.port }}"


