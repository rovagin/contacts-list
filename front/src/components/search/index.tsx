import * as React from 'react';
import './index.css';
import {runLastTimeout} from "../../utils/jobs";
import {Contact} from "../../models/contact";
import {FormControl} from "react-bootstrap";

interface Props {
    contacts: Array<Contact>;
    searchResult: (res: Contact|null) => void;
}

interface State {
    value: string;
}

export class Search extends React.Component<Props, State> {
    constructor(props: Props) {
        super(props);
        this.state = {value: ''};

        this.handleChange = this.handleChange.bind(this);
    }

     handleChange(value: any) {
         this.setState({value: value}, () => {
            runLastTimeout(this, 'search', () => {
                if (value == '') {
                    this.props.searchResult(null);
                    return
                }

                let res = this.search(value);
                this.props.searchResult(res);
            }, 600);
         });
     }

    search(value: string): Contact|null {
        let searchString = value.toLowerCase();

        let contacts = this.props.contacts;
        for (let c of contacts) {
            if (c.first_name.toLowerCase().match(searchString)) {
                return c
            }
            if (c.last_name.toLowerCase().match(searchString)) {
                return c
            }
            if (c.phone.match(searchString)) {
                return c
            }
        }

        return null
    }

    render() {
        return (
            <FormControl
                placeholder='Search'
                type="text"
                value={this.state.value}
                onChange={e => this.handleChange(e.target.value)}
                className='searchInput'
            />
        );
    }
}
