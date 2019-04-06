import React from 'react'
import superagent from 'superagent'
import CustomBreadcrumb from '../../../components/CustomBreadcrumb/index'
import {Card, Menu, Dropdown, Icon, Form, Input, Button, Col, BackTop, message} from 'antd'
import { EditorState, convertToRaw } from 'draft-js';
//import ContentState from 'draft-js'
import { Editor } from 'react-draft-wysiwyg';
import draftToHtml from 'draftjs-to-html';
//import htmlToDraft from 'html-to-draftjs';
import 'react-draft-wysiwyg/dist/react-draft-wysiwyg.css';
import {MEDIA_UPLOAD_URL, POSTER_UPLOAD_URL, DTYPE_IMAGE, DTYPE_DRAFT} from '../../../utils/Constants'
import { isAuthenticated } from '../../../utils/Session'
import { calculateWidth } from '../../../utils/utils'
import PromptBox from '../../../components/PromptBox'

import './style.css'

const content = {"entityMap":{},"blocks":[{"key":"637gr","text":"Initialized from content state.","type":"unstyled","depth":0,"inlineStyleRanges":[],"entityRanges":[],"data":{}}]};

@Form.create()
class Draft extends React.Component{
  state = {
    editorState: EditorState.createEmpty(),
    contentState: content,
    docFromState: '请选择',
    docCateState: '请选择',
  }

  onEditorStateChange = (editorState) => {
    this.setState({
      editorState,
    });
  };
  onContentStateChange =  (contentState) => {
    this.setState({
      contentState,
    });
  };

  uploadImageCallBack = (file)=>{
    return new Promise(
      (resolve, reject) => {
        function callback(error, response) {
          if (!error && (response.statusCode === 401 || response.statusCode === 403)) {
            message.error('系统错误');
            reject('系统错误');
          } else if (!error && response.statusCode === 200) {
            const jsonRes = JSON.parse(response.text);
            if (jsonRes.code === 0) {
              message.success(jsonRes.info);
              resolve(JSON.parse(jsonRes.info));
            } else {
              message.error(jsonRes.info);
              reject(jsonRes.info);
            }
          } else {
            message.error('网络错误');
            reject("网络错误");
          }
        }

        superagent.post(MEDIA_UPLOAD_URL)
        .field('name', isAuthenticated())
        .field('dtype', DTYPE_IMAGE)
        .field('file', file)
        .withCredentials()
        .set('Accept', 'application/json')
        //.set('Content-Type', 'multipart/form-data')
        .end(callback);
      }
    );
  };

  handleMenuDocFromClick = (e)=>{
    console.log(this)
    this.setState({
      docFromState: e.item.props.children
    })
  };

  handleMenuDocCateClick = (e)=>{
    this.setState({
      docCateState: e.item.props.children
    })
  };

  handleSubmitClick = (e)=>{
    const { editorState, docFromState, docCateState } = this.state;
    e.preventDefault()

    this.props.form.validateFields((err, values) => {
      if (!err) {
        if (values.draftTitle.length === 0) {
          this.props.form.setFields({
            draftTitle: {
              value: values.draftTitle,
              errors: [new Error('请输入文章标题')]
            }
          })
          return
        }

        this.uploadDraftAsync(values, docFromState, docCateState, draftToHtml(convertToRaw(editorState.getCurrentContent())));
      }
    })
  };

  handleCancelClick = ()=>{
    const {from} = this.props.location.state || {from: {pathname: '/'}}
    this.props.history.push(from)
  };

  uploadDraftAsync(values, docFromState, docCateState, editorContent) {
    function callback(error, response) {
      if (!error && (response.statusCode === 401 || response.statusCode === 403)) {
        message.error('系统错误');
      } else if (!error && response.statusCode === 200) {
        const jsonRes = JSON.parse(response.text);
        if (jsonRes.code === 0) {
          message.success(jsonRes.info);
        } else {
          message.error(jsonRes.info);
        }
      } else {
        message.error("网络错误");
      }
    }

    superagent.post(POSTER_UPLOAD_URL)
    .field('name', isAuthenticated())
    .field('title', values.draftTitle)
    .field('refer_url', values.draftReferUrl)
    .field('from', docFromState)
    .field('category', docCateState)
    .field('dtype', DTYPE_DRAFT)
    .field('content', editorContent)
    .withCredentials()
    .set('Accept', 'application/json')
    //.set('Content-Type', 'multipart/form-data')
    .end(callback);
  }

  render(){
    const cardContent = `您可以编辑个性化文案，图文并茂的方式展现您的作品。具体使用可以参考<a href="https://github.com/lawenliu/beego-demo">编辑器使用</a>`
    const { editorState, docFromState, docCateState } = this.state;
    const {getFieldDecorator, getFieldError} = this.props.form
    const menuDocFrom = (
      <Menu onClick={this.handleMenuDocFromClick}>
        <Menu.Item key="1">原创</Menu.Item>
        <Menu.Item key="2">转发</Menu.Item>
        <Menu.Item key="3">翻译</Menu.Item>
      </Menu>
    )
    const menuDocCate = (
      <Menu onClick={this.handleMenuDocCateClick}>
        <Menu.Item key="1">旅游</Menu.Item>
        <Menu.Item key="2">风景</Menu.Item>
        <Menu.Item key="3">游戏</Menu.Item>
        <Menu.Item key="4">影视</Menu.Item>
        <Menu.Item key="5">音乐</Menu.Item>
        <Menu.Item key="6">笑话</Menu.Item>
        <Menu.Item key="7">学习</Menu.Item>
        <Menu.Item key="8">人生</Menu.Item>
        <Menu.Item key="9">运动</Menu.Item>
      </Menu>
    )
    return (
      <div>
        <CustomBreadcrumb arr={['其它','文案编辑器']}/>
        <Form onSubmit={this.handleSubmitClick} name='draftForm'>
          <Form.Item help={getFieldError('draftTitle') && <PromptBox info={getFieldError('draftTitle')}
                                                                           width={calculateWidth(getFieldError('draftTitle'))}/>}>
            {getFieldDecorator('draftTitle', {
              validateFirst: true,
              rules: [
                {required: true, message: '文档标题不能为空'},
              ]
            })(
              <Input
                onFocus={() => this.setState({focusItem: 0})}
                onBlur={() => this.setState({focusItem: -1})}
                maxLength={50}
                placeholder='文章标题'/>
            )}
          </Form.Item>
          <Form.Item help={getFieldError('draftReferUrl') && <PromptBox info={getFieldError('draftReferUrl')}
                                                                           width={calculateWidth(getFieldError('draftReferUrl'))}/>}>
            {getFieldDecorator('draftReferUrl', {
              validateFirst: true,
              rules: [
                {required: false, message: '外链地址不对'},
              ]
            })(
              <Input
                onFocus={() => this.setState({focusItem: 0})}
                onBlur={() => this.setState({focusItem: -1})}
                maxLength={255}
                placeholder='外链地址'/>
            )}
          </Form.Item>
          <Card bordered={false} className='card-item'>
            <Editor
              editorState={editorState}
              onEditorStateChange={this.onEditorStateChange}
              onContentStateChange={this.onContentStateChange}
              wrapperClassName="wrapper-class"
              editorClassName="editor-class"
              toolbarClassName="toolbar-class"
              localization={{ locale: 'zh'}}
              toolbar={{
                image: { uploadCallback: this.uploadImageCallBack, alt: { present: true, mandatory: true }, previewImage: true },
              }}
            />
          </Card>
          <Form.Item>
            <Col>
              <label>文章来源：</label>
              <Dropdown overlay={menuDocFrom}><Button>{docFromState}<Icon type='down'/></Button></Dropdown>
            </Col>
            <Col>
              <label>文章类别：</label>
              <Dropdown overlay={menuDocCate}><Button>{docCateState}<Icon type='down'/></Button></Dropdown>
            </Col>
          </Form.Item>
          <Form.Item>
            <input className="ant-btn ant-btn-primary" type='submit' value="提交" />&emsp;
            <input className="ant-btn ant-btn-default" type="button" onClick={this.handleCancelClick} value="返回" />
          </Form.Item>
        </Form>
        <BackTop visibilityHeight={200} style={{right: 50}}/>
      </div>
    )
  }
}
export default Draft