# Set 1 Notes

## Challenge 1

- first make string into a set of bytes
- encode bytes using standard encoding
- copy to and from dest

1. Convert the value of the input from a string to byte array, simple as ascii table to values
2. base64 the byte array
   1. Basically base64 takes each individual number, converts to binary representation lined up against each other and takes in 6-bit increments and evaluates the integer value
   2. That integer is then converted to Ascii and then shifts to the next 6 bits until end.
   3. end is padded with "=" for anything left over


## Challenege 4

