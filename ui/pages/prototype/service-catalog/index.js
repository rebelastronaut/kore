import React from 'react'
import Link from 'next/link'
import { Button, Card, Col, Divider, Form, Icon, Input, Row, Select, Switch, Tag, Tooltip, Typography } from 'antd'
const { Paragraph, Text, Title } = Typography
import { intersection } from 'lodash'

import Breadcrumb from '../../../lib/components/layout/Breadcrumb'
import ServiceCatalogData from '../../../lib/prototype/utils/dummy-service-catalog-data'

class ServiceCatalogIndex extends React.Component {

  state = {
    loading: true,
    serviceCategoryFilter: [],
    serviceNameFilter: ''
  }

  fetchComponentData = async () => {
    const services = await Promise.resolve(ServiceCatalogData.services)
    return { services: services.items }
  }

  componentDidMount() {
    this.fetchComponentData().then((data) => this.setState({ ...data, loading: false }))
  }

  iconType = (category) => ({
    'Application': 'appstore',
    'Security': 'lock',
    'Monitoring': 'monitor'
  }[category])

  renderServiceCard = (service) => {
    return (
      <Col key={service.metadata.name} span={12} xl={8}>
        <Card>
          <Row>
            <Col span={20}>
              <Title level={4}>{service.spec.name}</Title>
            </Col>
            <Col span={4} style={{ textAlign: 'right' }}>
              {service.spec.icon ? <img src={`/static/images/${service.spec.icon}`} height={28} /> : <Icon type={this.iconType(service.spec.category)} style={{ fontSize: '28px' }} />}
            </Col>
          </Row>
          <Paragraph ellipsis={{ rows: 3 }}>{service.spec.description}</Paragraph>
          <Paragraph><a target="_blank" rel="noopener noreferrer" href={service.spec.documentationURL} style={{ textDecoration: 'underline' }}>Documentation</a></Paragraph>
          <Paragraph>
            {service.spec.category.map(c => <Tag key={c}>{c}</Tag>)}
            {service.spec.prereqs ? (
              <>
                <Divider type="vertical" style={{ marginLeft: '2px' }} />
                <Tooltip title={service.spec.prereqs.map(p => p.name).join(', ')}>
                Prerequisites
                </Tooltip>
              </>
            ) : null}
          </Paragraph>
          <Paragraph style={{ marginBottom: 0 }}>
            {service.spec.plans && service.spec.plans.length > 0 ? (
              <Row>
                <Col span={6}>
                  <Text>
                    <Link href="/prototype/service-catalog/[name]" as={`/prototype/service-catalog/${service.metadata.name}`}>
                      <Button>Manage</Button>
                    </Link>
                  </Text>
                </Col>
                <Col span={18} style={{ paddingTop: '6px', textAlign: 'right' }}>
                  <Tag color="green"><Icon type="check-circle" style={{ marginRight: '5px' }} />Configured</Tag>
                  <Tag color="orange" style={{ marginRight: 0 }}><Icon type="up-circle" style={{ marginRight: '5px' }}/>Update available</Tag>
                </Col>
              </Row>
            ) : (
              <Row>
                <Col>
                  <Text>
                    <Link href="/prototype/service-catalog/[name]" as={`/prototype/service-catalog/${service.metadata.name}`}>
                      <Button>Configure</Button>
                    </Link>
                  </Text>
                </Col>
              </Row>
            )}
          </Paragraph>
        </Card>
      </Col>
    )
  }

  filteredServices = () => {
    const { serviceNameFilter, serviceCategoryFilter, showConfiguredOnly } = this.state
    const categoryMatch = (service) => (serviceCategoryFilter.length === 0 || intersection(serviceCategoryFilter, service.spec.category).length)
    const nameMatch = (service) => (serviceNameFilter === '' || service.spec.name.toLowerCase().indexOf(serviceNameFilter.toLowerCase()) !== -1)
    const descriptionMatch = (service) => (serviceNameFilter === '' || service.spec.description.toLowerCase().indexOf(serviceNameFilter.toLowerCase()) !== -1)
    const configuredMatch = (service) => !showConfiguredOnly || (service.spec.plans && service.spec.plans.length > 0)
    return this.state.services.filter((service) => categoryMatch(service) && (nameMatch(service) || descriptionMatch(service)) && configuredMatch(service))
  }

  render() {

    return (
      <>
        <Breadcrumb items={[ { text: 'Service Catalog' } ]} />
        <Title level={2}>Service Catalog</Title>
        <Card style={{ marginBottom: '20px' }}>
          <Form.Item labelAlign="left" labelCol={{ span: 3 }} wrapperCol={{ span: 21 }} label={<Text strong>Name</Text>} style={{ marginBottom: '10px' }}>
            <Input onChange={(e) => this.setState({ serviceNameFilter: e.target.value })} value={this.state.serviceNameFilter} placeholder="Search by name or description"/>
          </Form.Item>
          <Form.Item labelAlign="left" labelCol={{ span: 3 }} wrapperCol={{ span: 21 }} label={<Text strong>Category</Text>} style={{ marginBottom: '10px' }}>
            <Select
              mode="multiple"
              style={{ width: '100%' }}
              placeholder="Please select"
              value={this.state.serviceCategoryFilter}
              onChange={(selected) => this.setState({ serviceCategoryFilter: selected })}
            >
              <Select.Option key="Application">Application</Select.Option>
              <Select.Option key="Monitoring">Monitoring</Select.Option>
              <Select.Option key="Security">Security</Select.Option>
            </Select>
          </Form.Item>
          <Form.Item labelAlign="left" labelCol={{ span: 3 }} wrapperCol={{ span: 21 }} label={<Text strong>Configured only</Text>} style={{ marginBottom: 0 }}>
            <Switch checked={this.state.showConfiguredOnly} onChange={(showConfiguredOnly) => this.setState({ showConfiguredOnly })} />
          </Form.Item>
          <a style={{ display: 'block', marginTop: '10px', marginBottom: '5px', textDecoration: 'underline' }} onClick={() => this.setState({ serviceNameFilter: '', serviceCategoryFilter: [], showConfiguredOnly: false })}>Clear filters</a>
        </Card>
        {this.state.loading ? <Icon type="loading" /> : (
          <>
            <Paragraph style={{ fontWeight: 600, fontSize: '17px', fontStyle: 'italic' }}>1 service configured. 4 services available.</Paragraph>
            <Row gutter={[16, 16]}>
              {this.filteredServices().map(service => this.renderServiceCard(service))}
            </Row>
          </>
        )}
      </>
    )
  }
}

export default ServiceCatalogIndex
