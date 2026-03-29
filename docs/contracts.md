# RideCore — shared contracts

These rules are agreed upon by both engineers. Do not change without a PR discussion.

## H3 resolution
Always `8`. Defined in `pkg/models/zone.go` as `H3Resolution`. Never hardcode `8` elsewhere.

## Redis key namespaces
| Key pattern         | Owner    | Purpose                        |
|---------------------|----------|--------------------------------|
| `drivers:geo`       | Person A | GEOADD sorted set for radius queries |
| `driver:{id}`       | Person A | HSET of full driver state      |
| `lock:driver:{id}`  | Person B | Redlock distributed lock keys  |

## Timestamps
All Kafka event structs use `int64` Unix milliseconds. Never embed `time.Time` in event structs.

## Sentinel errors
| Error                          | Package           | HTTP mapping |
|-------------------------------|-------------------|--------------|
| `ErrLockNotAcquired`          | internal/store    | 409          |
| `ErrNoDriversAvailable`       | internal/matching | 503          |
| `ErrLockContention`           | internal/matching | 409          |

## Kafka topics
Defined in `pkg/events/topics.go`. Never use raw strings for topic names.
