// Copyright (C) 2019-2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package merkledb

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ava-labs/avalanchego/database"
	"github.com/ava-labs/avalanchego/database/memdb"
	"github.com/ava-labs/avalanchego/utils/maybe"
)

// Tests:
// * Putting a key-node pair in the database
// * Getting a key-node pair from the cache and from the base db
// * Deleting a key-node pair from the database
// * Evicting elements from the cache
// * Flushing the cache
func TestIntermediateNodeDB(t *testing.T) {
	require := require.New(t)

	cacheSize := 100
	evictionBatchSize := cacheSize
	baseDB := memdb.New()
	db := newIntermediateNodeDB(
		baseDB,
		&sync.Pool{
			New: func() interface{} { return make([]byte, 0) },
		},
		&mockMetrics{},
		cacheSize,
		evictionBatchSize,
	)

	// Put a key-node pair
	key := newPath([]byte{0x01})
	node1 := &node{
		dbNode: dbNode{
			value: maybe.Some([]byte{0x01}),
		},
	}
	require.NoError(db.Put(key, node1))

	// Get the key-node pair from cache
	node1Read, err := db.Get(key)
	require.NoError(err)
	require.Equal(node1, node1Read)

	// Overwrite the key-node pair
	node1Updated := &node{
		dbNode: dbNode{
			value: maybe.Some([]byte{0x02}),
		},
	}
	require.NoError(db.Put(key, node1Updated))

	// Assert the key-node pair was overwritten
	node1Read, err = db.Get(key)
	require.NoError(err)
	require.Equal(node1Updated, node1Read)

	// Delete the key-node pair
	require.NoError(db.Delete(key))
	_, err = db.Get(key)

	// Assert the key-node pair was deleted
	require.Equal(database.ErrNotFound, err)

	// Put elements in the cache until it is full.
	expectedSize := cacheEntrySize(key, nil)
	added := 0
	for {
		key := newPath([]byte{byte(added)})
		node := &node{
			dbNode: dbNode{
				value: maybe.Some([]byte{byte(added)}),
			},
		}
		newExpectedSize := expectedSize + cacheEntrySize(key, node)
		if newExpectedSize > cacheSize {
			// Don't trigger eviction.
			break
		}

		require.NoError(db.Put(key, node))
		expectedSize = newExpectedSize
		added++
	}

	// Assert cache has expected number of elements
	require.Equal(added, db.nodeCache.fifo.Len())

	// Put one more element in the cache, which should trigger an eviction
	// of all but 2 elements. 2 elements remain rather than 1 element because of
	// the added key prefix increasing the size tracked by the batch.
	key = newPath([]byte{byte(cacheSize)})
	node := &node{
		dbNode: dbNode{
			value: maybe.Some([]byte{byte(cacheSize)}),
		},
	}
	require.NoError(db.Put(key, node))

	// Assert cache has expected number of elements
	require.Equal(2, db.nodeCache.fifo.Len())
	gotKey, _, ok := db.nodeCache.fifo.Oldest()
	require.True(ok)
	require.Equal(newPath([]byte{byte(added - 1)}), gotKey)

	// Get a node from the base database (not cache)
	nodeRead, err := db.Get(newPath([]byte{0x03}))
	require.NoError(err)
	require.Equal(maybe.Some([]byte{0x03}), nodeRead.value)

	// Flush the cache.
	require.NoError(db.Flush())

	// Assert the cache is empty
	require.Zero(db.nodeCache.fifo.Len())

	// Assert the evicted cache elements were written to disk with prefix.
	it := baseDB.NewIteratorWithPrefix(intermediateNodePrefix)
	defer it.Release()

	count := 0
	for it.Next() {
		count++
	}
	require.NoError(it.Error())
	require.Equal(added+1, count)
}