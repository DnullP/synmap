// Writen by GPT
package syncmap_test

import (
	syncmap "syncMap"
	"testing"
	"time"
)

// 测试 Store 和 Load 功能
func TestSyncMapChannel_StoreAndLoad(t *testing.T) {
	syncMap := syncmap.NewSyncMapChannel[string, int]()

	// 测试存储键值对
	syncMap.Store("key1", 100)
	syncMap.Store("key2", 200)

	// 测试读取键值对
	if val := syncMap.Load("key1"); val != 100 {
		t.Errorf("expected 100, got %d", val)
	}

	if val := syncMap.Load("key2"); val != 200 {
		t.Errorf("expected 200, got %d", val)
	}

	// 测试读取不存在的键
	if val := syncMap.Load("key3"); val != 0 {
		t.Errorf("expected 0 for non-existing key, got %d", val)
	}

	// 关闭 SyncMap
	syncMap.Close()
}

// 测试 Reset 功能
func TestSyncMapChannel_Reset(t *testing.T) {
	syncMap := syncmap.NewSyncMapChannel[string, int]()

	// 存储一些键值对
	syncMap.Store("key1", 100)
	syncMap.Store("key2", 200)

	// 确保键值存在
	if val := syncMap.Load("key1"); val != 100 {
		t.Errorf("expected 100, got %d", val)
	}

	// 重置 SyncMap
	syncMap.Reset()

	// 确保数据被清空
	if val := syncMap.Load("key1"); val != 0 {
		t.Errorf("expected 0 after reset, got %d", val)
	}

	if val := syncMap.Load("key2"); val != 0 {
		t.Errorf("expected 0 after reset, got %d", val)
	}

	// 关闭 SyncMap
	syncMap.Close()
}

// 测试 Close 功能
func TestSyncMapChannel_Close(t *testing.T) {
	syncMap := syncmap.NewSyncMapChannel[string, int]()

	// 启动一个协程尝试存储数据，确保不会造成死锁
	go func() {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("unexpected panic: %v", r)
			}
		}()
		syncMap.Store("key1", 100)
	}()

	// 关闭 SyncMap
	syncMap.Close()

	// 等待协程处理
	time.Sleep(100 * time.Millisecond)
}
