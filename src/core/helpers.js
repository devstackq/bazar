import { API } from '../config';

export const addItems = (item, next) => {
    let cart = [];
    if (typeof window !== 'undefined') {
        if (localStorage.getItem('cart')) {
            cart = JSON.parse(localStorage.getItem('cart'));
        }
        cart.push({
            ...item,
            count: 1
        });

        cart = Array.from(new Set(cart.map(product => product._id))).map(id => {
            return cart.find(product => product._id === id);
        });
        localStorage.setItem('cart', JSON.stringify(cart));
        next();
    }
};

export const itemTotal = () => {
    if (typeof window !== 'undefined') {
        if (localStorage.getItem('cart')) {
            return JSON.parse(localStorage.getItem('cart')).length;
        }
    }
    return 0;
};

export const getCart = () => {
    if (typeof window !== "undefined") {
        if (localStorage.getItem("cart")) {
            return JSON.parse(localStorage.getItem("cart"));
        }
    }
    return [];
};


export const updateItem = (productId, count) => {
    let cart = [];
    if (typeof window !== 'undefined') {
        if (localStorage.getItem('cart')) {
            cart = JSON.parse(localStorage.getItem('cart'))
        }
        cart.map((product, i) => {
            if (product._id === productId) {
                cart[i].count = count
            }
        }

        )
        localStorage.setItem('cart', JSON.stringify(cart))
    }
};


export const removeItem = productId => {
    let cart = [];
    if (typeof window !== 'undefined') {
        if (localStorage.getItem('cart')) {
            cart = JSON.parse(localStorage.getItem('cart'))
        }
        cart.map((product, i) => {
            if (product._id === productId) {
                cart.splice(i, 1)
            }
        }

        )
        localStorage.setItem('cart', JSON.stringify(cart))
    }
    return cart;
};


export const prices = [
    {
        _id: 0,
        name: 'Any',
        array: []
    },
    {
        _id: 1,
        name: '$0 to 300',
        array: [0, 300]
    },
    {
        _id: 2,
        name: '$300 to $1000 ',
        array: [300, 1000]
    },
    {
        _id: 3,
        name: '$1000 to $2000 ',
        array: [1000, 2000]
    },
    {
        _id: 4,
        name: 'more than $2000',
        array: [2000, 100000]
    }
];


export const emptyCart = next => {
    if(typeof window !== 'undefined') {
        localStorage.removeItem('cart');
        next();
    }
};

export const getMachineByID = id  => {

    return fetch(`${API}/machine/${id}`, {
        method: 'GET'
    })
        .then(response => {
            return response.json();
        })
        .catch(err => { console.log(err) });
} 