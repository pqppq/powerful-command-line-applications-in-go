unzip test data for benchmark
```
tar -xzvf colStatsBenchmarkData.tar.gz -C testdata
```

benchmark
```
go test -bench . -run ^$
```

cpu profile
```
go test -bench . -benchtime=10x  -run ^$ -cpuprofile cpu.prof
```
