import React, { useState } from 'react';

const Checkbox = ({ categories, handleFilters }) => {

    const [checked, setChecked] = useState([]);


    const handleToggle = category => () => {
        const currentCategoryId = checked.indexOf(category);
        const newCategoryId = [...checked];
        // if currently checked was not already in checked state > push
        // else pull/take off
        if (currentCategoryId === -1) {
            newCategoryId.push(category);
        } else {
            newCategoryId.splice(currentCategoryId, 1);
        }
        // console.log(newCheckedCategoryId);
        setChecked(newCategoryId);
        handleFilters(newCategoryId);

    };

    return categories.map((category, i) => (
        <span key={i}>
            <li className='list-unstyled'>
                <input onChange={handleToggle(category._id)}
                    value={checked.indexOf(category._id === -1)}
                    type='checkbox'
                    className="form-check-input" />
                <label className="form-check-label">{category.name}</label>
            </li>
        </span>


    ));
};

export default Checkbox;