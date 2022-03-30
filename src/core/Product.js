import React, { useState, useEffect } from 'react';
import Layout from './Layout';
import { read, listRelated } from './index';
import Card from './Card';

const Product = props => {

    const [product, setProduct] = useState({});
    const [related, setRelated] = useState([]);
    const [error, setError] = useState(false);

    const loadSingleProduct = productId => {
        read(productId).then(data => {
            if (data.error) {
                setError(data.error);
            } else {
                setProduct(data.Data);
            }
        });
    };

    useEffect(() => {
        const productId = props.match.params.id;
        loadSingleProduct(productId);
    }, [props]);

    return (
        <Layout title={product && product.name} description={`$ ${product.price} `} className="container-fluid">

            <div className="row">
                <div className='col-8'>
                    {
                        product && product.description &&
                        <div className='col-16 mb-3'> <Card product={product} showViewProductButton={false} machineByID={true} />  </div>
                    }
                </div>
            </div>

        </Layout>
    )
}

export default Product;