import React, { useState, useEffect } from 'react';
import Layout from './Layout';
import { getCategories, getFilteredProducts } from './index';
import Checkbox from './Checkbox';
import { prices } from './helpers';
import RadioBox from './RadioBox';
import Card from './Card';

const Shop = () => {

    const [clientFilters, setClientFilters] = useState({
        filters: { category: [], price: [] }
    });
    const [categories, setCategories] = useState([]);

    const [error, setError] = useState(false);
    const [limit] = useState(6);
    const [skip, setSkip] = useState(0);
    const [size, setSize] = useState(0);
    const [filteredResults, setFilteredResults] = useState([]);

    const init = () => {
        getCategories().then(data => {
            if (data.error) {
                setError(data.error);
            } else {
                setCategories(data);
            }
        });
    };
    //act2
    const loadFilteredResults = newFilterFromState => { // newFilter.filters category: [], price: [300,1000]
        // console.log(JSON.stringify(newFilterFromState));
        getFilteredProducts(skip, limit, newFilterFromState) //query API for back end, State - limit=6, skip =0, category: [], price: [300,1000]
            .then(data => { //act 7
                if (data.error) {
                    setError(data.error);
                } else {
                    setFilteredResults(data.data);
                    setSize(data.size);
                    setSkip(0);
                }
            });
    };

    const loadMoreProducts = () => {
        let toSkip = skip + limit;
        // console.log(newFilters);
        getFilteredProducts(toSkip, limit, clientFilters.filteres)
            .then(data => {
                if (data.error) {
                    setError(data.error);
                } else {
                    setFilteredResults([...filteredResults, ...data.data]);
                    setSize(data.size);
                    setSkip(toSkip);
                }
            });
    };

    const loadMoreButton = () => {
        return (
            size > 0 && size >= limit && (
                <button onClick={loadMoreProducts} className="btn btn-warning mb-5">Load more</button>
            )
        );
    };

    useEffect(() => {
        init();
        loadFilteredResults(skip, limit, clientFilters.filters);
    }, []);
    //act 1
    const handleFilters = (filters, filterBy) => { // filters = [id - radiobox 2]
        const newFilters = { ...clientFilters }; // if filter more than 1, load all filters
        newFilters.filters[filterBy] = filters; // RadioBOx component - handleChange(input value),[id =2] user choice
        if (filterBy === "price") {  // filterBy === price
            let priceValues = handlePrice(filters); //call handlePrice[]  => return arrat: 300 -1000
            newFilters.filters[filterBy] = priceValues; // newFilters .filters[300, 1000]  set array = [300, 1000]
        }
        loadFilteredResults(clientFilters.filters); // loadFilteredResults(clientFilters.filters => newFilter.filters[300,1000] - query API for back end
        setClientFilters(newFilters); //  set state newFilter
    };

    const handlePrice = value => {
        const data = prices; // data all filters array[obj]
        let array = [];

        for (let key in data) { // iteration all data
            if (data[key]._id === parseInt(value)) { // if data id == arg(value) id 2(filter) === prices id 2(hardcode)
                array = data[key].array; // array set    array = array: [300, 1000]
            }
        }
        return array; //array :[300, 1000]
    };

    // 1 RadioBox - user choice sortBy - value - prices[id2] from 300 to 1000$ - UI shop
    //2 handleFilters(1 - filters, ''price'); 



    return (
        <Layout title='Shop page' description="search and find your product">

            <div className="row">
                <div className='col-3'>
                    <h4>Filter by category</h4>
                    <ul>
                        <Checkbox
                            categories={categories}
                            handleFilters={filters => handleFilters(filters, 'category')}
                        />
                    </ul>

                    <h4>Filter by price range</h4>
                    <div>
                        <RadioBox
                            prices={prices}
                            handleFilters={filters => handleFilters(filters, 'price')}
                        />
                    </div>
                </div>

                <div className='col-9'><h2 className='mb-4'>Products</h2>
                    <div className="row">
                        {filteredResults.map((product, i) => ( //act 8

                            <div key={i} className='col-4 mb-3'>
                                <Card product={product} />
                            </div>
                        ))}
                    </div>
                    <hr />
                    {loadMoreButton()}
                </div>
            </div>
        </Layout>
    );
};
export default Shop;

//{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJfaWQiOiI1ZDc2ODMzZmE1MTc1MjMzOTQ3YTE2ODYiLCJpYXQiOjE1NjgwNDgwMjZ9.6mXb_DgwRw2QDlIt6KMT5uc_yxEF-1nRDfP-9PlvN8Q","user":{"_id":"5d76833fa5175233947a1686","email":"admin2@mail.com","name":"Admin","role":1}}