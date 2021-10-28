# TwoStack library

TwoStack library is a Golang implementation of 2-dimentional double-ended stack data primitive. The core of this data structure is a double-ended stack With "Front" node and "Back" node. Each node of this stack, called a 
"Global stack" is a double-ended stack, again with "Front" and "Back". This stack is called a "Cell stack"

![TwoStack](Documentation/twostack.png)

## How two-dimentional stack works.

Two-dimentional stack is working exactly like a common stack that you are accustom to. Depending on the stack mode, that could be "Normal" or "Reverse" data is pushed to the back or the front of the Cell stack which are at the back or front of the Global stack. At each particular moment, you are working with a single Cell Stack, but you can rotate Global stack in ether left or right direction. .Set or .Put operation will store data to the Cell stack, .Add will add a new Cell stack to the Global stack, .Left or .Right will rotate global stack, .Get or .G will read data from the Cell stack and .Take or .T will take data from ether front or back of the stack.

## Data type layer of TwoStack

There are "raw" or "typed" access to data stored in TwoStack data structure. .Set/.Get/.Take operates with stack in the "raw" mode, so you can store any data type or structure that you like. .Put/.G/.T are used a "datatype layer" of the TwoStack and stored typed data, represented by special structure as a values in the stack. This structure and associated functions is just a "sugar" which takes away details of how you will handle a supported data types. Currently "data type layer" supports boolean, int64, int, float64, float32 and string. In addition to the data, "data type layer" also stored the name of the value and optional labels.

## Show me the code

### Create TwoStack, store and then retrive the raw values.

```
import (
    "fmt"
    . "github.com/vulogov/twostack"
)
func main() {
    ts := Init()
    ts.Set("42")
    r, err := ts.Get()
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    } else {
        fmt.Printf("Value: %v\n", r)
    }
}
```
In this sample, we are storing a raw string in the node of the Cell stack and then reading it. Operation .Get doesn't change the state of the stack (unlike .Take) and the node storing that string still stays at the end of the stack (since our stack is in .Normal mode). As we see, the .Get or .Take operations will return an error as well as a value.

### Show me how to work with stack mode

```
import (
    "fmt"
    . "github.com/vulogov/twostack"
)
func main() {
    ts := Init()
	ts.Set("42")
	ts.Set("41")
    // Pushing data to stack, 42 followed by 41
	ts.Reverse()  // Changing mode. Will read from the front
	r, _ := ts.Get()
    // Now, we will expect 42 since we are reading from the front of the stack
	ts.Normal()   // Changing mode. Will read from the back
	r, _ = ts.Get()
    // Now we will expect 41 as we are reading from the back.
}
```
While looking at this sample, you have to keep in mind few things:

- TwoStack is in Normal mode by default. Cell stack is accessing the back of the Cell stack and Global stack is also accessing the back node.
- .Normal and .Reverse change TwoStack behavor globally
- TwoStack is essentially LIFO stack.

### Show me how to operate with different Cell stacks

```
import (
    "fmt"
    . "github.com/vulogov/twostack"
)
func main() {
    ts := Init()
    ts.Set("42")     // Pushing data to the Cell stack
	ts.Add()         // Adding new Cell stack to the back of global stack
	ts.Set("41")     // Pushing data to the second Cell stack
	ts.Left()        // Rotating global stack to the left
	r, _ := ts.Get()    // Expecting to read "42" here since after rotation, Cell stack with "42" become at the back of the global stack
	ts.Right()          // Rotating global stack to the right. Cell stack with "41" becomes the back node with global stack
	r, _ = ts.Get()	 // Now, we are expecting to have "41"
}
```
Here, we haved to note a few things:
- .Left and .Right will rotate global stack. Cell stacks will remain unchanged.
- Rotation of the stack doesn't effect the TwoStack mode.
- .Add have effect over global stack. Global stack is a storage for the Cell stacks. Cell stack is a storage for the data.

### Show me, how to operate with "data type layer"

As I've mentioned before, "data type layer" is just a "sugar" designed to simplify handling of the typed data, by taking all conversions to and from interface{} (that is what Cell stack storing as data).

```
import (
    "fmt"
    . "github.com/vulogov/twostack"
)
func main() {
    ts := Init()
    if !ts.Put(3.14, "pi") {
		// Putting float value with name "pi" had failed
	}
	r, err := ts.G() // Expecting to have 3.14   
}
```
You may note a few things here:

- .Put is different from .Set as it will store the data alongside with necessary information for conversion
- "data type layer" do have a name for each value. If you do not care about name, pass an empty string ""
- If you have a string representation of your data, "data type layer" provides you with .MakeInt, .MakeBool, MakeFloat, .MakeString primitives.
- If you know data type at the top of the stack, you can use .GetBool, .GetInt, .GetFloat and .GetString primitives.

```
import (
    "fmt"
    . "github.com/vulogov/twostack"
)
func main() {
    ts := Init()
    ts.MakeInt("42", "answer")    // Making value from the string and storing it in TwoStack 
	res, err := ts.GetInt()       // We do know what data type are at the top of the stack
	// Expecting to have 42
}
```
## Applying functions to data in Cell Stack

TwoStack provides you with convinient interface for applying custom functions to the data stored in Cell Stack. There are three types of functions:

- Generator. This function doesn't require any data to be in stack, but will add single element to the stack.
- Function. This function takes on element from the stack, perform operation with this element and push result back to stack as single element.
- Operator. This function takes two elements from stack, perform operations with those two elements and push result as single element back to stack.

### Show me the code

```
import (
    "fmt"
    . "github.com/vulogov/twostack"
)
func main() {
    ts := Init()
	ts.Put(42)
	_, err := ts.ApplyFun(NumericIncrease)  // Applying function NumericIncrease to the value on top of the stack. This function will convert numeric values to float and increase value to 1.0 
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	res, _ := ts.G()
	// Expecing res to be 43.0
}
```

If you want to apply functions to values in Global stack, you can perform .Gzip() operation.

## What is practical application for TwoStack data structure ?

Two-dimentional stack designed to overcome a few issues (as I see them) with one-dimentional stacks.

- Data isolation. Every time when you are storing data for the later use them in the stack, you can end-up in the situation, when you badly-behaving code will not only get an access to the data that it should, but beyond that. TwoStack solved that problem. You can store the data in different Cell stacks for a different computations.
- Arguments isolation. With languages of the FORT family, defining the arity of the function is very important step, as you do not want that function takes a deeper dive to the stack than it should. This is an extention of the Data isolation.
- More natural representation of the matrixes or . You can have a separate stack with data for each row of the data.

## Installation.

Just do it:

```
go get github.com/vulogov/twostack
```

Also, feel free to clone/branch source code:

```
git clone https://github.com/vulogov/twostack
```

Running make in the cloned directory will download all required modules and run tests.
