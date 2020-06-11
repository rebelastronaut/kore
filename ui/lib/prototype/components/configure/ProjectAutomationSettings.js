import React from 'react'
import { Alert, Card, Typography, Tooltip, Icon, Row, Col, Button, Form, Input, Radio, List, Drawer, Popover, Modal } from 'antd'
const { Title, Text, Paragraph } = Typography

import KoreApi from '../../../kore-api'
import copy from '../../../utils/object-copy'
import canonical from '../../../utils/canonical'
import PlanViewer from '../../../components/plans/PlanViewer'

// prototype components
import AutomatedProjectForm from './AutomatedProjectForm'
import { successMessage } from '../../../utils/message'

class ProjectAutomationSettings extends React.Component {

  static initialProjectList = [{
    code: 'not-production',
    title: 'Not production',
    description: 'To be used for all environments except production',
    prefix: 'kore',
    suffix: 'not-prod',
    plans: ['gke-development']
  }, {
    code: 'production',
    title: 'Production',
    description: 'Project just for the production environment',
    prefix: 'kore',
    suffix: 'prod',
    plans: ['gke-production']
  }]

  state = {
    gcpManagementType: 'KORE',
    gcpProjectAutomationType: 'CUSTOM',
    gcpProjectList: ProjectAutomationSettings.initialProjectList,
    addProject: false,
    associatePlanVisible: false
  }

  async fetchComponentData() {
    const api = await KoreApi.client()
    const planList = await api.ListPlans()
    return planList.items
  }

  componentDidMount() {
    this.fetchComponentData().then(plans => {
      this.setState({ plans })
    })
  }

  selectGcpManagementType = e => this.setState({ gcpManagementType: e.target.value })
  selectGcpProjectAutomationType = e => this.setState({ gcpProjectAutomationType: e.target.value })

  deleteGcpProject = (code) => {
    return () => {
      this.setState({
        gcpProjectList: this.state.gcpProjectList.filter(p => p.code !== code)
      })
      successMessage('GCP automated project removed')
    }
  }

  onChange = (code, property) => {
    return (value) => {
      const gcpProjectList = copy(this.state.gcpProjectList)
      gcpProjectList.find(p => p.code === code)[property] = value
      this.setState({
        gcpProjectList
      })
    }
  }

  showPlanDetails = (plan) => {
    return () => {
      Modal.info({
        title: (<><Title level={4}>{plan.spec.description}</Title><Text>{plan.spec.summary}</Text></>),
        content: <PlanViewer plan={plan} />,
        width: 700,
        onOk() {}
      })
    }
  }

  handleAssociatePlanVisibleChange = (projectCode) => () => this.setState({ associatePlanVisible: projectCode })

  associatePlan = (projectCode, plan) => {
    return () => {
      const gcpProjectList = copy(this.state.gcpProjectList)
      gcpProjectList.find(p => p.code === projectCode).plans.push(plan)
      this.setState({ gcpProjectList })
      successMessage('Plan associated')
    }
  }

  closeAssociatePlans = () => this.setState({ associatePlanVisible: false })

  associatePlanContent = (projectCode, selectedCloud) => {
    const project = this.state.gcpProjectList.find(p => p.code === projectCode)
    const cloudPlans = (this.state.plans || []).filter(p => p.spec.kind === selectedCloud)
    const unassociatedPlans = cloudPlans.filter(p => !project.plans.includes(p.metadata.name))
    if (unassociatedPlans.length === 0) {
      return (
        <>
          <Alert style={{ marginBottom: '20px' }} message="All existing plans are already associated with this project type." />
          <Button type="primary" onClick={this.closeAssociatePlans}>Close</Button>
        </>
      )
    }
    return (
      <>
        <Alert style={{ marginBottom: '20px' }} message="Plans available to be associated with this project type." />
        <List
          dataSource={unassociatedPlans}
          renderItem={plan => <Paragraph>{plan.spec.description} <a style={{ textDecoration: 'underline' }} onClick={this.associatePlan(project.code, plan.metadata.name)}>Associate</a></Paragraph>}
        />
        <Button style={{ marginTop: '10px' }} type="primary" onClick={this.closeAssociatePlans}>Close</Button>
      </>
    )
  }

  unassociatePlan = (projectCode, plan) => {
    return () => {
      const gcpProjectList = copy(this.state.gcpProjectList)
      const project = gcpProjectList.find(p => p.code === projectCode)
      project.plans = project.plans.filter(p => p !== plan)
      this.setState({ gcpProjectList })
      successMessage('Plan unassociated')
    }
  }

  setGcpProjectsToDefault = () => this.setState({ gcpProjectList: ProjectAutomationSettings.initialProjectList })

  addProject = (enabled) => () => this.setState({ addProject: enabled })

  handleProjectAdded = (project) => {
    const code = canonical(project.title)
    this.setState({
      gcpProjectList: this.state.gcpProjectList.concat([{ code, plans: [], ...project }]),
      addProject: false
    })
    successMessage('GCP automated project added')
  }

  infoModal = ({ title, content }) => {
    Modal.info({
      title,
      content,
      onOk() {},
      width: 500
    })
  }

  projectAutomationHelp = () => {
    this.infoModal({
      title: 'This defines how Kore will automate GCP projects for teams',
      content: (
        <div>
          <p>When a team is created in Kore and a cluster is requested, Kore will ensure the GCP project is also created and the cluster placed inside it.</p>
          <p>You can specify how the GCP projects should be created for Kore teams.</p>
        </div>
      )
    })
  }

  projectAutomationOnePerClusterHelp = () => {
    this.infoModal({
      title: 'Project automation: One per cluster',
      content: 'For every cluster a team creates Kore will also create a GCP project and provision the cluster inside it. The GCP project will share the name given to the cluster.'
    })
  }

  projectAutomationCustomHelp = () => {
    this.infoModal({
      title: 'Project automation: Custom',
      content: (
        <div>
          <p>When a team is created in Kore and a cluster is requested, Kore will ensure the associated GCP project is also created and the cluster placed inside it.</p>
          <p>You must also specify the plans available for each type of project, this is to ensure the correct cluster specification is being used.</p>
        </div>
      )
    })
  }

  projectCredentialAccessHelp = () => {
    this.infoModal({
      title: 'Team access',
      content: (
        <div>
          <p>When using Kore with existing GCP projects, you must allocate the project credentials to teams in order for them to provision clusters within those projects.</p>
          <p>When a new team is created they may not have access to any project credentials, here you can provide an email address which will be displayed to a team in this situation, in order to request access to a GCP project through Kore.</p>
        </div>
      )
    })
  }

  IconTooltip = ({ icon, text }) => (
    <Tooltip title={text}>
      <Icon type={icon} theme="twoTone" />
    </Tooltip>
  )

  IconTooltipButton = ({ icon, text, onClick }) => (
    <Tooltip title={text}>
      <a style={{ marginLeft: '5px' }} onClick={onClick}><Icon type={icon} theme="twoTone" /></a>
    </Tooltip>
  )

  render() {
    const { gcpManagementType, gcpProjectAutomationType, gcpProjectList, addProject, plans, associatePlanVisible } = this.state

    return (
      <>
        <div style={{ marginBottom: '15px' }}>
          <Paragraph style={{ fontSize: '16px', fontWeight: '600' }}>How do you want Kore teams to integrate with GCP projects?</Paragraph>
          <Radio.Group onChange={this.selectGcpManagementType} value={gcpManagementType}>
            <Radio value={'KORE'} style={{ marginRight: '20px' }}>
              <Text style={{ fontSize: '16px', fontWeight: '600' }}>Kore managed projects <Text type="secondary">(recommended)</Text></Text>
              <Paragraph style={{ marginLeft: '24px', marginBottom: '0' }}>Kore will manage the GCP projects required for teams</Paragraph>
            </Radio>
            <Radio value={'EXTERNAL'}>
              <Text style={{ fontSize: '16px', fontWeight: '600' }}>Use existing projects</Text>
              <Paragraph style={{ marginLeft: '24px', marginBottom: '0' }}>Kore teams will use existing GCP projects that it&apos;s given access to</Paragraph>
            </Radio>
          </Radio.Group>
        </div>

        {gcpManagementType === 'EXTERNAL' ? (
          <>
            <Paragraph style={{ fontSize: '16px', fontWeight: '600' }}>Project credential access for teams <Icon style={{ marginLeft: '5px' }} type="info-circle" theme="twoTone" onClick={this.projectCredentialAccessHelp}/></Paragraph>
            <Form.Item labelAlign="left" labelCol={{ span: 2 }} wrapperCol={{ span: 6 }} label="Email" help="Email for teams who need access to project credentails">
              <Input placeholder="Title" />
            </Form.Item>
          </>
        ) : null}

        {gcpManagementType === 'KORE' ? (
          <>
            <div style={{ marginBottom: '15px' }}>
              <Paragraph style={{ fontSize: '16px', fontWeight: '600' }}>How do you want Kore to automate GCP projects for teams? <Icon style={{ marginLeft: '5px' }} type="info-circle" theme="twoTone" onClick={this.projectAutomationHelp}/></Paragraph>
              <Radio.Group onChange={this.selectGcpProjectAutomationType} value={gcpProjectAutomationType}>
                <Radio value={'CLUSTER'} style={{ marginRight: '20px' }}>
                  <Text style={{ fontSize: '16px', fontWeight: '600' }}>One per cluster <Icon style={{ marginLeft: '5px' }} type="info-circle" theme="twoTone" onClick={this.projectAutomationOnePerClusterHelp}/></Text>
                  <Paragraph style={{ marginLeft: '24px', marginBottom: '0' }}>Kore will create a GCP project for each cluster a team provisions</Paragraph>
                </Radio>
                <Radio value={'CUSTOM'}>
                  <Text style={{ fontSize: '16px', fontWeight: '600' }}>Custom <Icon style={{ marginLeft: '5px' }} type="info-circle" theme="twoTone" onClick={this.projectAutomationCustomHelp}/></Text>
                  <Paragraph style={{ marginLeft: '24px', marginBottom: '0' }}>Configure how Kore will create GCP projects for teams</Paragraph>
                </Radio>
              </Radio.Group>
            </div>

            {gcpProjectAutomationType === 'CUSTOM' ? (
              <>
                <div style={{ display: 'block', marginBottom: '20px' }}>
                  <Button type="primary" onClick={this.addProject(true)}>+ New</Button>
                  <Button style={{ marginLeft: '10px' }} onClick={this.setGcpProjectsToDefault}>Set to Kore defaults</Button>
                </div>
                {gcpProjectList.length === 0 ? (
                  <Paragraph>No automated projects configured, you can &apos;Set to Kore defaults&apos; and/or add new ones. </Paragraph>
                ) : (
                  <List
                    itemLayout="vertical"
                    bordered={true}
                    dataSource={gcpProjectList}
                    renderItem={project => (
                      <List.Item actions={[<a key="delete" onClick={this.deleteGcpProject(project.code)}><Icon type="delete" /> Remove</a>]}>
                        <List.Item.Meta
                          title={<Text editable={{ onChange: this.onChange(project.code, 'title') }} style={{ fontSize: '16px' }}>{project.title}</Text>}
                          description={<Text editable={{ onChange: this.onChange(project.code, 'description') }}>{project.description}</Text>}
                        />

                        <Row gutter={16}>
                          <Col span={8}>
                            <Card
                              title="Naming"
                              size="small"
                              bordered={false}
                            >
                              <Paragraph>The project will be named using the team name, with the prefix and suffix below</Paragraph>
                              <Row style={{ padding: '5px 0' }}>
                                <Col span={8}>
                            Prefix
                                </Col>
                                <Col span={16}>
                                  <Text editable={{ onChange: this.onChange(project.code, 'prefix') }}>{project.prefix}</Text>
                                </Col>
                              </Row>
                              <Row style={{ padding: '5px 0' }}>
                                <Col span={8}>
                            Suffix
                                </Col>
                                <Col span={16}>
                                  <Text editable={{ onChange: this.onChange(project.code, 'suffix') }}>{project.suffix}</Text>
                                </Col>
                              </Row>
                              <Row style={{ paddingTop: '15px' }}>
                                <Col span={8}>
                            Example
                                </Col>
                                <Col span={16}>
                                  <Text>{project.prefix}-<span style={{ fontStyle: 'italic' }}>team-name</span>-{project.suffix}</Text>
                                </Col>
                              </Row>
                            </Card>
                          </Col>
                          <Col span={8}>
                            <Card
                              title="Cluster plans"
                              size="small"
                              bordered={false}
                            >
                              <Paragraph>The cluster plans associated with this project.</Paragraph>
                              {project.plans.length === 0 ? <div style={{ padding: '5px 0' }}>No plans</div> : null}
                              {(plans || []).filter(p => project.plans.includes(p.metadata.name)).map((plan, i) => (
                                <div key={i} style={{ padding: '5px 0' }}>
                                  <Text style={{ marginRight: '10px' }}>{plan.spec.description}</Text>
                                  {this.IconTooltip({ icon: 'info-circle', text: plan.spec.summary })}
                                  {this.IconTooltipButton({ icon: 'eye', text: 'View plan', onClick: this.showPlanDetails(plan) })}
                                  {this.IconTooltipButton({ icon: 'delete', text: 'Unassociate plan', onClick: this.unassociatePlan(project.code, plan.metadata.name) })}
                                </div>
                              ))}
                              <div style={{ padding: '5px 0' }}>

                                <Popover
                                  content={this.associatePlanContent(project.code, 'GCP')}
                                  title={`${project.title}: Associate plans`}
                                  trigger="click"
                                  visible={associatePlanVisible === project.code}
                                  onVisibleChange={this.handleAssociatePlanVisibleChange(project.code)}
                                >
                                  <a>+ Associate plan</a>
                                </Popover>
                              </div>
                            </Card>
                          </Col>
                        </Row>

                      </List.Item>
                    )}
                  />
                )}
                {addProject ? (
                  <Drawer
                    title={<Title level={4}>New project</Title>}
                    visible={addProject}
                    onClose={this.addProject(false)}
                    width={700}
                  >
                    <AutomatedProjectForm handleSubmit={this.handleProjectAdded} handleCancel={this.addProject(false)} />
                  </Drawer>
                ) : null}
              </>

            ) : null}
          </>
        ) : null}
      </>
    )
  }

}

export default ProjectAutomationSettings
