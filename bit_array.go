package addstructs

import (
	"fmt"
	"sync"
)

type BitArray struct {
	array []byte
	size  int
	mu    *sync.RWMutex
}

func NewBitArray(size int) *BitArray {
	array := make([]byte, size/8+1)

	return &BitArray{
		array: array,
		size:  size,
		mu:    &sync.RWMutex{},
	}
}

func (arr *BitArray) Size() int {
	return arr.size
}

func (arr *BitArray) Clear() {
	arr.mu.Lock()
	defer arr.mu.Unlock()

	arr.array = make([]byte, arr.size/8+1)
}

func (arr *BitArray) isValidIndex(index int) error {
	if index >= arr.Size() {
		return &OutOfRange{
			index:   index,
			maxSize: arr.size,
		}
	}

	return nil
}

func (arr *BitArray) Set(index int) error {
	err := arr.isValidIndex(index)
	if err != nil {
		return err
	}

	arr.mu.RLock()
	defer arr.mu.RUnlock()

	arr.array[index/8] |= 1 << uint(index%8)
	return nil
}

func (arr *BitArray) Get(index int) (int, error) {
	err := arr.isValidIndex(index)
	if err != nil {
		return 0, err
	}
	arr.mu.Lock()
	defer arr.mu.Unlock()

	if (arr.array[index/8] & (1 << uint(index%8))) != 0 {
		return 1, nil
	}
	return 0, nil
}

func (arr *BitArray) Toggle(index int) error {
	err := arr.isValidIndex(index)
	if err != nil {
		return err
	}

	arr.mu.Lock()
	defer arr.mu.Unlock()

	arr.array[index/8] ^= 1 << uint(index%8)
	return nil
}

type OutOfRange struct {
	index   int
	maxSize int
}

func (e *OutOfRange) Error() string {
	return fmt.Sprintf("index %d out of range %d", e.index, e.maxSize)
}
