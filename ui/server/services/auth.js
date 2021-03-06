const config = require('../../config')

class AuthService {
  constructor(KoreApi) {
    this.KoreApi = KoreApi
  }

  async getApiClient(basicAuth = null) {
    const token = basicAuth === null ? config.api.token : null
    const api = await this.KoreApi.client({ id_token: token, basicAuth })
    return api
  }

  async getDefaultConfiguredIdp() {
    try {
      const result = await (await this.getApiClient()).GetDefaultIDP()
      return result
    } catch (err) {
      if (err.response && err.response.status === 404) {
        return null
      }
      console.error('Error getting auth providers from API', err)
      return Promise.reject(err)
    }
  }

  generateIDPClientResource() {
    // it's not currently possible to use models from lib/kore-api/model
    // this is due to it using the import syntax which cannot be used on the server-side
    return {
      apiVersion: 'core.kore.appvia.io/v1',
      kind: 'IDPClient',
      metadata: {
        name: 'default'
      },
      spec: {
        displayName: 'Kore UI',
        secret: config.auth.openid.clientSecret,
        id: config.auth.openid.clientID,
        redirectURIs: [`${config.kore.baseUrl}/auth/callback`]
      }
    }

  }

  generateIDPResource(spec) {
    // it's not currently possible to use models from lib/kore-api/model
    // this is due to it using the import syntax which cannot be used on the server-side
    return {
      apiVersion: 'core.kore.appvia.io/v1',
      kind: 'IDP',
      metadata: {
        name: 'default'
      },
      spec
    }
  }

  async setAuthClient() {
    try {
      await (await this.getApiClient()).UpdateIDPClient('kore-ui', this.generateIDPClientResource())
      console.log('Auth client created successfully')
    } catch (err) {
      console.error('Error setting auth client from API', err)
      return Promise.reject(err)
    }
  }

  async setDefaultConfiguredIdp(name, displayName, config) {
    try {
      const spec = {
        displayName,
        config: { [name]: config }
      }
      await (await this.getApiClient()).UpdateIDP('default', this.generateIDPResource(spec))
    } catch (err) {
      console.error('Error setting configured auth provider from API', err)
      return Promise.reject(err)
    }
  }

  async authenticateLocalUser({ username, password }) {
    try {
      const tokens = await (await this.KoreApi.client({ id_token: null })).Login(username, password)
      const user = await (await this.KoreApi.client({ id_token: tokens.token })).WhoAmI()
      console.log('Local user successfully logged in: ', username)
      return { ...user, ...tokens }
    } catch (err) {
      if (err.response && err.response.status === 401) {
        return Promise.reject({ status: 401 })
      }
      console.log('Error authenticating local user', err)
      return Promise.reject({ status: 500 })
    }
  }

  async localTokenRefresh(refreshToken) {
    try {
      const tokens = await (await this.KoreApi.client({ id_token: null })).RefreshToken(refreshToken)
      console.log('Local user token refreshed successfully')
      return tokens.token
    } catch (err) {
      console.log('Local user token refresh failed', err)
      return null
    }
  }
}

module.exports = AuthService
