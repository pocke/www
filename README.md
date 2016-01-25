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

![www-basic](http://cdn-ak.f.st-hatena.com/images/fotolife/P/Pocke/20160125/20160125120042.gif)


VS. python3 -m http.server
-------------------------

|              | python         | www |
| ---          | -----          | --- |
| port         | 8000(Optional) | Random |
| browser open | Manual         | automatically |
| command name | long           | short |


Re-open the browser
---------


![www-reload](http://cdn-ak.f.st-hatena.com/images/fotolife/P/Pocke/20160125/20160125120207.gif)

License
-------

These codes are licensed under CC0.

[![CC0](http://i.creativecommons.org/p/zero/1.0/88x31.png "CC0")](http://creativecommons.org/publicdomain/zero/1.0/deed.en)
