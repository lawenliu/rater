import React from 'react'
import superagent from 'superagent'
import { Icon, Badge, Dropdown, Menu, Modal, Upload } from 'antd'
import screenfull from 'screenfull'
import { inject, observer } from 'mobx-react'
import { Link, withRouter } from 'react-router-dom'
import { isAuthenticated } from '../../utils/Session'
import {MEDIA_UPLOAD_URL, POSTER_UPLOAD_URL, DTYPE_IMAGE, DTYPE_DRAFT} from '../../utils/Constants'

class CustomUpload extends React.Component {
  state = {
    isUploaded: false,
    imageUrl: '',
    loading: false,
  }

  customRequest = ({
    file
  }) => {
    const {action, dtype, handleChange} = this.props
    let self = this
    function callback(error, response) {
      if (!error && (response.statusCode === 401 || response.statusCode === 403)) {
        handleChange({status: 'error', message: '系统错误'});
      } else if (!error && response.statusCode === 200) {
        const jsonRes = JSON.parse(response.text);
        if (jsonRes.code === 0) {
          // Get this url from response in real world.
          const info = JSON.parse(jsonRes.info);
          self.setState({imageUrl: info.data.link})
          handleChange({status: 'done', imageUrl: info.data.link});
        } else {
          handleChange({status: 'error', message: jsonRes.info});
        }
      } else {
        handleChange({status: 'error', message: '网络错误'});
      }

      self.setState({loading: false, isUploaded: true})
    }

    superagent.post(action)
    .field('name', isAuthenticated())
    .field('dtype', dtype)
    .field('file', file)
    .withCredentials()
    .set('Accept', 'application/json')
    //.set('Content-Type', 'multipart/form-data')
    .end(callback);
    this.setState({loading: true});
    handleChange({status: 'uploading'})
  }

  render () {
    const {isUploaded, imageUrl, loading} = this.state
    const {showUploadedList, beforeUpload, action, dtype} = this.props

    return (
      <Upload
        name="avatar"
        showUploadList={showUploadedList}
        customRequest={this.customRequest}
        beforeUpload={beforeUpload}
        listType="picture-card">
        <img style={isUploaded ? styles.avatarVisible : styles.avatarHidden} src={imageUrl} alt=""/>
        <div style={isUploaded ? styles.avatarHidden : styles.avatarVisible}>
          <Icon type={loading ? 'loading' : 'plus'}/>
          <div className="ant-upload-text">Upload</div>
        </div>
      </Upload>
    )
  }
}

const styles = {
  avatarVisible: {
    display: 'block',
    width: '100px',
    height: '100px'
  },
  avatarHidden: {
    display: 'none',
    width: '100px',
    height: '100px'
  }
}

export default CustomUpload