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
// @Summary Create Subscription
// @Description Create Subscription
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param data body schemas.SubscriptionCreate true "data"
// @Success 200 {object} ts.IdResponse
// @Router /alert/subscriptions [post]
func createSubscription(c *gin.Context) {
	var err error

	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()

	var sub schemas.SubscriptionCreate
	if err = c.ShouldBindJSON(&sub); err != nil {
		return
	}
	newSub, err := alert_biz.NewSubscriptionService().CreateSubscription(&sub)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, ts.IdResponse{Id: newSub})
}

// @Tags Alert
// @Summary Get Subscription
// @Description Get Subscription
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "uuid formatted subscriptionId"
// @Success 200 {object} schemas.Subscription
// @Router /alert/subscriptions/{id} [get]
func getSubscription(c *gin.Context) {
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
	sub, err := alert_biz.NewSubscriptionService().GetById(id)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, sub)
}

// @Tags Alert
// @Summary Update Subscription
// @Description Update Subscription
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "uuid formatted subscriptionId"
// @Param data body schemas.SubscriptionUpdate true "data"
// @Success 200
// @Router /alert/subscriptions/{id} [put]
func updateSubscription(c *gin.Context) {
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
	var sub schemas.SubscriptionUpdate
	if err = c.ShouldBindJSON(&sub); err != nil {
		return
	}
	if err = alert_biz.NewSubscriptionService().UpdateSubscription(id, &sub); err != nil {
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Tags Alert
// @Summary List Subscriptions
// @Description List Subscriptions
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param object query schemas.SubscriptionQuery true "query"
// @Success 200 {object} ts.ListResponse{results=[]schemas.Subscription}
// @Router /alert/subscriptions [get]
func listSubscriptions(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()

	var query schemas.SubscriptionQuery
	if err = c.ShouldBindQuery(&query); err != nil {
		return
	}
	count, subs, err := alert_biz.NewSubscriptionService().ListSubscriptions(&query)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, ts.ListResponse{Total: count, Results: subs})
}

// @Tags Alert
// @Summary Delete Subscription
// @Description Delete Subscription
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "uuid formatted subscriptionId"
// @Success 200
// @Router /alert/subscriptions/{id} [delete]
func deleteSubscription(c *gin.Context) {
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
	if err = alert_biz.NewSubscriptionService().DeleteSubscription(id); err != nil {
		return
	}
	c.JSON(http.StatusOK, nil)
}
