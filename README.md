# Spanner ToStruct Benchmarking

This repository contains benchmarking tests that measure the performance of Google Cloud Spanner's `Row.ToStruct` and `Row.Columns` functions in Golang. The purpose is to compare the overhead when using struct tags in Google Cloud Spanner's Golang client library.

## Experiment Setup

- **OS**: Windows
- **Architecture**: amd64
- **Go Version**: 1.21
- **CPU**: 11th Gen Intel(R) Core(TM) i5-11400H @ 2.70GHz
- **Spanner library version**: v1.49.0

## Benchmark Results

| Benchmark Test         | Operations | Time (ns/op) | Bytes Allocated (B/op) | Memory Allocations (allocs/op) |
|------------------------|------------|--------------|-------------------------|--------------------------------|
| BenchmarkRowToStruct   | 2763956    | 436.7        | 112                     | 6                              |
| BenchmarkRowColumns    | 19261048   | 58.65        | 0                       | 0                              |
| BenchmarkRowToStructBig| 167869     | 7159         | 2887                    | 36                             |
| BenchmarkRowColumnsBig | 3732808    | 318.9        | 0                       | 0                              |

### Explanation of Results

- **BenchmarkRowToStruct**: Benchmarking the `ToStruct` method on smaller rows.
- **BenchmarkRowColumns**: Benchmarking the `Columns` method on smaller rows.
- **BenchmarkRowToStructBig**: Benchmarking the `ToStruct` method on larger rows.
- **BenchmarkRowColumnsBig**: Benchmarking the `Columns` method on larger rows.

## Observations

1. `Row.ToStruct` is considerably slower compared to `Row.Columns` and also has higher memory allocations. This suggests that `ToStruct` has a larger overhead, likely due to the use of reflection and additional memory allocations for the struct.

2. The time and memory overhead scales up as the row size increases (`BenchmarkRowToStructBig` vs `BenchmarkRowColumnsBig`).
