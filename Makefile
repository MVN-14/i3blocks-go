b-i3spotify:
	go build -C i3spotify -o i3spotify main.go

b-i3date:
	go build -C i3date -o i3date  main.go

b-i3time:
	go build -C i3time -o i3time  main.go

b-i3battery:
	go build -C i3battery -o i3battery  main.go

b-i3volume:
	go build -C i3volume -o i3volume  main.go

b-all:
	make b-i3date && make b-i3time && make b-i3battery && make b-i3volume	&& make b-i3spotify
