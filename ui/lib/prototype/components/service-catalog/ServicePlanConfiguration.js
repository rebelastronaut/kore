import * as React from 'react'
import PropTypes from 'prop-types'
import Link from 'next/link'
import { Alert, Button, Collapse, Form, Icon, Input, Modal, Table, Tag, Tooltip, Typography } from 'antd'
const { Paragraph, Text, Title } = Typography
const { Panel } = Collapse

import { clusterProviderIconSrcMap } from '../../../utils/ui-helpers'
import PlanOption from '../../../components/plans/PlanOption'
import IconTooltip from '../../../components/utils/IconTooltip'
import PlanViewer from '../../../components/plans/PlanViewer'

class ServicePlanConfiguration extends React.Component {
  static propTypes = {
    service: PropTypes.object.isRequired,
    plan: PropTypes.object,
    clusterPlans: PropTypes.array.isRequired,
    creating: PropTypes.bool
  }

  constructor(props) {
    super(props)
    this.state = {
      activeKeys: props.creating ? ['name', 'config', 'deployment'] : [],
      editInProgress: false,
      config: props.creating ? {} : props.plan.config,
      appliedPlans: props.creating ? [] : props.plan.appliedPlans
    }
  }

  updateConfigValue = (n, v) => {
    this.setState((state) => ({ config: { ...state.config, [n]: v } }))
  }

  renderServiceConfig = () => {
    const schema = this.props.service.spec.schema
    const config = this.state.config
    const editable = this.props.creating || this.state.editInProgress
    return (
      Object.keys(schema.properties).map((name) => {
        return <PlanOption
          manage={false}
          mode="edit"
          team="kore-admin"
          resourceType="service"
          kind="Service"
          key={name}
          name={name}
          property={schema.properties[name]}
          value={config[name]}
          editable={editable}
          forceShow={true}
          onChange={(n, v) => this.updateConfigValue(n, v)}
          validationErrors={null}
        />
      })
    )
  }

  viewPlan = (plan) => () => {
    Modal.info({
      title: (<><Title level={4}>{plan.spec.description}</Title><Text>{plan.spec.summary}</Text></>),
      content: <PlanViewer
        plan={plan}
        resourceType="cluster"
      />,
      width: 900,
      onOk() {}
    })
  }

  renderPlansTable = () => {
    const { clusterPlans, creating } = this.props
    const { editInProgress, appliedPlans } = this.state
    let filteredPlans = clusterPlans
    if (!creating && !editInProgress) {
      filteredPlans = clusterPlans.filter((p) => appliedPlans.includes(p.metadata.name))
    }
    const rowSelection = (creating || editInProgress) ? {
      getCheckboxProps: (record) => {
        return {
          checked: appliedPlans.includes(record.metadata.name)
        }
      },
      onSelect: (record, selected) => {
        if (selected) {
          this.setState((state) => ({ appliedPlans: [ ...state.appliedPlans, record.metadata.name] }))
        } else {
          this.setState((state) => ({ appliedPlans: state.appliedPlans.filter(p => p !== record.metadata.name) }))
        }
      },
      onSelectAll: (selected) => {
        if (selected) {
          this.setState({ appliedPlans: clusterPlans.map(p => p.metadata.name) })
        } else {
          this.setState({ appliedPlans: [] })
        }
      }
    } : undefined

    const columns = [
      {
        title: 'Cloud provider',
        key: 'cloud',
        dataIndex: 'spec.kind',
        render: (kind) => <img height={clusterProviderIconSrcMap[kind].indexOf('.svg') > 0 ? 25 : 35} src={clusterProviderIconSrcMap[kind]} />
      },
      {
        title: 'Name',
        key: 'name',
        dataIndex: 'spec.description',
        render: (desc, record) => <Text>{desc} <IconTooltip text={record.spec.summary} /></Text>
      },
      {
        title: 'Labels',
        key: 'labels',
        dataIndex: 'spec.labels',
        render: (tags) => {
          if (!tags) {
            return null
          }
          return <span>{Object.keys(tags).map(k => <Tag key={k}>{k.replace('kore.appvia.io/', '')}={tags[k]}</Tag>)}</span>
        }
      },
      {
        title: 'Actions',
        key: 'actions',
        render: (_, record) => (
          <span>
            <Tooltip title="View this plan">
              <a onClick={this.viewPlan(record)}><Icon type="eye" /></a>
            </Tooltip>
          </span>
        ),
      }
    ]

    return <Table rowSelection={rowSelection} columns={columns} dataSource={filteredPlans} />
  }

  render() {
    const { plan, service, creating } = this.props
    const { activeKeys, editInProgress } = this.state

    return (
      <>
        <Collapse style={{ marginTop: '20px' }} activeKey={activeKeys} onChange={(activeKeys) => this.setState({ activeKeys })}>
          <Panel header="Name" key="name">
            <Form.Item labelAlign="left" labelCol={{ span: 3 }} wrapperCol={{ span: 10 }} label="Name">
              <Input value={!creating ? plan.name : undefined} disabled={!creating && !editInProgress} />
            </Form.Item>
            <Form.Item labelAlign="left" labelCol={{ span: 3 }} wrapperCol={{ span: 16 }} label="Description">
              <Input value={!creating ? plan.description : undefined} disabled={!creating && !editInProgress} />
            </Form.Item>
          </Panel>
          <Panel header="Configuration" key="config">
            <Alert
              message="Some configuration may be required for Kore to know how to run the service"
              type="info"
              style={{ marginBottom: '20px' }}
            />
            {service.spec.schema ? this.renderServiceConfig() : <Paragraph style={{ fontStyle: 'italic', marginBottom: 0 }}>None required</Paragraph>}
          </Panel>
          <Panel header="Deployment" key="deployment">
            <Alert
              message="Choose how the service should be deployed"
              type="info"
              style={{ marginBottom: '20px' }}
            />
            {this.renderPlansTable()}
          </Panel>
        </Collapse>
        {creating ? (
          <Form.Item style={{ marginTop: '20px' }}>
            <Link href="/prototype/service-catalog/[name]" as={`/prototype/service-catalog/${service.metadata.name}`}>
              <Button type="primary">Save configuration</Button>
            </Link>
            <Link href="/prototype/service-catalog/[name]" as={`/prototype/service-catalog/${service.metadata.name}`}>
              <Button type="link">Cancel</Button>
            </Link>
          </Form.Item>
        ) : (
          <Form.Item style={{ marginTop: '20px' }}>
            {editInProgress ?
              <Button type="primary" onClick={() => this.setState({ editInProgress: false, activeKeys: [] })}>Save settings</Button> :
              <Button type="primary" onClick={() => this.setState({ editInProgress: true, activeKeys: ['name', 'config', 'deployment'] })}>Edit settings</Button>
            }
          </Form.Item>
        )}
      </>
    )

  }
}

export default ServicePlanConfiguration
