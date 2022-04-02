import React, { useState, useEffect } from 'react';
import Layout from './Layout';
import { getProducts } from './index';
import Card from './Card';
import Search from './Search';

const Home = () => {

    const [productsByFilter, setProductByFilter] = useState([]);
    const [productsBySell, setProductBySell] = useState([]);

    const [productsByArrival, setProductByArrival] = useState({});
    const [error, setError] = useState(false);

    //loadMachinesByDate
    //loadFilterComponents
    
    // const loadProductsBySell = () => {
    //     getProducts('sold')
    //         .then(data => {
    //             if (data.error) {
    //                 setError(data.error);
    //             } else {
    //                 setProductBySell(data)
    //             }
    //         })
    // };

    const loadProductsByArrival = () => {
        // getProducts('sort_created_at', 'asc')
        getProducts()
            .then(data =>  {
                if (data.error) {
                    setError(data.error);
                } else {
                    setProductByArrival(data)
                }
            })
    };


    useEffect(() => {
        loadProductsByArrival();
        // loadProductsBySell();
    }, []);

    return (
        <Layout title='Home page' description="Stay tuned, arrivals brand's clothes" className="container-fluid">
            <Search />
            <h2 className="mb-4">Last Arrivals</h2>
            <div className="row">
            {   productsByArrival != undefined && productsByArrival.Data != undefined  ?  (
                productsByArrival.Data.map((product, i) => (
                    <div key={i} className='col-3 mb-3'>
                        <Card product={product} machineByID = {false} />
                    </div>
                ))
               )  : null
            }
            </div>
{/* 
            <h2 className="mb-4">Best Sellers</h2>
            <div className="row">
                {productsBySell.map((product, i) => (
                    <div key={i} className='col-3 mb-3'>
                        <Card product={product}  machineByID = {false}/>
                    </div>
                ))}
            </div> */}
        </Layout>
    );
};
export default Home;