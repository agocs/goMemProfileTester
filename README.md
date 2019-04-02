# Testing `ReadMemStats` vs `pprof.Lookup`

This is a really simple webserver designed to test the 
performance of ReadMemStats vs pprof.Lookup. It exposes
a couple of endpoints that both generate a long, sorted
slice of float64s. It then reports on the allocated 
memory one of two ways.

This repo carries the Genuine "Written on an Airplane" seal of approval.

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
Duration      [total, attack, wait]    9.981021439s, 9.980178389s, 843.05µs
Latencies     [mean, 50, 95, 99, max]  1.076732ms, 991.45µs, 1.565569ms, 1.955077ms, 7.243589ms
Bytes In      [total, mean]            9866887, 19733.77
Bytes Out     [total, mean]            0, 0.00
Success       [ratio]                  100.00%
Status Codes  [code:count]             200:500
Error Set:
```

### ReadMemStats

```
Requests      [total, rate]            500, 50.09
Duration      [total, attack, wait]    9.982931537s, 9.982067082s, 864.455µs
Latencies     [mean, 50, 95, 99, max]  1.077519ms, 966.665µs, 1.586833ms, 1.97482ms, 6.859381ms
Bytes In      [total, mean]            9866533, 19733.07
Bytes Out     [total, mean]            0, 0.00
Success       [ratio]                  100.00%
Status Codes  [code:count]             200:500
Error Set:
```

### pprof.Lookup

```
Requests      [total, rate]            500, 50.09
Duration      [total, attack, wait]    9.983083109s, 9.981228359s, 1.85475ms
Latencies     [mean, 50, 95, 99, max]  1.637835ms, 1.557003ms, 2.705042ms, 3.478394ms, 7.55065ms
Bytes In      [total, mean]            9866530, 19733.06
Bytes Out     [total, mean]            0, 0.00
Success       [ratio]                  100.00%
Status Codes  [code:count]             200:500
Error Set:
```
