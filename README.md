This is an idea for a go Writer implementation that wraps another
Writer implementation, but upon construction takes a function as a
parameter. This function is called on the data each time the Write
method is called. This allows the user to 'spy' on the information
in order to come up with some kind of aggregate conclusion.

For example, one could spy on the data to count total bytes.

You need to be able to have enough memory to store what you read.
Keep in mind that the choice of function could impact the O(1)
memory consumption and O(N) write speed of the Writer.
