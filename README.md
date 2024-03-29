go-abspath
==========

`go-abspath` is a command line tool that prints the absolute paths of all
given files. File names can be piped via `STDIN` or given as arguments.
This implementation is in Golang. For a Python version, check out
[arne-cl/abspath](https://github.com/arne-cl/abspath).

Usage
-----

Print the absolute paths of all paths given as arguments:

    abspath file1.txt path/to/file2.pdf

Print the absolute paths of all files and directories in `./Desktop`:

    abspath Desktop/*

Print the absolute paths of all files and directories in `Desktop`
**and** its subdirectories:

    abspath -r Desktop

Print the absolute paths of all `.pdf` files in `~/Documents`

    find ~/Documents -name *.pdf | abspath


Installation
------------

    go get github.com/arne-cl/go-abspath
    go build -o abspath github.com/arne-cl/go-abspath
    
This will get you the binary, that you can call as `./abspath`.
You can make this available on your `$PATH`, e.g. like this:

    sudo mv abspath /usr/local/bin/


License
-------

MIT
