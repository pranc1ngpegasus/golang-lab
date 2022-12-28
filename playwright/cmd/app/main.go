package main

func main() {
	app, err := initialize()
	if err != nil {
		panic(err)
	}

	if err := app.server.ListenAndServe(); err != nil {
		app.logger.Error(app.ctx, "failed to start server", app.logger.Field("err", err))
	}

	if err := app.tracer.Shutdown(app.ctx); err != nil {
		app.logger.Error(app.ctx, "failed to shutdown tracer", app.logger.Field("err", err))
	}
}
