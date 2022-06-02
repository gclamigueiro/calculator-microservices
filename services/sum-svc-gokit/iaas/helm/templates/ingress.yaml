apiVersion: networking.k8s.io/v1
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
            pathType: Prefix
            backend:
              service:
                name: {{ .Values.name }}-svc
                port:
                  number: {{ .Values.config.port }}


