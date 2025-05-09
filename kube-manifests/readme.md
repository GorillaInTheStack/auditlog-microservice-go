# Cluster testing


if you would like to test this microservice in a cluster and observe the usage of a NOSQL DB, use
   
```bash
kubectl -f apply auditlog-deployment.yaml
```
    
This file will create the namespace demo and deploy all the necssary pods and services.
You can check everything with
    
```bash 
kubectl get pods,svc -n demo
```

In another terminal, you can run 
    
```bash
kubectl logs -n demo deployment.apps/auditlog-deployment -f
```

to see the microservice logs.

In another terminal, you need to port forward from host to pod using

```bash
kubectl port-forward deployment.apps/auditlog-deployment 6969:6969 -n demo
```

Now you are ready to test the application.

## Tests

### SUBMIT EVENTS

    Event 1:

```bash
TOKEN=$(curl -s  localhost:6969/generatetoken); curl -X POST -H "Content-Type: application/json" -H "Authorization: ${TOKEN}" -d '{ "SourceEventID": "987654", "SourceTimestamp": "2023-06-11T08:30:00Z", "CorrelationID": "a1b2c3d4", "SourceTimezone": "America/Los_Angeles", "SourceServiceName": "audit-service", "SourceServiceLocation": "San Francisco, USA", "SourceIpAddress": "192.168.0.1", "EventTags": { "env": "production", "category": "login" }, "EventDataHash": "1234567890abcdef", "EventDataVersion": "2.0", "EventData": { "username": "john_doe", "status": "success", "login_time": "2023-06-11T08:30:00Z", "ip_address": "192.168.0.100" } }' http://localhost:6969/events/submit;
```


    Event 2:

```bash
TOKEN=$(curl -s  localhost:6969/generatetoken); curl -X POST -H "Content-Type: application/json" -H "Authorization: ${TOKEN}" -d '{ "SourceEventID": "123456", "SourceTimestamp": "2023-06-12T12:00:00Z", "CorrelationID": "x1y2z3", "SourceTimezone": "Europe/London", "SourceServiceName": "audit-service", "SourceServiceLocation": "London, UK", "SourceIpAddress": "192.168.0.10", "EventTags": { "env": "development", "category": "logout" }, "EventDataHash": "abcdef1234567890", "EventDataVersion": "1.0", "EventData": { "username": "jane_doe", "status": "success", "login_time": "2023-06-12T12:00:00Z", "ip_address": "192.168.0.200" } }' http://localhost:6969/events/submit;
```

Feel free to make your own events. None of the fields are required for testing purposes.

### QUERY EVENTS 


Q1
    
Both events satsify this filter so both will be printed.
    
```bash 
curl -X GET -H "Content-Type: application/json" -H "Authorization: ${TOKEN}" "http://localhost:6969/events/query?SourceServiceName=audit-service"
```

Q2

Only one of the events will be returned.
    
```bash
curl -X GET -H "Authorization: ${TOKEN}" "http://localhost:6969/events/query?SourceEventID=987654"
```

Q3

Two values for the filter but they both belong to one event only.
    
```bash
curl -X GET -H "Content-Type: application/json" -H "Authorization: ${TOKEN}" "http://localhost:6969/events/query?SourceIpAddress=192.168.0.1&SourceEventID=987654"
```

Q4

Two values that will bring both the events in the system.

```bash
curl -X GET -H "Content-Type: application/json" -H "Authorization: ${TOKEN}" "http://localhost:6969/events/query?SourceServiceName=audit-service&SourceEventID=123456"
```
