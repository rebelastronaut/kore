import VerifiedAllocatedResourceForm from '../resources/VerifiedAllocatedResourceForm'
import KoreApi from '../../kore-api'
import { Form, Input, Alert, Card, Checkbox } from 'antd'
import AllocationHelpers from '../../utils/allocation-helpers'

class GKECredentialsForm extends VerifiedAllocatedResourceForm {

  getResource = async metadataName => {
    const api = await KoreApi.client()
    const gkeCredentialsResult = await api.GetGKECredential(this.props.team, metadataName)
    gkeCredentialsResult.allocation = await AllocationHelpers.getAllocationForResource(gkeCredentialsResult)
    return gkeCredentialsResult
  }

  putResource = async values => {
    const api = await KoreApi.client()
    values.name = this.getMetadataName(values)
    const secretName = values.name
    const teamResources = KoreApi.resources().team(this.props.team)

    if (!this.props.data || this.state.replaceKey) {
      const secretData = { service_account_key: btoa(values.account) }
      const secretResource = teamResources.generateSecretResource(secretName, 'gke-credentials', `GCP ${values.project} project Service Account`, secretData)
      await api.UpdateTeamSecret(this.props.team, values.name, secretResource)
    }
    const gkeCredResource = teamResources.generateGKECredentialsResource(values, secretName)
    const gkeResult = await api.UpdateGKECredential(this.props.team, values.name, gkeCredResource)
    gkeResult.allocation = await this.storeAllocation(gkeCredResource, values)
    return gkeResult
  }

  allocationFormFieldsInfo = {
    allocationMissing: {
      infoMessage: 'This project credential is not allocated to any teams',
      infoDescription: 'Give the project credential a Name and Description below and enter Allocated team(s) as appropriate. Once complete, click Save to allocate it."'
    },
    nameSection: {
      infoMessage: 'Help Kore teams understand this project credential',
      infoDescription: 'Give this project credential a name and description to help teams choose the correct one.',
      nameHelp: 'The name for the project credential eg. MyOrg project-one',
      descriptionHelp: 'A description of the project credential to help when choosing between them'
    },
    allocationSection: {
      infoMessage: 'Make this project credential available to teams in Kore',
      infoDescription: 'This will give teams the ability to create clusters within the project.',
      allTeamsWarningMessage: 'This project credential will be available to all teams',
      allTeamsWarningDescription: 'No teams exist in Kore yet, therefore currently this project credential will be available to all teams created in the future. If you wish to restrict this please return here and allocate to teams once they have been created.',
      allocateExtra: 'If nothing selected then this project will credential be available to ALL teams'
    }
  }

  resourceFormFields = () => {
    const { form, data } = this.props
    const { replaceKey } = this.state
    return (
      <Card style={{ marginBottom: '20px' }}>
        <Alert
          message="Project and service account"
          description="Retrieve these values from your GCP project. Providing these gives Kore the ability to create clusters within the project."
          type="info"
          style={{ marginBottom: '20px' }}
        />
        <Form.Item label="Project ID" validateStatus={this.fieldError('project') ? 'error' : ''} help={this.fieldError('project') || 'The GCP project ID that Kore will be able to build clusters within.'}>
          {form.getFieldDecorator('project', {
            rules: [{ required: true, message: 'Please enter your project ID!' }],
            initialValue: data && data.spec.project
          })(
            <Input placeholder="Project ID" name="project" />,
          )}
        </Form.Item>

        {data ? (
          <>
            <Alert
              message="For security reasons, the service account key cannot be retrieved after creation. If you need to replace it, tick the box below."
              type="warning"
              style={{ marginTop: '10px' }}
            />
            <Form.Item label="Replace key">
              <Checkbox id="gke_credentials_replace_key" onChange={(e) => this.setState({ replaceKey: e.target.checked })}></Checkbox>
            </Form.Item>
          </>
        ) : null}

        {!data || replaceKey ? (
          <>
            <Form.Item label="Service Account JSON" labelCol={{ span: 24 }} wrapperCol={{ span: 24 }} validateStatus={this.fieldError('account') ? 'error' : ''} help={this.fieldError('account') || 'The Service Account key in JSON format, with GKE admin permissions on the GCP project'}>
              {form.getFieldDecorator('account', {
                rules: [{ required: true, message: 'Please enter your Service Account key!' }]
              })(
                <Input.TextArea name="service_account_json" autoSize={{ minRows: 4, maxRows: 10  }} placeholder="Service Account JSON" />,
              )}
            </Form.Item>
          </>
        ) : null}
      </Card>
    )
  }
}

const WrappedGKECredentialsForm = Form.create({ name: 'gke_credentials' })(GKECredentialsForm)

export default WrappedGKECredentialsForm
