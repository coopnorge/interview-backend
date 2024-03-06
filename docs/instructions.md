# Interview Assignment - Coop Logistics API

**Implement Server side part that can handle client requests
with specific business logic.**

## Acceptance Criteria

- All three parts must be done.
- API Server must be capable of handling all requests sent by this client
  without dropping them.
- API Server - code design must be flexible for introducing new features.

## Assignment Appendix

Coop Logistics Engine simulates client activity about several cargo units that
deliver goods between different warehouses in the world. All this activity will
be reported as request to your API server.

For example
It sends `gRPC` or `HTTP` requests to `127.0.0.1:50051`.

You can use `env` variables to redefine values for configuration.

| ENV                   | Description                                                   |
|-----------------------|---------------------------------------------------------------|
| CLIENT_SERVICE_HOST   | Base host like 127.0.0.1                                      |
| CLIENT_SERVICE_PORT   | Server port like 50051, 8080                                  |
| CLIENT_TRANSPORT_TYPE | Protocol that client will use to send requests (gRPC or HTTP) |
| CLIENT_HTTP_SCHEME    | HTTP Scheme (http or https) if CLIENT_TRANSPORT_TYPE is HTTP  |

For reference, you can copy template
of [docker-compose](../docker-compose.yaml) "interview_backend_client" and
configure you `API server`.

If you are willing to compile it locally and execute client imitation, then
look for `cmd/logistics` where `main.go` is located.

## Assignment

The result of this task should be an API system that fulfils the criteria
described in three parts below.

*Don't hesitate to ask questions for clarification if you have them.*

**Good luck and we hope you will have fun!**
___

### Part One: Implement Backend API

Implement an API sever that provides the following service.

- If you know gRPC: [proto-file](../api/v1/logistics.proto)
- Alternative HTTP: [swagger-file](../api/v1/logistics.swagger.json)

> ! The solution must continuously output a log message to STDOUT every second.
> Each log message should include the total number of messages received during
> that second.

### Part Two: Store Delivery Units

In this part of the assignment, you are tasked with tracking cargo units as
they are transported between various warehouses around the world. You will be
capturing this delivery path information through the API you developed in Part
One.

Your objectives are to:

1. Implement storage for these delivery paths, with incoming data processed by
   your API.
2. Write at least one unit test to verify the functionality of your solution.
3. Your API Server (solution) must include the following outputs:

- Periodic Logging: Every second, output a log message to STDOUT that displays
  the count of messages received during that second.
- Summary Output: Provide a summary that includes:
  - The total number of delivery units.
  - A list of warehouses that have received supplies (i.e., units that have
    reached their destination).
  - The quantity of delivery units each warehouse has received.

### Part Three: Export Delivery Paths

Building on the delivery paths data stored in Part Two, your task is now to
implement an API endpoint. This endpoint should be capable of exporting data
concerning the warehouses and their suppliers.

Specifically, the API endpoint should:

- Retrieve and format data about each warehouse, including information about
  its suppliers.
- Ensure the data is structured in a way that clearly identifies the
  relationship between warehouses and the suppliers that deliver to them.
