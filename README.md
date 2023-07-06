# Interview Assignment - Coop Logistics API

Before jump in:

- [Golang documentation](https://go.dev/doc/)
- [gRPC documentation](https://grpc.io/docs/)
- [Buf tool to manage Protobuf files](https://buf.build/docs/introduction/)

The test assignment for a Software Engineer
-> [instructions](docs/instructions.md).

## Overview

![project_overview](docs/assets/overview.png)


## Repository Overview

### Folders

```text
├── api                                  // API definition with gRPC proto-files
├── docker-compose                       // Docker related files
├── docs                                 // Instructions
└── src                                  // Source code related to test assignment
    ├── client                           // Client that will send requests to API Server
    ├── generated                        // Shared code that can be used to implement gRPC
    │   └── logistics
    │       └── api
    │           └── v1                   // protoc-gen-go of API
    └── server                           // Your solution that provides API server
```

### Files

```text
├── api
│   └── v1
│       └── logistics.proto   // API definition that used in this test assignment
├── buf.*.yaml                // Files that used by "Buf" to managed generated code from Protobuf
├── catalog-info.yaml         // Internal Coop configuration of project, you can ignore it
└── docker-compose.yml        // As part of assignment to running multi-container Docker applications (Client and Server APIs)
```

## Example of client program output:

```text
2023/06/02 10:10:55 CargoUnit: Ferrari -  650ci Convertible moving to - Latitude:113, Longitude:31
2023/06/02 10:10:55 CargoUnit: Ligier - Milan moving to - Latitude:113, Longitude:30
2023/06/02 10:10:55 CargoUnit: Koenigsegg - New Beetle Convertible moving to - Latitude:113, Longitude:30 - Reached Objective.
2023/06/02 10:10:55 CargoUnit: Ferrari -  650ci Convertible moving to - Latitude:113, Longitude:30
2023/06/02 10:10:55 CargoUnit: Ligier - Milan moving to - Latitude:113, Longitude:30
2023/06/02 10:10:55 CargoUnit: Ligier - Milan moving to - Latitude:113, Longitude:30 - Reached Objective.
2023/06/02 10:10:55 CargoUnit: Ferrari -  650ci Convertible moving to - Latitude:113, Longitude:30
2023/06/02 10:10:56 CargoUnit: Ferrari -  650ci Convertible moving to - Latitude:113, Longitude:30 - Reached Objective.
2023/06/02 10:10:56 All delivery units reached warehouse...

Execution time: 1.202202348s
| Operation            | Count | Errors |
-----------------------------------------
| MoveUnit             | 13324 | 0      |
| UnitReachedWarehouse | 108   | 0      |
-----------------------------------------
```
