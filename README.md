Installation
---------------

```sh
go get github.com/pocke/www
```

You download binary from [latest release](https://github.com/pocke/www/releases/latest), and place it in `$PATH` directory.

Usage
---------

```sh
cd DOCUMENT_ROOT
www
```

HTTP Server start serving for static files at random port. And browser is automatically opened.

```sh
$ www --help
Usage of www:
  -p, --port=0: TCP port number
```

VS. python3 -m http.server
-------------------------

|              | python         | www |
| ---          | -----          | --- |
| port         | 8000(Optional) | Random |
| browser open | Manual         | automatically |
| command name | long           | short |

License
-------

These codes are licensed under CC0.

[![CC0](http://i.creativecommons.org/p/zero/1.0/88x31.png "CC0")](http://creativecommons.org/publicdomain/zero/1.0/deed.en)
