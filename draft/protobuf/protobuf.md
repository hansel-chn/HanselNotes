# Protobuf

## Protobuf更快更小的原因

* 用编号代替key值

```protobuf
syntax = "proto3";
message SearchRequest {
  string query = 1;
  int32 page_number = 2;
  int32 result_per_page = 3;
}
```