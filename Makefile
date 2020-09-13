.DEFAULT_GOAL := run
.SILENT: # silence!!

export PATH := ./bin:$(PATH)

clean:
	rm -rf ./public || true
	rm -rf content/posts/* static/public/images/* || true

setup:
	go mod tidy
	mkdir -p bin
	curl -sL https://install.goreleaser.com/github.com/ValeLint/vale.sh | bash
	curl -sL https://htmltest.wjdp.uk | bash
	chmod +x ./bin/*

run:
	hugo server --watch --buildFuture --cleanDestinationDir

ci: refresh
	hugo
	vale --glob='**/*.md' .
	htmltest -c .htmltest.yaml ./public

avatar:
	wget -O static/avatar.jpg https://github.com/caarlos0.png
	convert static/avatar.jpg \
		-bordercolor white -border 0 \
		\( -clone 0 -resize 16x16 \) \
		\( -clone 0 -resize 32x32 \) \
		\( -clone 0 -resize 48x48 \) \
		\( -clone 0 -resize 64x64 \) \
		-delete 0 -alpha off -colors 256 static/favicon.ico
	convert -resize x120 static/avatar.jpg static/apple-touch-icon.png

og:
	wget -O static/og-image.png "https://og.caarlos0.dev/Carlos%20Becker%20%7C%20**caarlos0**.png?theme=light&md=1&fontSize=100px&images=https://github.com/caarlos0.png"

refresh: clean
	go run cmd/notion/main.go
