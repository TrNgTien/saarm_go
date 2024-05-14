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
			return utilities.R403(c, "Forbidden")
		}

		return next(c)
	}
}

func LandlordPermission(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		role := c.Get("role").(string)

		adminRoles := append(common.FixedAllowedRoles, "landlord")

		if !VerifyByRole(adminRoles, role) {
			return utilities.R403(c, "Forbidden")
		}

		return next(c)
	}
}
