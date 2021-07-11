
all: linux osx windows file index

linux:
	go build -tags netgo -o www/i2p-gemini

osx:
	GOOS=darwin go build -tags netgo -o www/i2p-gemini-osx

windows:
	GOOS=windows go build -tags netgo -o www/i2p-gemini.exe

file:
	@echo 'Gemini Fileserver for I2P' > README.md
	@echo '=========================' >> README.md
	@echo '' >> README.md
	@echo 'This is a static fileserver that speaks the Gemini protocol' >> README.md
	@echo 'that defaults to only using I2P connections, while retaining' >> README.md
	@echo "all of Gemini's TLS-based authentication features. Just make" >> README.md
	@echo 'sure that it has a directory of files to share(`./www` by' >> README.md
	@echo 'default, pass the `-files` flag to change it) and i2p-gemini' >> README.md
	@echo 'will take care of the rest.' >> README.md
	@echo '' >> README.md
	@echo 'An instance of it is sometimes available at this address.' >> README.md
	@echo '' >> README.md
	@echo ' - `gemini://3sswn7xy4v4ev2q2o3r3cr5hxsoowbf5i6wksqxebpuv3zsva3nq.b32.i2p`' >> README.md
	@echo '' >> README.md
	@echo 'Anyone with a Go toolchain installed should be able to install' >> README.md
	@echo 'with `go get -u i2pgit.org/idk/i2p-gemini`' >> README.md
	@echo 'file www/i2p-gemini www/i2p-gemini-osx www/i2p-gemini.exe' >> README.md
	@echo '' >> README.md
	@echo '```' >> README.md
	./www/i2p-gemini -h 2>&1 | tee -a README.md
	@echo '```' >> README.md
	@echo '' >> README.md
	@echo '[The source code is a single `.go` file](server.go), requiring these' >> README.md
	@echo '[modules](go.mod). You can download a static binary for linux here:' >> README.md
	@echo '[i2p-gemini](i2p-gemini), for OSX here: [i2p-gemini-osx](i2p-gemini-osx)' >> README.md
	@echo 'and for Windows here [i2p-gemini.exe](i2p-gemini.exe)' >> README.md
	@echo '' >> README.md
	@echo '```' >> README.md
	file www/i2p-gemini >> README.md
	sha256sum www/i2p-gemini >> README.md
	@echo '' >> README.md
	file www/i2p-gemini-osx >> README.md
	sha256sum www/i2p-gemini-osx >> README.md
	@echo '' >> README.md
	file www/i2p-gemini.exe >> README.md
	sha256sum www/i2p-gemini.exe >> README.md
	@echo '```' >> README.md
	@echo '' >> README.md

index:
	mkdir -p www
	/home/idk/.local/bin/md2gemini README.md -l paragraph > www/index.gmi
	@echo "<!DOCTYPE html>" > www/index.html
	@echo "<html>" >> www/index.html
	@echo "<head>" >> www/index.html
	@echo "  <title>I2P Gemini Server</title>" >> www/index.html
	@echo "  <link rel=\"stylesheet\" type=\"text/css\" href =\"home.css\" />" >> www/index.html
	@echo "</head>" >> www/index.html
	@echo "<body>" >> www/index.html
	pandoc README.md >> www/index.html
	@echo "</body>" >> www/index.html
	@echo "</html>" >> www/index.html
	cp www/index.html index.html
	cp home.css www/home.css
	cp *.* www


