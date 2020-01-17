# Letter ID

This will generate a series of IDs given A-Z or a-z ASCII Input.

```
letterid.NextId("a124") // Error
letterid.NextId("a")    // B
letterid.NextId("A")    // B
letterid.NextId("Z")    // AA
letterid.NextId("ZA")   // AB
letterid.NextId("ZZ")   // AAA
letterid.NextId("ZAA")  // ABA
letterid.NextId("ZZA")  // AAB
```