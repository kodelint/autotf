### autotf

`autotf` is a demo `tool` created as POC for [Blog](https://medium.com/@email2sroy/terraform-solve-dynamic-backend-problem-with-golang-85d381bc48b5). Please read the blog for more context

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