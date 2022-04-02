// import React, { Fragment } from 'react';
import { Link, withRouter } from 'react-router-dom';
import { signout, isAuthenticated } from '../auth/index';
import { itemTotal } from './helpers';
import React, { useState, useEffect } from 'react';


const isActive = (history, path) => {
    if (history.location.pathname === path) {
        return { color: 'green' };
    } else {
        return { color: '#fff' };
    }
};

const Navbar = ({ history }) => {

    const [modalState, setModalState] = useState(false)
    
    const changeState = ()=> {
        console.log(modalState, 'change')
        setModalState(!modalState)
    }

    return (

        <div>

            <ul className="nav nav-tabs bg-warning">
                <li className="nav-item">
                    <Link className="nav-link" style={isActive(history, '/')} to="/"> Home</Link>
                </li>

                <li className="nav-item">
                    <Link className="nav-link" style={isActive(history, '/about')} to="/about"> About us</Link>
                </li>

                {/* profile page ? */}
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
                            style={isActive(history, "/profile")}
                            to="/profile"
                        >
                            Profile
                        </Link>
                    </li>
                )}

                {!isAuthenticated() &&  (
                    <div className='navbar-center'>
                        <li className="nav-item">
                            <button onClick={changeState} > To Sale !</button>
                        </li>
                    </div>
                )}

                { !isAuthenticated() &&  modalState ? (

                    <div className='navbar-right'>
                        <li className="nav-item">
                            <Link className="nav-link" style={isActive(history, '/signin')} to="/signin"> Signin</Link>
                        </li>
                        <li className="nav-item">
                            <Link className="nav-link" style={isActive(history, '/signup')} to="/signup"> Signup</Link>
                        </li>
                    </div>
                ) : null }

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
    )
}

export default withRouter(Navbar);