import * as React from 'react';
import './index.css';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faPlus } from '@fortawesome/free-solid-svg-icons';
import {Search} from "../../components/search";
import {ContactCard} from "../../components/contact-card";
import {ContactList} from "../../components/contact-list";
import {Contact} from "../../models/contact";


interface Props {
    contacts: Array<Contact>;
}


export class ContactsPage extends React.Component<Props, {}> {
    constructor(props: Props) {
        super(props);
    }

    render() {
        return (
            <div>
                <div>
                    <div className='test groups'>
                        <a href="#">
                            Groups
                        </a>
                    </div>
                    <div className='test create-contact'>
                        <a href="https://google.com" target='_blank'>
                            <FontAwesomeIcon icon={faPlus} />
                        </a>
                    </div>
                </div>
                <div>
                    <h1 className='header'>Contacts</h1>
                </div>
                <Search/>
                <ContactCard/>
                <ContactList contacts={this.props.contacts}/>
            </div>
        );
    }
}
