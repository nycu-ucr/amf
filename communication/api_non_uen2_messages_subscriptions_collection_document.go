/*
 * Namf_Communication
 *
 * AMF Communication Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package communication

import (
	"github.com/nycu-ucr/gonet/http"

	"github.com/gin-gonic/gin"

	"github.com/nycu-ucr/amf/logger"
)

// NonUeN2InfoSubscribe - Namf_Communication Non UE N2 Info Subscribe service Operation
func HTTPNonUeN2InfoSubscribe(c *gin.Context) {
	logger.CommLog.Warnf("Handle Non Ue N2 Info Subscribe is not implemented.")
	c.JSON(http.StatusOK, gin.H{})
}
