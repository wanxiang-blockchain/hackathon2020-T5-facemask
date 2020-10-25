package middlewares

import (
	"data-manager/web/session"

	"github.com/gin-gonic/gin"
)

func SessionMiddleware(mgrObj session.Mgr) gin.HandlerFunc {

	return func(c *gin.Context) {
		SessionID, err := c.Cookie(session.SessionCookieName)

		var sd session.SessionData
		if err != nil {
			SessionID, sd = mgrObj.CreatSession()

		} else {
			sd, err = mgrObj.GetSessionData(SessionID)

			if err != nil {
				SessionID, sd = mgrObj.CreatSession()

				SessionID = sd.GetID() //这个session id用于回写coookie

			}
		}

		c.Set(session.SessionContextName, sd)

		c.SetCookie(session.SessionCookieName, SessionID, 3600, "/", "", false, false)

		c.Next()
	}
}
