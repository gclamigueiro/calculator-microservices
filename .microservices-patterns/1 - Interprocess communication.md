# Interprocess communication (IPC)

In a microservice architecture services must often collaborate in order to handle
a request. Because service instances are typically processes running on multiple
machines, they must interact using IPC.
## - Communicating using  the synchronous Remote Procedure Invocation (RPI) pattern

A client invokes a service using a synchronous, remote procedure invocation-based
protocol, such as REST.  
[RPI Pattern](https://microservices.io/patterns/communication-style/rpi.html)
### Synchronous Request/response interaction style
In this case the client sends a request and asumme that a response will arrive in a timely fashion. the two most common mechanisms could be an HTTP-based REST or gRPC.

### Drawbacks
In a distributed system, whenever a service makes a synchronous request to another
service, there is an ever-present risk of partial failure. And, the client is blocked waiting for a response, and the danger could cascade to the client and so on cause an outage.
### Related Patterns
#### Circuit breaker
The Circuit Breaker pattern, can prevent an application from repeatedly trying to execute an operation that's likely to fail.  
[Circuit breaker](./1.1%20-%20Circuit%20Breaker%20Pattern%20.md)  

#### Service Discovery Mechanism
A problem using synchronous Remote procedure invocation is that in order
for one service to invoke another service it needs to know the network
location of a service instance. So, you need to use a Service Discovery Mechanism.  
[Service Discovery Mechanisms](./1.2%20-Service%20Discovery.md)

### Examples
In the service [multiple-operation-svc-go-kit](./../services/multiple-operation-svc-go-kit/) we called to another services using HTTP tranports, so our procces is blocked while we are waiting for a response.   
It is used platform-provided service discovery patterns specifically Kubernetes service discovery mechanisms 

## - Communicating using the Asynchronous messaging pattern
When using messaging, services communicate by asynchronously exchanging messages.
[Messaging Pattern](http://microservices.io/patterns/communication-style/messaging.html)
### One-to-one
#### Asynchronous request/response interaction style

A service client sends a request to a service, which replies asynchronously. The client doesn’t block while waiting, because the service might not send the response for a long time.
#### One-way notifications

Service client sends a request to a service, but no reply is expected or sent
### One-to-many
#### Publish/subscribe
A client publishes a notification message, which is consumed by zero or more interested services.
#### Publish/async responses
A client publishes a request message and then waits for a certain amount of time for responses from interested services.