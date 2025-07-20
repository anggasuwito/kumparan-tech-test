run:
	go run .

generate-mock:
	mockgen -package=mockrepository -source=./internal/repository/article.go -destination=mock/repository/mock_article.go
	mockgen -package=mockrepository -source=internal/repository/author.go -destination=mock/repository/mock_author.go

	mockgen -package=mockusecase -source=internal/usecase/article.go -destination=mock/usecase/mock_article.go

test:
	go test ./...