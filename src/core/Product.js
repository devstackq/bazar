import React, { useState, useEffect } from 'react';
import Layout from './Layout';
import { read, listRelated } from './index';
import Card from './Card';

const Product = props => {

    const [product, setProduct] = useState({});
    // const [related, setRelated] = useState([]);
    const [error, setError] = useState(false);

    const loadSingleProduct = productId => {
        read(productId).then(data => {
            if (data.error) {
                setError(data.error);
                return
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

        product != null  ? (
        <Layout title={product && product.name} description={`$ ${product.price} `} className="container-fluid">
            <div className="row">
                <div className='col-12'>
                    {
                        product != null && product.description &&
                        <div className='col-12 p-2'> <Card product={product} showViewProductButton={false} machineByID={true} />  </div>
                    }
                </div>
            </div>

        </Layout>
        ): console.log('not found ')
    )
}

export default Product;