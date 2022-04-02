import React, { useState, useEffect } from 'react';
import { Link, Redirect } from "react-router-dom";
import ShowImage from './Image';
import Slider from './slider/Slider'
import AliceCarousel from 'react-alice-carousel';
import "react-alice-carousel/lib/alice-carousel.css";
import '../style.css';
import moment from 'moment';
import { addItems, updateItem, removeItem } from './helpers';
import { API } from '../config';


const Card = ({ product, showViewProductButton = true, showAddToCartButton = true, cartUpdate = false, showRemoveProductButton = false, setRun = f => f, run = undefined, machineByID = false }) => {

    const [redirect, setRedirect] = useState(false);
    const [count, setCount] = useState(product.count);
    const [images, setImages] = useState([])

    const [imagesElem, setImagesElems] = useState([])


    const showViewButton = showViewProductButton => {
        return (
            showViewProductButton && (
                <Link to={`/machine/${product.id}`}>
                    <button className="btn btn-outline-primary mt-2 mb-2">
                        View Product
                    </button>
                </Link>
            )
        );
    };

    //todo whatsapp
    const addToCart = () => {
        alert(`call by phone : redirect whatsapp ${product.creator.phone}`)
    };

    const shouldRedirect = redirect => {
        if (redirect) {
            return <Redirect to='/cart' />;
        }
    };

    const showAddToCart = showAddToCartButton => {
        return (
            showAddToCartButton && (
                <button onClick={addToCart}
                    className="btn btn-outline-warning mt-2 mb-2 ml-2">
                    Write to Message
                </button>
            )
        );
    };

    const showRemoveButton = showRemoveProductButton => {
        return (
            showRemoveProductButton && (
                <button
                    onClick={() => {
                        removeItem(product._id)
                        setRun(!run)
                    }}

                    className="btn btn-outline-danger mt-2 mb-2">
                    Remove product
                </button>
            )
        )
    };

    const showCartUpdateOptions = cartUpdate => {
        return (
            cartUpdate && (
                <div>
                    <div className="input-group mb-3">
                        <div className="input-group-prepend">
                            <span className="input-group-text">
                                Adjust Quantity
                            </span>
                        </div>
                        <input
                            type="number"
                            className="form-control"
                            value={count}
                            onChange={handleChange(product._id)}
                        />
                    </div>
                </div>
            )
        );
    };
    const handleChange = productId => event => {
        setRun(!run); // run useEffect in parent Cart
        setCount(event.target.value < 1 ? 1 : event.target.value);
        if (event.target.value >= 1) {
            updateItem(productId, event.target.value);
        }
    };

    const showStock = quantity => {
        return quantity > 0 ? (
            <span className="badge badge-primary badge-pill">In Stock</span>
        ) : (
            <span className="badge badge-danger badge-pill">Out of Stock</span>
        );
    };

    // const readMore= product => {
    //     return (
    //     <p className="lead mt-2">{product.substring(0, 400)}</p>
    //     );
    // };

    const loadImages = paths => {

        let seq = []
        if (paths != null) {
            for (let i = 0; i < paths.length; i++) {
                seq.push(`${API}${paths[i]}`)
            }
            setImages(seq)
        }
        return
    }

    useEffect(() => {
        // loadImages(product.images)
        renderImages(product.images)
    }, []);


    const renderImages = (images) => {

        let temp = []
        if (images != null) {
        for (let i = 0; i < images.length; i++) {
            let img = document.createElement('img')
            img.src = `${API}${images[i]}`
            temp.push(img)
        }
        setImagesElems(temp)
    }
    }

    return (
        <div className="card">
            <div className="card-header name">{product.name}</div>
            <div className="card-body">
                {shouldRedirect(redirect)}

                {!machineByID ? <ShowImage item={product} url='machine' /> : null}
                
                {machineByID && product.images.length > 0 ? (
                    <div>
                        <AliceCarousel autoPlay autoPlayInterval="3000">
                        (
                            <div>{
                                //render 1 time
                                product.images.forEach(element => {
                                    let img = document.createElement('img')
                                    img.src = `${API}${element}`
                                })
                            }</div>
                            )
                        </AliceCarousel>
                    </div>
                    ): null}
                {/* {machineByID && images.length > 0 ? <Slider slides={images} />    : null} */}

                <p className="lead mt-2">Title: {product.title}</p>
                <p className="lead mt-2">Description: {product.description.substring(0, 150)}</p>
                <p className="lead mt-2"> VIN: {product.vin}</p>
                <p className="lead mt-2">Odometer: {product.odometer}</p>

                {machineByID ? (
                    <div>
                        <p className='black-10 text-danger'> transmissin: {product.trans.Name}</p>
                        <p className='black-10 text-danger'> brand: {product.brand.Name}</p>
                        <p className='black-10 text-danger'> model: {product.model.Name}</p>

                    </div>
                ) : null}

                <p className='black-10 text-danger'>${product.price}</p>
                {/* <p className='black-9'>Category: {product.category && product.category.name}</p> */}
                {/* <p className='black-8'> Added on {moment(product.createdAt).fromNow()}</p> */}

                {/* {showStock(product.quantity)} */}

                <br />
                {showViewButton(showViewProductButton)}
                {showAddToCart(showAddToCartButton)}
                {showRemoveButton(showRemoveProductButton)}
                {showCartUpdateOptions(cartUpdate)}

            </div>
        </div>
    );
};

export default Card;
