const { testUrl, timeout, expectTimeout } = require('../config')
import { setDefaultOptions } from 'expect-puppeteer'
import { waitForDrawerOpenClose } from './utils'

setDefaultOptions({ timeout: expectTimeout })
jest.setTimeout(timeout)

export class BasePage {
  constructor(p) {
    this.p = p
  }

  getFileName() {
    if (this.pagePath === '/') {
      return 'index'
    }
    return this.pagePath.replace('/', '').replace(/\//g,'-')
  }

  verifyPageURL() {
    expect(this.p.url()).toBe(`${testUrl}${this.pagePath}`)
  }

  /**
   * Close any open notifications/messages (e.g. success / error messages from saving resources).
   * 
   * Useful to call from beforeEach.
   */
  async closeAllNotifications() {
    const notifs = await this.p.$$('a.ant-notification-notice-close')
    await Promise.all(notifs.map(async (n) => {
      try {
        await n.click()
        await waitForDrawerOpenClose(this.p)
      } catch (err) {
        // Sometimes these randomly go out of scope while we're clicking, ignore errors
        // in that case.
      }
    }))
  }

  /**
   * Close any open drawer.
   */
  async closeDrawer() {
    await this.p.click('button.ant-drawer-close')
    await waitForDrawerOpenClose(this.p)
  }

  async visitPage(query = '') {
    await this.p.goto(`${testUrl}${this.pagePath}${query}`)
    await this.p.waitForSelector('body')
  }

  async getHeading() {
    return await this.p.$eval('h1', el => el.innerHTML)
  }

  async clickPrimaryButton(options) {
    options = options || { waitForNav: true }
    if (options.waitForNav) {
      await Promise.all([
        this.p.waitForNavigation(),
        this.p.click('.ant-btn-primary')
      ])
    } else {
      await this.p.click('.ant-btn-primary')
    }
  }
}
