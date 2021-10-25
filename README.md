# TwoStack library

TwoStack library is a Golang implementation of 2-dimentional double-ended stack data primitive. The core of this data structure is a double-ended stack With "Front" node and "Back" node. Each node of this stack, called a 
"Global stack" is a double-ended stack, again with "Front" and "Back". This stack is called a "Cell stack"

![TwoStack](Documentation/twostack.png)

## How two-dimentional stack works.

Two-dimentional stack is working exactly like a common stack that you are accustom to. Depending on the stack mode, that could be "Normal" or "Reverse" data is pushed to the back or the front of the Cell stack which are at the back or front of the Global stack. At each particular moment, you are working with a single Cell Stack, but you can rotate Global stack in ether left or right direction. .Set or .Put operation will store data to the Cell stack, .Add will add a new Cell stack to the Global stack, .Left or .Right will rotate global stack, .Get or .G will read data from the Cell stack and .Take or .T will take data from ether front or back of the stack.

## Data type layer of TwoStack

There are "raw" or "typed" access to data stored in TwoStack data structure. .Set/.Get/.Take operates with stack in the "raw" mode, so you can store any data type or structure that you like. .Put/.G/.T are used a "datatype layer" of the TwoStack and stored typed data, represented by special structure as a values in the stack. This structure and associated functions is just a "sugar" which takes away details of how you will handle a supported data types. Currently "data type layer" supports boolean, int64, int, float64, float32 and string. In addition to the data, "data type layer" also stored the name of the value and optional labels.

## Show me the code

### Create two stack, store and then retrive the raw values.

```
ts := Init()
ts.Set("42")
r, _ := ts.Get()
```
