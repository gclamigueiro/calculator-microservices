apiVersion: extensions/v1beta1
kind: Ingress
metadata:
  name: {{ .Values.name }}-ing
  annotations:
    ingress.kubernetes.io/rewrite-target: /

spec:
  rules:
    - host: localhost
      http:
        paths:
          - path: {{ .Values.path }}
            backend:
              serviceName: {{ .Values.name }}-v1-svc
              servicePort: {{ .Values.config.port }}


