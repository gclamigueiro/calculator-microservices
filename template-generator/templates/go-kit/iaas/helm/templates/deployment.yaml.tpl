
apiVersion: apps/v1
kind: Deployment
metadata:
  name: {{ .Values.name }}-deploy
  namespace: {{ .Values.namespace }}  
spec:
  replicas: {{ .Values.replicaCount }}
  selector:
    matchLabels:
      app: {{ .Values.name }}
  strategy:
    rollingUpdate:
      maxSurge: 25%
      maxUnavailable: 25%
    type: RollingUpdate
  template:
    metadata:
      labels:
        app.kubernetes.io/name: {{ .Values.name }}
        app: {{ .Values.name }}
    spec:
      containers:
      - image: gcamps/{{ .Values.name }}:latest
        imagePullPolicy: IfNotPresent
        name: {{ .Values.name }}
        resources:
          requests:
            cpu: 50m
            memory: 256Mi