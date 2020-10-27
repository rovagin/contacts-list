import * as React from 'react';
import './index.css';
import {FontAwesomeIcon} from '@fortawesome/react-fontawesome';
import {faPlus} from '@fortawesome/free-solid-svg-icons';
import {Search} from "../../components/search";
import {ContactList} from "../../components/contact-list";
import {Contact} from "../../models/contact";
import api from "../../utils/api";
import {Col, Container, Row} from 'react-bootstrap';
import {Link} from 'react-router-dom';


interface Props {
}

interface State {
    searchResult: Contact|null;
    isLoading: boolean;
    contacts: Array<Contact>|null;
    userID: number;
}

export class ContactsPage extends React.Component<Props, State> {
    constructor(props: Props, state: State) {
        super(props, state);

        this.updateSearchResult = this.updateSearchResult.bind(this);

        this.state = {searchResult: null, contacts: null, isLoading: true, userID: 0};
    }

    async componentDidMount() {
        await this.loadContacts()
    }

    async loadContacts() {
        let contactsList = await api.get(`/${this.state.userID}/contacts`);

        this.setState({contacts: contactsList.data.payload});
        this.setState({isLoading: false});
    }

    updateSearchResult(contact: Contact|null) {
        this.setState({searchResult: contact});
    }

    render() {
        let contacts;

        if (this.state.searchResult != null) {
            contacts = [this.state.searchResult]
        } else {
            contacts = this.state.contacts as Array<Contact>
        }
        const contactsList = <ContactList
            reloadHandler={this.loadContacts}
            userID={this.state.userID}
            contacts={contacts}
            isLoading={this.state.isLoading}/>

        return (
            <Container fluid={"md"} className='contactsPage'>
                <Row>
                    <Col className='test groups'>
                        <Link to={'/'}>
                            Groups
                        </Link>
                    </Col>
                    <Col>
                        <Link className={'create-contact'}
                            to={{pathname: '/new', state: {userID: this.state.userID, newContactID: this.state.contacts?.length}}}>
                            <FontAwesomeIcon icon={faPlus} />
                        </Link>
                    </Col>
                </Row>
                <Row>
                    <h1 className='header'>Contacts</h1>
                </Row>
                <Row>
                    <Search searchResult={this.updateSearchResult} contacts={this.state.contacts as Array<Contact>}/>
                </Row>
                <Row>
                    <Col>
                        <p className='contact-avatar'>Pic</p>
                    </Col>
                    <Col >
                        <p className={'contact-name'}>My card</p>
                    </Col>
                </Row>
                <Row>
                    <Col>
                        {contactsList}
                    </Col>
                </Row>
            </Container>
        );
    }
}
