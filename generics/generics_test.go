package generics

import "testing"

func TestStack(t *testing.T) {
	t.Run("integer stack", func(t *testing.T) {
		stackOfInts := new(Stack[int])

		AssertTrue(t, stackOfInts.IsEmpty())

		stackOfInts.Push(123)
		AssertFalse(t, stackOfInts.IsEmpty())

		stackOfInts.Push(1337)
		value, _ := stackOfInts.Pop()
		AssertEqual(t, value, 1337)
		value, _ = stackOfInts.Pop()
		AssertEqual(t, value, 123)
		AssertTrue(t, stackOfInts.IsEmpty())

		stackOfInts.Push(1)
		stackOfInts.Push(2)
		firstNum, _ := stackOfInts.Pop()
		secondNum, _ := stackOfInts.Pop()
		AssertEqual(t, firstNum+secondNum, 3)
	})
	t.Run("string stack", func(t *testing.T) {
		stackOfInts := new(Stack[string])

		AssertTrue(t, stackOfInts.IsEmpty())

		stackOfInts.Push("123")
		AssertFalse(t, stackOfInts.IsEmpty())

		stackOfInts.Push("1337")
		value, _ := stackOfInts.Pop()
		AssertEqual(t, value, "1337")
		value, _ = stackOfInts.Pop()
		AssertEqual(t, value, "123")
		AssertTrue(t, stackOfInts.IsEmpty())
	})
}
