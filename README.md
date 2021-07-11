Gemini Fileserver for I2P
=========================

This is a static fileserver that speaks the Gemini protocol
that defaults to only using I2P connections, while retaining
all of Gemini's TLS-based authentication features. Just make
sure that it has a directory of files to share(`./www` by
default, pass the `-files` flag to change it) and i2p-gemini
will take care of the rest.

Anyone with a Go toolchain installed should be able to install
with `go get -u i2pgit.org/idk/i2p-gemini`

        Usage of ./i2p-gemini:
          -certs string
            Directory where server certificates(TLS) will be stored (default "var/lib/gemini/certs")
          -files string
            Directory of files to serve up with Gemini (default "www")
          -name string
            Name of the service to pass to SAM (default "i2pgemini")
          -sam string
            SAM API to connect to and user (default "127.0.0.1:7656")
