import React from 'react'
import { Icon, Badge, Dropdown, Menu, Modal } from 'antd'
import screenfull from 'screenfull'
import { inject, observer } from 'mobx-react'
import { Link, withRouter } from 'react-router-dom'
import { isAuthenticated } from '../../utils/Session'

//withRouter一定要写在前面，不然路由变化不会反映到props中去
@withRouter @inject('appStore') @observer
class HeaderBar extends React.Component {
  state = {
    icon: 'arrows-alt',
    count: 100,
    avatar: require('./img/04.jpg')
  }

  componentDidMount () {
    screenfull.onchange(() => {
      this.setState({
        icon: screenfull.isFullscreen ? 'shrink' : 'arrows-alt'
      })
    })
  }

  componentWillUnmount () {
    screenfull.off('change')
  }

  toggle = () => {
    this.props.onToggle()
  }
  screenfullToggle = () => {
    if (screenfull.enabled) {
      screenfull.toggle()
    }
  }
  logout = () => {
    this.props.appStore.toggleLogin(false)
    this.props.history.push(this.props.location.pathname)
  }

  personal = () => {
    this.props.history.push('/manage/profile/info')
  }

  setting = () => {
    this.props.history.push('/manage/profile/chpwd')
  }

  help = () => {
    this.props.history.push('/manage/info/help')
  }

  render () {
    const {icon, count, avatar} = this.state
    const {appStore, collapsed, location} = this.props
    const notLogin = (
      <div>
        <Link to={{pathname: '/login', state: {from: location}}} style={{color: 'rgba(0, 0, 0, 0.65)'}}>登录</Link>&nbsp;
        <img src={require('../../assets/img/defaultUser.jpg')} alt=""/>
      </div>
    )
    const menu = (
      <Menu className='menu'>
        <Menu.ItemGroup title={`你好${isAuthenticated()}`} className='menu-group'>
          <Menu.Item><Icon type='setting'/><span onClick={this.personal}>个人中心</span></Menu.Item>
          <Menu.Item><Icon type='setting'/><span onClick={this.setting}>修改密码</span></Menu.Item>
        </Menu.ItemGroup>
        <Menu.ItemGroup title='设置中心' className='menu-group'>
          <Menu.Item><Icon type='setting'/><span onClick={this.help}>帮助文档</span></Menu.Item>
          <Menu.Item><Icon type='setting'/><span onClick={this.logout}>退出登录</span></Menu.Item>
        </Menu.ItemGroup>
      </Menu>
    )
    const login = (
      <Dropdown overlay={menu}>
        <img src={avatar} alt=""/>
      </Dropdown>
    )
    return (
      <div id='headerbar'>
        <Icon
          type={collapsed ? 'menu-unfold' : 'menu-fold'}
          className='trigger'
          onClick={this.toggle}/>
        <div style={{lineHeight: '64px', float: 'right'}}>
          <ul className='header-ul'>
            <li><Icon type={icon} onClick={this.screenfullToggle}/></li>
            <li onClick={() => this.setState({count: 0})}>
              <Badge count={appStore.isLogin ? count : 0} overflowCount={99} style={{marginRight: -17}}>
                <Icon type="notification"/>
              </Badge>
            </li>
            <li>
              {appStore.isLogin ? login : notLogin}
            </li>
          </ul>
        </div>

      </div>
    )
  }
}

export default HeaderBar