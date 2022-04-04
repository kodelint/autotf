### autotf

`autotf` is a demo `tool` created as POC for [Blog](). Please read the blog for more context

#### `Init` and `Plan`
```bash
./autotf verify stage/s3/autotf-testing01.tfvars
```

#### `Init`, `Plan` and `Apply`
```bash
./autotf deploy stage/s3/autotf-testing01.tfvars
```

#### `Init` and `Plan -destroy`
```bash
./autotf deploy stage/s3/autotf-testing01.destroy
```