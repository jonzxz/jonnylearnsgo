# Strings

## Related Types
- byte -> synonym for uint8
- rune -> synonym for int32 for characters
- string -> **immutable** sequence of characters
  - physically sequence of bytes in UTF-8
  - logically sequence of unicode runes

- Rune is a Go equilvane of a "character" / or a "wide" character
  - the 4 byte is big enough to represent any unicode printable character
  - runes / character in GOlang is in single quote
