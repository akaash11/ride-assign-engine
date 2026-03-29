# RideCore

A backend dispatch simulation system modeled after Uber/Lyft internals.
Built with Go, Kafka (Redpanda), Redis, and PostgreSQL.

## Architecture

- **GPS Simulator** — generates realistic driver location events to Kafka
- **Location Consumer** — consumes events, updates Redis GEO index + H3 cell state
- **Matching Engine** — radius queries Redis, scores drivers, uses Redlock for assignment
- **Zone Service** — aggregates H3 hex cell stats, feeds demand/supply metrics
- **Admin API** — REST API for zone stats, live driver state, simulation control

## Quick start

```bash
# Start all infrastructure
make infra-up

# Create Kafka topics
make topics

# Build all services
make build
```

## Infrastructure
| Service     | Port  |
|-------------|-------|
| Redpanda    | 9092  |
| Redis       | 6379  |
| PostgreSQL  | 5432  |
| Prometheus  | 9090  |
| Grafana     | 3000  |

## Docs
See [docs/contracts.md](docs/contracts.md) for shared interface contracts.
