import React from 'react'
import PropTypes from 'prop-types'
import { Divider, Icon, Table, Tag, Typography } from 'antd'
const { Title } = Typography

import Breadcrumb from '../../../../lib/components/layout/Breadcrumb'
import ResourceStatusTag from '../../../../lib/components/resources/ResourceStatusTag'
import { clusterProviderIconSrcMap } from '../../../../lib/utils/ui-helpers'
import KoreApi from '../../../../lib/kore-api'

import ServiceCatalogData from '../../../../lib/prototype/utils/dummy-service-catalog-data'
import ServiceHeader from '../../../../lib/prototype/components/service-catalog/ServiceHeader'
import ServicePlanConfiguration from '../../../../lib/prototype/components/service-catalog/ServicePlanConfiguration'

class ConfigureServicePlan extends React.Component {
  static propTypes = {
    plan: PropTypes.object.isRequired,
    service: PropTypes.object.isRequired
  }

  constructor(props) {
    super(props)
    this.state = {
      plan: props.plan,
      config: props.plan.config || {},
      loadingPlans: true,
      clusterPlans: [],
      editInProgress: false
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
    const plan = service.spec.plans.find(p => p.id === ctx.query.plan)
    return { service, plan }
  }

  renderDeploymentsTable = () => {
    const { plan } = this.props

    const columns = [
      {
        title: 'Team ID',
        key: 'team',
        dataIndex: 'metadata.namespace'
      },
      {
        title: 'Cluster',
        key: 'cluster',
        dataIndex: 'metadata.name',
        render: (name, record) => <><img height={clusterProviderIconSrcMap[record.spec.kind].indexOf('.svg') > 0 ? 25 : 35} src={clusterProviderIconSrcMap[record.spec.kind]} /> {name}</>
      },
      {
        title: 'Namespace',
        key: 'namespace ',
        dataIndex: 'spec.serviceStatus.namespace'
      },
      {
        title: 'Status',
        key: 'status',
        render: (_, record) => (
          <span>
            <ResourceStatusTag resourceStatus={record.spec.serviceStatus} />
            {record.spec.serviceStatus.updateable ? (
              <Tag color="orange"><Icon type="up-circle" style={{ marginRight: '5px' }} />Update available</Tag>
            ) : null}
          </span>
        ),
      }
    ]

    return <Table columns={columns} dataSource={plan.clusters} />
  }

  render() {
    const service = this.props.service
    const { plan, clusterPlans } = this.state

    return (
      <>
        <Breadcrumb items={[
          { text: 'Service Catalog', link: '/prototype/service-catalog', href: '/prototype/service-catalog' },
          { text: service.spec.name, link: `/prototype/service-catalog/${service.metadata.name}`, href: '/prototype/service-catalog/[name]' },
          { text: plan.name }
        ]} />
        <ServiceHeader service={service} />
        <Divider />
        <ServicePlanConfiguration service={service} plan={plan} clusterPlans={clusterPlans}/>
        <Divider />
        <Title level={4} style={{ marginBottom: '20px' }}>Deployments</Title>
        {this.renderDeploymentsTable()}
      </>
    )
  }
}

export default ConfigureServicePlan
