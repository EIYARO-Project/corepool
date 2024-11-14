EYPOOL
======

## 1 Install

### 1) Get the source code

```
$ git clone https://github.com/EIYARO/corepool.git $GOPATH/src/github.com/EIYARO/corepool
```

### 2) Build source code

```
$ cd $GOPATH/src/github.com/EIYARO/corepool/stratum/ey/cmd
$ go build -o stratum_ey
```

## 2 Run

### 1) Configurate parameters

```
$ cd $GOPATH/src/github.com/EIYARO/corepool/stratum/ey/conf
$ vim prod.yml
```

set `node.url` with the EIYARO node url, then leave other parameters with default value.

### 2) Run corepool

```
$ cd $GOPATH/src/github.com/EIYARO/corepool/stratum/ey/cmd
$ ./stratum_ey -config=../conf/prod.yml
```

## 3 Parameter interpretation

```
mode: prod # run mode, defines logger level and so on

# server
stratum.id: 0 # session offset id for different miner
stratum.port: 9119 # miner connection
stratum.max_conn: 32768 # max connection of miner
stratum.default_ban_period: 10m # ban malicious miner, 0s means disable

# session
session_timeout: 5m # connection timeout
session.sched_interval: 0 # work braodcast interval, 0 means braodcast when new work coming
session.diff: 1050000 # diff for miner

# node
node.url: http://127.0.0.1:9888 # EIYARO node url
node.name: ey_mainnet # EIYARO node name, set with default
node.sync_interval: 100ms # interval of getting work from node

service.port: 11002 # gin server port
```
