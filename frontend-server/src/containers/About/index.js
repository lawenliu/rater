import React from 'react'
import CustomBreadcrumb from '../../components/CustomBreadcrumb/index'
import TypingCard from '../../components/TypingCard'

export default class About extends React.Component{
  render(){
    return (
      <div>
        <CustomBreadcrumb arr={['关于']}/>
        <TypingCard source={'让生活变成美好体验...'} title='关于' />
      </div>
    )
  }
}