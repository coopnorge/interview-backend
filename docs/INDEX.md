# Interview Assignment - Coop Logistics API

## Acceptance Criteria

- All three parts must be done (One, Two, Three).
- API must be sustainable to handle requests without dropping them.

## Assignment Appendix

Coop Logistics Engine, simulates several cargo units that
deliver goods between different warehouses in the world.

Coop Logistics Engine sends `gRPC` requests to `127.0.0.1:50051`.
Checkout service in
[docker-compose](Dockerfile) "interview_backend_coop_logistics"

## Assignment Details

Develop an API system that matches the criteria described in parts,
here is [Swagger documentation with descriptions.](api/logistics.swagger.json)

Don't hesitate to ask questions for clarification if you have them.

*Good luck and we hope you will have fun!*
___

### Part One: Implement Backend API

Your first task is to implement API server
for that [proto-file](api/v1/logistics.proto).

The solution should output a log message to STDOUT once per second with:

- Number of received messages / second.

### Part Two: Store Delivery Units

The cargo units run between the different warehouses in the world.

Your task is to store these delivery paths,
data comes into the API you implemented in *Part One*.
Write at least one unit test that verifies that your solution works.

The solution should output a log message to STDOUT once per second with:

- Number of received messages / second.
- Write summary
  - Total delivery units.
  - Warehouses that has been supplied.
  - How many delivery units have warehouses.

### Part Three: Export Delivery Paths

Given the delivery paths stored in *Part Two*.

Implement an API endpoint that can export data about warehouses and it's
suppliers.
