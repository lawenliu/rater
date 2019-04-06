import React from 'react'
import superagent from 'superagent'
import CustomBreadcrumb from '../../../components/CustomBreadcrumb/index'
import {Card, List, Tooltip, Button, Form, Modal, Row, Col, BackTop, Upload, Icon} from 'antd'
import { EditorState, convertToRaw } from 'draft-js';
//import htmlToDraft from 'html-to-draftjs';
import 'react-draft-wysiwyg/dist/react-draft-wysiwyg.css';
import {MEDIA_UPLOAD_URL, DTYPE_AVATAR} from '../../../utils/Constants'
import { isAuthenticated } from '../../../utils/Session'
import CustomUpload from '../../../components/CustomUpload'

import './style.css'

const data = [
  '用户名：',
  '电话：',
  '邮箱：',
  '性别：',
  '年龄：',
  '地址：',
  '注册日期：',
];

@Form.create()
class Profile extends React.Component{
  state = {
    visibleAvatarUpload: false,
    confirmLoading: false,
    count: 100,
    avatar: require('../../../assets/img/defaultUser.jpg')
  }

  closeModal(a) {
    this.setState({
      [a]: false
    })
  }

  asynModalOnOk = (a) => {
    this.setState({
      confirmLoading: true,
    })
    setTimeout(() => this.setState({
      [a]: false,
      confirmLoading: false,
    }), 2000)
  }

  beforeUpload(file, fileList) {
    //限制图片 格式、size、分辨率
    const isJPG = file.type === 'image/jpeg';
    const isJPEG = file.type === 'image/jpeg';
    const isGIF = file.type === 'image/gif';
    const isPNG = file.type === 'image/png';
    if (!(isJPG || isJPEG || isGIF || isPNG)) {
      Modal.error({
        title: '只能上传JPG 、JPEG 、GIF、 PNG格式的图片~',
      });
      return false;
    }
    const isLt2M = file.size / 1024 / 1024 < 2;
    if (!isLt2M) {
      Modal.error({
        title: '超过2M限制 不允许上传~',
      });
      return false;
    }
    return (isJPG || isJPEG || isGIF || isPNG) && isLt2M;
  }

  handleChange = (info) => {
    if (info.status === 'uploading') {
      this.setState({loading: true});
      return;
    }
    if (info.status === 'done') {
      // Get this url from response in real world.
    } else if (info.status === 'error') {
      Modal.error({
        title: `文件上传失败（${info.message}）`
      });
      this.setState({
        loading: false
      })
    }
  }

  render(){
    const {avatar, visibleAvatarUpload, confirmLoading} = this.state
    return (
      <div>
        <CustomBreadcrumb arr={['个人','个人信息']}/>
        <Card bordered={false} className='card-item'>
          <Row type='flex' align='middle'>
            <Col type='flex'>
            <div className='avatar-bar'>
              <img className="avatar" src={avatar} alt=""/><br/><br/>
              <Tooltip title='上传头像，让更多人认识你' placement='right'>
                <Button onClick={() => this.setState({visibleAvatarUpload: true})}>上传头像</Button>
              </Tooltip>
              <Modal
                visible={visibleAvatarUpload}
                title='上传头像'
                onOk={() => this.asynModalOnOk('visibleAvatarUpload')}
                onCancel={() => this.closeModal('visibleAvatarUpload')}
                footer={
                  <div>
                    <Button onClick={() => this.closeModal('visibleAvatarUpload')}>取消</Button>
                    <Button type="primary" loading={confirmLoading} onClick={() => this.asynModalOnOk('visibleAvatarUpload')}>
                      提交
                    </Button>
                  </div>}
              >
                <Row type='flex' align='middle'>
                  <Col span={8}>
                    <CustomUpload
                      action={MEDIA_UPLOAD_URL}
                      dtype={DTYPE_AVATAR}
                      showUploadedList={false}
                      beforeUpload={this.beforeUpload}
                      handleChange={this.handleChange} />
                  </Col>
                  <Col span={12}>
                    只能上传JPG 、JPEG 、GIF、 PNG格式的不大于2M的图片
                  </Col>
                </Row>
              </Modal>
            </div>
            </Col>
            <Col span={12} align="left">
              <List dataSource={data}
                className='no-border'
                bordered={false}
                size='default'
                renderItem={item => (<List.Item>{item}</List.Item>)}/>
            </Col>
          </Row>
        </Card>
        <BackTop visibilityHeight={200} style={{right: 50}}/>
      </div>
    )
  }
}

export default Profile