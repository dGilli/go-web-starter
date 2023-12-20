run:
	@tailwindcss -i assets/css/styles.css -o tmp/css/styles.css 
	@templ generate
	@go run main.go
