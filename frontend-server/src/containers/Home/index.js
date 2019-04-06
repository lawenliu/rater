import React from 'react'
import CustomCarousel from '../../components/CustomCarousel/index'
import Gallery from '../../components/Gallery/index'
import './style.css'

class Home extends React.Component {
  render() {
    return (
      <div>
        <CustomCarousel arrows effect='fade' className='size' />
        <Gallery />
      </div>
    )
  }
}

export default Home