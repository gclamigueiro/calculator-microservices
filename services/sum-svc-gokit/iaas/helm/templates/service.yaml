apiVersion: v1
kind: Service
metadata:
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