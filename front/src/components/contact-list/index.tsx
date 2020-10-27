import * as React from 'react';
import './index.css';
import {Contact} from "../../models/contact";
import {Button, ListGroup} from "react-bootstrap";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {faTrash} from "@fortawesome/free-solid-svg-icons/faTrash";
import api from '../../utils/api';

interface Props {
    contacts: Array<Contact>;
    isLoading: boolean;
    userID: number;
    reloadHandler: () => void;
}

interface State {
}

export class ContactList extends React.Component<Props> {
    constructor(props: Props) {
        super(props);

        this.handleRemove = this.handleRemove.bind(this);
    }

    private async handleRemove(id: any) {
        await api.delete(`/${this.props.userID}/contact/${id}`);

        this.props.reloadHandler();
    }

    render() {
        let content;

        if (this.props.isLoading) {
            content = <div>Loading</div>
        } else {
            content = (this.props.contacts.map((contact) =>
                <ListGroup.Item key={contact.id}>
                    <div className={'name'}>
                        <p>{contact.first_name + ' ' + contact.last_name}</p>
                    </div>
                    <Button variant={'danger'} className='removeBtn' type={"submit"} onClick={e => this.handleRemove(contact.id)}>
                        <FontAwesomeIcon icon={faTrash} />
                    </Button>
                </ListGroup.Item>)
            )
        }

        return (
            <ListGroup variant='flush'>
                {content}
            </ListGroup>
        );
    }
}
