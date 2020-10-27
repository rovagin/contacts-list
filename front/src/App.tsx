import React, {Component} from 'react';
import {BrowserRouter, Route} from 'react-router-dom';
import './App.css';
import {ContactsPage} from "./pages/contacts";
import {NewContactPage} from "./pages/new-contact";

class App extends Component<any, any>{
    render() {
        return (
            <BrowserRouter>
                <div className="App">
                    <Route  path='/' exact component={ContactsPage} />
                    <Route  path='/new' component={NewContactPage} />
                </div>
            </BrowserRouter>
        );
    }
}

export default App;
