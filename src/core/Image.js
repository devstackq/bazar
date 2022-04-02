import React from 'react';
import { API } from '../config';

const ShowImage = ({ item, url }) => (

<div className="product-img">
        {
        item.main_image != "" ? (
        <img src={`${API}${item.main_image}`} alt={item.name} className="card-img mb-2"
            style={{ maxHeight: '100px', maxWidth: '200px' }} />
        ):
        <img src={`${API}/images/noimage.png`} alt={item.name} className="card-img mb-2"
            style={{ maxHeight: '100px', maxWidth: '200px' }} />
        }
    </div>
)
//100%
export default ShowImage;