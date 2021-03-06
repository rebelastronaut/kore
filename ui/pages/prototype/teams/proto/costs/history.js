import React from 'react'
import Link from 'next/link'
import { Collapse, Typography } from 'antd'
const { Paragraph, Text, Title } = Typography

import Breadcrumb from '../../../../../lib/components/layout/Breadcrumb'
import TeamMonthlyCostTable from '../../../../../lib/prototype/components/costs/TeamMonthlyCostTable'

class TeamCostsHistory extends React.Component {

  render() {
    return (
      <>
        <Breadcrumb items={[{ text: 'Proto' }, { text: 'Team costs history' }]}/>

        <Title level={3}>Historical team costs</Title>
        <Paragraph>
          <Link href="/prototype/teams/proto/costs">
            <a style={{ fontSize: '14px', textDecoration: 'underline' }}>See current cost</a>
          </Link>
        </Paragraph>

        <Collapse bordered={false}>
          <Collapse.Panel className="enlarged-header" header="May 2020" extra={<Text>£765.43</Text>}>
            <TeamMonthlyCostTable />
          </Collapse.Panel>
          <Collapse.Panel className="enlarged-header" header="April 2020" extra={<Text>£734.14</Text>}>
            <TeamMonthlyCostTable />
          </Collapse.Panel>
          <Collapse.Panel className="enlarged-header" header="March 2020" extra={<Text>£695.97</Text>}>
            <TeamMonthlyCostTable />
          </Collapse.Panel>
        </Collapse>
      </>
    )
  }
}

export default TeamCostsHistory
