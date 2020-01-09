// Copyright (c) 2015-2016 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package txscript

import (
	"sync"

	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
)

// sigCacheEntry represents an entry in the SigCache. Entries within the
// SigCache are keyed according to the sigHash of the signature. In the
// scenario of a cache-hit (according to the sigHash), an additional comparison
// of the signature, and public key will be executed in order to ensure a complete
// match. In the occasion that two sigHashes collide, the newer sigHash will
// simply overwrite the existing entry.
type sigCacheEntry struct {
	sig    *btcec.Signature
	pubKey *btcec.PublicKey
}

// SigCache implements an ECDSA signature verification cache with a randomized
// entry eviction policy. Only valid signatures will be added to the cache. The
// benefits of SigCache are two fold. Firstly, usage of SigCache mitigates a DoS
// attack wherein an attack causes a victim's client to hang due to worst-case
// behavior triggered while processing attacker crafted invalid transactions. A
// detailed description of the mitigated DoS attack can be found here:
// https://bitslog.wordpress.com/2013/01/23/fixed-bitcoin-vulnerability-explanation-why-the-signature-cache-is-a-dos-protection/.
// Secondly, usage of the SigCache introduces a signature verification
// optimization which speeds up the validation of transactions within a block,
// if they've already been seen and verified within the mempool.
//
// SigCache 使用随机条目逐出策略实现 ECDSA 签名验证缓存. 只有有效的签名会被添加到缓存中.
// SigCache 的好处有两方面. 首先, 使用 SigCache 可以缓解 DoS 攻击, 其中,
// 由于在处理攻击者制作的无效交易时触发的最坏情况的行为, 攻击导致受害者的客户端挂起.
// 可在以下位置找到缓解的 DoS 攻击的详细说明:
// https://bitslog.wordpress.com/2013/01/23/fixed-bitcoin-vulnerability-explanation-why-the-signature-cache-is-a-dos-protection/.
// 其次, 如果已在内存池中看到并验证了交易, 则 SigCache 的使用会引入签名验证优化功能,
// 从而加快块内交易的验证速度.
type SigCache struct {
	sync.RWMutex
	validSigs  map[chainhash.Hash]sigCacheEntry
	maxEntries uint
}

// NewSigCache creates and initializes a new instance of SigCache. Its sole
// parameter 'maxEntries' represents the maximum number of entries allowed to
// exist in the SigCache at any particular moment. Random entries are evicted
// to make room for new entries that would cause the number of entries in the
// cache to exceed the max.
func NewSigCache(maxEntries uint) *SigCache {
	return &SigCache{
		validSigs:  make(map[chainhash.Hash]sigCacheEntry, maxEntries),
		maxEntries: maxEntries,
	}
}

// Exists returns true if an existing entry of 'sig' over 'sigHash' for public
// key 'pubKey' is found within the SigCache. Otherwise, false is returned.
//
// NOTE: This function is safe for concurrent access. Readers won't be blocked
// unless there exists a writer, adding an entry to the SigCache.
func (s *SigCache) Exists(sigHash chainhash.Hash, sig *btcec.Signature, pubKey *btcec.PublicKey) bool {
	s.RLock()
	entry, ok := s.validSigs[sigHash]
	s.RUnlock()

	return ok && entry.pubKey.IsEqual(pubKey) && entry.sig.IsEqual(sig)
}

// Add adds an entry for a signature over 'sigHash' under public key 'pubKey'
// to the signature cache. In the event that the SigCache is 'full', an
// existing entry is randomly chosen to be evicted in order to make space for
// the new entry.
//
// NOTE: This function is safe for concurrent access. Writers will block
// simultaneous readers until function execution has concluded.
func (s *SigCache) Add(sigHash chainhash.Hash, sig *btcec.Signature, pubKey *btcec.PublicKey) {
	s.Lock()
	defer s.Unlock()

	if s.maxEntries <= 0 {
		return
	}

	// If adding this new entry will put us over the max number of allowed
	// entries, then evict an entry.
	if uint(len(s.validSigs)+1) > s.maxEntries {
		// Remove a random entry from the map. Relying on the random
		// starting point of Go's map iteration. It's worth noting that
		// the random iteration starting point is not 100% guaranteed
		// by the spec, however most Go compilers support it.
		// Ultimately, the iteration order isn't important here because
		// in order to manipulate which items are evicted, an adversary
		// would need to be able to execute preimage attacks on the
		// hashing function in order to start eviction at a specific
		// entry.
		for sigEntry := range s.validSigs {
			delete(s.validSigs, sigEntry)
			break
		}
	}
	s.validSigs[sigHash] = sigCacheEntry{sig, pubKey}
}
