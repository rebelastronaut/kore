import PropTypes from 'prop-types'
import moment from 'moment'
import { List, Avatar, Icon, Typography, Tooltip } from 'antd'
const { Text } = Typography

import ResourceVerificationStatus from '../resources/ResourceVerificationStatus'
import AutoRefreshComponent from '../teams/AutoRefreshComponent'
import { successMessage, errorMessage } from '../../utils/message'
import { getProviderCloudInfo } from '../../utils/cloud'
import IconTooltip from '../utils/IconTooltip'

class Credentials extends AutoRefreshComponent {
  static propTypes = {
    provider: PropTypes.oneOf(['GKE', 'EKS', 'AKS']),
    identifierKey: PropTypes.string.isRequired,
    credentials: PropTypes.object.isRequired,
    allTeams: PropTypes.array.isRequired,
    editCredential: PropTypes.func.isRequired,
    deleteCredential: PropTypes.func.isRequired
  }

  cloudInfo = getProviderCloudInfo(this.props.provider)

  componentDidUpdate(prevProps) {
    const prevStatus = prevProps.credentials.status && prevProps.credentials.status.status
    const newStatus = this.props.credentials.status && this.props.credentials.status.status
    if (prevStatus !== newStatus) {
      this.startRefreshing()
    }
  }

  stableStateReached({ state }) {
    const { credentials, identifierKey } = this.props
    if (state === AutoRefreshComponent.STABLE_STATES.SUCCESS) {
      return successMessage(`${this.cloudInfo.cloud} credentials for ${this.cloudInfo.accountNoun} "${credentials.spec[identifierKey]}" verified successfully`)
    }
    if (state === AutoRefreshComponent.STABLE_STATES.FAILURE) {
      return errorMessage(`${this.cloudInfo.cloud} credentials for ${this.cloudInfo.accountNoun} "${credentials.spec[identifierKey]}" could not be verified`)
    }
  }

  actions = () => {
    const { provider, credentials, editCredential, deleteCredential } = this.props
    return [
      <ResourceVerificationStatus key="verification_status" resourceStatus={credentials.status} />,
      <Text key="delete_creds">
        <Tooltip title="Delete this credential">
          <a id={`${provider.toLowerCase()}creds_del_${credentials.metadata.name}`} onClick={deleteCredential(credentials)}><Icon type="delete" /></a>
        </Tooltip>
      </Text>,
      <Text key="show_creds">
        <Tooltip title="Edit this credential">
          <a id={`${provider.toLowerCase()}creds_edit_${credentials.metadata.name}`} onClick={editCredential(credentials)}><Icon type="edit" /></a>
        </Tooltip>
      </Text>
    ]
  }

  render() {
    const { provider, identifierKey, credentials, allTeams } = this.props
    const created = moment(credentials.metadata.creationTimestamp).fromNow()

    const displayAllocations = () => {
      if (!credentials.allocation) {
        return <Text>No teams <Tooltip title={`This ${this.cloudInfo.accountNoun} is not allocated to any teams, click edit to fix this.`}><Icon type="warning" theme="twoTone" twoToneColor="orange" /></Tooltip> </Text>
      }
      const allocatedTeams = allTeams.filter(team => credentials.allocation.spec.teams.includes(team.metadata.name)).map(team => team.spec.summary)
      return allocatedTeams.length > 0 ? allocatedTeams.join(', ') : 'All teams'
    }

    return (
      <List.Item id={`${provider.toLowerCase()}creds_${credentials.metadata.name}`} key={credentials.metadata.name} actions={this.actions()}>
        <List.Item.Meta
          avatar={<Avatar icon="project" />}
          title={
            <>
              <Text style={{ display: 'inline', marginRight: '15px', fontSize: '16px', fontWeight: '600' }}>{credentials.spec[identifierKey]}</Text>
              <Text style={{ marginRight: '5px' }}>{credentials.allocation ? credentials.allocation.spec.name : null}</Text>
              <IconTooltip text={credentials.allocation ? credentials.allocation.spec.summary : null} />
            </>
          }
          description={<Text>Allocated to: {displayAllocations()}</Text>}

        />
        <Text type='secondary'>Created {created}</Text>
      </List.Item>
    )
  }

}

export default Credentials
