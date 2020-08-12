import PropTypes from 'prop-types'
import { Icon, Tooltip } from 'antd'

const IconTooltip = ({ icon, text, color, placement, onClick }) => {
  color = color || '#13c2c2' // info color from antd-var-overrides.less
  icon = icon || 'info-circle'
  return (
    <Tooltip title={text} placement={placement}>
      {onClick ? <a style={{ marginLeft: '5px' }} onClick={onClick}><Icon type={icon} theme="filled" style={{ color }} /></a> : <Icon type={icon} theme="filled" style={{ color }} /> }
    </Tooltip>
  )
}

IconTooltip.propTypes = {
  icon: PropTypes.string.isRequired,
  text: PropTypes.string.isRequired,
  color: PropTypes.string,
  placement: PropTypes.string,
  onClick: PropTypes.func
}

export default IconTooltip
