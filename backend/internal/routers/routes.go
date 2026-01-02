package routers

import (
	"database/sql"
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux, db *sql.DB) {
	RegisterMemberRoutes(mux, db)
	RegisterStaffRoutes(mux, db)
	RegisterMessageRoutes(mux, db)
	RegisterTalkUserRoutes(mux, db)
	RegisterTalkUserMemberRoutes(mux, db)
	RegisterTemplateRoutes(mux, db)
	RegisterFanletterRoutes(mux, db)
	RegisterBlogRoutes(mux, db)
	RegisterNotificationRoutes(mux, db)
	RegisterOfficialNewsRoutes(mux, db)
	registerMessageReadRoutes(mux, db)
}
