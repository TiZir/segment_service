wait-for "${DB_HOST}:${DB_PORT}" -- "$@"

# Watch your .go files and invoke go build if the files changed.
CompileDaemon --build="go build -o main main.go"  --command=./main

# if docker volume inspect segment_service_db_data > /dev/null 2>&1; then
#     docker volume rm segment_service_db_data
# fi