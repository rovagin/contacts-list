import React, {Component} from 'react';
import {BrowserRouter, Route} from 'react-router-dom';
import './App.css';
import {ContactsPage} from "./pages/contacts";
import {NewContactPage} from "./pages/new-contact";
import {ContactPage} from "./pages/contact";

class App extends Component<any, any>{
    render() {
        return (
            <BrowserRouter>
                <div className="App">
                    <Route  path='/' exact component={ContactsPage} />
                    <Route  path='/new' component={NewContactPage} />
                    <Route  path='/:userID/contact/:id' component={ContactPage} />
                </div>
            </BrowserRouter>
        );
    }
}

export default App;
