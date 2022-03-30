import React, { useState } from 'react';
import Layout from '../core/Layout';
import { Redirect } from 'react-router-dom';
import { signup, authenticate, signin, isAuthenticated } from '../auth/index';


const Signup = () => {

    const [values, setValues] = useState({
        name: '',
        email: '',
        password: '',
        error: '',
        success: false,
        redirectToRefer: false
    });

    const { name, email, password, success, error, redirectToRefer } = values;
    const { user } = isAuthenticated();

    const handleChange = name => event => {
        setValues({ ...values, error: false, [name]: event.target.value })
    }


    const handleSubmit = event => {
        event.preventDefault();
        setValues({ ...values, error: false })
        signup({ name, email, password })
            .then(data => {
                if (data.error) {
                    setValues({ ...values, error: data.error, success: false })
                } else {
                    setValues({
                        ...values,
                        name: '',
                        email: '',
                        password: '',
                        error: '',
                        success: true
                    })
                    signin({ email, password })
                        .then(data => {
                            if (data.error) {
                                setValues({ ...values, error: data.error });
                            } else {
                                authenticate(data, () => {
                                    setValues({
                                        ...values,
                                        redirectToReferrer: true
                                    });

                                });

                            }

                        });
                }
            });
    };

    const signUpForm = () => (
        <form>
            <div className="form-group">
                <label className="text-muted">Name</label>
                <input onChange={handleChange('name')} type='text' className="form-control" value={name} required />
            </div>

            <div className="form-group">
                <label className="text-muted">Email</label>
                <input onChange={handleChange('email')} type='email' className="form-control" value={email} required />
            </div>
            <div className="form-group">
                <label className="text-muted">Password</label>
                <input onChange={handleChange('password')} type='password' className="form-control" value={password} required />
            </div>
            <button onClick={handleSubmit} type='submit' className="btn btn-primary">Register</button>

        </form>
    );

    const showError = () => (
        <div className="alert alert-danger" style={{ display: error ? '' : 'none' }} > {error}</div>
    )

    const redirectUser = () => {
        if (redirectToRefer) {
            if (user && user.role === 1) {
                return <Redirect to="/admin/dashboard" />;
            } else {
                return <Redirect to="/user/dashboard" />;
            }
        }
        if (isAuthenticated()) {
            return <Redirect to="/" />;
        }
    };


    const showSuccess = () => (
        <div className="alert alert-info" style={{ display: success ? '' : 'none' }} > New account created</div>
    )

    return (
        <Layout title='sign up' description="signiup bla bla" className='container col-md-8 offset-md-2'>
            {showSuccess()}
            {redirectUser()}
            {showError()}
            {signUpForm()}

        </Layout>
    )

}

export default Signup;