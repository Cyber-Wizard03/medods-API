package server



func Start() {
	
	// Create default gin Router , And has attached Logger and Recovery middleware
	router := setRouter()

	// log.Fatal(http.ListenAndServe(":"+port, router))

	router.Run("127.0.0.1:5000")
}
