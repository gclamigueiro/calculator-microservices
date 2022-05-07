apiVersion: v1
kind: Service
metadata:
  labels:
    app: {{ .Values.name }}
    run:  {{ .Values.name }}
    app.kubernetes.io/name: {{.Values.name}}
    app.kubernetes.io/version: "1.0.0"
    app.kubernetes.io/component: services
    app.kubernetes.io/part-of: calculator
  name: {{.Values.name}}-svc
  namespace: {{ .Values.namespace }} 
spec:
  ports:
    - name: http
      port: {{ .Values.config.port }}
      protocol: TCP
      targetPort: {{ .Values.config.port }}
  selector:
    app: {{ .Values.name }}
    run:  {{ .Values.name }}
    app.kubernetes.io/name: {{.Values.name}}
    app.kubernetes.io/version: "1.0.0"
    app.kubernetes.io/component: services
    app.kubernetes.io/part-of: calculator