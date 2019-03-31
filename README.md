# Testing `ReadMemStats` vs `pprof.Lookup`

This is a really simple webserver designed to test the 
performance of ReadMemStats vs pprof.Lookup. It exposes
a couple of endpoints that both generate a long, sorted
slice of float64s. It then reports on the allocated 
memory one of two ways.

## Why

I want to look into how much recording heap utilization will
impact the performance of my webservers. ReadMemStats is fast,
but stops the world. pprof.Lookup is about 100x slower, but 
shouldn't stop the world.

## How

This exposes three endpoints:

- `localhost:8080/readMem0/` -- Control. Doesn't read memory.
- `localhost:8080/readMem1/` -- ReadMemStats
- `localhost:8080/readMem2/` -- pprof.Lookup

I tested this using https://github.com/tsenart/vegeta . 

```
echo "GET http://localhost:8080/readMem2/" | vegeta attack -duration 10s | vegeta report
```

## First pass results

### Control

```
Requests      [total, rate]            500, 50.10
Duration      [total, attack, wait]    9.981827624s, 9.980855699s, 971.925µs
Latencies     [mean, 50, 95, 99, max]  1.12808ms, 1.083607ms, 1.527106ms, 1.83792ms, 4.748087ms
Bytes In      [total, mean]            9866209, 19732.42
Bytes Out     [total, mean]            0, 0.00
Success       [ratio]                  100.00%
Status Codes  [code:count]             200:500
Error Set:
```

### ReadMemStats

```
Requests      [total, rate]            500, 50.10
Duration      [total, attack, wait]    9.981129903s, 9.980053077s, 1.076826ms
Latencies     [mean, 50, 95, 99, max]  1.11529ms, 1.068754ms, 1.536628ms, 1.730011ms, 5.803736ms
Bytes In      [total, mean]            9866236, 19732.47
Bytes Out     [total, mean]            0, 0.00
Success       [ratio]                  100.00%
Status Codes  [code:count]             200:500
Error Set:
```

### pprof.Lookup

```
Requests      [total, rate]            500, 50.10
Duration      [total, attack, wait]    10.002298725s, 9.980712303s, 21.586422ms
Testing `ReadMemStats` vs `pprof.Lookup`
Latencies     [mean, 50, 95, 99, max]  14.205203ms, 13.191055msTesting `ReadMemStats` vs `pprof.Lookup`
, 23.668592ms, 61.845205ms, 81.412335ms

Bytes In      [total, mean]            9866039, 19732.08
Bytes Out     [total, mean]            0, 0.00
Success       [ratio]                  100.00%
Status Codes  [code:count]             200:500
Error Set:
```
