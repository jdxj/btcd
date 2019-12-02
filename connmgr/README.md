connmgr
=======

[![Build Status](http://img.shields.io/travis/btcsuite/btcd.svg)](https://travis-ci.org/btcsuite/btcd)
[![ISC License](http://img.shields.io/badge/license-ISC-blue.svg)](http://copyfree.org)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/btcsuite/btcd/connmgr)

Package connmgr implements a generic Bitcoin network connection manager.

connmgr 包实现了一个通用的比特币网络连接管理器.

## Overview

Connection Manager handles all the general connection concerns such as
maintaining a set number of outbound connections, sourcing peers, banning,
limiting max connections, tor lookup, etc.

连接管理器处理所有常规连接问题, 例如维护一定数量的出站连接, 获取对等点, 禁止, 限制最大连接数, tor 查找等.

The package provides a generic connection manager which is able to accept
connection requests from a source or a set of given addresses, dial them and
notify the caller on connections. The main intended use is to initialize a pool
of active connections and maintain them to remain connected to the P2P network.

该包提供了一个通用的连接管理器, 该管理器能够接受来自源或一组给定地址的连接请求,
对其进行拨号并在连接时通知调用方. 主要用途是初始化活动连接池, 并使它们保持连接到 P2P 网络.

In addition the connection manager provides the following utilities:

另外, 连接管理器提供以下实用程序:

- Notifications on connections or disconnections
- Handle failures and retry new addresses from the source
- Connect only to specified addresses
- Permanent connections with increasing backoff retry timers
- Disconnect or Remove an established connection

- 有关连接或断开的通知
- 处理故障并从源中重试新地址
- 仅连接到指定地址
- 具有增加的退避重试计时器的永久连接
- 断开或删除已建立的连接

## Installation and Updating

```bash
$ go get -u github.com/btcsuite/btcd/connmgr
```

## License

Package connmgr is licensed under the [copyfree](http://copyfree.org) ISC License.
