async function retryWithTokenRefresh(err, req, authService, retryWhat) {
  if (err.response && err.response.status === 401 && req.session.localUser && req.session.localUserRefreshToken) {
    console.log('attempting token refresh')
    // attempt a token refresh for local users and retry once if successful
    const token = await authService.localTokenRefresh(req.session.localUserRefreshToken)
    if (token) {
      /* eslint-disable-next-line require-atomic-updates */
      req.session.passport.user.id_token = token
      return new Promise(function(resolve, reject) {
        req.session.save((err) => {
          if (err) {
            reject(err)
          }
          return resolve(retryWhat())
        })
      })
    }
  }
  // Re-throw if we've not refreshed
  throw err
}

module.exports = {
  retryWithTokenRefresh
}