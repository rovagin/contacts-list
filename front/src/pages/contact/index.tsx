import React from "react";
import './index.css';
import {Contact} from "../../models/contact";
import {ContactForm} from "../../components/contact-form";
import api from "../../utils/api";

interface State {
    contact: Contact;
}

export class ContactPage extends React.Component<any, State> {
    constructor(props: any) {
        super(props);

        console.log(props);

        this.state = {contact: props.location.state.contact}
    }

    async submitForm(id: number, contact: Contact): Promise<boolean> {
        console.log("submit form update", id, contact);

        await api.patch(`/${id}/contact`, {payload: {contact}});

        return true;
    }

    render() {
        return (
            <ContactForm
                submitText={'Update'}
                submitForm={this.submitForm}
                userID={0}
                newContactID={0}
                contact={this.state.contact}
            />
    );
    }
}