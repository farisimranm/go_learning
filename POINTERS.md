# Go Pointer Cheat Sheet

## 1. Syntax & Operation Matrix

| Operation | Syntax | What it does |
| --- | --- | --- |
| **Pointer Type** | `*MyStruct` | Declares a variable holds a memory address. |
| **Address-of** | `&myStruct` | Generates the memory address of an existing value. |
| **Implicit Direct** | `myStruct.Field` | Automatically resolves a pointer to access its fields. |
| **Explicit Dereference** | `*myStruct` | Reaches past the address to manipulate the underlying data block. |

---

## 2. Calling a Function: Passing `&myStruct` vs `myStruct`

```plaintext
PASSING BY VALUE (myStruct)        PASSING BY POINTER (&myStruct)
    [Main Memory]                      [Main Memory]
    user := User{...}                  user := User{...} 
         │                                  ▲
         ▼ (Deep copy)                      │ (Memory Address)
    func(u User){}                     func(u *User){}
```

### Pass the Address (`&myStruct`)

* **When the function signature requires a pointer:** `func Process(u *User)`
* **Use Cases:**
  * **Mutation:** You need the function to modify the caller's original data.
  * **Efficiency:** The struct is large. Passing an 8-byte pointer avoids copying large data blocks into memory.
  * **Safety:** The struct contains a synchronization primitive (e.g., `sync.Mutex`). **Must** be passed by pointer; copying a mutex breaks its state.

### Pass Directly (`myStruct`)

* **When the function signature requires a value:** `func Process(u User)`
* **Use Cases:**
  * **Immutability:** You want to guarantee the function absolutely cannot modify your original data.
  * **Small Objects:** The struct is tiny (e.g., `time.Time` or `Point{X, Y int}`). It is faster for the CPU to copy these on the stack than to resolve a pointer address.

---

## 3. Inside a Function: Using `u.Field` vs `*u`

When your function already accepts a pointer parameter (`u *User`):

### Use Directly (`u.Name`)

* **When to use:** 99% of the time for reading or modifying individual fields, or calling methods.
* **Why:** Go supports automatic dereferencing. Writing `u.Name = "Alice"` automatically compiles to `(*u).Name = "Alice"`. Manual `*` syntax is redundant here.

### Explicitly Dereference (`*u`)

* **When to use:** Only in two rare structural scenarios:

1. **Whole-Object Assignment:** You want to overwrite the entire memory block with a new struct at once: `*u = User{Name: "New"}`.

2. **Cloning:** You want to copy the pointer's data out into a standalone, independent variable: `localCopy := *u`.

---

## 4. Pointers Haram List

In Go: slices, maps, and interfaces are specialized descriptors. **Do not use their pointers as function arguments**

### ❌ Never pass `*[]Type` (Pointer to Slice)

* **Standard Practice:** Pass slices directly as values (`mySlice []string`).
* **Why:** A slice is already a lightweight, 24-byte header containing a pointer to the underlying data array. Passing `*[]string` creates a redundant pointer-to-a-pointer layout.

### ❌ Never pass `*map[K]V` (Pointer to Map)

* **Standard Practice:** Pass maps directly as values (`myMap map[string]int`).
* **Why:** A map variable in Go is inherently a pointer to a runtime hash table descriptor under the hood. Passing a pointer to a map adds useless overhead.

### ❌ Never pass `*Interface` (Pointer to Interface)

* **Standard Practice:** Pass interfaces directly as values (`r io.Reader`).
* **Why:** An interface in Go is a two-word structure designed to hold a pointer to the underlying concrete value and its type information. Passing `*io.Reader` boxes the wrapper itself, breaking Go's type-assertion system.
