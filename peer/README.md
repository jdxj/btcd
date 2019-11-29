peer
====

[![Build Status](http://img.shields.io/travis/btcsuite/btcd.svg)](https://travis-ci.org/btcsuite/btcd)
[![ISC License](http://img.shields.io/badge/license-ISC-blue.svg)](http://copyfree.org)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/btcsuite/btcd/peer)

Package peer provides a common base for creating and managing bitcoin network
peers.

peer 包为创建和管理比特币网络对等点提供了通用基础.

This package has intentionally been designed so it can be used as a standalone
package for any projects needing a full featured bitcoin peer base to build on.

该软件包经过精心设计, 因此可以用作需要完整功能的比特币同等基础的任何项目的独立软件包.

## Overview

This package builds upon the wire package, which provides the fundamental
primitives necessary to speak the bitcoin wire protocol, in order to simplify
the process of creating fully functional peers.  In essence, it provides a
common base for creating concurrent safe fully validating nodes, Simplified
Payment Verification (SPV) nodes, proxies, etc.

该程序包建立在 wire 包的基础上, wire 包提供了讲 bitcoin wire 协议所必需的基本原语,
以简化创建全功能的对等点的过程. 本质上, 它为创建并发安全的全验证节点,
简化付款验证 (SPV) 节点, 代理等提供了通用基础.

A quick overview of the major features peer provides are as follows:

peer 提供的主要功能的简要概述如下:

 - Provides a basic concurrent safe bitcoin peer for handling bitcoin
   communications via the peer-to-peer protocol
 - Full duplex reading and writing of bitcoin protocol messages
 - Automatic handling of the initial handshake process including protocol
   version negotiation
 - Asynchronous message queueing of outbound messages with optional channel for
   notification when the message is actually sent
 - Flexible peer configuration
   - Caller is responsible for creating outgoing connections and listening for
     incoming connections so they have flexibility to establish connections as
     they see fit (proxies, etc)
   - User agent name and version
   - Bitcoin network
   - Service support signalling (full nodes, bloom filters, etc)
   - Maximum supported protocol version
   - Ability to register callbacks for handling bitcoin protocol messages
 - Inventory message batching and send trickling with known inventory detection
   and avoidance
 - Automatic periodic keep-alive pinging and pong responses
 - Random nonce generation and self connection detection
 - Proper handling of bloom filter related commands when the caller does not
   specify the related flag to signal support
   - Disconnects the peer when the protocol version is high enough
   - Does not invoke the related callbacks for older protocol versions
 - Snapshottable peer statistics such as the total number of bytes read and
   written, the remote address, user agent, and negotiated protocol version
 - Helper functions pushing addresses, getblocks, getheaders, and reject
   messages
   - These could all be sent manually via the standard message output function,
     but the helpers provide additional nice functionality such as duplicate
     filtering and address randomization
 - Ability to wait for shutdown/disconnect
 - Comprehensive test coverage
 
- 提供一个基本的并发安全比特币对等点, 用于通过点对点协议处理比特币通信
- 全双工读取和写入比特币协议消息
- 自动处理初始握手过程, 包括协议版本协商
- 实际发送消息时, 带有可选 channel 的出站消息的异步消息队列, 以进行通知
- 灵活的对等点配置
    - 调用者负责创建传出连接并监听传入连接, 因此他们可以灵活地建立自己认为合适的连接 (代理等)
    - 用户代理名称和版本
    - 比特币网络
    - 服务支持信号 (完整节点, 布隆过滤器等)
    - 最高支持的协议版本
    - 能够注册回调以处理比特币协议消息
- 库存消息批处理和发送带有已知库存检测的 trickling 和预防
- 自动定期 keep-alive 的 pinging 和 pong 响应
- 随机 nonce 生成和自连接检测
- 当调用方未指定相关标志来表示支持时, 将正确处理与 bloom filter 相关的命令
    - 协议版本足够高时断开对等点
    - 不为较早的协议版本调用相关的回调
- 可快照的对等点统计信息, 例如, 读取和写入的字节总数, 远程地址, 用户代理和协商的协议版本
- 辅助函数推送地址, getblock, getheader 和拒绝消息
    - 这些都可以通过标准消息输出函数手动发送, 但是帮助程序提供了其他不错的功能, 例如重复过滤和地址随机化
    - 能够等待 shutdown/disconnect
- 全面的测试范围

## Installation and Updating

```bash
$ go get -u github.com/btcsuite/btcd/peer
```

## Examples

* [New Outbound Peer Example](https://godoc.org/github.com/btcsuite/btcd/peer#example-package--NewOutboundPeer)  
  Demonstrates the basic process for initializing and creating an outbound peer.
  Peers negotiate by exchanging version and verack messages.  For demonstration,
  a simple handler for the version message is attached to the peer.
  
* 演示了初始化和创建出站对等点的基本过程. 对等点通过交换版本和 Verack 消息进行协商.
  为了演示, 将版本消息的简单处理程序附加到对等点.

## License

Package peer is licensed under the [copyfree](http://copyfree.org) ISC License.
