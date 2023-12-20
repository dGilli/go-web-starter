run:
	@tailwindcss -i assets/css/styles.css -o tmp/css/styles.css 
	@templ generate
	@go run main.go

watch: ; ${MAKE} -j3 watch-templ watch-tailwind watch-main
watch-templ:
	@templ generate --watch
watch-tailwind:
	@tailwindcss -i assets/css/styles.css -o tmp/css/styles.css --watch
watch-main:
	@air --build.cmd "go build -o tmp/main ." --build.bin "./tmp/main"
