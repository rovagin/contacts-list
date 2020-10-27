import React from "react";
import Button from "react-bootstrap/Button";
import {Col, Container, Form, FormControl, FormGroup, Row} from "react-bootstrap";
import {Link, Redirect} from "react-router-dom";
import './index.css';
import api from "../../utils/api";
import {Contact} from "../../models/contact";
// @ts-ignore
import phone from 'phone';

interface Props {
    userID: number;
    newContactID: number;
}

interface State {
    redirectToHome: boolean;
    contact: Contact;
    submitSuccess?: boolean;
}

export class NewContactPage extends React.Component<Props, State> {
    constructor(props: Props) {
        super(props);

        this.state = {redirectToHome: false, contact: {id: this.props.newContactID+1, phone: '', email: '', first_name: '', last_name: '', note: ''}};

        this.handleSubmit = this.handleSubmit.bind(this);
        this.handleInputChange = this.handleInputChange.bind(this);
    }

    private async handleSubmit(e: React.FormEvent<HTMLFormElement>): Promise<void> {
        e.preventDefault();

        if (this.validateForm()) {
            const submitSuccess: boolean = await NewContactPage.submitForm(this.props.userID, this.state.contact);
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
        if (this.state.contact.first_name === '' || this.state.contact.last_name === '') {
            return false;
        }
        if (!this.validateEmail(this.state.contact.email)) {
            return false;
        }
        let res = phone(this.state.contact.phone)
        console.log(res);

        if (res.length ==0) {
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

    private static async submitForm(id: number, contact: Contact): Promise<boolean> {
        console.log(id, contact);

        await api.post(`/${id}/contact`, {payload: {contact: contact}});

        return true;
    }

    render() {
        if (this.state.redirectToHome) {
            console.log("redirecting to home");
            return <Redirect push to={"/"}/>
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
                            <Button variant="primary" type="submit">Create</Button>
                        </Form>
                    </Col>
                </Row>
            </Container>
    );
    }
}