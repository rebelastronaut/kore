import * as React from 'react'
import { Form, Icon, List, Button, Drawer, Input, Descriptions, InputNumber, Collapse, Modal, Alert, Switch, Typography } from 'antd'
const { Paragraph, Text } = Typography

import copy from '../../../utils/object-copy'
import PlanOptionBase from '../PlanOptionBase'
import ConstrainedDropdown from './ConstrainedDropdown'
import PlanOption from '../PlanOption'
import PlanOptionClusterMachineType from './PlanOptionClusterMachineType'
import NodePoolCost from '../../costs/NodePoolCost'
import IconTooltip from '../../utils/IconTooltip'

const imageTypes = [
  { value: 'Linux', display: 'Linux' },
  { value: 'Windows', display: 'Windows' }
]

const modes = [
  { value: 'System', display: 'System' },
  { value: 'User', display: 'User' }
]

export default class PlanOptionAKSNodePools extends PlanOptionBase {
  state = {
    selectedIndex: -1,
    prices: null
  }

  addNodePool = () => {
    if (!this.props.editable || !this.props.onChange) {
      return
    }

    // Create the default from the defaults on the plan schema
    const newNodePool = {
      name: ''
    }
    const properties = this.props.property.items.properties
    Object.keys(properties).forEach((k) => {
      if (properties[k].default !== undefined) {
        newNodePool[k] = copy(properties[k].default)
      }
    })

    // Need to handle the value being undefined in the case where this is a new plan or no
    // node pools are defined yet.
    let newValue
    if (this.props.value) {
      newValue = [ ...this.props.value, newNodePool ]
    } else {
      newValue = [ newNodePool ]
    }

    this.props.onChange(this.props.name, newValue)

    // Open the drawer to immediately edit the new node pool:
    this.setState({
      selectedIndex: newValue.length - 1
    })
  }

  removeNodePool = (idx) => {
    if (!this.props.editable || !this.props.onChange) {
      return
    }

    Modal.confirm({
      title: `Are you sure you want to remove node pool ${idx + 1} (${this.props.value[idx].name})?`,
      okText: 'Yes',
      okType: 'danger',
      cancelText: 'No',
      onOk: () => {
        this.setState({
          selectedIndex: -1
        })

        this.props.onChange(
          this.props.name,
          this.props.value.filter((_, i) => i !== idx)
        )
      }
    })
  }

  setNodePoolProperty = (idx, property, value) => {
    if (!this.props.editable || !this.props.onChange) {
      return
    }

    this.props.onChange(
      this.props.name,
      this.props.value.map((ng, i) => i !== idx ? ng : { ...ng, [property]: value })
    )
  }

  viewEditNodePool = (idx) => {
    this.setState({ selectedIndex: idx, prices: null })
  }

  closeNodePool = () => {
    this.setState({ selectedIndex: -1, prices: null })
  }

  nodePoolActions = (idx, editable) => {
    const actions = [
      <a id={`plan_nodepool_${idx}_viewedit`} key="viewedit" onClick={() => this.viewEditNodePool(idx)}><Icon type={editable ? 'edit' : 'eye'}></Icon></a>
    ]

    // Only show delete if we have more than one node pool
    if (editable && this.props.value && this.props.value.length > 1) {
      actions.push(<a id={`plan_nodepool_${idx}_del`} key="delete" onClick={() => this.removeNodePool(idx)}><Icon type="delete"></Icon></a>)
    }
    return actions
  }

  isEditable = (property) => {
    // always allow editing if the node pool is not part of the original pre-edited plan
    if (this.props.originalPlan && !this.props.originalPlan.nodePools[this.state.selectedIndex]) {
      return true
    }
    return super.isEditable(property)
  }

  render() {
    const { name, editable, property, plan, manage } = this.props
    const { displayName, valueOrDefault } = this.prepCommonProps(this.props, [])
    const { selectedIndex, prices, showReadOnly } = this.state
    const id_prefix = 'plan_nodepool'
    const selected = selectedIndex >= 0 ? valueOrDefault[selectedIndex] : null
    const description = manage ? 'Default node pools for clusters created from this plan' : null

    let ngNameClash = false, nodePoolCloseable = true
    if (selected) {
      // we have duplicate names if another node pool with a different index has the same name as this one
      ngNameClash = selected.name && selected.name.length > 0 && valueOrDefault.some((v, i) => i !== selectedIndex && v.name === selected.name)
      nodePoolCloseable = !ngNameClash && selected.name && selected.name.match(property.items.properties.name.pattern)
    }

    return (
      <>
        <Form.Item label={displayName} help={description}>
          <List id={`${id_prefix}s`} dataSource={valueOrDefault} renderItem={(ng, idx) => {
            return (
              <List.Item actions={this.nodePoolActions(idx, editable)}>
                <List.Item.Meta
                  title={<a id={`${id_prefix}_${idx}_viewedittitle`} onClick={() => this.viewEditNodePool(idx)}>{`Node Pool ${idx + 1} (${ng.name})`}</a>}
                  description={ng.enableAutoscaler ?
                    `Size: min=${ng.minSize} initial=${ng.size} max=${ng.maxSize} | Node type: ${ng.machineType}`
                    :
                    `Size: ${ng.size} | Node type: ${ng.machineType}`
                  }
                />
                {!this.hasValidationErrors(`${name}[${idx}]`, false) ? null : <Alert type="error" message="Validation errors - please edit and resolve" />}
              </List.Item>
            )
          }} />
          {!editable ? null : <Button id={`${id_prefix}_add`} onClick={this.addNodePool} disabled={!(plan.region)}>Add node pool {plan.region ? null : '(choose region first)'}</Button>}
          {this.validationErrors(name, true)}
        </Form.Item>
        <Drawer
          title={`Node Pool ${selected ? selectedIndex + 1 : ''}`}
          visible={Boolean(selected)}
          closable={nodePoolCloseable}
          maskClosable={nodePoolCloseable}
          onClose={() => this.closeNodePool()}
          width={800}
        >
          {!selected ? null : (
            <>
              {manage ? null : (
                <Paragraph>
                  <Text strong style={{ marginRight: '10px' }}>Show read-only parameters <IconTooltip text="Parameters may be read-only as defined by the plan policy or may be editable on cluster creation only. Turn this on to see these parameters." icon="info-circle" placement="bottom" /></Text>
                  <Switch checked={showReadOnly} onChange={(showReadOnly) => this.setState({ showReadOnly })} />
                </Paragraph>
              )}
              <Collapse defaultActiveKey={['basics','compute','metadata']}>
                <Collapse.Panel key="basics" header="Basic Configuration (name, versions, sizing)">
                  {this.isEditable(property.items.properties.name) || showReadOnly ? (
                    <Form.Item label="Name" help="Unique name for this group within the cluster">
                      <Input id={`${id_prefix}_name`} pattern={property.items.properties.name.pattern} value={selected.name} onChange={(e) => this.setNodePoolProperty(selectedIndex, 'name', e.target.value)} disabled={!this.isEditable(property.items.properties.name)} />
                      {this.validationErrors(`${name}[${selectedIndex}].name`)}
                      {!ngNameClash ? null : <Alert type="error" message="This name is already used by another node pool, it must be changed." />}
                      {selected.name && selected.name.match(property.items.properties.name.pattern) ? null : <Alert type="error" message="Name must be minimum 2, maximum 40 alpha-numeric characters and hyphens" />}
                    </Form.Item>
                  ) : null}

                  {this.isEditable(property.items.properties.mode) || showReadOnly ? (
                    <Form.Item label="Mode" help="Type of the node pool. System node pools serve the primary purpose of hosting critical system pods such as CoreDNS and tunnelfront. User node pools serve the primary purpose of hosting your application pods.">
                      <ConstrainedDropdown id={`${id_prefix}_mode`} allowedValues={modes} value={selected.mode} readOnly={!this.isEditable(property.items.properties.mode)} onChange={(v) => this.setNodePoolProperty(selectedIndex, 'mode', v)} />
                      {this.validationErrors(`${name}[${selectedIndex}].mode`)}
                    </Form.Item>
                  ) : null}

                  <PlanOption id={`${id_prefix}_version`} {...this.props} forceShow={showReadOnly} displayName="Version" name={`${name}[${selectedIndex}].version`} property={property.items.properties.version} value={selected.version} disabled={!this.isEditable(property.items.properties.version)} onChange={(_, v) => this.setNodePoolProperty(selectedIndex, 'version', v)} />

                  <PlanOption id={`${id_prefix}_enableAutoscaler`} {...this.props} forceShow={showReadOnly} displayName="Auto-scale" name={`${name}[${selectedIndex}].enableAutoscaler`} property={property.items.properties.enableAutoscaler} value={selected.enableAutoscaler} disabled={!this.isEditable(property.items.properties.enableAutoscaler)} onChange={(_, v) => this.setNodePoolProperty(selectedIndex, 'enableAutoscaler', v)} />
                  <Form.Item label="Pool size">
                    <Descriptions layout="horizontal" size="small">
                      {!selected.enableAutoscaler ? null : <Descriptions.Item label="Minimum">
                        <InputNumber id={`${id_prefix}_minSize`} value={selected.minSize} size="small" min={property.items.properties.minSize.minimum} max={selected.maxSize} disabled={!this.isEditable(property.items.properties.minSize)} onChange={(v) => this.setNodePoolProperty(selectedIndex, 'minSize', v)} />
                        {this.validationErrors(`${name}[${selectedIndex}].minSize`)}
                      </Descriptions.Item>}
                      <Descriptions.Item label={selected.enableAutoscaler ? 'Initial size' : null}>
                        <InputNumber id={`${id_prefix}_size`} value={selected.size} size="small" min={selected.enableAutoscaler ? selected.minSize : 1} max={selected.enableAutoscaler ? selected.maxSize : 99999} disabled={!this.isEditable(property.items.properties.size)} onChange={(v) => this.setNodePoolProperty(selectedIndex, 'size', v)} />
                        {this.validationErrors(`${name}[${selectedIndex}].size`)}
                      </Descriptions.Item>
                      {!selected.enableAutoscaler ? null : <Descriptions.Item label="Maximum">
                        <InputNumber id={`${id_prefix}_maxSize`} value={selected.maxSize} size="small" min={selected.minSize} disabled={!this.isEditable(property.items.properties.maxSize)} onChange={(v) => this.setNodePoolProperty(selectedIndex, 'maxSize', v)} />
                        {this.validationErrors(`${name}[${selectedIndex}].maxSize`)}
                      </Descriptions.Item>}
                    </Descriptions>
                  </Form.Item>
                  <NodePoolCost prices={prices} nodePool={selected} help="Adjust pool size and machine type to see the cost impacts" />
                  <PlanOption id={`${id_prefix}_maxPodsPerNode`} {...this.props} forceShow={showReadOnly} displayName="Max pods per node" name={`${name}[${selectedIndex}].maxPodsPerNode`} property={property.items.properties.maxPodsPerNode} value={selected.maxPodsPerNode} editable={this.isEditable(property.items.properties.maxPodsPerNode)} onChange={(_, v) => this.setNodePoolProperty(selectedIndex, 'maxPodsPerNode', v)} />
                </Collapse.Panel>
                <Collapse.Panel key="compute" header="Compute Configuration (image type, machine type, disk size)">
                  {this.isEditable(property.items.properties.imageType) || showReadOnly ? (
                    <Form.Item label="Image Type" help="The image type used by the nodes">
                      <ConstrainedDropdown id={`${id_prefix}_imageType`} allowedValues={imageTypes} value={selected.imageType} readOnly={!this.isEditable(property.items.properties.imageType)} onChange={(v) => this.setNodePoolProperty(selectedIndex, 'imageType', v)} />
                      {this.validationErrors(`${name}[${selectedIndex}].imageType`)}
                    </Form.Item>
                  ) : null}
                  <PlanOptionClusterMachineType id={`${id_prefix}_machineType`} {...this.props} forceShow={showReadOnly} editable={this.isEditable(property.items.properties.machineType)} displayName="Machine Type" name={`${name}[${selectedIndex}].machineType`} property={property.items.properties.machineType} value={selected.machineType} onChange={(_, v) => this.setNodePoolProperty(selectedIndex, 'machineType', v )} nodePriceSet={(prices) => this.setState({ prices })} />
                  <PlanOption id={`${id_prefix}_diskSize`} {...this.props} displayName="Instance Root Disk Size (GiB)" name={`${name}[${selectedIndex}].diskSize`} property={property.items.properties.diskSize} value={selected.diskSize} editable={this.isEditable(property.items.properties.diskSize)} onChange={(_, v) => this.setNodePoolProperty(selectedIndex, 'diskSize', v)} />
                </Collapse.Panel>
                <Collapse.Panel key="metadata" header="Labels & Taints">
                  <PlanOption id={`${id_prefix}_labels`} {...this.props} forceShow={showReadOnly} displayName="Labels" help="Labels help kubernetes workloads find this group" name={`${name}[${selectedIndex}].labels`} property={property.items.properties.labels} value={selected.labels} editable={this.isEditable(property.items.properties.labels)} onChange={(_, v) => this.setNodePoolProperty(selectedIndex, 'labels', v)} />
                  <PlanOption id={`${id_prefix}_taints`} {...this.props} forceShow={showReadOnly} displayName="Taints" help="Taints help kubernetes make scheduling decisions against nodepools" name={`${name}[${selectedIndex}].taints`} property={property.items.properties.taints} value={selected.taints} editable={this.isEditable(property.items.properties.taints)} onChange={(_, v) => this.setNodePoolProperty(selectedIndex, 'taints', v)} />
                </Collapse.Panel>
              </Collapse>
              <Form.Item style={{ marginTop: '20px' }}>
                <Button id={`${id_prefix}_close`} type="primary" disabled={!nodePoolCloseable} onClick={() => this.closeNodePool()}>{nodePoolCloseable ? 'Close' : 'Node pool not valid - please correct errors'}</Button>
              </Form.Item>
            </>
          )}
        </Drawer>
      </>
    )
  }
}
