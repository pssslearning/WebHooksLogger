// routes.go

package main

func initializeRoutes() {

	// Handle the index route
	router.GET("/", showRootGET)

	// Handle the index route
	router.POST("/", showRootPOST)

	// Handle the PAY-IN Webhook route
	router.POST("/webhook/event", showWebhookEventPOST)
}
