# Профилирование

## В **соседней** консоли команду:

go tool pprof -http=":9090" -seconds=30 http://localhost:8080/debug/pprof/profile
