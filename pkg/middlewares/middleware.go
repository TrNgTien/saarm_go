package middlewares

import (
	"saarm/pkg/common"
	"saarm/pkg/utilities"

	"github.com/labstack/echo/v4"
)

func VerifyByRole(allowRoles []string, role string) bool {
	return utilities.ArrayIncludeString(allowRoles, role)
}

func AdminPermission(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		role := c.Get("role").(string)

		adminRoles := common.FixedAllowedRoles

		if !VerifyByRole(adminRoles, role) {
			return utilities.R403(c)
		}

		return next(c)
	}
}

func LandlordPermission(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		role := c.Get("role").(string)

		allowRoles := append(common.FixedAllowedRoles, common.LANDLORD_ROLE)

		if !VerifyByRole(allowRoles, role) {
			return utilities.R403(c)
		}

		return next(c)
	}
}
