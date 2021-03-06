import PropTypes from 'prop-types'
import moment from 'moment'
import { List, Avatar, Icon, Typography, Tooltip } from 'antd'
const { Text } = Typography

import ResourceVerificationStatus from '../resources/ResourceVerificationStatus'
import AutoRefreshComponent from '../teams/AutoRefreshComponent'
import { successMessage, errorMessage } from '../../utils/message'
import IconTooltip from '../utils/IconTooltip'

class GCPOrganization extends AutoRefreshComponent {
  static propTypes = {
    organization: PropTypes.object.isRequired,
    allTeams: PropTypes.array.isRequired,
    editOrganization: PropTypes.func.isRequired,
    deleteOrganization: PropTypes.func.isRequired
  }

  componentDidUpdate(prevProps) {
    const prevStatus = prevProps.organization.status && prevProps.organization.status.status
    const newStatus = this.props.organization.status && this.props.organization.status.status
    if (prevStatus !== newStatus) {
      this.startRefreshing()
    }
  }

  stableStateReached({ state }) {
    const { organization } = this.props
    if (state === AutoRefreshComponent.STABLE_STATES.SUCCESS) {
      return successMessage(`GCP organization "${organization.allocation.spec.name}" created successfully`)
    }
    if (state === AutoRefreshComponent.STABLE_STATES.FAILURE) {
      return errorMessage(`GCP organization "${organization.allocation.spec.name}" failed to be created`)
    }
  }

  actions = () => {
    const { organization, editOrganization, deleteOrganization } = this.props
    return [
      <ResourceVerificationStatus key="verification_status" resourceStatus={organization.status} />,
      <Text key="delete_org">
        <Tooltip title="Delete this organization">
          <a id={`gcporg_del_${organization.metadata.name}`}  onClick={deleteOrganization(organization)}><Icon type="delete" /></a>
        </Tooltip>
      </Text>,
      <Text key="edit">
        <Tooltip title="Edit this organization">
          <a id={`gcporg_edit_${organization.metadata.name}`} onClick={editOrganization(organization)}><Icon type="edit" /></a>
        </Tooltip>
      </Text>
    ]
  }

  render() {
    const { organization, allTeams } = this.props
    const created = moment(organization.metadata.creationTimestamp).fromNow()

    const displayAllocations = () => {
      if (!organization.allocation) {
        return <Text>No teams <Tooltip title="This organization is not allocated to any teams, click edit to fix this."><Icon type="warning" theme="twoTone" twoToneColor="orange" /></Tooltip> </Text>
      }
      const allocatedTeams = allTeams.filter(team => organization.allocation.spec.teams.includes(team.metadata.name)).map(team => team.spec.summary)
      return allocatedTeams.length > 0 ? allocatedTeams.join(', ') : 'All teams'
    }

    return (
      <List.Item id={`gcporg_${organization.metadata.name}`} key={organization.metadata.name} actions={this.actions()}>
        <List.Item.Meta
          avatar={<Avatar icon="cloud" />}
          title={
            <>
              <Text style={{ display: 'inline', marginRight: '15px', fontSize: '20px', fontWeight: '600' }}>{organization.spec.parentID}</Text>
              <Text style={{ marginRight: '5px' }}>{organization.allocation ? organization.allocation.spec.name : organization.metadata.name}</Text>
              <IconTooltip text={organization.allocation ? organization.allocation.spec.summary : organization.spec.summary} />
            </>
          }
          description={
            <Text>Allocated to: {displayAllocations()}</Text>
          }
        />
        <Text type='secondary'>Created {created}</Text>
      </List.Item>
    )
  }

}

export default GCPOrganization
