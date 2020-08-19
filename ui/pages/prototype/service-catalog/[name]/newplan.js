import React from 'react'
import PropTypes from 'prop-types'
import { Divider } from 'antd'

import Breadcrumb from '../../../../lib/components/layout/Breadcrumb'
import KoreApi from '../../../../lib/kore-api'

import ServiceCatalogData from '../../../../lib/prototype/utils/dummy-service-catalog-data'
import ServiceHeader from '../../../../lib/prototype/components/service-catalog/ServiceHeader'
import ServicePlanConfiguration from '../../../../lib/prototype/components/service-catalog/ServicePlanConfiguration'

class ConfigureServicePlan extends React.Component {
  static propTypes = {
    service: PropTypes.object.isRequired
  }

  constructor(props) {
    super(props)
    this.state = {
      config: {},
      loadingPlans: true,
      clusterPlans: []
    }
  }

  async fetchComponentData() {
    const api = await KoreApi.client()
    const plans = await api.ListPlans()
    return { clusterPlans: plans.items, loadingPlans: false }
  }

  componentDidMount() {
    this.fetchComponentData().then((data) => this.setState({ ...data }))
  }

  static getInitialProps = async (ctx) => {
    const service = ServiceCatalogData.services.items.find(s => s.metadata.name === ctx.query.name)
    return { service }
  }

  render() {
    const service = this.props.service
    const { plan, clusterPlans } = this.state

    return (
      <>
        <Breadcrumb items={[
          { text: 'Service Catalog', link: '/prototype/service-catalog', href: '/prototype/service-catalog' },
          { text: service.spec.name, link: `/prototype/service-catalog/${service.metadata.name}`, href: '/prototype/service-catalog/[name]' },
          { text: 'New plan' }
        ]} />
        <ServiceHeader service={service} />
        <Divider />
        <ServicePlanConfiguration service={service} plan={plan} clusterPlans={clusterPlans} creating={true} />
      </>
    )
  }
}

export default ConfigureServicePlan
