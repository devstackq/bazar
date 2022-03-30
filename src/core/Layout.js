import React from 'react';
import Navbar from '../core/Navbar';

const Layout = ({ title = 'title', description = 'description', className, children }) => (

    <div>
        <Navbar />
        <div className="jumbotron">
            <h2>{title}</h2>
            <p className="lead">{description}</p>

        </div>
        <div className={className}> {children} </div>
    </div>
)

export default Layout;