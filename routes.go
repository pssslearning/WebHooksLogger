// routes.go

package main

func initializeRoutes() {

	// Handle the index route
	router.GET("/", showRootGET)
	router.POST("/", showRootPOST)

	// Handle the API samples route
	router.GET("/api/samples", showApiSamplesGET)
	router.POST("/api/samples", showApiSamplesPOST)

	// Handle the PAY-IN Webhook route
	router.POST("/webhook/event", showWebhookEventPOST)
}
