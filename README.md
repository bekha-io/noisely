# Noisely

High-performance synthetic data generation API based on dynamic CRUD schemas.

## Description

### Supported Data Generation Types

- **Chaos** — pseudo-random data generation using Go's `math/rand` package.
- **Perlin noise** — smooth, continuous data generation based on the Perlin noise algorithm.
- **Gauss** — sharp noise and abrupt changes using Gaussian (normal) distribution.
- **Timeseries** - time-based data generation with configurable trends, seasonality, and cycles.

Planning to add GAN models in the future

### Features

- Flexible and dynamic CRUD schema configuration via API.
- Multiple noise and pattern types for realistic data simulation.
- High performance, suitable for large-scale data generation.
- RESTful API for easy integration with any technology stack.
- Extensible architecture for custom data generators.

### Use Cases

- Testing and load simulation.
- Generating datasets for machine learning and data analysis.
- Simulating real-world application scenarios.
- Prototyping and validating data-driven systems.

## Getting Started

### Run with Docker

```bash
docker run -p 8080:8080 yourusername/noisely
```

### Or build from source

```bash
git clone https://github.com/yourusername/noisely.git
cd noisely
go build -o noisely
./noisely
```

## Project Structure

The project uses a multilayer (clean) architecture and MongoDB as the main database. Main layers:

- `domain/entities` — business entities (e.g., Schema, Dataset)
- `internal/usecase` — business logic and application services
- `internal/repository` — interfaces and implementations for data access (MongoDB)
- `internal/service` — additional domain services
- `api/handler` — HTTP handlers (REST API)
- `infrastructure/mongo` — MongoDB connection and repository implementations
- `config/` — configuration files (YAML, env, etc.)

## API Documentation

API docs are available as an OpenAPI file in `/docs` directory

## Concepts

### `Schema`

`Schema` is a versioned JSON document that describes the data model used for further data generation. It includes:

- Name and version
- Field descriptions and their data types
- Validation rules
- Generation parameters for each field
- Field relationships (if any)
- Schema metadata

#### Versioning

Each time you create or edit schema with the same name, you create a new version of this schema, which increments the `version` field by 1

Querying schemas can be done specifying either:

- version number, e.g. `1`, `v1`
- `latest` keyword

`latest` will return the maximum value for the `version` field in the database. This functionality added in order to abstract and simplify interaction with Noisely.

### `Dataset`

`Dataset` is a collection of generated data points based on a specific `Schema`. Each dataset has:

- Unique identifier
- Reference to schema version
- Generation parameters
- Metadata about generation process
- Generated data points
