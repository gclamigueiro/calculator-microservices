apiVersion: extensions/v1beta1
kind: Ingress
metadata:
  name: {{ .Values.name }}-ing
  annotations:
    ingress.kubernetes.io/rewrite-target: /

spec:
  rules:
    - http:
        paths:
          - path: {{ .Values.ingress.path }}
            backend:
              serviceName: {{ .Values.name }}-svc
              servicePort: {{ .Values.config.port }}


