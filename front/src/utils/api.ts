import axios from "axios";

export default axios.create({
    baseURL: "http://localhost/",
    responseType: "json",
    headers: {"Content-Type": "application/json"}
});