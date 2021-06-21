// SPDX-FileCopyrightText: 2020 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
//
// SPDX-License-Identifier: Apache-2.0

package rest

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/osrgroup/product-model-toolkit/pkg/services/querying"
	"github.com/osrgroup/product-model-toolkit/pkg/services/version"
	"github.com/osrgroup/product-model-toolkit/pkg/server/services"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

// Handler handle all request for the given route group.
func Handler(g *echo.Group, srv services.Service) {
	// default routes
	g.GET("/", handleEntryPoint)
	g.GET("/health", handleHealth)

	// crawler routes
	g.GET("/license", findAllLicenses(srv))
	g.GET("licenses/:id", findLicenseById(srv))
	g.POST("licenses/add", addLicense(srv))

	// diff routes
	g.POST("/diff/id", diffById(srv))
	g.POST("/diff/path", diffByPath(srv))
	// export routes
	g.POST("/export", exportFile(srv))

	// import routes
	g.POST("/import", importFile(srv))

	// merge routes
	g.POST("/merge", mergeFiles(srv))

	// product routes
	g.GET("/products", findAllProducts(srv))
	g.GET("/products/:id", findProductByID(srv))
	g.POST("/products/import/:scanner", importFromScanner(srv))

	// search routes
	g.GET("/search", searchData(srv))

	// version routes
	g.GET("/version", handleVersion)
}