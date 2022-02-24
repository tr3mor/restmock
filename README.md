# Restmock
Simplest way to mock http response. Pretty useful to troubleshoot Ingress/Istio and other routing issues.  
It logs all available information to simplify troubleshooting.

Config example:
```yaml
---
interactions:
  - request:
      path: /
      method: GET
    response:
      statusCode: 200
      type: plain
      body: 'test plain response'
  - request:
      path: /json
      method: GET
    response:
      statusCode: 200
      type: json
      body: '{"name": "test json response"}'
```