import React, { useState, useEffect } from 'react';
import Layout from '../core/Layout';
import { isAuthenticated } from '../auth/index';
import { listOrder, getStatusValues } from './index';
import moment from 'moment';

const Orders = () => {

    const [orders, setOrders] = useState([]);
    const [statusValues, setStatus] = useState([]);

    const { user, token } = isAuthenticated();


    const loadOrders = () => {
        listOrder(user._id, token)
            .then(data => {
                if (data.error) {
                    console.log(data.error)
                } else {
                    setOrders(data)
                }
            }
            );
    };


    const loadStatusValues = () => {
        getStatusValues(user._id, token)
            .then(data => {
                if (data.error) {
                    console.log(data.error)
                } else {
                    setStatus(data)
                }
            }
            );
    }

    useEffect(() => {
        loadOrders();
        loadStatusValues();
    }, [])


    const showOrdersLength = () => {
        if (orders.length > 0) {
            return (
                <h1 className='text-danger display-2'
                >Total orders: {orders.length}</h1>)
        } else {
            return <h1 className='text-danger'> No orders</h1>
        }
    };

    const handleStatusChange = (event, orderId) => {
        console.log('update order status');
    };

    
    const showInput = (key, value) => (
        <div className='input-group mb-2 mr-sm-2'>
            <div className='input-group-prepend'>
                <div className='input-group-text'>
                    {key}
                </div>

            </div>
            <input type='text' value={value} className='form-control' readOnly />
        </div>
    );

    const showStatus = (order) => (
        <div className='form-group'>
            <h3 className='mark mb-4'>Status: {order.status}</h3>
            <select
                className='form-control'
                onChange={(event) => handleStatusChange(event, order._id)}>
                <option>
                    Update status
                </option>
                {statusValues.map((item, index) => (
                    <option key={index}
                        value={item} >
                            {item}
                    </option>
                ))}
            </select>
        </div>
    );


    return (
        <Layout
            title="Orders"
            description={`have nice day ${user.name}!`}
        >
            <div className="row">

                <div className="col-md-8 offset-md-2">
                    {showOrdersLength()}
                    {orders.map((order, index) => {
                        return (
                            <div className='mt-5' key={index}>
                                <h2 className='mb-5'>
                                    <span className='bg-primary'> Order ID {order._id}</span>

                                </h2>
                                <ul className='list-group mb-2'>
                                    <li className='list-group-item'>
                                        Status: {showStatus(order)}
                                    </li>
                                    <li className='list-group-item'>
                                        Transaction ID: {order.transaction_id}
                                    </li>
                                    <li className='list-group-item'>
                                        Amount: ${order.amount}
                                    </li>
                                    <li className='list-group-item'>
                                        Ordered by: {order.user.name}
                                    </li>
                                    <li className='list-group-item'>
                                        Ordered on: {moment(order.createdAt).fromNow()}
                                    </li>
                                    <li className='list-group-item'>
                                        Delivery address: {order.address}
                                    </li>

                                </ul>
                                <h3 className='mt-4 mb-4 font-italic'>
                                    Total products in the order: {order.products.length}
                                </h3>

                                {order.products.map((product, index) => (
                                    <div className='mb-4' key={index} style={{
                                        padding: '20px',
                                        border: '1px solid indigo'
                                    }}>
                                        {showInput('Product name', product.name)}
                                        {showInput('Product price', product.price)}
                                        {showInput('Product total', product.count)}
                                        {showInput('Product ID', product._id)}
                                    </div>
                                ))}
                            </div>
                        )
                    })}
                </div>
            </div>
        </Layout>
    )
}


export default Orders;