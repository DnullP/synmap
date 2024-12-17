package syncmap_test

import (
	syncmap "syncMap"
	"testing"
)

func TestSynMapRWLock(t *testing.T) {
	testMap := syncmap.NewSyncMapRWLock[string, string]()

	testMap.Store("foo", "bar")
	t.Log(testMap.Load("foo"))

	testMap.Reset()
}
