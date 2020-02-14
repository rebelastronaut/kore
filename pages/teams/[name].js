import React from 'react'
import PropTypes from 'prop-types'
import axios from 'axios'
import moment from 'moment'
import Link from 'next/link'
import { Typography, Card, List, Tag, Button, Avatar, Icon, Popconfirm, message, Select, Drawer, Badge, Modal } from 'antd'
const { Paragraph, Text } = Typography
const { Option } = Select

import Breadcrumb from '../../lib/components/Breadcrumb'
import NamespaceClaimForm from '../../lib/components/forms/NamespaceClaimForm'
import apiRequest from '../../lib/utils/api-request'
import copy from '../../lib/utils/object-copy'
import asyncForEach from '../../lib/utils/async-foreach'

class TeamDashboard extends React.Component {
  static propTypes = {
    team: PropTypes.object.isRequired,
    members: PropTypes.object.isRequired,
    user: PropTypes.object.isRequired,
    clusters: PropTypes.object.isRequired,
    namespaceClaims: PropTypes.object.isRequired,
    available: PropTypes.object.isRequired
  }

  static staticProps = {
    title: 'Team dashboard'
  }

  constructor(props) {
    super(props)
    this.state = {
      teamName: props.team.metadata.name,
      members: props.members,
      allUsers: [],
      membersToAdd: [],
      clusters: props.clusters,
      createNamespace: false,
      namespaceClaims: props.namespaceClaims
    }
  }

  static async getTeamDetails(req, name) {
    const getTeam = () => apiRequest(req, 'get', `/teams/${name}`)
    const getTeamMembers = () => apiRequest(req, 'get', `/teams/${name}/members`)
    const getTeamClusters = () => apiRequest(req, 'get', `/teams/${name}/clusters`)
    const getNamespaceClaims = () => apiRequest(req, 'get', `/teams/${name}/namespaceclaims`)
    const getAvailable = () => apiRequest(req, 'get', `/teams/${name}/allocations?assigned=true`)

    return axios.all([getTeam(), getTeamMembers(), getTeamClusters(), getNamespaceClaims(), getAvailable()])
      .then(axios.spread(function (team, members, clusters, namespaceClaims, available) {
        return { team, members, clusters, namespaceClaims, available }
      }))
      .catch(err => {
        throw new Error(err.message)
      })
  }

  static getInitialProps = async ctx => {
    const teamDetails = await TeamDashboard.getTeamDetails(ctx.req, ctx.query.name)
    return teamDetails
  }

  getAllUsers = async () => {
    const users = await apiRequest(null, 'get', '/users')
    if (users.items) {
      return users.items.map(user => user.spec.username).filter(user => user !== 'admin')
    }
    return []
  }

  componentDidMount() {
    return this.getAllUsers()
      .then(users => {
        const state = copy(this.state)
        state.allUsers = users
        this.setState(state)
      })
  }

  componentDidUpdate(props, state) {
    if (this.props.team.metadata.name !== state.teamName) {
      const state = copy(this.state)
      state.teamName = this.props.team.metadata.name
      state.members = props.members
      state.clusters = props.clusters
      state.namespaceClaims = props.namespaceClaims
      this.getAllUsers()
        .then(users => {
          state.allUsers = users
          this.setState(state)
        })
    }
  }

  addTeamMembersUpdated = membersToAdd => {
    const state = copy(this.state)
    state.membersToAdd = membersToAdd
    this.setState(state)
  }

  addTeamMembers = async () => {
    const state = copy(this.state)
    const members = state.members

    await asyncForEach(this.state.membersToAdd, async member => {
      await apiRequest(null, 'put', `/teams/${this.props.team.metadata.name}/members/${member}`)
      message.success(`Team member added: ${member}`)
      members.items.push(member)
    })

    state.membersToAdd = []
    this.setState(state)
  }

  deleteTeamMember = member => {
    return async () => {
      const team = this.props.team.metadata.name
      try {
        await apiRequest(null, 'delete', `/teams/${team}/members/${member}`)
        const state = copy(this.state)
        const members = state.members
        members.items = members.items.filter(m => m !== member)
        this.setState(state)
        message.success(`Team member deleted: ${member}`)
      } catch (err) {
        console.error('Error deleting team member', err)
        message.error('Error deleting team member, please try again.')
      }
    }
  }

  deleteCluster = cluster => {
    return async () => {
      const clusterNamespaces = this.state.namespaceClaims.items.filter(nc => nc.spec.cluster.name === cluster.metadata.name)
      if (clusterNamespaces.length > 0) {
        return Modal.warning({
          title: 'Warning: cluster cannot be deleted',
          content: (
            <div>
              <Paragraph strong>The cluster namespaces must be deleted first</Paragraph>
              <List
                size="small"
                dataSource={clusterNamespaces}
                renderItem={ns => <List.Item>{ns.metadata.name}</List.Item>}
              />
            </div>
          ),
          onOk() {}
        })
      }

      const team = this.props.team.metadata.name
      try {
        await apiRequest(null, 'delete', `/teams/${team}/clusters/${cluster.metadata.name}`)
        await apiRequest(null, 'delete', `/teams/${team}/gkes/${cluster.metadata.name}`)
        const state = copy(this.state)
        state.clusters.items = state.clusters.items.filter(c => c.metadata.name !== cluster.metadata.name)
        this.setState(state)
        message.loading(`Cluster deletion requested: ${cluster.metadata.name}`)
      } catch (err) {
        console.error('Error deleting cluster', err)
        message.error('Error deleting cluster, please try again.')
      }
    }
  }

  createNamespace = value => {
    return () => {
      const state = copy(this.state)
      state.createNamespace = value
      this.setState(state)
    }
  }

  handleNamespaceCreated = namespaceClaim => {
    const state = copy(this.state)
    state.createNamespace = false
    state.namespaceClaims.items.push(namespaceClaim)
    this.setState(state)
    message.success(`Namespace "${namespaceClaim.spec.name}" created on cluster "${namespaceClaim.spec.cluster.name}"`)
  }

  deleteNamespace = namespace => {
    return async () => {
      const team = this.props.team.metadata.name
      try {
        await apiRequest(null, 'delete', `/teams/${team}/namespaceclaims/${namespace.metadata.name}`)
        const state = copy(this.state)
        state.namespaceClaims.items = state.namespaceClaims.items.filter(nc => nc.metadata.name !== namespace.metadata.name)
        this.setState(state)
        message.success(`Namespace deleted: ${namespace.spec.name}`)
      } catch (err) {
        console.error('Error deleting namespace', err)
        message.error('Error deleting namespace, please try again.')
      }
    }
  }

  render() {
    const { members, namespaceClaims, allUsers, membersToAdd, createNamespace, clusters } = this.state
    const { team, user, available } = this.props
    const teamMembers = ['ADD_USER', ...members.items]

    const memberActions = member => {
      const deleteAction = (
        <Popconfirm
          key="delete"
          title="Are you sure you want to delete this user?"
          onConfirm={this.deleteTeamMember(member)}
          okText="Yes"
          cancelText="No"
        >
          <a>Remove</a>
        </Popconfirm>
      )
      if (member !== user.username) {
        return [deleteAction]
      }
      return []
    }

    const memberName = member => (
      <Text>{member} {member === user.username ? <Tag>You</Tag>: null}</Text>
    )

    const membersAvailableToAdd = allUsers.filter(user => !members.items.includes(user))

    const clusterActions = cluster => {
      const actions = []
      const status = cluster.status.status || 'Pending'
      if (status === 'Success') {
        const deleteAction = (
          <Popconfirm
            key="delete"
            title="Are you sure you want to delete this cluster?"
            onConfirm={this.deleteCluster(cluster)}
            okText="Yes"
            cancelText="No"
          >
            <a><Icon type="delete" /></a>
          </Popconfirm>
        )
        actions.push(deleteAction)
      }
      actions.push(<Tag color="#5cdbd3">{status}</Tag>)
      return actions
    }

    const namespaceClaimActions = namespaceClaim => {
      const actions = []
      const status = namespaceClaim.status.status || 'Pending'
      if (status === 'Success') {
        const deleteAction = (
          <Popconfirm
            key="delete"
            title="Are you sure you want to delete this namespace?"
            onConfirm={this.deleteNamespace(namespaceClaim)}
            okText="Yes"
            cancelText="No"
          >
            <a><Icon type="delete" /></a>
          </Popconfirm>
        )
        actions.push(deleteAction)
      }
      actions.push(<Tag color="#5cdbd3">{status}</Tag>)
      return actions
    }

    const clusterProviderIconSrcMap = {
      'GKECredentials': '/static/images/GKE.png',
      'EKSCredentials': '/static/images/EKS.png'
    }

    return (
      <div>
        <Breadcrumb items={[{text: team.spec.summary}]} />
        <Paragraph strong>{team.spec.description}</Paragraph>
        <Card
          title={<div><Text style={{ marginRight: '10px' }}>Team members</Text><Badge style={{ backgroundColor: '#1890ff' }} count={members.items.length} /></div>}
          style={{ marginBottom: '16px' }}
          className="team-members"
        >
          <List
            dataSource={teamMembers}
            renderItem={m => {
              if (m === 'ADD_USER') {
                return <List.Item style={{ paddingTop: '0' }} actions={[<Button key="add" type="primary" onClick={this.addTeamMembers}>Add</Button>]}>
                  <List.Item.Meta
                    title={
                      <Select
                        mode="multiple"
                        placeholder="Add existing users to this team"
                        onChange={this.addTeamMembersUpdated}
                        style={{ width: '100%' }}
                        value={membersToAdd}
                      >
                        {membersAvailableToAdd.map((user, idx) => (
                          <Option key={idx} value={user}>{user}</Option>
                        ))}
                      </Select>
                    }
                  />
                </List.Item>
              } else {
                return <List.Item actions={memberActions(m)}>
                  <List.Item.Meta avatar={<Avatar icon="user" />} title={memberName(m)} />
                </List.Item>
              }
            }}
          >
          </List>
        </Card>
        <Card
          title={<div><Text style={{ marginRight: '10px' }}>Clusters</Text><Badge style={{ backgroundColor: '#1890ff' }} count={clusters.items.length} /></div>}
          style={{ marginBottom: '20px' }}
          extra={
            <Button type="primary">
              <Link href="/teams/[name]/clusters/new" as={`/teams/${team.metadata.name}/clusters/new`}>
                <a>+ New</a>
              </Link>
            </Button>
          }
        >
          <List
            dataSource={clusters.items}
            renderItem={cluster => {
              const provider = available.items.find(a => a.metadata.name === cluster.spec.provider.name)
              const created = moment(cluster.metadata.creationTimestamp).fromNow()
              return (
                <List.Item actions={clusterActions(cluster)}>
                  <List.Item.Meta
                    avatar={<img src={clusterProviderIconSrcMap[provider.spec.resource.kind]} height="32px" />}
                    title={<Text>{provider.spec.name} <Text style={{ fontFamily: 'monospace', marginLeft: '15px' }}>{cluster.metadata.name}</Text></Text>}
                    description={<Text type='secondary'>Created {created}</Text>}
                  />
                </List.Item>
              )
            }}
          >
          </List>
        </Card>

        <Card
          title={<div><Text style={{ marginRight: '10px' }}>Namespaces</Text><Badge style={{ backgroundColor: '#1890ff' }} count={namespaceClaims.items.length} /></div>}
          style={{ marginBottom: '20px' }}
          extra={clusters.items.length > 0 ? <Button type="primary" onClick={this.createNamespace(true)}>+ New</Button> : null}
        >

          <List
            dataSource={namespaceClaims.items}
            renderItem={namespaceClaim => {
              const clusterName = namespaceClaim.spec.cluster.name
              const created = moment(namespaceClaim.metadata.creationTimestamp).fromNow()
              return (
                <List.Item actions={namespaceClaimActions(namespaceClaim)}>
                  <List.Item.Meta
                    avatar={<Avatar icon="block" />}
                    title={<Text>{namespaceClaim.metadata.name} <Text style={{ fontFamily: 'monospace', marginLeft: '15px' }}>{clusterName}</Text></Text>}
                    description={<Text type='secondary'>Created {created}</Text>}
                  />
                </List.Item>
              )
            }}
          >
          </List>
        </Card>

        <Drawer
          title="Create namespace"
          placement="right"
          closable={false}
          onClose={this.createNamespace(false)}
          visible={createNamespace}
          width={700}
        >
          <NamespaceClaimForm team={team.metadata.name} clusters={clusters} handleSubmit={this.handleNamespaceCreated} handleCancel={this.createNamespace(false)}/>
        </Drawer>

      </div>
    )
  }
}

export default TeamDashboard
