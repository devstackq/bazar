import React, { Fragment } from 'react';
import { Link, withRouter } from 'react-router-dom';
import { signout, isAuthenticated } from '../auth/index';
import { itemTotal } from './helpers';


const isActive = (history, path) => {
    if (history.location.pathname === path) {
        return { color: 'green' };
    } else {
        return { color: '#fff' };
    }
};

const Navbar = ({ history }) => (

    <div>

        <ul className="nav nav-tabs bg-warning">
            <li className="nav-item">
                <Link className="nav-link" style={isActive(history, '/')} to="/"> Home</Link>
            </li>
            <li className="nav-item">
                <Link className="nav-link" style={isActive(history, '/shop')} to="/shop"> Shop</Link>
            </li>
            <li className="nav-item">
                <Link
                    className="nav-link"
                    style={isActive(history, "/cart")}
                    to="/cart"
                >
                    Cart{" "}
                    <sup>
                        <small className="cart-badge">{itemTotal()}</small>
                    </sup>
                </Link>
            </li>

            {isAuthenticated() && isAuthenticated().user.role === 0 && (
                <>
         
                    <li className="nav-item">
                        <Link
                            className="nav-link"
                            style={isActive(history, "/user/dashboard")}
                            to="/user/dashboard"
                        >
                            Dashboard
                    </Link>
                    </li>
        

                </>
            )}

            {isAuthenticated() && isAuthenticated().user.role === 1 && (
                <li className="nav-item">
                    <Link
                        className="nav-link"
                        style={isActive(history, "/admin/dashboard")}
                        to="/admin/dashboard"
                    >
                        ADmin
                    </Link>
                </li>
            )}




            {!isAuthenticated() && (
                
                    <div className='navbar-right'>
                    <li className="nav-item">
                        <Link className="nav-link" style={isActive(history, '/signin')} to="/signin"> Signin</Link>
                    </li>
                    <li className="nav-item">
                        <Link className="nav-link" style={isActive(history, '/signup')} to="/signup"> Signup</Link>
                    </li>
                    </div>
                
            )}

            {isAuthenticated() && (
                <li className="nav-item navbar-right">
                    <span className="nav-link"
                        style={{ cursor: 'pointer', color: '#fff' }}
                        onClick={() => signout(() => {
                            history.push('/');
                        })
                        }
                    > Sign out
                    </span>
                </li>
            )}
        </ul>

    </div>
);

export default withRouter(Navbar);