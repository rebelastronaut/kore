import * as React from 'react'
import { Form, Icon, List, Button, Drawer, Input, Descriptions, InputNumber, Checkbox, Collapse, Radio, Modal, Alert } from 'antd'

import PlanOptionBase from '../PlanOptionBase'
import PlanOption from '../PlanOption'
import copy from '../../../utils/object-copy'
import NodePoolCost from '../../costs/NodePoolCost'
import PlanOptionClusterMachineType from './PlanOptionClusterMachineType'

export default class PlanOptionEKSNodeGroups extends PlanOptionBase {
  constructor(props) {
    super(props)
  }

  static AMI_TYPE_GENERAL = 'AL2_x86_64'
  static AMI_TYPE_GPU = 'AL2_x86_64_GPU'

  state = {
    selectedIndex: -1,
    prices: null
  }

  addNodeGroup = () => {
    if (!this.props.editable || !this.props.onChange) {
      return
    }

    // Create the default from the defaults on the plan schema
    const newNodeGroup = {}
    const properties = this.props.property.items.properties
    Object.keys(properties).forEach((k) => {
      if (properties[k].default !== undefined) {
        newNodeGroup[k] = copy(properties[k].default)
      }
    })

    // Need to handle the value being undefined in the case where this is a new plan or no
    // node groups are defined yet.
    let newValue
    if (this.props.value) {
      newValue = [ ...this.props.value, newNodeGroup ]
    } else {
      newValue = [ newNodeGroup ]
    }

    this.props.onChange(this.props.name, newValue)

    // Open the drawer to immediately edit the new node group:
    this.setState({
      selectedIndex: newValue.length - 1
    })
  }

  removeNodeGroup = (idx) => {
    if (!this.props.editable || !this.props.onChange) {
      return
    }

    Modal.confirm({
      title: `Are you sure you want to remove node group ${idx + 1} (${this.props.value[idx].name})?`,
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

  setAmiType = (idx, value) => {
    if (!this.props.editable || !this.props.onChange) {
      return
    }

    this.props.onChange(
      this.props.name,
      this.props.value.map((ng, i) => i !== idx ? ng : { ...ng, amiType: value, instanceType: null })
    )
  }

  onReleaseVersionChecked = (idx, checked) => {
    if (!this.props.editable || !this.props.onChange) {
      return
    }

    const releaseVersion = !checked ? `${this.props.plan.version}.` : undefined

    this.props.onChange(
      this.props.name,
      this.props.value.map((ng, i) => i !== idx ? ng : { ...ng, releaseVersion })
    )
  }

  setNodeGroupProperty = (idx, property, value) => {
    if (!this.props.editable || !this.props.onChange) {
      return
    }

    this.props.onChange(
      this.props.name,
      this.props.value.map((ng, i) => i !== idx ? ng : { ...ng, [property]: value })
    )
  }

  viewEditNodeGroup = (idx) => {
    this.setState({ selectedIndex: idx, prices: null })
  }

  closeNodeGroup = () => {
    this.setState({ selectedIndex: -1, prices: null })
  }

  nodeGroupActions = (idx, editable) => {
    const actions = [
      <a id={`plan_nodegroup_${idx}_viewedit`} key="viewedit" onClick={() => this.viewEditNodeGroup(idx)}><Icon type={editable ? 'edit' : 'eye'}></Icon></a>
    ]

    // Only show delete if we have more than one node group
    if (editable && this.props.value && this.props.value.length > 1) {
      actions.push(<a id={`plan_nodegroup_${idx}_del`} key="delete" onClick={() => this.removeNodeGroup(idx)}><Icon type="delete"></Icon></a>)
    }
    return actions
  }

  isEditable = (property) => {
    // always allow editing if the node group is not part of the original pre-edited plan
    if (this.props.originalPlan && !this.props.originalPlan.nodeGroups[this.state.selectedIndex]) {
      return true
    }
    return super.isEditable(property)
  }

  render() {
    const { name, editable, property, plan } = this.props
    const { displayName, valueOrDefault } = this.prepCommonProps(this.props, [])
    const { selectedIndex } = this.state
    const id_prefix = 'plan_nodegroup'
    const selected = selectedIndex >= 0 ? valueOrDefault[selectedIndex] : null
    const description = this.props.manage ? 'Set default node groups for clusters created from this plan' : 'Manage node groups for this cluster'

    let amiType = null
    let releaseVersionSet = false, ngNameClash = false, nodeGroupCloseable = true, instanceTypeFilter = null
    if (selected) {
      amiType = selected.amiType || 'AL2_x86_64'
      releaseVersionSet = selected.releaseVersion && selected.releaseVersion.length > 0
      // we have duplicate names if another node group with a different index has the same name as this one
      ngNameClash = selected.name && selected.name.length > 0 && valueOrDefault.some((v, i) => i !== selectedIndex && v.name === selected.name)
      nodeGroupCloseable = !ngNameClash && selected.name && selected.name.length > 0
      instanceTypeFilter = amiType === PlanOptionEKSNodeGroups.AMI_TYPE_GENERAL ? (c) => c !== 'GPU instance' : (c) => c === 'GPU instance'
    }

    return (
      <>
        <Form.Item label={displayName} help={description}>
          <List id={`${id_prefix}s`} dataSource={valueOrDefault} renderItem={(ng, idx) => {
            return (
              <List.Item actions={this.nodeGroupActions(idx, editable)}>
                <List.Item.Meta
                  title={<a id={`${id_prefix}_${idx}_viewedittitle`} onClick={() => this.viewEditNodeGroup(idx)}>{`Node Group ${idx + 1} (${ng.name})`}</a>}
                  description={ng.enableAutoscaler ?
                    `Size: min=${ng.minSize} initial=${ng.desiredSize} max=${ng.maxSize} | Node type: ${ng.instanceType}`
                    :
                    `Desired Size: ${ng.desiredSize} | Node type: ${ng.instanceType}`
                  }
                />
                {!this.hasValidationErrors(`${name}[${idx}]`, false) ? null : <Alert type="error" message="Validation errors - please edit and resolve" />}
              </List.Item>
            )
          }} />
          {!editable ? null : <Button id={`${id_prefix}_add`} onClick={this.addNodeGroup} disabled={!(plan.region)}>Add node group {plan.region ? null : '(choose region first)'}</Button>}
          {this.validationErrors(name, true)}
        </Form.Item>
        <Drawer
          title={`Node Group ${selected ? selectedIndex + 1 : ''}`}
          visible={Boolean(selected)}
          closable={!ngNameClash}
          maskClosable={!ngNameClash}
          onClose={() => this.closeNodeGroup()}
          width={800}
        >
          {!selected ? null : (
            <>
              <Collapse defaultActiveKey={['basics','compute','metadata']}>
                <Collapse.Panel key="basics" header="Basic Configuration (name, sizing)">
                  <Form.Item label="Name" help="Unique name for this group within the cluster">
                    <Input id={`${id_prefix}_name`} value={selected.name} onChange={(e) => this.setNodeGroupProperty(selectedIndex, 'name', e.target.value)} disabled={!this.isEditable(property.items.properties.name)} />
                    {this.validationErrors(`${name}[${selectedIndex}].name`)}
                    {!ngNameClash ? null : <Alert type="error" message="This name is already used by another node group, it must be changed." />}
                    {selected.name && selected.name.length > 0 ? null : <Alert type="error" message="Name must be set" />}
                  </Form.Item>
                  <PlanOption id={`${id_prefix}_enableAutoscaler`} {...this.props} hideNonEditable={false} displayName="Auto-scale" name={`${name}[${selectedIndex}].enableAutoscaler`} property={property.items.properties.enableAutoscaler} value={selected.enableAutoscaler} editable={this.isEditable(property.items.properties.enableAutoscaler)} onChange={(_, v) => this.setNodeGroupProperty(selectedIndex, 'enableAutoscaler', v)} />
                  <Form.Item label="Group Size">
                    <Descriptions layout="horizontal" size="small">
                      {!selected.enableAutoscaler ? null : <Descriptions.Item label="Minimum">
                        <InputNumber id={`${id_prefix}_minSize`} value={selected.minSize} size="small" min={property.items.properties.minSize.minimum} max={selected.maxSize} disabled={!this.isEditable(property.items.properties.minSize)} onChange={(v) => this.setNodeGroupProperty(selectedIndex, 'minSize', v)} />
                        {this.validationErrors(`${name}[${selectedIndex}].minSize`)}
                      </Descriptions.Item>}
                      <Descriptions.Item label={selected.enableAutoscaler ? 'Initial size' : null}>
                        <InputNumber id={`${id_prefix}_desiredSize`} value={selected.desiredSize} size="small" min={selected.enableAutoscaler ? selected.minSize : 1} max={selected.enableAutoscaler ? selected.maxSize : undefined} disabled={!this.isEditable(property.items.properties.desiredSize)} onChange={(v) => this.setNodeGroupProperty(selectedIndex, 'desiredSize', v)} />
                        {this.validationErrors(`${name}[${selectedIndex}].desiredSize`)}
                      </Descriptions.Item>
                      {!selected.enableAutoscaler ? null : <Descriptions.Item label="Maximum">
                        <InputNumber id={`${id_prefix}_maxSize`} value={selected.maxSize} size="small" min={selected.minSize} editable={this.isEditable(property.items.properties.maxSize)} onChange={(v) => this.setNodeGroupProperty(selectedIndex, 'maxSize', v)} />
                        {this.validationErrors(`${name}[${selectedIndex}].maxSize`)}
                      </Descriptions.Item>}
                    </Descriptions>
                  </Form.Item>
                  <NodePoolCost prices={this.state.prices} nodePool={selected} help="Adjust group size and instance type to see the cost impacts" />
                </Collapse.Panel>
                <Collapse.Panel key="compute" header="Compute Configuration (instance type, GPU or regular workload)">
                  <Form.Item label={property.items.properties.amiType.title} help={property.items.properties.amiType.description}>
                    <Radio.Group id={`${id_prefix}_amiType`} value={amiType} onChange={(v) => this.setAmiType(selectedIndex, v.target.value)} disabled={!this.isEditable(property.items.properties.amiType)}>
                      <Radio value={PlanOptionEKSNodeGroups.AMI_TYPE_GENERAL}>General Purpose</Radio>
                      <Radio value={PlanOptionEKSNodeGroups.AMI_TYPE_GPU}>GPU</Radio>
                    </Radio.Group>
                    {this.validationErrors(`${name}[${selectedIndex}].amiType`)}
                  </Form.Item>
                  <Form.Item label="AWS AMI Version" help={!releaseVersionSet ? undefined : <><b>Must</b> be for Kubernetes <b>{plan.version}</b> and <b>{amiType === PlanOptionEKSNodeGroups.AMI_TYPE_GPU ? 'GPU' : 'general'}</b> workloads. Find <a target="_blank" rel="noopener noreferrer" href="https://docs.aws.amazon.com/eks/latest/userguide/eks-linux-ami-versions.html">supported versions</a> in AWS documentation.</>}>
                    <Checkbox id={`${id_prefix}_releaseVersion_latest`} checked={!releaseVersionSet} disabled={!this.isEditable(property.items.properties.releaseVersion)} onChange={(v) => this.onReleaseVersionChecked(selectedIndex, v.target.checked)}/> Use latest (<b>recommended</b>)
                    {!releaseVersionSet ? null : <Input id={`${id_prefix}_releaseVersion_custom`} value={selected.releaseVersion} placeholder={this.describe(property.items.properties.releaseVersion)} disabled={!this.isEditable(property.items.properties.releaseVersion)} onChange={(e) => this.setNodeGroupProperty(selectedIndex, 'releaseVersion', e.target.value)} />}
                    {this.validationErrors(`${name}[${selectedIndex}].releaseVersion`)}
                  </Form.Item>
                  <PlanOptionClusterMachineType filterCategories={instanceTypeFilter} id={`${id_prefix}_instanceType`} {...this.props} editable={this.isEditable(property.items.properties.instanceType)} displayName="AWS Instance Type" name={`${name}[${selectedIndex}].instanceType`} property={property.items.properties.instanceType} value={selected.instanceType} onChange={(_, v) => this.setNodeGroupProperty(selectedIndex, 'instanceType', v )} nodePriceSet={(prices) => this.setState({ prices })} />
                  <PlanOption {...this.props} hideNonEditable={false} id={`${id_prefix}_instanceType`} displayName="Instance Root Disk Size (GiB)" name={`${name}[${selectedIndex}].diskSize`} property={property.items.properties.diskSize} value={selected.diskSize} editable={this.isEditable(property.items.properties.diskSize)} onChange={(_, v) => this.setNodeGroupProperty(selectedIndex, 'diskSize', v)} />
                </Collapse.Panel>
                <Collapse.Panel key="metadata" header="Metadata (labels, tags, etc)">
                  <PlanOption {...this.props} hideNonEditable={false} id={`${id_prefix}_labels`} displayName="Labels" help="Labels help kubernetes workloads find this group" name={`${name}[${selectedIndex}].labels`} property={property.items.properties.labels} value={selected.labels} editable={this.isEditable(property.items.properties.labels)} onChange={(_, v) => this.setNodeGroupProperty(selectedIndex, 'labels', v)} />
                  <PlanOption {...this.props} hideNonEditable={false} id={`${id_prefix}_tags`} displayName="Tags" help="AWS tags to apply to the node group" name={`${name}[${selectedIndex}].tags`} property={property.items.properties.tags} value={selected.tags} editable={this.isEditable(property.items.properties.tags)} onChange={(_, v) => this.setNodeGroupProperty(selectedIndex, 'tags', v)} />
                </Collapse.Panel>
                <Collapse.Panel key="ssh" header="SSH Connectivity (keys, security groups)">
                  <PlanOption {...this.props} hideNonEditable={false} id={`${id_prefix}_eC2SSHKey`} displayName="EC2 SSH Key" name={`${name}[${selectedIndex}].eC2SSHKey`} property={property.items.properties.eC2SSHKey} value={selected.eC2SSHKey} editable={this.isEditable(property.items.properties.eC2SSHKey)} onChange={(_, v) => this.setNodeGroupProperty(selectedIndex, 'eC2SSHKey', v)} />
                  <PlanOption {...this.props} hideNonEditable={false} id={`${id_prefix}_sshSourceSecurityGroups`} displayName="SSH Security Groups" name={`${name}[${selectedIndex}].sshSourceSecurityGroups`} property={property.items.properties.sshSourceSecurityGroups} value={selected.sshSourceSecurityGroups} editable={this.isEditable(property.items.properties.sshSourceSecurityGroups)} onChange={(_, v) => this.setNodeGroupProperty(selectedIndex, 'sshSourceSecurityGroups', v)} />
                </Collapse.Panel>
              </Collapse>
              <Form.Item>
                <Button type="primary" id={`${id_prefix}_close`} disabled={!nodeGroupCloseable} onClick={() => this.closeNodeGroup()}>{nodeGroupCloseable ? 'Close' : 'Node group not valid - correct errors'}</Button>
              </Form.Item>
            </>
          )}
        </Drawer>
      </>
    )
  }
}
