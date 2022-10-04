/*
 * Namf_Location
 *
 * AMF Location Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package location

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/free5gc/amf/logger"
	"github.com/free5gc/amf/producer"
	"github.com/free5gc/http_wrapper"
	"github.com/nycu-ucr/openapi"
	"github.com/nycu-ucr/openapi/models"
)

// ProvideLocationInfo - Namf_Location ProvideLocationInfo service Operation
func HTTPProvideLocationInfo(c *gin.Context) {
	var requestLocInfo models.RequestLocInfo

	requestBody, err := c.GetRawData()
	if err != nil {
		problemDetail := models.ProblemDetails{
			Title:  "System failure",
			Status: http.StatusInternalServerError,
			Detail: err.Error(),
			Cause:  "SYSTEM_FAILURE",
		}
		logger.LocationLog.Errorf("Get Request Body error: %+v", err)
		c.JSON(http.StatusInternalServerError, problemDetail)
		return
	}

	err = openapi.Deserialize(&requestLocInfo, requestBody, "application/json")
	if err != nil {
		problemDetail := "[Request Body] " + err.Error()
		rsp := models.ProblemDetails{
			Title:  "Malformed request syntax",
			Status: http.StatusBadRequest,
			Detail: problemDetail,
		}
		logger.LocationLog.Errorln(problemDetail)
		c.JSON(http.StatusBadRequest, rsp)
		return
	}

	req := http_wrapper.NewRequest(c.Request, requestLocInfo)
	req.Params["ueContextId"] = c.Params.ByName("ueContextId")

	rsp := producer.HandleProvideLocationInfoRequest(req)

	responseBody, err := openapi.Serialize(rsp.Body, "application/json")
	if err != nil {
		logger.CommLog.Errorln(err)
		problemDetails := models.ProblemDetails{
			Status: http.StatusInternalServerError,
			Cause:  "SYSTEM_FAILURE",
			Detail: err.Error(),
		}
		c.JSON(http.StatusInternalServerError, problemDetails)
	} else {
		c.Data(rsp.Status, "application/json", responseBody)
	}
}

// ProvidePositioningInfo - Namf_Location ProvidePositioningInfo service Operation
func HTTPProvidePositioningInfo(c *gin.Context) {
	logger.LocationLog.Warnf("Handle Provide Positioning Info is not implemented.")
	c.JSON(http.StatusOK, gin.H{})
}
