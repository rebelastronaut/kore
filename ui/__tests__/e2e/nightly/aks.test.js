const { TeamCluster } = require('./team-cluster')

describe('AKS end-to-end', () => {
  new TeamCluster({
    provider: 'AKS',
    plan: 'AKS Development Cluster',
    timeouts: {
      // 15 minutes
      create: 15 * 60 * 1000,
      // 20 minutes
      delete: 20 * 60 * 1000,
    }
  }).run()
})
