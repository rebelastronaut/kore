import { ConfigureCloudPage } from '../configure-cloud'
import { waitForDrawerOpenClose, clearFillTextInput, modalYes } from '../../../utils'

export class ConfigureCloudAWSAccounts extends ConfigureCloudPage {
  constructor(p) {
    super(p)
    this.pagePath = '/configure/cloud/AWS/accounts'
  }

  async openTab() {
    await this.selectCloud('aws')
    await this.selectSubTab('Account credentials', 'AWS/accounts')
  }

  async checkAccountListed(name) {
    await expect(this.p).toMatchElement(`#ekscreds_${name}`)
  }

  async add() {
    await expect(this.p).toClick('button', { text: '+ New' })
    await waitForDrawerOpenClose(this.p)
    await expect(this.p).toMatch('New AWS account')
  }

  async edit(name, accountID) {
    await this.p.click(`a#ekscreds_edit_${name}`)
    await waitForDrawerOpenClose(this.p)
    await expect(this.p).toMatch(`AWS account: ${accountID}`)
  }

  async populate({ accountID, accessKeyID, secretAccessKey, name, summary }) {
    await clearFillTextInput(this.p, 'eks_credentials_accountID', accountID)
    await clearFillTextInput(this.p, 'eks_credentials_accessKeyID', accessKeyID)
    await clearFillTextInput(this.p, 'eks_credentials_secretAccessKey', secretAccessKey)
    await clearFillTextInput(this.p, 'eks_credentials_name', name)
    await clearFillTextInput(this.p, 'eks_credentials_summary', summary)
  }

  async replaceKey(accessKeyID, secretAccessKey) {
    await this.p.type('input#eks_credentials_replace_key',' ')
    // Wait for service account text field to be shown:
    await expect(this.p).toMatch('Access key ID')
    await clearFillTextInput(this.p, 'eks_credentials_accessKeyID', accessKeyID)
    await clearFillTextInput(this.p, 'eks_credentials_secretAccessKey', secretAccessKey)
  }

  async saveButtonDisabled() {
    return (await this.p.$('button#save[disabled]')) !== null
  }

  async save() {
    await this.p.click('button#save')
    await waitForDrawerOpenClose(this.p)
  }

  async delete(name) {
    await this.p.click(`a#ekscreds_del_${name}`)
  }

  async confirmDelete() {
    await modalYes(this.p, 'Are you sure you want to delete the credentials')
  }
}
