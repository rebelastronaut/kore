import * as React from 'react'
import PropTypes from 'prop-types'
import { Switch, Typography } from 'antd'
const { Paragraph, Text } = Typography

import PlanOption from './PlanOption'
import IconTooltip from '../utils/IconTooltip'

/**
 * PlanViewEdit is the underlying component which handles all plan viewing and editing. Most likely, you
 * want to use UsePlanForm or ManagePlanForm instead of using this directly.
 */
export default class PlanViewEdit extends React.Component {
  static propTypes = {
    resourceType: PropTypes.oneOf(['cluster', 'service', 'servicecredential', 'monitoring']).isRequired,
    mode: PropTypes.oneOf(['create', 'edit', 'view']).isRequired,
    manage: PropTypes.bool,
    team: PropTypes.object,
    kind: PropTypes.string.isRequired,
    plan: PropTypes.object.isRequired,
    originalPlan: PropTypes.object,
    schema: PropTypes.object.isRequired,
    editableParams: PropTypes.array.isRequired,
    onPlanValueChange: PropTypes.func,
    validationErrors: PropTypes.array
  }

  constructor(props) {
    super(props)
    this.state = { showReadOnly: props.manage === true }
  }

  componentDidUpdate(prevProps) {
    if (this.props.mode !== prevProps.mode) {
      this.setState({ showReadOnly: this.props.manage === true })
    }
  }

  render() {
    const { resourceType, mode, manage, team, kind, plan, originalPlan, schema, editableParams, onPlanValueChange, validationErrors } = this.props
    const showReadOnly = this.state.showReadOnly

    return (
      <>
        {manage ? null : (
          <Paragraph>
            <Text strong style={{ marginRight: '10px' }}>Show read-only parameters <IconTooltip text="Parameters may be read-only as defined by the plan policy or may only be editable on cluster creation. Turn this on to see these parameters." icon="info-circle" /></Text>
            <Switch checked={showReadOnly} onChange={(showReadOnly) => this.setState({ showReadOnly })} />
          </Paragraph>
        )}

        {Object.keys(schema.properties).map((name) => {
          const editable = mode !== 'view' &&
            (editableParams.includes('*') || editableParams.includes(name)) &&
            (schema.properties[name].const === undefined || schema.properties[name].const === null) &&
            (mode === 'create' || manage || !schema.properties[name].immutable) // Disallow editing of params which can only be set at create time when in 'use' mode
          // always show properties that are editable according to the policy, even when in view mode
          // properties not editable by the policy can be shown by enabling showReadOnly
          const forceShow = showReadOnly || (mode === 'view' && (editableParams.includes('*') || editableParams.includes(name)))

          return (
            <PlanOption
              manage={manage}
              mode={mode}
              team={team}
              resourceType={resourceType}
              kind={kind}
              plan={plan}
              originalPlan={originalPlan}
              key={name}
              name={name}
              property={schema.properties[name]}
              value={plan[name]}
              forceShow={forceShow}
              editable={editable}
              onChange={(n, v) => onPlanValueChange(n, v)}
              validationErrors={validationErrors} />
          )
        })}
      </>
    )
  }
}
