# Chaotic App

A tiy application based on [fiber](https://github.com/gofiber/fiber) to experiment with [Prometheus Blackbox Exporter]()

This application exposes 3 endpoints:

- `/hello`: always returns a `200 OK` response with a nice "Hello, World 👋!" message
- `/_heatlh/ready`: always returns a `200 OK` response
- `/_heatlh/alive`: randomly returns a `200 OK` response or a `500 Internal Server Error` response

## Building 

Run the following command (requires [task](https://taskfile.dev/)) to run the application locally:

```
task run
```

Run the following command (requires [task](https://taskfile.dev/)) to build and push to quay.io (make sure you update the repository when building and pushing the image, first!):

```
task push-to-quay
```

## Deploying on OpenShift

Run the following commands to deploy the application:
```
oc create ns app
oc apply -f scripts/chaotic-app.yaml
```

Run the following commands to deploy Backbox Exporter:
```
oc create ns blackbox-exporter
oc apply -f scripts/blackbox-exporter.yaml
```

Run the following commands to enable User Workload Monitoring:
```
oc apply -f scripts/openshift-user-workload-monitoring.yaml
```

## License

chaotic-app is free and open-source software licensed under the [MIT License](LICENSE).