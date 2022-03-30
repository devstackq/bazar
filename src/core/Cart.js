import React, { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import Layout from './Layout';
import { getCart } from './helpers';
import Card from './Card';
import Checkout from './Checkout'


const Cart = () => {

    const [items, setItems] = useState([]);
    const [run, setRun] = useState(false);

    useEffect(() => {
        setItems(getCart());
    }, [run]);

;


    const showItems = items => {
        return (
            <div>
                <h2> You cart has {`${items.length}`} items</h2>
                <hr />
                {items.map((product, i) => (
                    <Card key={i}
                        product={product}
                        showAddToCartButton={false}
                        cartUpdate={true}
                        showRemoveProductButton={true}
                        setRun={setRun}
                        run={run}
                       
                    />
                ))}
            </div>
        );
    };


    const noItemMessage = () => (
        <h2> You cart is empty <br /> <Link to='/shop'>Continue shopping </Link>  </h2>
    );

    return (
        <Layout title='Shopping cart' description="Buy and enjoy" className="container-fluid">
<>
            <div className="row">
                <div className="col-6">
                    {items.length > 0 ? showItems(items) : noItemMessage()}
                </div>

                <div className="col-6">
                    <h2 className="mb-4">Your cart summary</h2>
                    <hr />
                    <Checkout products={items} />
                </div>
            </div>j
            </>j
        </Layout>
    );

};
export default Cart;