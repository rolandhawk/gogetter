# gogetter
Like Getter functions of generated protobuf code? You can simply create yours

## Installation
```bash
go install github.com/rolandhawk/gogetter@v0.0.1
```

## Usage
```
//go:generate gogetter StructName1 StructName2
```

Flags:
| Name | Default | Description |
|------|---------|-------------|
| out | `{source}_getter.go` | Custom output file name. Format: `*.go`. Default: `{source}_getter.go` |
| dry | false | Dry run. If true, it will print the output to stdout instead of write to file. |

See [example](https://github.com/rolandhawk/gogetter/blob/master/example/example.go#L1) for actual use.

## Limitation

Currently `gogetter` is not smart enough to know underlying type of named type. So it will always assume
underlying type of named type is struct. It become a problem because `gogetter` will generate wrong zero value
for the getter method. For now, you can work around this problem by specifying custom zero value map using `-custom` 
flag options.