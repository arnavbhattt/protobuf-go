# Protocol Buffers

- JSON: used for when you can't control clients (because it's more standardized)
- Protobuf: where you control clients, because it's more productive
  - type safety, prevents schema violations, fast serialization
- It's essentially good for communicating w/ two systems
- We can define objects in Protobuf and compile it into Go code
- This gives us consistent schemas -> general data model used for rest of system
- Also allows us to maintain versioning and backwards compat. (due to field numbering and reserved type)

- Summary: Less boilerplate, Extensibility, Language Agnosticism, Performance, Used by gRPC
