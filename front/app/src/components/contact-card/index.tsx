import * as React from 'react';
import './index.css';

export class ContactCard extends React.Component {
    render() {
        return (
            <div className='component'>
                <div className='block contact-avatar'>
                </div>
                <div className='block'>
                    <p>My card</p>
                </div>
            </div>
        );
    }
}
