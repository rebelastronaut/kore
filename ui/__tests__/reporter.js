const page = global.page

async function takeScreenshot(name) {
  await page.screenshot({
    fullPage: true,
    path: `__tests__/e2e/screenshots/${name}.png`
  })
}

const reporter = {
  specDone: async (result) => {
    if (result.status === 'failed') {
      try {
        await takeScreenshot(result.fullName)
      } catch (e) {
        // just don't take a screenshot
      }
    }
  },
}

/* eslint-disable no-undef */
jasmine.getEnv().addReporter(reporter)
