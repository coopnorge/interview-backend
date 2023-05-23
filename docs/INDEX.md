# Interview Assignment - Coop Logistics API

The test assignment for a Software Engineer
focusing on involves developing a functional
API that enables efficient management of logistics operations.
This assignment requires implementing key features
such as collecting delivery routes, collect metrics
and to be able for integrating with external systems.

> The goal is to assess the candidate's proficiency
> In designing and implementing API endpoints,
> handling data processing and validation,
> and ensuring robustness and scalability in a logistics context.

## Acceptance Criteria

- All three parts must be done (
  [One](INDEX.md#part-oneimplement-backend-api),
  [Two](INDEX.md#part-twostore-delivery-paths),
  [Three](INDEX.md#part-threeexport-delivery-paths)
  )
- API must be sustainable to handle requests without dropping them.

## Assignment Details

Coop Logistics Engine, simulates several cargo units that
deliver cargo between different storages in the world.
Start the Coop Logistics Engine executable delivered in this package.

- Coop Logistics Engine binary
  - TODO [Win64](#)
  - TODO [Linux64](#)
  - TODO [Darwin64](#)
- Assigment [instructions](docs/INDEX.md)


### Part One:Implement Backend API

The executable sends requests to http://localhost:8080
> TODO Add a link to API Docs

Your first task is to implement a backend API that handles messages.
The solution should output a log message to STDOUT once per second with:

- Number of received messages / second.

### Part Two:Store Delivery Paths

The cargo units run between the different storages in the world.

Your task is to store these delivery paths,
data comes into the API you implemented in *Part One*.
Write at least one unit test that verifies that your solution works.

The solution should output a log message to STDOUT once per second with:

- Number of received messages / second.
- Total number of delivery paths it has found.

#### Notes:

- There are a limited number of Storage.
- Storage has IDs that start with 0.
- StorageID in CargoReachedDestination equals DestinationStorage in StoragePort

### Part Three:Export Delivery Paths

Given the delivery paths stored in *Part Two*.

Implement an API endpoint that can export data with unique delivery paths.
