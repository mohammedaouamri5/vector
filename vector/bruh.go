// vector/bruh.go
package vector

import "fmt"

type Vec[T any] struct {
	Capacity int
	Updator  float32
	len      int
	Data     *[]T
}

func New[T any](p_AdditionalCapacity int, p_Updator float32, p_Data []T) Vec[T] {
	// Allocate the underlying array on the heap
	data := make([]T, len(p_Data), len(p_Data)+p_AdditionalCapacity)
	copy(data, p_Data)

	return Vec[T]{
		Capacity: len(data),
		Updator:  p_Updator,
		len:      len(p_Data),
		Data:     &data, // Data is a pointer to the heap-allocated array
	}
}

func (me *Vec[T]) Size() int {
	return me.len
}

/* Push adds a new element to the vector.
The element to be added is passed as the parameter v.
No return value.
EX:

 bruh.Push(69)
 bruh.Push(MyStruct{name: "mohammed" , age: 20})

*/
func (me *Vec[T]) Push(v T) {
	me.Extend(0)                   // Extend by 1 for the new element
	*me.Data = append(*me.Data, v) // Append to the slice
	me.len++
}

/*
	#

Extend extends the capacity of the vector.

	`by` : is optional and specifies the amount by which the capacity should be increased.
	If 'by' is 0, the capacity will be increased by its current value multiplied by the Updator.
	No return value.
*/
func (me *Vec[T]) Extend(by int) {
	if me.len == me.Capacity {
		if by == 0 {
			by = int(float32(me.Capacity) * me.Updator)
		}
		me.Capacity += by

		newData := make([]T, me.len, me.Capacity)
		copy(newData, *me.Data) // Copy the old data to the new slice
		me.Data = &newData
	}
}

func (me *Vec[T]) RemoveByIndex(index int) {
	if index < 0 || index >= me.len {
		return
	}
	me.len--
	copy((*me.Data)[index:], (*me.Data)[index+1:])
}

func (me *Vec[T]) String() string {
	return fmt.Sprintf("Capacity: %d, Updator: %.2f, len: %d, Data: %v", me.Capacity, me.Updator, me.len, *me.Data)
}

func (me *Vec[T]) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, "Vec[%d/%d] = ", me.len, me.Capacity)
	for i := 0; i < me.len; i++ {
		fmt.Fprintf(f, "%v ", (*me.Data)[i])
	}
}
