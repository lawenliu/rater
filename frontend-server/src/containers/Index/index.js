import React from 'react'
import {Layout} from 'antd'
import HeaderBar from '../../components/HeaderBar'
import CustomCarousel from '../../components/CustomCarousel/index'
import Gallery from '../../components/Gallery/index'

const {Header, Content, Footer} = Layout


class Index extends React.Component{
  state = {
    collapsed: false
  }

  toggle = () => {
    // console.log(this)  状态提升后，到底是谁调用的它
    this.setState({
      collapsed: !this.state.collapsed
    })
  }
  render() {
    // 设置Sider的minHeight可以使左右自适应对齐
    return (
      <div id='page'>
        <Layout>
          <Header style={{background: '#fff', padding: '0 16px'}}>
            <HeaderBar collapsed={this.state.collapsed} onToggle={this.toggle}/>
          </Header>
          <Content>
            <CustomCarousel arrows effect='fade' className='size' />
            <Gallery />
          </Content>
          <Footer style={{textAlign: 'center'}}>
            <p>Lazy Bones ©2019 Created by lwc541117@hotmail.com </p>
            <a target='_blank' rel='noopener noreferrer' href='https://github.com/lawenliu'>github地址</a>
          </Footer>
        </Layout>
      </div>
    );
  }
}
export default Index