package alert_api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	alert_biz "github.com/wangxin688/narvis/server/features/alert/biz"
	"github.com/wangxin688/narvis/server/features/alert/schemas"
	"github.com/wangxin688/narvis/server/tools/errors"
	"github.com/wangxin688/narvis/server/tools/helpers"
	ts "github.com/wangxin688/narvis/server/tools/schemas"
)

// @Tags Alert
// @Summary Create Maintenance
// @X-func {"name": "CreateMaintenance"}
// @Description Create Maintenance
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param data body schemas.MaintenanceCreate true "data"
// @Success 200 {object} ts.IdResponse
// @Router /alert/maintenances [post]
func createMaintenance(c *gin.Context) {
	var err error

	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()

	var maintenance schemas.MaintenanceCreate
	if err = c.ShouldBindJSON(&maintenance); err != nil {
		return
	}
	newMt, err := alert_biz.NewMaintenanceService().CreateMaintenance(&maintenance)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, ts.IdResponse{Id: newMt})
}

// @Tags Alert
// @Summary Get Maintenance
// @X-func {"name": "GetMaintenance"}
// @Description Get Maintenance
// @Security BearerAuth
// @Accept json
// @Produce json
// @Success 200 {object} schemas.Maintenance
// @Router /alert/maintenances/{id} [get]
func getMaintenance(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	id := c.Param("id")
	if err = helpers.ValidateUuidString(id); err != nil {
		return
	}

	mts, err := alert_biz.NewMaintenanceService().GetById(id)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, mts)
}

// @Tags Alert
// @Summary List Maintenances
// @X-func {"name": "ListMaintenances"}
// @Description List Maintenances
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param object query schemas.MaintenanceQuery true "query"
// @Success 200 {object} ts.ListResponse{results=[]schemas.Maintenance}
// @Router /alert/maintenances [get]
func listMaintenances(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()

	var query schemas.MaintenanceQuery
	if err = c.ShouldBindQuery(&query); err != nil {
		return
	}
	count, mts, err := alert_biz.NewMaintenanceService().ListMaintenances(&query)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, ts.ListResponse{Total: count, Results: mts})
}

// @Tags Alert
// @Summary Update Maintenance
// @X-func {"name": "UpdateMaintenance"}
// @Description Update Maintenance
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param data body schemas.MaintenanceUpdate true "data"
// @Success 200 {object} ts.IdResponse
// @Router /alert/maintenances/{id} [put]
func updateMaintenance(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	id := c.Param("id")
	if err = helpers.ValidateUuidString(id); err != nil {
		return
	}
	var maintenance schemas.MaintenanceUpdate
	if err = c.ShouldBindJSON(&maintenance); err != nil {
		return
	}
	err = alert_biz.NewMaintenanceService().UpdateMaintenance(id, &maintenance)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, ts.IdResponse{Id: id})
}

// @Tags Alert
// @Summary Delete Maintenance
// @X-func {"name": "DeleteMaintenance"}
// @Description Delete Maintenance
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} ts.IdResponse
// @Router /alert/maintenances/{id} [delete]
func deleteMaintenance(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	id := c.Param("id")
	if err = helpers.ValidateUuidString(id); err != nil {
		return
	}
	err = alert_biz.NewMaintenanceService().DeleteMaintenance(id)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, ts.IdResponse{Id: id})
}
