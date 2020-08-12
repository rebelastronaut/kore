const { retryWithTokenRefresh } = require('../lib/auth-helpers')

const Router = require('express').Router

function getSessionUser(orgService, authService) {
  return async(req, res) => {
    const user = req.session.passport.user
    try {
      await orgService.refreshUser(user)
      return res.json(user)
    } catch (firstErr) {
      let err = firstErr
      try {
        await retryWithTokenRefresh(err, req, authService, async function () { 
          return await orgService.refreshUser(user) 
        })
        return res.json(user)
      } catch (innerErr) {
        err = innerErr
      }
      console.log('Failed to refresh user in /session/user', err)
      return res.status(err.statusCode || 500).send()
    }
  }
}

function initRouter({ ensureAuthenticated, ensureUserCurrent, persistRequestedPath, orgService, authService }) {
  const router = Router()
  router.get('/session/user', ensureAuthenticated, ensureUserCurrent, persistRequestedPath, getSessionUser(orgService, authService))
  return router
}

module.exports = {
  initRouter
}
