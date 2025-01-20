package fphash

import (
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestContextPool(t *testing.T) {
	// Number of goroutines to run concurrently
	const numGoroutines = 200
	const iterationsPerGoroutine = 10

	// Create a wait group to synchronize goroutines
	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	// Run concurrent Hash operations
	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			input := []uint8{1, 2, 3, 4, 5}
			for j := 0; j < iterationsPerGoroutine; j++ {
				Hash(input)
			}
		}()
	}

	// Wait for all goroutines to complete
	wg.Wait()

	// Force garbage collection multiple times to trigger finalizers
	for i := 0; i < 5; i++ {
		runtime.GC()
		time.Sleep(100 * time.Millisecond) // Give finalizers time to run
	}

	// Get stats after test
	created, finalized := GetCtxStats()
	totalOperations := numGoroutines * iterationsPerGoroutine

	// Verify the number of contexts created
	// We expect significantly fewer contexts than operations due to pooling
	if int(created) >= totalOperations {
		t.Errorf("Expected context pool to reuse contexts, but got %d creations for %d operations", created, totalOperations)
	}

	// We expect at least a few contexts to be created for concurrent operations
	if created < 1 {
		t.Error("Expected at least one context to be created")
	}

	t.Logf("Created %d contexts, %d finalizers called for %d operations",
		created, finalized, totalOperations)

	// We expect some finalizers to be called after forcing GC
	if finalized == 0 {
		t.Error("Expected some finalizers to be called after garbage collection")
	}

	// We expect finalizer count to not exceed creation count
	if finalized > created {
		t.Errorf("Finalizer count (%d) exceeds context creation count (%d)",
			finalized, created)
	}
}
