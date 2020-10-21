import * as React from 'react';
import './index.css';
import {Contact} from "../../models/contact";

interface Props {
    contacts: Array<Contact>;
}

export class ContactList extends React.Component<Props, {}> {
    constructor(props: Props) {
        super(props);
    }

    render() {
        return (
            <div className='component'>
                <ul>
                    {this.props.contacts.map((contact) =>
                        <li key={contact.id}>
                            {contact.first_name + ' ' + contact.last_name}
                        </li>
                    )}
                </ul>
            </div>
        );
    }
}
