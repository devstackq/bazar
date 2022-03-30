import React, { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import { isAuthenticated } from '../auth/index';
import { emptyCart } from './helpers';
import { getBraintreeClientToken, proccessPayment, createOrder } from './index';
import DropIn from 'braintree-web-drop-in-react';

const Checkout = ({ products }) => {

    const [data, setData] = useState({
        success: false,
        loading: false,
        clientToken: null,
        error: '',
        instance: {},
        address: ''
    });

    const userId = isAuthenticated() && isAuthenticated().user._id;
    const token = isAuthenticated() && isAuthenticated().token;


    const getToken = (userId, token) => {
        getBraintreeClientToken(userId, token).then(data => {
            if (data.error) {
                setData({ ...data, error: data.error });
            } else {
                setData({ clientToken: data.clientToken });
            }
        });
    };

    useEffect(() => {
        getToken(userId, token);
    }, []);

    const handleAddress = event => {
        setData({ ...data, address: event.target.value });
    };

    const getTotal = () => {
        return products.reduce((acc, nextValue) => {
            return acc + nextValue.count * nextValue.price;
        }, 0);
    };

    const showCheckout = () => {
        return isAuthenticated() ? (
            <div> {showDropIn()}} </div>
        ) : (
                <Link to='/signin'>
                    <button className='btn btn-primary'> Sign in to checkout </button>
                </Link>
            );
    };

    const deliveryAddress = data.address;

    const byuProducts = () => {
        setData({ loading: true });
        //send the nonce to tourserver
        //nonce  = data..instance.requestPaymentMethod
        let nonce;
        let getNonce = data.instance
            .requestPaymentMethod()
            .then(data => {
                // console.log(data);
                nonce = data.nonce;
                   // once you have nonce (card type, card number) send nonce as 'paymentMethodNonce'
                // and also total to be charged
                // console.log(
                //     "send nonce and total to process: ",
                //     nonce,
                //     getTotal(products)
                // );

                const paymentData = {
                    paymentMethodNonce: nonce,
                    amount: getTotal(products)
                };
                proccessPayment(userId, token, paymentData)
                    .then(response => {
                        console.log(response);
                         // empty cart
                        // create order
                        const orderData = {
                            products: products,
                            transaction_id: response.transaction.id,
                            amount: response.transaction.amount,
                            address: deliveryAddress
                        }
                        createOrder(userId, token, orderData)
                            .then(response => {
                                emptyCart(() => {
                                    console.log("payment success and empty cart");
                                    setData({ loading: false, success: true });
                                });
                            })
                            .catch(error => {
                                console.log(error);
                                setData({ loading: false });
                            });
                    })
                    .catch(error => {
                        console.log(error);
                        setData({ loading: false });
                    });
            })
            .catch(error => {
                // console.log("dropin error: ", error);
                setData({ ...data, error: error.message });
            });
    };


    const showError = error => (
        <div
            className="alert alert-danger"
            style={{ display: error ? "" : "none" }}
        >
            {error}
        </div>
    );


    const showSuccess = success => (
        <div
            className="alert alert-info"
            style={{ display: success ? "" : "none" }}
        >
            Thanl u! you payment was succesfull!
        </div>
    );


    const showLoading = loading =>
        loading && <h2 className="text-danger">Loading...</h2>;  

    const showDropIn = () => (
        <div onBlur={() => setData({ ...data, error: "" })}>

            {data.clientToken !== null && products.length > 0 ? (
                <div>
                    <div className='gorm-group mb-3'>
                        <label className='text-muted'> Delivery address: </label>
                        <textarea
                            onChange={handleAddress}
                            className='form-control'
                            value={data.address}
                            placeholder='type delivery address'
                        />
                    </div>
                    <DropIn options={{
                        authorization: data.clientToken

                    }} onInstance={instance => (data.instance = instance)} />
                    <button onClick={byuProducts} className='btn btn-success btn-block'>Pay </button>
                </div>
            ) : null}
        </div>
    );


    return (
        <div>
            <h2>Total: ${getTotal()}</h2>
            {showLoading(data.loading)}
            {showSuccess(data.success)}
            {showError(data.error)}
            {showCheckout()}

        </div>
    );
};

export default Checkout;