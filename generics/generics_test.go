package generics

import "testing"

func TestStack(t *testing.T) {
	t.Run("test stack of integers", func(t *testing.T) {
		intStack := StackOfInts{}

		AssertTrue(t, intStack.IsEmpty())

		intStack.Push(42)
		AssertFalse(t, intStack.IsEmpty())

		intStack.Push(1337)
		value, _ := intStack.Pop()
		AssertEqual(t, value, 1337)

		value, _ = intStack.Pop()
		AssertEqual(t, value, 42)
		AssertTrue(t, intStack.IsEmpty())
	})
	t.Run("string stack", func(t *testing.T) {
		myStackOfStrings := StackOfStrings{}

		AssertTrue(t, myStackOfStrings.IsEmpty())

		myStackOfStrings.Push("123")
		AssertFalse(t, myStackOfStrings.IsEmpty())

		myStackOfStrings.Push("456")
		value, _ := myStackOfStrings.Pop()
		AssertEqual(t, value, "456")
		value, _ = myStackOfStrings.Pop()
		AssertEqual(t, value, "123")
		AssertTrue(t, myStackOfStrings.IsEmpty())
	})
}
