### Introduction
JavaScript Object Notation (full form of JSON) is a standard file format used to interchange data.
The data objects are stored and transmitted using key-value pairs and array data types.

JSON is a minimal, readable format for this structured data. It is extremely lightweight and
thus is used in the interaction between server and client (HTTP requests).
JSON is an alternative to XML, the latter being more complex.
Using JSON, data can be managed and handled logically. The extension of a JSON file is .json

### Data Types for JSON:
- Number (integer or decimal)
- String (should be in double quotes)
- Boolean (i.e. true or false)
- Null
- Objects
- Arrays
  An object is a collection of key-value pairs defining a data set; while an array is an ordered
  list of values. JSON objects can contain multiple JSON arrays and vice-versa.

### Rules for valid JSON:
- All JSON data is written inside curly braces
- JSON data is represented as key-value pairs
- Key should always be enclosed in double quotes
- Always separate Key and value with a colon (:)
- Values should be written appropriately; for String double quotes should be used while for Numbers, Boolean quotes should not be used.
- To separate data commas (,) should be used
- For array square brackets and objects, curly braces should be used

### Serialization
Serialization is a mechanism that converts the data into a byte stream (or byte string).
Once the other machine receives this byte stream it needs to "deserialize" the byte stream back
to its data so that it can be used further.

So, who exactly does this serialization-deserialization?
It is the programming language you use that takes care of this process. Go natively supports
unmarshalling and marshalling of JSON.

Data in APIs are commonly transferred in the JSON format because “It is easy for humans to read and write”
and “It is easy for machines to parse and generate” because JSON data is a light weight and can be parsed
quickly by the server.