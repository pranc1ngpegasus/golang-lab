package main

func main() {
	app, err := initialize()
	if err != nil {
		panic(err)
	}

	if err := app.server.ListenAndServe(); err != nil {
		app.logger.Error("failed to boot server", err)
	}
}
