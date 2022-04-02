import { API } from '../config';
import queryString from "query-string";

// return fetch(`${API}/filter?${key}=${val}`, {

export const getProducts = (key, val) => {
    return fetch(`${API}/machine`, {
        method: 'GET'
    })
        .then(response => {
            return response.json()
        })
        .catch(err => { console.log(err) });
};

export const getCategories = () => {
    return fetch(`${API}/category`, {
        method: 'GET'
    })
        .then(response => {
            return response.json();
        })
        .catch(err => { console.log(err) });
};

export const getStates = () => {
    return fetch(`${API}/state`, {
        method: 'GET'
    })
        .then(response => {
            return response.json();
        })
        .catch(err => { console.log(err) });
};

export const read = productId => {
    console.log(productId, 'id', API)
    return fetch(`${API}/machine/${productId}`, {
        method: 'GET'
    })
        .then(response => {
            return response.json();
        })
        .catch(err => { console.log(err) });
};

//api for backe end
//act 3
export const getFilteredProducts = (skip, limit, filters = {}) => { //request body, data from UI user choice

    const data = {
        limit, skip, filters
    };
    // console.log(JSON.stringify( data));
    return fetch(`${API}/products/by/search`, {
        method: 'POST',
        headers: {
            Accept: 'application/json',
            "Content-Type": 'application/json'
        },
        body: JSON.stringify(data)   // data send back end
    })
        .then(response => {                //
            return response.json();
        })
        .catch(err => { console.log(err) });
};

//search request backend - search - cola, category Node
export const list = params => {
    const query = queryString.stringify(params); // query string
    console.log("query", query);
    return fetch(`${API}/products/search?${query}`, {
        method: 'GET'
    })
        .then(response => {
            return response.json();
        })
        .catch(err => console.log(err));
};



export const listRelated = productId => {
    return fetch(`${API}/products/related/${productId}`, {
        method: 'GET'
    })
        .then(response => {
            return response.json();
        })
        .catch(err => { console.log(err) });
};


export const getBraintreeClientToken = (userId, token) => {
    return fetch(`${API}/braintree/getToken/${userId}`, {
        method: 'GET',
        headers: {
            Accept: 'application/json',
            "Content-Type": 'application/json',
            Authorization: `Bearer ${token}`
        },
    })
        .then(response => {
            return response.json();
        })
        .catch(err => { console.log(err) });
};

export const proccessPayment = (userId, token, paymentData) => {
    return fetch(`${API}/braintree/payment/${userId}`, {
        method: 'POST',
        headers: {
            Accept: 'application/json',
            "Content-Type": 'application/json',
            Authorization: `Bearer ${token}`
        },
        body: JSON.stringify(paymentData)
    })
        .then(response => {
            return response.json();
        })
        .catch(err => console.log(err));
};


export const createOrder = (userId, token, orderData) => {
    return fetch(`${API}/order/create/${userId}`, {
        method: 'POST',
        headers: {
            Accept: 'application/json',
            "Content-Type": 'application/json',
            Authorization: `Bearer ${token}`
        },
        body: JSON.stringify({ order: orderData })
    })
        .then(response => {
            return response.json();
        })
        .catch(err => console.log(err));
};

