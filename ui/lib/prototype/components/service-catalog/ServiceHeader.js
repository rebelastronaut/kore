import PropTypes from 'prop-types'
import { Divider, Icon, Tag, Tooltip, Typography } from 'antd'
const { Paragraph, Title } = Typography

const ServiceHeader = ({ service }) => {

  const iconType = (category) => ({
    'Application': 'appstore',
    'Security': 'lock',
    'Monitoring': 'monitor'
  }[category])

  return (
    <>
      <div style={{ float: 'left' }}>
        <div style={{ display: 'inline-block', float: 'left' }}>
          {service.spec.icon ?
            <img src={`/static/images/${service.spec.icon}`} height={50} style={{ marginRight: '10px' }} /> :
            <Icon type={iconType(service.spec.category)} style={{ fontSize: '50px', marginRight: '10px' }} />
          }
        </div>
        <Title level={2} style={{ display: 'inline-block', marginTop: '3px', marginBottom: '25px' }}>{service.spec.name}</Title>
      </div>
      <div style={{ float: 'left', paddingTop: '14px', marginLeft: '15px' }}>
        {service.spec.category.map(c => <Tag key={c}>{c}</Tag>)}
      </div>
      <Paragraph style={{ clear: 'both' }}>{service.spec.description}</Paragraph>
      <Paragraph>
        <a target="_blank" rel="noopener noreferrer" href={service.spec.documentationURL} style={{ textDecoration: 'underline' }}>Documentation</a>
        {service.spec.prereqs ? (
          <>
            <Divider type="vertical"/>
            <Tooltip title={service.spec.prereqs.map(p => p.name).join(', ')}>
            Prerequisites
            </Tooltip>
          </>
        ) : null}
      </Paragraph>
    </>
  )
}

ServiceHeader.propTypes = {
  service: PropTypes.object.isRequired
}

export default ServiceHeader
