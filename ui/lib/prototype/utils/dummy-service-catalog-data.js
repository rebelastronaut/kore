const services = {
  items: [
    {
      apiVersion: 'services.catalog.kore.appvia.io/v1',
      kind: 'Service',
      metadata: {
        name: 'cert-manager-service',
        namespace: 'kore-admin',
      },
      spec: {
        name: 'Cert manager',
        description: 'cert-manager is a Kubernetes add-on to automate the management and issuance of TLS certificates from various issuing sources. It will ensure certificates are valid and up to date periodically, and attempt to renew certificates at an appropriate time before expiry.',
        documentationURL: 'https://github.com/helm/charts/tree/master/stable/cert-manager',
        category: ['Security'],
        icon: 'cert-manager.svg',
        configured: false,
        schema: {
          'definitions': {},
          '$schema': 'http://json-schema.org/draft-07/schema#',
          '$id': 'https://example.com/object1597325830.json',
          'title': 'Root',
          'type': 'object',
          'required': [
            'name'
          ],
          'properties': {
            'name': {
              '$id': '#root/name',
              'title': 'Name',
              'type': 'string',
              'default': '',
              'pattern': '^.*$'
            }
          }
        },
        prereqs: [{ name: 'Service A' }, { name: 'Service B' }],
        plans: [{
          id: 'dev',
          name: 'Development',
          description: 'This is the development plan for this service',
          config: {
            name: 'some value'
          },
          appliedPlans: ['gke-development', 'eks-development', 'aks-development'],
          clusters: [
            {
              'kind': 'Cluster',
              'apiVersion': 'clusters.compute.kore.appvia.io/v1',
              'metadata': {
                'name': 'project-a-notprod',
                'namespace': 'project-a',
                'selfLink': '/apis/clusters.compute.kore.appvia.io/v1/namespaces/devx/clusters/devx-gke-development',
                'uid': 'cd043e08-c516-432d-997a-43e1403c5ade',
                'resourceVersion': '193740917',
                'generation': 2,
                'creationTimestamp': '2020-06-01T20:18:01Z',
                'labels': {
                  'kore.appvia.io/clusterid': 'bsfhl6phnemjqvqdjcvg',
                  'kore.appvia.io/koreid': 'bsgsh9k6bqqlh5r4nr3g',
                  'kore.appvia.io/teamid': 'bsgsh9s6bqqlh5r4nr50'
                },
                'finalizers': [
                  'cluster.clusters.kore.appvia.io'
                ]
              },
              'spec': {
                'kind': 'GKE',
                'plan': 'gke-development',
                'serviceStatus': {
                  'namespace': 'redis',
                  'status': 'Success',
                  'updateable': true
                }
              }
            },
            {
              'kind': 'Cluster',
              'apiVersion': 'clusters.compute.kore.appvia.io/v1',
              'metadata': {
                'name': 'project-b-notprod',
                'namespace': 'project-b',
                'selfLink': '/apis/clusters.compute.kore.appvia.io/v1/namespaces/devx/clusters/devx-gke-development',
                'uid': 'cd043e08-c516-432d-997a-43e1403c5ade',
                'resourceVersion': '193740917',
                'generation': 2,
                'creationTimestamp': '2020-06-01T20:18:01Z',
                'labels': {
                  'kore.appvia.io/clusterid': 'bsfhl6phnemjqvqdjcvg',
                  'kore.appvia.io/koreid': 'bsgsh9k6bqqlh5r4nr3g',
                  'kore.appvia.io/teamid': 'bsgsh9s6bqqlh5r4nr50'
                },
                'finalizers': [
                  'cluster.clusters.kore.appvia.io'
                ]
              },
              'spec': {
                'kind': 'GKE',
                'plan': 'gke-development',
                'serviceStatus': {
                  'namespace': 'redis',
                  'status': 'Pending',
                  'updateable': false
                }
              }
            }
          ]
        }, {
          id: 'prod',
          name: 'Production',
          description: 'This is the production plan for this service',
          config: {
            name: 'some other value'
          },
          appliedPlans: ['gke-production', 'eks-production', 'aks-production'],
        }]
      }
    }, {
      apiVersion: 'services.catalog.kore.appvia.io/v1',
      kind: 'Service',
      metadata: {
        name: 'nginx-ingress-service',
        namespace: 'kore-admin',
      },
      spec: {
        name: 'nginx-ingress',
        description: 'nginx-ingress is an Ingress controller that uses ConfigMap to store the nginx configuration. To use, add the kubernetes.io/ingress.class: nginx annotation to your Ingress resources.',
        documentationURL: 'https://github.com/helm/charts/tree/master/stable/nginx-ingress',
        category: ['Application'],
        configured: false,
        schema: null
      }
    }, {
      apiVersion: 'services.catalog.kore.appvia.io/v1',
      kind: 'Service',
      metadata: {
        name: 'prometheus-service',
        namespace: 'kore-admin',
      },
      spec: {
        name: 'Prometheus',
        description: 'Prometheus, a Cloud Native Computing Foundation project, is a systems and service monitoring system. It collects metrics from configured targets at given intervals, evaluates rule expressions, displays the results, and can trigger alerts if some condition is observed to be true.',
        documentationURL: 'https://github.com/helm/charts/tree/master/stable/prometheus',
        icon: 'prometheus.svg',
        category: ['Monitoring', 'Application'],
        configured: false,
        schema: null
      }
    },
    {
      apiVersion: 'services.catalog.kore.appvia.io/v1',
      kind: 'Service',
      metadata: {
        name: 'redis-service',
        namespace: 'kore-admin',
      },
      spec: {
        name: 'Redis',
        description: 'Redis is an advanced key-value cache and store. It is often referred to as a data structure server since keys can contain strings, hashes, lists, sets, sorted sets, bitmaps and hyperloglogs.',
        documentationURL: 'https://github.com/helm/charts/tree/master/stable/redis',
        category: ['Application'],
        icon: 'redis-small.png',
        configured: true,
        schema: {
          'definitions': {},
          '$schema': 'http://json-schema.org/draft-07/schema#',
          '$id': 'https://example.com/object1597325830.json',
          'title': 'Root',
          'type': 'object',
          'required': [
            'name'
          ],
          'properties': {
            'name': {
              '$id': '#root/name',
              'title': 'Name',
              'type': 'string',
              'default': '',
              'pattern': '^.*$'
            }
          }
        }
      }
    }
  ]
}

class ServiceCatalogData {
  static services = services
}

export default ServiceCatalogData