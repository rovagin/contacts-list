import React from "react";
import Button from "react-bootstrap/Button";
import {Alert, Col, Container, Form, FormControl, FormGroup, Row} from "react-bootstrap";
import {Link, Redirect} from "react-router-dom";
import './index.css';
import {Contact} from "../../models/contact";
// @ts-ignore
import phone from 'phone';

interface Props {
    userID: number;
    newContactID: number;
    contact?: Contact;
    submitText: string;
    submitForm: (id: number, contact: Contact) => Promise<boolean>;
}

interface State {
    redirectToHome: boolean;
    contact: Contact;
    submitSuccess?: boolean;
    validationFail: boolean;
    invalidField: string;
}

export class ContactForm extends React.Component<Props, State> {
    constructor(props: Props) {
        super(props);

        let contact: Contact = {id: this.props.newContactID+1, phone: '', email: '', first_name: '', last_name: '', note: ''};
        if (props.contact != null) {
            let c = props.contact;
            contact = {
                id: c.id,
                phone: c.phone,
                email: c.email,
                first_name: c.first_name,
                last_name: c.last_name,
                note: c.note,
            }
        }

        console.log(contact);

        this.state = {redirectToHome: false, validationFail: false, contact: contact, invalidField: ''};

        this.handleSubmit = this.handleSubmit.bind(this);
        this.handleInputChange = this.handleInputChange.bind(this);
    }

    private async handleSubmit(e: React.FormEvent<HTMLFormElement>): Promise<void> {
        e.preventDefault();

        if (this.validateForm()) {
            const submitSuccess: boolean = await this.props.submitForm(this.props.userID, this.state.contact);
            this.setState({ redirectToHome: submitSuccess });
        }
    };

    handleInputChange(event: any) {
        const target = event.target;
        const field = target.name;
        const value = target.value;

        let currentState: any = this.state.contact;

        currentState[field] = value;

        this.setState({contact: currentState});
    }

    private validateForm(): boolean {
        if (this.state.contact.first_name === '') {
            this.setState({invalidField: 'first name', validationFail: true})
            return false;
        }
        if (this.state.contact.last_name === '') {
            this.setState({invalidField: 'last name', validationFail: true})
            return false;
        }
        if (!this.validateEmail(this.state.contact.email)) {
            this.setState({invalidField: 'email', validationFail: true})
            return false;
        }

        let res = phone(this.state.contact.phone)
        if (res.length ==0) {
            this.setState({invalidField: 'phone', validationFail: true})
            return false;
        }

        let contact = this.state.contact;
        contact.phone = res[0]

        this.setState({contact: contact});

        return true;
    }

    validateEmail(email: string) {
        const re = /^(([^<>()[\]\\.,;:\s@\"]+(\.[^<>()[\]\\.,;:\s@\"]+)*)|(\".+\"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
        return re.test(email);
    }

    render() {
        if (this.state.redirectToHome) {
            console.log("redirecting to home");
            return <Redirect push to={"/"}/>
        }

        let alert;
        if (this.state.validationFail) {
            alert = (<Alert variant={'info'} onClose={() => {this.setState({validationFail: false})}} dismissible={true}>
                Please check <b>{this.state.invalidField}</b> field for errors.
            </Alert>)
        }

        return (
            <Container fluid={"md"} className='newContactPage'>
                <Row className='navigate'>
                    <Col>
                        <Link to={'/'}>
                            Back
                        </Link>
                    </Col>
                </Row>
                <Row className='contactForm'>
                    <Col>
                        {alert}
                        <Form  onSubmit={this.handleSubmit}>
                            <FormGroup>
                                <FormControl
                                    className={'formElement'}
                                    type="email"
                                    placeholder="name@example.com"
                                    value={this.state.contact.email}
                                    name='email'
                                    onChange={this.handleInputChange}
                                />
                                <FormControl
                                    className={'formElement'}
                                    placeholder="First name"
                                    value={this.state.contact.first_name}
                                    name='first_name'
                                    onChange={this.handleInputChange}
                                />
                                <FormControl
                                    className={'formElement'}
                                    placeholder="Last name"
                                    value={this.state.contact.last_name}
                                    name='last_name'
                                    onChange={this.handleInputChange}
                                />
                                <FormControl
                                    className={'formElement'}
                                    placeholder='Phone'
                                    type="tel"
                                    value={this.state.contact.phone}
                                    name='phone'
                                    onChange={this.handleInputChange}
                                />
                                <FormControl
                                    className={'formElement'}
                                    placeholder='Note'
                                    value={this.state.contact.note}
                                    name='note'
                                    onChange={this.handleInputChange}
                                />
                            </FormGroup>
                            <Button variant="primary" type="submit">{this.props.submitText}</Button>
                        </Form>
                    </Col>
                </Row>
            </Container>
    );
    }
}