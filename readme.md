# subs

Grab valid subdomains from files!

### Installation
##### Linux and macOS

Run the ./install.sh script

```bash
dpkg-deb --build .build/binaries .
```

### Usage
```bash
subs [input_file] [flags]
```

### Flags
  -d, --domains   Get domains only
  -h, --help      Help
  -u, --unique    Only get unique entries

### Example

```bash
❯ subs test.txt
subdomain1.example.com
subdomain2.example.com
subdomain3.example.com
subdomain4.example.com
...
```

---

Copyright (c) 2024  Owais Shaikh

Licensed under [GNU GPL 3.0](LICENSE)