import React, { useState } from 'react';
import Layout from '../core/Layout';
import { isAuthenticated } from '../auth/index';
import { Link } from 'react-router-dom';
import { createCategory } from './index';

const AddCategory = () => {
    const [name, setName] = useState('');
    const [error, setError] = useState(false);
    const [success, setSuccess] = useState(false);

    //descture user and token take  from LS
    const { user, token } = isAuthenticated();

    const handleChange = (e) => {
        setError('')
        setName(e.target.value); //value user input 1(Kids)
    };

    const handleSubmit = (e) => {
        e.preventDefault();
        setError('')
        setSuccess(false)
        //make request from to api
        createCategory(user._id, token, { name })  //, user_id - take - isAuthenticated/ => когда signin, 
            //authenticate function , setItem (jwt token) local storage, isAuthenticated - user, token данные хранятся (jwt - token, user all data)
            // categeory name - handleChange(women) - данные идут в функцию createCategory ()
            .then(data => {
                if (data.error) {
                    setError(true);
                } else {
                    setError('');
                    setSuccess(true);
                }
            })
    };

    const newCategory = () => (
        <form onSubmit={handleSubmit}>
            <div className="form-group">
                <label className="text-muted">
                    Name
                </label>
                <input type="text" className="form-control" onChange={handleChange} value={name} autoFocus required />

            </div>
            <button className="btn btn-outline-primary ">create category</button>
        </form>
    );

    const showSuccess = () => {
        if (success) {
            return <h3 className="text-success">category  {name} is created </h3>
        }
    };
    const showError = () => {
        if (error) {
            return <h3 className="text-danger">category busy,  should be unique name </h3>
        }
    };

    const goBack = () => (
        <div className="mt-5">
            <Link to='/admin/dashboard' className="text-warning"> back to admin </Link>
        </div>
    );

    return (
        <Layout
            title="Add new category"
            description={`have nice day ${user.name}!`}
        >
            <div className="row">

                <div className="col-md-8 offset-md-2">
                    {showError()}
                    {showSuccess()}
                    {newCategory()}
                    {goBack()}
                </div>
            </div>
        </Layout>
    )
}

export default AddCategory;