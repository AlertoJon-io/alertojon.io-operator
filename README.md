# WIP ⚠️ 🚧 🏗 
This repo is currentlly in early development stages and for sure will change. 
***dont use this yet***.
<br/>

[![Artifact Hub](https://img.shields.io/endpoint?url=https://artifacthub.io/badge/repository/alertojon-operator)](https://artifacthub.io/packages/search?repo=alertojon-operator)
![Make Install](https://github.com//AlertoJon-io/alertojon.io-operator/actions/workflows/make-install.yaml/badge.svg) 
![Build image](https://github.com//AlertoJon-io/alertojon.io-operator/actions/workflows/build-image.yaml/badge.svg)

## alertojon.io-operator

![alt text](https://drive.google.com/uc?id=1qa2i_2pggjIookzToHlNDU_HNr453YMd) 


The aim of alertojon.io is to provide abstraction of incident management 🚨 services like pagerduty and opsgenie by leveraging k8s operators.

This repo is the main operator that will get you going with you favorite incident management provider.

CR example :
```yaml
apiVersion: alertojon.io/v1
kind: IncidentManagementProvider
metadata:
  name: incidentmanagementprovider-sample
spec:
  type: Pagerduty
  apikey: "YOUR_API_KEY" // will be depricated in the future
  apiKeySecretRef:
    secretName: pagerduty-api-key-secret
    secretKey: "YOUR_SERVICE_KEY"
```

This IncidentManagementProvider will be used to create a new alertojon.io Pagerduty provider.

Read about alertojon.io [pagerduty provider](https://github.com/AlertoJon-io/alertojon.io-pagerduty-operator)

Documentation :
[AlertoJon.io](https:://alertojon.io)