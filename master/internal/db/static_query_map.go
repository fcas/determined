package db

import (
	"sync"

	"github.com/determined-ai/determined/master/pkg/etc"
)

// StaticQueryMap caches static sql files.
type StaticQueryMap struct {
	queries map[string]string
	sync.Mutex
}

// GetOrLoad fetches static sql from the cache or loads them from disk.
func (q *StaticQueryMap) GetOrLoad(queryName string) string {
	q.Lock()
	defer q.Unlock()
	if q.queries == nil {
		q.queries = make(map[string]string)
	}

	query, ok := q.queries[queryName]
	if !ok {
		query = string(etc.MustStaticFile(queryName + ".sql"))
		q.queries[queryName] = query
	}
	return query
}
