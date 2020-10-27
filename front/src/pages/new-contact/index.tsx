import React from "react";
import './index.css';
import {ContactForm} from "../../components/contact-form";
import {Contact} from "../../models/contact";
import api from "../../utils/api";

interface Props {
    userID: number;
    newContactID: number;
}

interface State {
}

export class NewContactPage extends React.Component<Props, State> {
    constructor(props: Props) {
        super(props);
    }

    async submitForm(id: number, contact: Contact): Promise<boolean> {
        console.log("submit form update", id, contact);

        await api.post(`/${id}/contact`, {payload: {contact: contact}});

        return true;
    }

    render() {
        return (
            <ContactForm
                submitText={'Save'}
                submitForm={this.submitForm}
                newContactID={this.props.newContactID+1}
                userID={this.props.userID}
            />
        );
    }
}