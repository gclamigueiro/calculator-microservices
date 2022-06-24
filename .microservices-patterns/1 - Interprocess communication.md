# Interprocess communication


## Communicating using the synchronous Remote procedure invocation pattern

### Synchronous Request/response

In this case the client sends a request and asumme that a response will arrive in a timely fashion. the two most common mechanisms could be an HTTP-based REST or gRPC.

### Drawbacks

In a distributed system, whenever a service makes a synchronous request to another
service, there is an ever-present risk of partial failure. And, the client is blocked waiting for a response, and the danger could cascade to the client and so on cause an outage.

### Example

In the service ```multiple-operation-svc-go-kit``` we called to another services using HTTP tranports, so our procces is blocked while we are waiting for a response. 

### Related Patterns
#### Circuit breaker

The Circuit Breaker pattern, can prevent an application from repeatedly trying to execute an operation that's likely to fail.

[Circuit breaker](./1.1%20-%20Circuit%20Breaker%20Pattern%20.md)  

## Communicating using the Asynchronous messaging pattern

### One-to-one

#### Asynchronous request/response

Asynchronous request/response—A service client sends a request to a service, which replies asynchronously. The client doesn’t block while waiting, because the service might not send the response for a long time.

#### One-way notifications

Service client sends a request to a service, but no reply is expected or sent

### One-to-many

#### Publish/subscribe

A client publishes a notification message, which is consumed by zero or more interested services.

#### Publish/async responses

A client publishes a request message and then waits for a certain amount of time for responses from interested services.