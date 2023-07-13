package linux

// func InitLinuxRoutes(routeGroup *gin.RouterGroup) {
// 	r := routeGroup.Group("/linux").Use(middleware.Auth())
// 	r.POST("/command/execute", sendCommand)
// 	r.POST("/script/execute", executeScript)
// }

//Run command on instance
// func sendCommand(c *gin.Context) {
// 	logger.Log.Info("IN:sendCommand")
// 	input := model.RunCommand{}
// 	err := c.Bind(&input)
// 	if err != nil {
// 		logger.Log.Sugar().Errorf("Error binding data", err)
// 		c.JSON(http.StatusExpectationFailed, err)
// 	}
// 	out, err := sendCommandService(input)
// 	if err != nil {
// 		logger.Log.Sugar().Errorf("Error executing command on instance", err)
// 		c.JSON(http.StatusExpectationFailed, err)
// 	}
// 	logger.Log.Info("OUT:sendCommand")
// 	c.JSON(http.StatusOK, out)
// }

//Ececute shell script on instance
// func executeScript(c *gin.Context) {
// 	logger.Log.Info("IN:executeScript")

// 	var input model.Executable

// 	err := c.Bind(&input.Script)
// 	if err != nil {
// 		logger.Log.Sugar().Errorf("error binding data", err)
// 		c.JSON(http.StatusExpectationFailed, err)
// 		return
// 	}
// 	out, err := executeScriptService(input)
// 	if err != nil {
// 		logger.Log.Sugar().Errorf("Error executing command on instance", err)
// 		c.JSON(http.StatusExpectationFailed, err)
// 	}
// 	logger.Log.Info("OUT:executeScript")
// 	c.JSON(http.StatusOK, out.Output)
// }
