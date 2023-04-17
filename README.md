
# About

1) Golang can arguably be a good drop-in replacement for python for many
scripting-esque problems.

2) Using Golang results in significant performance gains for not much less readability and dev speed.

3) One point in favour of still using Python however is the 'easy to reach' builtin data structures.

go-builtin-collections attempts to provide a similar set of helper data-structures that can speed up casual scripting.

## OK truthfully this was just a experiment to try out generics

If you have any ideas on useful things worth adding, please do suggest.

# Main additions currently

1. Deque ()

2. Heap (both min and max via a comparison function)

<code>
H := NewHeap[int](func(a, b int) bool { return a < b })

H.Push(100)

H.Push(1)

H.Push(18)

H.Push(1000)

fmt.Println(H.Heap) 

a, _ := H.Pop()

fmt.Println(H.Heap)

b, _ := H.Pop()

fmt.Println(H.Heap)

c, _ := H.Pop()

fmt.Println(H.Heap)
</code>