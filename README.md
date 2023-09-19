# Spanner ToStruct Benchmarking

This repository contains benchmarking tests that measure the performance of Google Cloud Spanner's `Row.ToStruct` and `Row.Columns` functions in Golang. The purpose is to compare the overhead when using struct tags in Google Cloud Spanner's Golang client library.

## Experiment Setup

- **OS**: Windows
- **Architecture**: amd64
- **Go Version**: 1.21
- **CPU**: 11th Gen Intel(R) Core(TM) i5-11400H @ 2.70GHz
- **Spanner library version**: v1.49.0

## Benchmark Results

| Benchmark Test            | Operations  | Time (ns/op) | Bytes Allocated (B/op) | Memory Allocations (allocs/op) |
|---------------------------|-------------|--------------|-------------------------|--------------------------------|
| BenchmarkRowToStruct-12   | 2708936     | 446.5        | 112                     | 6                              |
| BenchmarkRowColumns-12    | 19216356    | 59.30        | 0                       | 0                              |
| BenchmarkColumnVariadic-12| 19895943    | 60.29        | 0                       | 0                              |
| BenchmarkRowToStructBig-12| 166470      | 7147         | 2886                    | 36                             |
| BenchmarkRowColumnsBig-12 | 3624980     | 327.3        | 0                       | 0                              |
| BenchmarkColumnVariadicBig-12 | 2563435  | 471.5        | 512                     | 1                              |

### Explanation of Results

- **BenchmarkRowToStruct-12**: Benchmarking the `ToStruct` method on smaller rows.
- **BenchmarkRowColumns-12**: Benchmarking the `Columns` method on smaller rows.
- **BenchmarkColumnVariadic-12**: Benchmarking the `Columns` method with variadic inputs on smaller rows.
- **BenchmarkRowToStructBig-12**: Benchmarking the `ToStruct` method on larger rows.
- **BenchmarkRowColumnsBig-12**: Benchmarking the `Columns` method on larger rows.
- **BenchmarkColumnVariadicBig-12**: Benchmarking the `Columns` method with variadic inputs on larger rows.

## Observations

1. `Row.ToStruct` is considerably slower compared to `Row.Columns` and also has higher memory allocations. This suggests that `ToStruct` has a larger overhead, likely due to the use of reflection and additional memory allocations for the struct.
  
2. The time and memory overhead scales up as the row size increases (`BenchmarkRowToStructBig-12` vs `BenchmarkRowColumnsBig-12`).

3. Adding variadic inputs to `Row.Columns` (`BenchmarkColumnVariadic-12` and `BenchmarkColumnVariadicBig-12`) has minimal impact on performance.

