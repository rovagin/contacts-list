import React from 'react';
import './App.css';
import {ContactsPage} from "./pages/contacts";
import {stub} from './utils/stub';

function App() {
  return (
    <div className="App">
      <ContactsPage contacts={stub}/>
    </div>
  );
}

export default App;
