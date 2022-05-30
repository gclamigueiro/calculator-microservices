
apiVersion: apps/v1
kind: Deployment
metadata:
  name: {{ .Values.name }}-deploy
  namespace: {{ .Values.namespace }} 
  labels:
    app: {{ .Values.name }}   
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
        app: {{ .Values.name }}
    spec:
      containers:
      - name: {{ .Values.name }}
        image: {{ .Values.image.name }}:{{ .Values.image.tag }}
        imagePullPolicy: IfNotPresent
        ports:
        - containerPort: {{ .Values.config.port }}
        resources:
          requests:
            cpu: 50m
            memory: 256Mi