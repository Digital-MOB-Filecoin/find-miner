# Find Miner CLI Tool


The use case for find-miner is to enable a user of the Lotus CLI or API to select a miner for a given type of storage deal. The CLI tool uses data from filrep.io Filecoin Reputation System.


### Usage:
```
  find-miner [flags]
```

### Flags:
```
  --h, --help                          Help for find-miner
  --fastRetrieval string               Fast Retrieval (true/false)
  --region string                      Miner's region : ap|cn|na|eu
  --size int                           Deal size in bytes
  --skip-miners int                    The first N miners that would normally be returned are skipped
  --verified string                    Verified (true/false)
  --verified-storage-price-limit int   Maximum acceptable verified storage price (in FIL)
```
      
### Build
```
go build
```
