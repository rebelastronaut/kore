import React from 'react'
import PropTypes from 'prop-types'
import moment from 'moment'
import { List, Avatar, Icon, Typography, Tooltip } from 'antd'
const { Text } = Typography
import { pluralize, titleize } from 'inflect'

import IconTooltip from '../utils/IconTooltip'
import { isReadOnlyCRD } from '../../utils/crd-helpers'
import { warningMessage } from '../../utils/message'
import { getProviderCloudInfo } from '../../utils/cloud'

class PlanItem extends React.Component {
  static propTypes = {
    plan: PropTypes.object.isRequired,
    viewPlan: PropTypes.func.isRequired,
    editPlan: PropTypes.func.isRequired,
    deletePlan: PropTypes.func.isRequired,
    copyPlan: PropTypes.func.isRequired,
    displayUnassociatedPlanWarning: PropTypes.bool.isRequired
  }

  cloudInfo = getProviderCloudInfo(this.props.plan.spec.kind)

  actions = () => {
    const readonly = isReadOnlyCRD(this.props.plan)
    const actions = []
    if (this.props.displayUnassociatedPlanWarning) {
      actions.push(<IconTooltip key="warning" icon="warning" color="orange" text={`This plan not associated with any ${this.cloudInfo.cloud} automated ${pluralize(this.cloudInfo.accountNoun)} and will not be available for teams to use. Edit this plan or go to ${titleize(this.cloudInfo.accountNoun)} automation settings to review this.`}/>)
    }

    return [
      ...actions,
      <Text key="view_plan">
        <Tooltip title="View this plan">
          <a id={`plans_view_${this.props.plan.metadata.name}`} onClick={this.props.viewPlan(this.props.plan)}><Icon type="eye" /></a>
        </Tooltip>
      </Text>,
      <Text key="edit_plan">
        <Tooltip title="Edit this plan">
          <a id={`plans_edit_${this.props.plan.metadata.name}`} onClick={readonly ? () => warningMessage('Read Only', { description: 'This plan is read-only. Create a new plan if this built-in plan does not meet your needs.' }) : this.props.editPlan(this.props.plan)} style={{ color: readonly ? 'lightgray' : null }}><Icon type="edit" /></a>
        </Tooltip>
      </Text>,
      <Text key="delete_plan">
        <Tooltip title="Delete this plan">
          <a id={`plans_delete_${this.props.plan.metadata.name}`} onClick={readonly ? () => warningMessage('Read Only', { description: 'This plan is read-only and cannot be deleted. To prevent teams using this plan, remove the allocation.' }) : this.props.deletePlan(this.props.plan)} style={{ color: readonly ? 'lightgray' : null }}><Icon type="delete" /></a>
        </Tooltip>
      </Text>,
      <Text key="copy_plan">
        <Tooltip title="Copy this plan">
          <a id={`plans_copy_${this.props.plan.metadata.name}`} onClick={this.props.copyPlan(this.props.plan)}><Icon type="copy" /></a>
        </Tooltip>
      </Text>
    ]
  }

  render() {
    const plan = this.props.plan
    const created = moment(plan.metadata.creationTimestamp).fromNow()

    return (
      <List.Item key={plan.metadata.name} actions={this.actions()}>
        <List.Item.Meta
          avatar={<Avatar icon="build" />}
          title={plan.spec.description}
          description={plan.spec.summary}
        />
        <Text type='secondary'>Created {created}</Text>
      </List.Item>
    )
  }

}

export default PlanItem
