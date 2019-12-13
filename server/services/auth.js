const axios = require('axios')
const AuthProvider = require('../../lib/crd/AuthProvider')

class AuthService {
  constructor(hubApi, baseUrl) {
    this.hubApi = hubApi
    this.baseUrl = baseUrl
    this.requestOptions = {
      headers: {
        'X-Identity': 'admin'
      }
    }
  }

  async getConfiguredAuthProvider() {
    try {
      const result = await axios.get(`${this.hubApi.url}/auth`, this.requestOptions)
      const providers = result.data.items
      return providers.find(p => p.spec.clientID)
    } catch (err) {
      console.error('Error getting auth providers from API', err)
      return Promise.reject(err)
    }
  }

  async setConfiguredAuthProvider(data) {
    try {
      const authProvider = await AuthProvider(data.name, data.config)
      await axios.put(`${this.hubApi.url}/auth/${data.name}`, authProvider, this.requestOptions)
      const config = {
        clientID: data.config.clientID,
        clientSecret: data.config.clientSecret,
        ...this.providerSpecificConfig(data.name)
      }
      require('../lib/passport')(data.name)(config)
    } catch (err) {
      console.error('Error setting configured auth provider from API', err)
      return Promise.reject(err)
    }
  }

  providerSpecificConfig(provider) {
    return {
      'github': {
        callbackUrl: `${this.baseUrl}/login/github/callback`
      }
    }[provider]
  }
}

module.exports = AuthService
