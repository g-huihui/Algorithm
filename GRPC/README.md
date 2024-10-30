### proto生成
```cd proto```

```protoc --go_out=../ ./response.proto```

### 生成triple文件
```protoc --go_out=../ --go-triple_out=../ ./response.proto```

### 生成grpc文件
```protoc --go_out=../ --go-grpc_out=../ ./response.proto```