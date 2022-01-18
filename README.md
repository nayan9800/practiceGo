# practiceGo
## Go basics
in this the Go basics are covered are following

|No | Topic name     |Commit hash |
|--|-----------------|---------|
|1|Functions|13ad9c84f928def348c629872cc4eb3824e75e3c|
|2| Data types, type casting and constant |0b69127016db4ab066a75f6749388da7e0e4a3bd
|3|Control flow|a7e746c722634f73f616b3421f567e72fe254f59
|4|Pointers|2afab6a10205df5d4fd19c879fcc1d7a8c3cef54
|5|More types in Go|2a5802f5062187651e01b02974ca6e0289293b74
|6|Methods and Interfaces|c1ae5679e59d88ec13db269b2951ec31c1f0a24e
|7|Go Concurrency|762d749a270631a8dbdbe51bb0cd4310c8ea89a1
|8|Go Sync package|c441c385d64dc71315707c6665b56d2f6cd83ab7

## Go chitChat server
gochitchat server is simple tcp server build in go. in which **net** package is used to create the server. **map** is used to store current logged in users and the connection are handled concurrenctly using 
**goroutines**. to run server uncomment the last line in main.go **syncpract.Startserver()** 