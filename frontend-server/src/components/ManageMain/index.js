import React from 'react'
import { withRouter, Switch, Redirect } from 'react-router-dom'
import LoadableComponent from '../../utils/LoadableComponent'
import PrivateRoute from '../PrivateRoute'

// Profile Manage
const Profile = LoadableComponent(()=>import('../../containers/Manage/Profile/index'))
const ProfileChpsw = LoadableComponent(()=>import('../../containers/Manage/Profile/Chpsw/index'))

// Resource Manage
const Content = LoadableComponent(()=>import('../../containers/Manage/Content/index'))
const ContentImage = LoadableComponent(()=>import('../../containers/Manage/Content/Image/index'))
const ContentVideo = LoadableComponent(()=>import('../../containers/Manage/Content/Video/index'))
const ContentDraft = LoadableComponent(()=>import('../../containers/Manage/Content/Draft/index'))

//关于
const Help = LoadableComponent(()=>import('../../containers/Help/index'))
const About = LoadableComponent(()=>import('../../containers/About/index'))

@withRouter
class ContentMain extends React.Component {
  render () {
    return (
      <div style={{padding: 16, position: 'relative'}}>
        <Switch>
          <PrivateRoute exact path='/manage/profile' component={Profile}/>
          <PrivateRoute exact path='/manage/profile/chpwd' component={ProfileChpsw}/>

          <PrivateRoute exact path='/manage/content' component={Content}/>
          <PrivateRoute exact path='/manage/content/image' component={ContentImage}/>
          <PrivateRoute exact path='/manage/content/video' component={ContentVideo}/>
          <PrivateRoute exact path='/manage/content/draft' component={ContentDraft}/>

          <PrivateRoute exact path='/about' component={About}/>
          <PrivateRoute exact path='/help' component={Help}/>

          <Redirect exact from='/' to='/home'/>
        </Switch>
      </div>
    )
  }
}

export default ContentMain