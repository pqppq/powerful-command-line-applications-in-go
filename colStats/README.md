unzip test data for benchmark
```
tar -xzvf colStatsBenchmarkData.tar.gz -C testdata
```

benchmark
```
go test -bench . -run ^$
```
