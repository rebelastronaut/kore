import React from 'react'
import PropTypes from 'prop-types'
import Link from 'next/link'
import { Button, Divider, Icon, List, Tooltip, Typography } from 'antd'
const { Paragraph, Title } = Typography

import Breadcrumb from '../../../../lib/components/layout/Breadcrumb'

import ServiceCatalogData from '../../../../lib/prototype/utils/dummy-service-catalog-data'
import ServiceHeader from '../../../../lib/prototype/components/service-catalog/ServiceHeader'

class ListServicePlans extends React.Component {
  static propTypes = {
    service: PropTypes.object.isRequired
  }

  static getInitialProps = async (ctx) => {
    const service = ServiceCatalogData.services.items.find(s => s.metadata.name === ctx.query.name)
    return { service }
  }

  render() {
    const { service } = this.props

    return (
      <>
        <Breadcrumb items={[
          { text: 'Service Catalog', link: '/prototype/service-catalog', href: '/prototype/service-catalog' },
          { text: service.spec.name },
          { text: 'Plans' }
        ]} />

        <ServiceHeader service={service} />
        <Divider />

        <Title level={4} style={{ marginBottom: '20px' }}>Plans</Title>
        <Link href="/prototype/service-catalog/[name]/newplan" as={`/prototype/service-catalog/${service.metadata.name}/newplan`}>
          <Button type="primary" style={{ marginBottom: '20px' }}>+ New plan</Button>
        </Link>
        {!service.spec.plans || service.spec.plans.length === 0 ? <Paragraph>No service plans found</Paragraph> : (
          <List
            dataSource={service.spec.plans}
            renderItem={plan => (
              <List.Item actions={[
                <Link key="edit" href="/prototype/service-catalog/[name]/[plan]" as={`/prototype/service-catalog/${service.metadata.name}/${plan.id}`}>
                  <Tooltip title="Manage this plan">
                    <a><Icon type="edit" /></a>
                  </Tooltip>
                </Link>,
                <Tooltip key="delete" title="Delete this plan">
                  <a><Icon type="delete" /></a>
                </Tooltip>
              ]}>
                <List.Item.Meta
                  title={plan.name}
                  description={plan.description}
                />
              </List.Item>
            )}
          />
        )}
      </>
    )
  }
}

export default ListServicePlans
