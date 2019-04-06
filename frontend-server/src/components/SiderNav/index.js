import React from 'react'
import CustomMenu from "../CustomMenu/index";
import { withRouter } from 'react-router-dom'

const menus = [
  {
    title: '首页',
    icon: 'home',
    key: '/home'
  },
  {
    title: '个人',
    icon: 'laptop',
    key: '/home/profile',
    subs: [
      {key: '/manage/profile', title: '个人信息', icon: '',},
      {key: '/manage/profile/chpwd', title: '修改密码', icon: '',},
    ]
  },
  {
    title: '分享',
    icon: 'bars',
    key: '/home/cotent',
    subs: [
      {key: '/manage/content', title: '历史记录', icon: '',},
      {key: '/manage/content/image', title: '分享图片', icon: '',},
      {key: '/manage/content/video', title: '分享视频', icon: '',},
      {key: '/manage/content/draft', title: '分享文章', icon: '',},
      {key: '/manage/content/shop', title: '分享好货', icon: '',},
    ]
  },
  {
    title: '帮助文档',
    icon: 'info-circle-o',
    key: '/manage/info',
    subs: [
      {key: '/manage/info/help', title: '帮助文档', icon: ''},
      {key: '/manage/info/about', title: '关于我们', icon: ''},
    ]
  }
]

@withRouter
class SiderNav extends React.Component {
  
  gobackHome = () => {
    this.props.history.push('/')
  }

  render() {
    return (
      <div style={{height: '100vh',overflowY:'scroll'}}>
        <div style={styles.logo} onClick={this.gobackHome}>Lazy Bones</div>
        <CustomMenu menus={menus}/>
      </div>
    )
  }
}

const styles = {
  logo: {
    height: '32px',
    background: 'rgba(255, 255, 255, .2)',
    margin: '16px',
    color: '#ffffff',
    padding: '4px',
    fontWeight: 'bold',
    textAlign: 'center',
    cursor: 'pointer'
  }
}

export default SiderNav