const axios = require('axios')
const Router = require('express').Router

const PATH_BLACKLIST = ['/auth']

function apiProxy(koreApiUrl) {
  return async (req, res) => {
    const method = req.method.toLowerCase()
    const apiUrlPath = req.originalUrl.replace('/apiproxy', '')
    const options = {
      headers: {
        'Authorization': `Bearer ${req.session.passport.user.id_token}`
      }
    }
    try {
      const result = await axios[method](
        `${koreApiUrl}${apiUrlPath}`,
        ['get', 'delete'].includes(method) ? options : req.body,
        ['get', 'delete'].includes(method) ? undefined : options
      )
      return res.json(result.data)
    } catch (err) {
      const status = (err.response && err.response.status) || 500
      if (status === 400 || status === 409) {
        console.log(`Validation error for ${apiUrlPath}`, err.response.data)
        return res.status(status).json(err.response.data).send()
      }
      const message = (err.response && err.response.data && err.response.data.message) || err.message
      console.error(`Error making request to API with path ${apiUrlPath}`, status, message)
      return res.status(status).send({ message })
    }
  }
}

function checkBlacklist(req, res, next) {
  const apiUrlPath = req.originalUrl.replace('/apiproxy', '')
  if (PATH_BLACKLIST.includes(apiUrlPath)) {
    return res.status(404).send()
  }
  next()
}

function initRouter({ ensureAuthenticated, ensureUserCurrent, koreApiUrl }) {
  const router = Router()
  router.use('/apiproxy/*', ensureAuthenticated, checkBlacklist, ensureUserCurrent, apiProxy(koreApiUrl))
  return router
}

module.exports = {
  initRouter
}
