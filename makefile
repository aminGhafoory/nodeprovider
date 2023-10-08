run:
	go build -o out && ./out
sqlc:
	sqlc generate
gooseup:
	cd sql/schema && goose postgres postgres://postgres:amin235711@amin-laptop.local:5432/nodeProvider up
goosedown:
	cd sql/schema && goose postgres postgres://postgres:amin235711@amin-laptop.local:5432/nodeProvider down