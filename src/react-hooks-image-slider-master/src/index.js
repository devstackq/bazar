import React from 'react'
import ReactDOM from 'react-dom'
import Slider from './components/slider'
import images from './images'

ReactDOM.render(<Slider slides={images} />, document.querySelector('.main'))
