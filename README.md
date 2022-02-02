# WIP ⚠️ 🚧 🏗 
This repo is currentlly in early development stages and for sure will change. 
***dont use this yet***.
<br/>

[![Artifact Hub](https://img.shields.io/endpoint?url=https://artifacthub.io/badge/repository/alertojon-operator)](https://artifacthub.io/packages/search?repo=alertojon-operator)
![Make Install](https://github.com//AlertoJon-io/alertojon.io-operator/actions/workflows/make-install.yaml/badge.svg) 
![Build image](https://github.com//AlertoJon-io/alertojon.io-operator/actions/workflows/build-image.yaml/badge.svg)

## alertojon.io-operator

The aim of alertojon.io is to provide abstraction to alert 🚨 management Third-party services like pagerduty and opsgenie by leveraging k8s operators.

This repo is the main operator that will get you going with you favorit provider.

CR example :
```yaml
apiVersion: alertojon.io/v1
kind: Pagerduty
metadata:
  name: pagerduty-sample
spec:
  apiToken: ""
```

Apply this yaml to the cluster and a pagerduty provider will init.

Documentation :
[AlertoJon.io](https:://alertojon.io)