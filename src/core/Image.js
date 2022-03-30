import React from 'react';
import { API } from '../config';

const ShowImage = ({ item, url }) => (
//todo like - carousel, more 1 image

<div className="product-img">
        {item.main_image != "" ? (
        <img src={`${API}/images/${item.main_image}`} alt={item.name} className="card-img mb-2"
            style={{ maxHeight: '90%', maxWidth: '100%' }} />
        ):
        <img src={`${API}/images/noimage.png`} alt={item.name} className="card-img mb-2"
            style={{ maxHeight: '90%', maxWidth: '100%' }} />
        }
    </div>
)
export default ShowImage;