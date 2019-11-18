# booked

**booked** is a small tool that offers visualization of Facebook export data.

To use **booked** you'll need a copy of your Facebook data. Follow [these
instructions](https://www.facebook.com/help/1701730696756992) and download a
copy of your data in the JSON format. Currently only visualization for
messages is implemented so you can select only Messages for export to reduce
the size of the download (which may be large).

To use **booked** you'll need a working [Go](https://golang.org/)
installation. Follow the instructions at
[golang.org](https://golang.org/doc/install) to install it. With a working
Go installation, **booked** can be installed by running the following on the
command line:

```bash
go get go.tmthrgd.dev/booked
```

Once installed, **booked** can be run by running the following on the command
line:

```bash
booked <path-to-facebook-json.zip>
```

**booked** must be passed the path to the Facebook export data zip you
downloaded earlier. All data will be extracted from the zip file. Once started,
the UI will be launched in your browser where you can navigate through your
data.

## License

Unless otherwise noted, the spork source files are distributed under the
Modified BSD License found in the LICENSE file.
